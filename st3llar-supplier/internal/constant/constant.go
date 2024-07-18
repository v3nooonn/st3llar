package constant

type StatusCode int

const (
	Success StatusCode = iota
	ErrReadPassword
)

func (c StatusCode) Int() int {
	return int(c)
}
