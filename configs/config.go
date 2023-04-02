package configs

import (
	"fmt"

	logger "github.com/Vlad1slavZhuk/grpc-postgresql/pkg/log"
	"github.com/spf13/viper"
)

type Config struct {
	GRPCSettings       `mapstructure:"grpc"`
	PostgreSQLSettings `mapstructure:"postgresql"`
}

type GRPCSettings struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type PostgreSQLSettings struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func ReadConfig() *Config {
	log := logger.GetLoggerInstance()
	config := new(Config)

	config.setDefault()
	log.Warn().Msg("set default parameter")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")     // root project
	viper.AddConfigPath("/app/") // docker place
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("fatal error config file")
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal().Err(err).Msg("fatal error unmarshaling")
	}

	return config
}

func (c *Config) setDefault() {
	c.GRPCSettings.Host = "localhost"
	c.GRPCSettings.Port = 443

	c.PostgreSQLSettings.Database = "postgres"
	c.PostgreSQLSettings.Username = "postgres"
	c.PostgreSQLSettings.Password = "password"
}

func (p *PostgreSQLSettings) DSN() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		p.Host, p.Username, p.Password, p.Database, p.Port,
	)
}
