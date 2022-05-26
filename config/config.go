package config

import "github.com/caarlos0/env"

// App contains needed envs to run service
type AppConfig struct {
	Host string `env:"HOST" envDefault:"127.0.0.1"`
	Port string `env:"PORT" envDefault:"8089"`
	KafkaPort string `env:"KAFKA_PORT" envDefault:"9092"`
	ZookeeperPort string `env:"ZOOKEEPER_PORT" envDefault:"2181"`
}

// App parses envs and constructs the config
func NewAppConfig() (*AppConfig, error) {
	var initiatorConfig AppConfig

	if err := env.Parse(&initiatorConfig); err != nil {
		return nil, err
	}

	return &initiatorConfig, nil
}
