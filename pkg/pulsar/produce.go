package pulsar

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/xuexiangyou/code-art/config"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/fx"
	"time"
)

type TencentPulsarClient struct {
	Config *config.Config
	Client pulsar.Client
	ResourceManager *ResourceManager
}

//NewTencentPulsarClient 初始化链接
func NewTencentPulsarClient(lc fx.Lifecycle, config *config.Config) (*TencentPulsarClient, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               config.PulsarConfig.EndPoint,
		Authentication:    pulsar.NewAuthenticationToken(config.PulsarConfig.Token),
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		logx.Error("Could not instantiate Pulsar client: %v", err)
		return nil, err
	}

	tencentPulsarClient := &TencentPulsarClient{
		Config: config,
		Client: client,
		ResourceManager: NewResourceManager(),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			tencentPulsarClient.Close()
			return nil
		},
	})

	return tencentPulsarClient, nil
}

//PulsarProducer 发送队列消息
func (p *TencentPulsarClient) PulsarProducer(topic string, data string, delay int64) (pulsar.MessageID, error) {
	producer, err := p.getProducerClient(p.Config.PulsarConfig.ColonyId + topic)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	msgId, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload:      []byte(data),
		DeliverAfter: time.Duration(delay) * time.Second,
	})
	//logx.Infof("msgId type：%T\n", msgId)
	//defer producer.Close()

	if err != nil {
		logx.Error("Failed to publish message", err)
		return nil, err
	}

	return msgId, nil
}

//getProducerClient 获取Producer链接
func (p *TencentPulsarClient) getProducerClient(topic string) (pulsar.Producer, error) {
	val, err := p.ResourceManager.GetResource(topic, func() (pulsar.Producer, error) {
		return p.newProducerClient(topic)
	})

	if err != nil {
		return nil, err
	}

	return val, nil
}

//newProducerClient 初始化Producer
func (p *TencentPulsarClient) newProducerClient(topic string) (pulsar.Producer, error) {
	producer, err := p.Client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})

	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return producer, nil
}

//Close 关闭链接
func (p *TencentPulsarClient) Close() {
	p.Client.Close()
	p.ResourceManager.Close()
}

