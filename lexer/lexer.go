package lexer

import (
	"testing"

	"github.com/minju-kim98/monkey/token"
)

type Lexer struct {
	input        string
	posision     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests [%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, &tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests [%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, &tok.Literal)
		}
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.posision = l.readPosition
	l.readPosition += 1
}
