package config

//todo bring config from yaml

/* MarketApiServer */

func GetTrade() Application {
	return Application{
		Http: Http{
			Port:    "8090",
			URLPath: "/api/trade",
		},
		DataBase: Mongo{
			DBName:   "market",
			Host:     "172.200.0.10",
			Port:     "27017",
			User:     "root",
			Password: "secret",
		},
		LogLevel: "info",
	}
}

func GetStatistic() Application {
	return Application{
		Http: Http{
			Port:    "8092",
			URLPath: "/api/statistic",
		},
		DataBase: Mongo{
			DBName:   "market",
			Host:     "172.200.0.10",
			Port:     "27017",
			User:     "root",
			Password: "secret",
		},
		LogLevel: "debug",
	}
}

type Application struct {
	Http     Http
	DataBase Mongo
	LogLevel string
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

/* FutureFlag */

func GetFutureFlag() FutureFlag {
	return FutureFlag{
		statisticAsDomain: false,
	}
}

type FutureFlag struct {
	//statisticAsDomain is responsible for switching that statistic will be calculated as domain.Trade or it owns internal representation as independent module
	statisticAsDomain bool
}
