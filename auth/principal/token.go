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
