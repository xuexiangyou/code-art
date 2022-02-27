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
			DataSource: "",
		},
		Redis: struct {
			Addr string
			Pass string
		}{Addr: "", Pass: ""},
	}
}
