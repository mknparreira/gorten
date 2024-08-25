package config

type MongoConfig struct {
	DBName        string `mapstructure:"db_name"`
	ConnectionURL string `mapstructure:"connection_url"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type AppConfig struct {
	Mongo  MongoConfig  `mapstructure:"mongo"`
	Server ServerConfig `mapstructure:"server"`
}
