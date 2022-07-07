package token

import "fmt"

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLIGAL = "ILLIGAL"
	EOF     = "EOF"

	// identifier + Literal
	IDENT = "IDENT"
	INT   = "INT"

	// operator
	ASSIGN = "="
	PLUS   = "+"

	//delimiter
	COMMA     = ","
	SEMICOLON = ";"

	// symbols
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTON"
	LET      = "LET"
)

func PrintHello() {
	fmt.Println("Hello, Monkey!")
}
