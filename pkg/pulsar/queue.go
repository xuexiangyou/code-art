package pulsar

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/xuexiangyou/code-art/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/queue"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/threading"
	//"gitlab.yunxi.tv/yxgo/yxgo-queue/common"
	"log"
	"time"
)

type CallHandler func(pulsar.Message) (interface{}, error)

type (
	PulsarQueues struct {
		c      *config.Config
		client pulsar.Client
		queues []queue.MessageQueue
		group  *service.ServiceGroup
	}

	PulsarQueue struct {
		c config.PulsarTopic

		routerMap map[string]CallHandler

		routerType       string
		channel          chan pulsar.Message
		consumer         pulsar.Consumer
		producerRoutines *threading.RoutineGroup
		consumerRoutines *threading.RoutineGroup
	}
)

func MustNewQueues(c *config.Config, routerMap map[string]CallHandler) *PulsarQueues {
	q, err := newPulsarQueues(c, routerMap)
	if err != nil {
		log.Fatal(err)
	}
	return q
}

func newPulsarQueues(c *config.Config, routerMap map[string]CallHandler) (*PulsarQueues, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               c.PulsarConfig.EndPoint,
		Authentication:    pulsar.NewAuthenticationToken(c.PulsarConfig.Token),
		OperationTimeout:  time.Duration(c.PulsarConfig.OperationTimeout) * time.Second,
		ConnectionTimeout: time.Duration(c.PulsarConfig.ConnectionTimeout) * time.Second,
	})
	if err != nil {
		return nil, err
	}

	q := &PulsarQueues{
		group:  service.NewServiceGroup(),
		client: client,
		c:      c,
	}

	//循环处理逻辑
	for _, topic := range c.PulsarConfig.Topics {
		for _, subscription := range topic.Subscriptions {
			pLink, err := newPulsarQueue(client, topic, subscription, routerMap)
			if err != nil {
				logx.Errorf("newPulsarQueueFail, %q", err.Error())
				continue
			}
			q.queues = append(q.queues, pLink)
		}
	}
	return q, err
}

func (p *PulsarQueues) Start() {
	for _, each := range p.queues {
		p.group.Add(each)
	}
	p.group.Start()
}

func (p *PulsarQueues) Stop() {
	p.group.Stop()
	p.client.Close() //关闭pulsar的client链接
}

func newPulsarQueue(client pulsar.Client, c config.PulsarTopic, subscription config.PulsarSubscription, routerMap map[string]CallHandler) (queue.MessageQueue, error) {
	options := pulsar.ConsumerOptions{
		Topic:            c.TopicUrl,
		SubscriptionName: subscription.Name,
		Type:             pulsar.Shared,
	}
	consumer, err := client.Subscribe(options)
	if err != nil {
		logx.Errorf("newPulsarSubscribeFail, %s", err.Error())
		return nil, err
	}

	return &PulsarQueue{
		c:                c,
		routerMap:        routerMap,
		routerType:       subscription.Router,
		consumer:         consumer,
		channel:          make(chan pulsar.Message),
		producerRoutines: threading.NewRoutineGroup(),
		consumerRoutines: threading.NewRoutineGroup(),
	}, nil
}

func (p *PulsarQueue) Start() {
	p.startConsumers()
	p.startProducers()
	p.producerRoutines.Wait()
	close(p.channel)
	p.consumerRoutines.Wait()
}

func (p *PulsarQueue) Stop() {
	p.consumer.Close()
	//logx.Close() //todo 日志关闭
}

func (p *PulsarQueue) startConsumers() {
	for i := 0; i < p.c.Processors; i++ {
		p.consumerRoutines.Run(func() {
			for msg := range p.channel {
				callHandler, ok := p.routerMap[p.routerType]
				if !ok {
					log.Println("不存在改方法路由", p.routerType)
					continue
				}
				response, err := callHandler(msg)
				if err != nil {
					log.Print("消费失败", err)
					continue
				}
				fmt.Println("-------", string(msg.Payload()), "----", msg.ID(), "----", response)
				p.consumer.Ack(msg)
			}
		})
	}
}

func (p *PulsarQueue) startProducers() {
	for i := 0; i < p.c.Consumers; i++ {
		p.producerRoutines.Run(func() {
			for {
				msg, err := p.consumer.Receive(context.Background())
				if value, ok := err.(*pulsar.Error); ok { //判断是否关闭了链接则直接退出
					if value.Result() == pulsar.ConsumerClosed {
						//fmt.Println("consumer closed: ConsumerClosed", err)
						return
					}
				}
				if err != nil {
					logx.Errorf("pulsarReceiveMsgFail,%s", err.Error())
					continue
				}
				p.channel <- msg
			}
		})
	}
}
