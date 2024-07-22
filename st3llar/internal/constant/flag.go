package constant

type FlagName string

const (
	FLogLevel     FlagName = "log-level"
	FEnvPrefix    FlagName = "env-prefix"
	FEnvironment  FlagName = "environment"
	FOrganization FlagName = "organization"
	FCredential   FlagName = "credential"
)

func (f FlagName) ValStr() string {
	return string(f)
}
