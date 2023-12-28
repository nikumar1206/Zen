package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	// bad code
	ILLEGAL TokenType = "ILLEGAL"
	// operators
	ASSIGN       = "="
	ADD          = "+"
	SUBTRACT     = "-"
	MULTIPLY     = "*"
	DIVIDE       = "/"
	FLOOR_DIVIDE = "//"
	NOT          = "!"
	LT           = "<"
	GT           = ">"
	LTE          = "<="
	GTE          = ">="
	EQUALITY     = "=="
	NOT_EQUAL    = "!="

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	ELIF     = "ELIF"
	LOG      = "LOG"
	// definitions
	LITERAL = "LITERAL"
	NUMBER  = "NUMBER"
	EOF     = "EOF"
	RETURN  = "RETURN"

	// semantics
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	LBRACKET  = "LBRACKET"
	RBRACKET  = "RBRACKET"
	SEMICOLON = ";"
)

func NewToken(t TokenType, v string) *Token {
	return &Token{
		Type:  t,
		Value: v,
	}
}
