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
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"

	//identifier, literal
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	//symbols
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	BANG     = "!"
	SLASH    = "/"
	EQ       = "=="
	NOT_EQ   = "!="

	LT       = "<"
	GT       = ">"
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	COMMA     = ","
	SEMICOLON = ";"

	//necessarily needed
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
