package config

type Config struct {
	HTTP    HTTP
	MongoDB MongoDB
}

type MongoDB struct {
	DBName string
	Host   string
	Port   int
}

type HTTP struct {
	Host string
	Port int
}

func New() Config {
	return Config{
		HTTP: HTTP{
			Host: "",
			Port: 8080,
		},
		MongoDB: MongoDB{
			DBName: "medods",
			Host:   "mongodb",
			Port:   27017,
		},
	}
}
