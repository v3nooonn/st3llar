package config

type St3llarConfig struct {
	Environment  string `yaml:"environment"`
	EnvVarPrefix string `yaml:"env_var_prefix"`
}
