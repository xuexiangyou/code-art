package config

type Config struct {
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Addr string
		Pass string
	}
}

//NewConfig 初始化配置
func NewConfig() *Config {
	return &Config{
		Mysql: struct {
			DataSource string
		}{
			DataSource: "whj:)OKM9ijn@tcp(rm-bp18h27wrt8mhty46o.mysql.rds.aliyuncs.com)/activity?charset=utf8mb4&parseTime=true",
		},
		Redis: struct {
			Addr string
			Pass string
		}{Addr: "r-bp17bf1b1a9f3824pd.redis.rds.aliyuncs.com:6379", Pass: "Zhenzhen123"},
	}
}
