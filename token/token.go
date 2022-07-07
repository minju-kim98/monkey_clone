package token

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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
