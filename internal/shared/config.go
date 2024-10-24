package shared

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Port         int    `envconfig:"PORT" default:"5000"`
	GithubApiURI string `envconfig:"GITHUB_API_URI" default:"https://api.github.com"`
	Workers      int    `envconfig:"WORKERS" default:"5"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to build config from env")
	}
	return &cfg, nil
}
