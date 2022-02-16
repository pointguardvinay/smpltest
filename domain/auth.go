package domain

type Auth struct {
	id       string
	password string
	role     string
}

type authrepository interface {
	authenticate(aut Auth)
}
