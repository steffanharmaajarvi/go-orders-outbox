package koanf

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type ExternalConfig struct {
	wrapped *koanf.Koanf
}

func (e ExternalConfig) GetString(path string) string {
	return e.wrapped.String(path)
}

func (e ExternalConfig) MapKeys(path string) []string {
	return e.wrapped.MapKeys(path)
}

func (e ExternalConfig) GetInt(path string) int {
	return e.wrapped.Int(path)
}

func New() *ExternalConfig {

	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		panic("Cant determine current working directory!")
	}

	var configFiles = []string{}

	// Default framework config
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Cant determine default config path!")
	}
	configFiles = append(configFiles, fmt.Sprintf("%s/%s", filepath.Dir(filename), "default.yml"))

	// Default kubernetes config
	configFiles = append(configFiles, fmt.Sprintf("%s/default.yml", wd))

	// Local dev config
	config_env := os.Getenv("CONFIG_ENV")

	if config_env != "" {
		configFiles = append(configFiles, fmt.Sprintf("%s/configs/.env.%s.yml", wd, config_env))
	} else {
		configFiles = append(configFiles, fmt.Sprintf("%s/configs/.env.yml", wd))
	}

	koanfInstance := koanf.New(".")
	for _, configFile := range configFiles {
		if _, err := os.Stat(configFile); err == nil {
			if err := koanfInstance.Load(file.Provider(configFile), yaml.Parser()); err != nil {
				log.Error().Err(err).Msgf("Error read config file: %s", configFile)
			} else {
				log.Info().Msgf("Loaded config file: %s", configFile)
			}
		} else {
			log.Info().Msgf("Config file not found: %s", configFile)
		}
	}

	// Load env variables
	if err := koanfInstance.Load(env.Provider("APP_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "APP_")), "__", ".", -1)
	}), nil); err != nil {
		log.Fatal().Err(err).Msg("Error loading env variables")
	}

	return &ExternalConfig{
		wrapped: koanfInstance,
	}
}
