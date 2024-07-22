package config

import "github.com/v3nooom/st3llar/internal/constant"

type (
	St3llarConfig struct {
		Environment  string `yaml:"environment"`
		LogLevel     string `yaml:"log_level"`
		EnvPrefix    string `yaml:"env_prefix"`
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
		sc.LogLevel = constant.Warn.ValStr()
		sc.EnvPrefix = constant.EnvPrefix.ValStr()
		sc.Organization = constant.Organization.ValStr()
	}
}

func WithEnvironment(environment string) St3llarConfigOpt {
	return func(sc *St3llarConfig) {
		sc.Environment = environment
	}
}

func WithLogLevel() St3llarConfigOpt {
	return func(sc *St3llarConfig) {
		sc.LogLevel = constant.Warn.ValStr()
	}
}

func WithEnvPrefix() St3llarConfigOpt {
	return func(sc *St3llarConfig) {
		sc.EnvPrefix = constant.EnvPrefix.ValStr()
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
