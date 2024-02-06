package principal

type TokenType string

const (
	TokenTypeSession = TokenType("session")
	TokenTypeAccount = TokenType("account")
)

func (t TokenType) Header() string {
	switch t {
	case TokenTypeAccount:
		return "x-account"
	case TokenTypeSession:
		return "x-session"
	}
	return ""
}

type Token struct {
	T TokenType
	V string
}

func (t Token) Type() TokenType {
	return t.T
}

func (t Token) Value() string {
	return t.V
}
