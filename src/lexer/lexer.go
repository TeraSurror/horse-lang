package lexer

import (
	"horse-lang/src/token"
)

type Lexer struct {
	source  string
	Tokens  []token.Token
	start   int
	current int
	line    int
}

func New(source string) *Lexer {
	return &Lexer{
		source:  source,
		Tokens:  []token.Token{},
		start:   0,
		current: 0,
		line:    1,
	}
}

func (lexer *Lexer) ScanTokens() {
	for !lexer.isAtEnd() {
		lexer.start = lexer.current
		lexer.scanToken()
	}

	lexer.Tokens = append(lexer.Tokens, token.Token{
		Token:   token.EOF,
		Lexeme:  "",
		Literal: "",
		Line:    lexer.line,
	})
}

func (lexer *Lexer) isAtEnd() bool {
	return lexer.current >= len(lexer.source)
}

func (lexer *Lexer) scanToken() {
	c := lexer.advance()
	switch c {
	case '(':
		lexer.addToken(token.LEFT_PAREN)
	case ')':
		lexer.addToken(token.RIGHT_PAREN)
	case '{':
		lexer.addToken(token.LEFT_BRACE)
	case '}':
		lexer.addToken(token.RIGHT_BRACE)
	case ',':
		lexer.addToken(token.COMMA)
	case '.':
		lexer.addToken(token.DOT)
	case '-':
		lexer.addToken(token.MINUS)
	case '+':
		lexer.addToken(token.PLUS)
	case ';':
		lexer.addToken(token.SEMICOLON)
	case '*':
		lexer.addToken(token.STAR)
	default:

	}
}

func (lexer *Lexer) advance() byte {
	index := lexer.current
	lexer.current += 1
	return lexer.source[index]
}

func (lexer *Lexer) addToken(tokenType token.TokenType) {
	lexer.addTokenToList(tokenType, "")
}

func (lexer *Lexer) addTokenToList(tokenType token.TokenType, literal string) {
	text := lexer.source[lexer.start:lexer.current]
	lexer.Tokens = append(lexer.Tokens, token.Token{
		Token:   tokenType,
		Lexeme:  text,
		Literal: literal,
		Line:    lexer.line,
	})
}
