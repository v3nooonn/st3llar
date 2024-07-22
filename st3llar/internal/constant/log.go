package constant

type LogLevel string

const (
	Debug LogLevel = "Debug"
	Info  LogLevel = "Info"
	Warn  LogLevel = "Warn"
	Error LogLevel = "Error"
)

func (ll LogLevel) ValStr() string {
	return string(ll)
}
