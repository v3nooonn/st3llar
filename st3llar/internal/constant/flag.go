package constant

type FlagName string

const (
	FlagLogLevel     FlagName = "log-level"
	FlagEnvPrefix    FlagName = "env-prefix"
	FlagEnvironment  FlagName = "environment"
	FlagOrganization FlagName = "organization"
	FlagCredential   FlagName = "credential"
)

func (f FlagName) ValStr() string {
	return string(f)
}
