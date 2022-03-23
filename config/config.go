package config

//Config 定义配置变量
type Config struct {
	Env 	 string
	Host     string
	Port     int
	Log struct {
		WebLogPath string
		AppLogPath string
	}
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Addr string
		Pass string
	}
}