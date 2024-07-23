package config

import (
	"fmt"
	"os"

	"github.com/v3nooom/st3llar/internal/constant"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type (
	St3llarConfig struct {
		Environment  string `yaml:"environment"`
		LogLevel     string `yaml:"log-level"`
		EnvPrefix    string `yaml:"env-prefix"`
		Organization string `yaml:"organization"`
		Credential   string `yaml:"credential"`
	}
	St3llarConfigOpt func(sc *St3llarConfig)
)

func Build(opts ...St3llarConfigOpt) *St3llarConfig {
	sc := new(St3llarConfig)

	for _, opt := range opts {
		opt(sc)
	}

	return sc
}

func WithDefault() St3llarConfigOpt {
	return func(sc *St3llarConfig) {
		sc.Environment = constant.Environment.ValStr()
		sc.LogLevel = constant.Info.ValStr()
		sc.EnvPrefix = constant.EnvPrefix.ValStr()
		sc.Organization = constant.Organization.ValStr()
	}
}

func WithEnvironment(environment string) St3llarConfigOpt {
	return func(sc *St3llarConfig) {
		sc.Environment = environment
	}
}

func WithLogLevel(logLevel string) St3llarConfigOpt {
	return func(sc *St3llarConfig) {
		sc.LogLevel = logLevel
	}
}

func WithEnvPrefix(prefix string) St3llarConfigOpt {
	return func(sc *St3llarConfig) {
		sc.EnvPrefix = prefix
	}
}

func WithOrganization(organization string) St3llarConfigOpt {
	return func(sc *St3llarConfig) {
		sc.Organization = organization
	}
}

func WithCredential(credential string) St3llarConfigOpt {
	return func(sc *St3llarConfig) {
		sc.Credential = credential
	}
}

func Home() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return home
}

func WriteConfig(cfg *St3llarConfig, path string) error {
	yamlBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshalling default config error: %w", err)
	}

	if err := os.WriteFile(path, yamlBytes, 0666); err != nil {
		return fmt.Errorf("writing config file error, %w: %s\n", err, path)
	}

	return nil
}

func ReadConfig(path string) (*St3llarConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := new(St3llarConfig)

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
