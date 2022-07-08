package token

type TokenType string

type Token struct {
	Type    TokenType
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
