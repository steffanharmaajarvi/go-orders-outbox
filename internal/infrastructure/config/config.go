package config

import "orders.go/m/internal/infrastructure/config/koanf"

type ConfigProvider string

const (
	ConfigProviderKoanfs ConfigProvider = "koanf"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	MapKeys(path string) []string
}

func NewConfig(provider ConfigProvider) Config {
	switch provider {
	case ConfigProviderKoanfs:
		return koanf.New()
	default:
		return nil
	}
}
