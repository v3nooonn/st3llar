package repr

type Login struct {
	Account      string `json:"account"`
	Organization string `json:"organization"`
	Password     string `json:"password"`
}
