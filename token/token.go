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
