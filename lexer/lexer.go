package lexer

import (
	"horse-lang/token"
	"horse-lang/utils"
)

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

// Gives the token from the input string (it is like Scanner.next() in Java)
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

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
	default:
		if utils.IsLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdentifier(tok.Literal)
			return tok
		} else if utils.IsNumber(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// Return the identifier literal
func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for utils.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

// Return number identifier
func (l *Lexer) readNumber() string {
	startPosition := l.position
	for utils.IsNumber(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

// Skips over all the white-space between tokens
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
