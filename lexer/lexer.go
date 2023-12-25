package lexer

import "horse-lang/token"

type Lexer struct {
	input        string // the entire input string to be parsed
	position     int    // Current position in input (points to current character)
	readPosition int    // Points the character after the current character
	ch           byte   // current character
}

// "Converts" input string to Lexer data type
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar() // read the first character
	return l
}

// Gives the next character and advances the position of the pointer
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}
