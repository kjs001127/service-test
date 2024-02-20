package principal

type Token interface {
	Value() string
	Type() string
}
