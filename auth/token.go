package auth

type Token interface {
	Value() string
}
