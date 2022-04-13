package config

//Config 定义配置变量
type Config struct {
	Env 	 string
	Host     string
	Port     int
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Addr string
		Pass string
	}

	PulsarConfig struct {
		ColonyId		  string
		EndPoint          string //pulsar 队列接入地址
		Token             string //pulsar 接入token
		OperationTimeout  int    `json:",default=30"` //操作超时时间
		ConnectionTimeout int    `json:",default=30"` //链接超时时间

		Topics []PulsarTopic //pulsar topic 列表
	}
}


type PulsarTopic struct {
	TopicUrl          	string   				//topic地址
	Subscriptions 	  	[]PulsarSubscription 	//topic订阅名列表
	Type              	int      				//分享类型
	Consumers  		  	int 		`json:",default=8"`
	Processors 		  	int 		`json:",default=8"`
}

type PulsarSubscription struct {
	Name 			 	string
	Router     			string
}