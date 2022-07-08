package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	//keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"

	//identifier, literal
	IDENT = "IDENT"
	INT   = "INT"

	//symbols
	ASSIGN    = "="
	PLUS      = "+"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","
	SEMICOLON = ";"

	//necessarily needed
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
