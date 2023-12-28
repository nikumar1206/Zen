package lexer

import (
	"fmt"
	"strconv"

	"interpreter/tokens"
)

type Lexer struct {
	Buffer       string
	Tokens       []*token.Token
	Position     int
	PeakPosition int
	CurrChar     byte
	BuffLength   int
}

func NewLexer(b string) *Lexer {
	lex := &Lexer{
		Buffer:     b,
		Tokens:     make([]*token.Token, 0),
		BuffLength: len(b),
	}
	lex.readChars()
	return lex
}

func (l *Lexer) NextToken() (*token.Token, error) {
	var newToken *token.Token

	l.skipWhitespace()
	b := l.CurrChar

	switch b {
	case '+':
		newToken = token.NewToken(token.ADD, string(b))
	case '-':
		newToken = token.NewToken(token.SUBTRACT, string(b))
	case '*':
		newToken = token.NewToken(token.MULTIPLY, string(b))
	case '/':
		newToken = token.NewToken(token.DIVIDE, string(b))
	case '=':
		if l.peakChar() == '=' {
			newToken = token.NewToken(token.EQUALITY, "==")
		} else {

			newToken = token.NewToken(token.ASSIGN, string(b))
		}
	case '!':
		if l.peakChar() == '=' {
			newToken = token.NewToken(token.NOT_EQUAL, "!=")
		} else {
			newToken = token.NewToken(token.NOT, string(b))
		}
	case '<':
		if l.peakChar() == '=' {
			newToken = token.NewToken(token.LTE, "<=")
		} else {
			newToken = token.NewToken(token.LT, string(b))
		}
	case '>':
		if l.peakChar() == '=' {
			newToken = token.NewToken(token.GTE, ">=")
		} else {
			newToken = token.NewToken(token.GT, string(b))
		}
	case '(':
		newToken = token.NewToken(token.LPAREN, string(b))
	case ')':
		newToken = token.NewToken(token.RPAREN, string(b))
	case '{':
		newToken = token.NewToken(token.LBRACKET, string(b))
	case '}':
		newToken = token.NewToken(token.RBRACKET, string(b))
	case ';':
		newToken = token.NewToken(token.SEMICOLON, string(b))
	case '\'', '"':
		l.readChars()
		literal, err := l.readLiteralIdentifier(b)
		if err != nil {
			return nil, err
		}
		newToken = token.NewToken(token.LITERAL, literal)
	case 0:
		newToken = token.NewToken(token.EOF, string(b))

	default:
		if l.isAlpha(b) {
			poss_token := l.readStringIdentifier()
			switch poss_token {

			case "return":
				newToken = token.NewToken(token.RETURN, poss_token)
			case "log":
				newToken = token.NewToken(token.LOG, poss_token)
			case "func", "fn":
				newToken = token.NewToken(token.FUNCTION, poss_token)
			case "let":
				newToken = token.NewToken(token.LET, poss_token)
			case "const":
				newToken = token.NewToken(token.CONST, poss_token)
			case "if":
				newToken = token.NewToken(token.IF, poss_token)
			case "else":
				newToken = token.NewToken(token.ELSE, poss_token)
			case "elif":
				newToken = token.NewToken(token.ELIF, poss_token)
			case "true":
				newToken = token.NewToken(token.TRUE, poss_token)
			case "false":
				newToken = token.NewToken(token.FALSE, poss_token)
			default:
				newToken = token.NewToken(token.LITERAL, poss_token)
			}
			return newToken, nil
		} else if l.isNumeric(b) {
			poss_token := l.readNumericIdentifier()

			newToken = token.NewToken(token.NUMBER, poss_token)
			return newToken, nil
		} else {
			newToken = token.NewToken(token.ILLEGAL, "")
		}
	}
	if newToken.Type == token.ILLEGAL {
		return nil, fmt.Errorf("Token %q is an illegal token. Please verify your syntax on index %d", b, l.Position)
	}

	l.readChars()

	return newToken, nil
}

func (l *Lexer) Tokenize() ([]*token.Token, error) {
	for {
		next_token, err := l.NextToken()

		if err != nil {
			return nil, err
		}

		if next_token.Type == token.EOF {
			break
		}
		l.Tokens = append(l.Tokens, next_token)
	}
	return l.Tokens, nil
}

func (l *Lexer) readChars() {

	if l.PeakPosition < l.BuffLength {
		l.Position = l.PeakPosition
		l.PeakPosition += 1
		l.CurrChar = l.Buffer[l.Position]
	} else if l.PeakPosition >= l.BuffLength {
		l.CurrChar = 0
		l.Position = l.PeakPosition
		l.PeakPosition += 1
	}
}

func (l *Lexer) peakChar() byte {
	if l.PeakPosition >= l.BuffLength {
		return 0
	} else {
		return l.Buffer[l.PeakPosition]
	}
}

func (l *Lexer) isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func (l *Lexer) isNumeric(char byte) bool {
	_, err := strconv.ParseFloat(string(char), 64)
	return err == nil
}

func (l *Lexer) readStringIdentifier() string {
	currPos := l.Position
	for l.isAlpha(l.CurrChar) {
		l.readChars()
	}
	return l.Buffer[currPos:l.Position]
}

func (l *Lexer) readNumericIdentifier() string {
	currPos := l.Position
	for l.isNumeric(l.CurrChar) {
		l.readChars()
	}
	return l.Buffer[currPos:l.Position]
}

func (l *Lexer) readLiteralIdentifier(sID byte) (string, error) {
	currPos := l.Position
	for l.CurrChar != sID {

		if l.PeakPosition > l.BuffLength {
			return "", fmt.Errorf("String literal is not closed.")
		}
		l.readChars()
	}
	return l.Buffer[currPos:l.Position], nil

}

func (l *Lexer) skipWhitespace() {
	for l.CurrChar == ' ' || l.CurrChar == '\t' || l.CurrChar == '\n' || l.CurrChar == '\r' {
		l.readChars()
	}
}
