package token

type tokenType string

type Token struct {
	Type    tokenType
	Literal string
}

const (
	//keywords
	LET      = "LET"
	FUNCTION = "fn"

	//identifier
	IDENT = "IDENT"

	//symbols
	ASSIGN    = "="
	PLUS      = "+"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","
	SEMICOLON = ";"
)
