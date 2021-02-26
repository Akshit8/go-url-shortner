package config

import "github.com/spf13/viper"

// AppConfig stores all configuration of the application.
type AppConfig struct {
	RestPort      int    `mapstructure:"REST_PORT"`
	GraphqlPort   int    `mapstructure:"GRAPHQL_PORT"`
	Host          string `mapstructure:"HOST"`
	RepoType      string `mapstructure:"REPO_TYPE"`
	DbName        string `mapstructure:"DB_NAME"`
	URLTable      string `mapstructure:"URL_TABLE"`
	MongoURI      string `mapstructure:"MONGO_URI"`
	RedisURI      string `mapstructure:"REDIS_URI"`
	CassandraHost string `mapstructure:"CASSANDRA_HOST"`
	CassandraPort int    `mapstructure:"CASSANDRA_PORT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (AppConfig, error) {
	var config AppConfig
	var err error

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)

	return config, err
}
