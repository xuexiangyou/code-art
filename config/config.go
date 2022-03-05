package config

//定义配置变量
type Config struct {
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