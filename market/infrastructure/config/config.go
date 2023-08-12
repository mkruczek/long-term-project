package config

//todo bring config from yaml

func GetMarket() Market {
	return Market{
		Http: Http{
			Port:    "8090",
			URLPath: "/api/market",
		},
		Mongo: Mongo{
			DBName:   "market",
			Host:     "172.200.0.10",
			Port:     "27017",
			User:     "root",
			Password: "secret",
		},
	}
}

type Market struct {
	Http  Http
	Mongo Mongo
}

type Http struct {
	Port    string
	URLPath string
}

type Mongo struct {
	DBName   string
	Host     string
	Port     string
	User     string
	Password string
}
