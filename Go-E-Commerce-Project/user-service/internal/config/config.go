package config

import (
	"shared/pkg/config"
)

type Config = config.Config

func Load() Config {
	return config.Load()
}
