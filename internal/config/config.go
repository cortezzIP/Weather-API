package config

import (
	"os"
)

type Config struct {
	WeatherServiceConfig
	RedisConfig
	ListenAddress string
}

type WeatherServiceConfig struct {
	APIBaseUrl string
	APIKey     string
}

type RedisConfig struct {
	ConnectionString string
}

func MustLoad() *Config {
	listenAddress := os.Getenv("LISTEN_ADDR")
	if listenAddress == "" {
		panic("failed to get env var")
	}

	apiBaseURL := os.Getenv("WEATHER_API_BASE_URL")
	if apiBaseURL == "" {
		panic("failed to get env var")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		panic("failed to get env var")
	}
	
	redisConnectionString := os.Getenv("REDIS_CONN_STRING")
	if redisConnectionString == "" {
		panic("failed to get env var")
	}

	return &Config{
		ListenAddress: listenAddress,
		RedisConfig: RedisConfig{
			ConnectionString: redisConnectionString,
		},
		WeatherServiceConfig: WeatherServiceConfig{
			APIBaseUrl: apiBaseURL,
			APIKey:     apiKey,
		},
	}
}
