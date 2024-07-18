package oauth

type OAuth interface {
	Login(o, u, p string) error
}
