package lexer

import (
	"horse-lang/src/error"
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
	case '!':
		if lexer.match('=') {
			lexer.addToken(token.BANG_EQUAL)
		} else {
			lexer.addToken(token.BANG)
		}
	case '=':
		if lexer.match('=') {
			lexer.addToken(token.EQUAL_EQUAL)
		} else {
			lexer.addToken(token.EQUAL)
		}
	case '<':
		if lexer.match('=') {
			lexer.addToken(token.LESS_EQUAL)
		} else {
			lexer.addToken(token.LESS)
		}
	case '>':
		if lexer.match('=') {
			lexer.addToken(token.GREATER_EQUAL)
		} else {
			lexer.addToken(token.GREATER)
		}
	case '/':
		if lexer.match('/') {
			for lexer.peek() != '\n' && !lexer.isAtEnd() {
				lexer.advance()
			}
		} else {
			lexer.addToken(token.SLASH)
		}
	case '\n':
		lexer.line += 1
	case '"':
		lexer.string()
	case ' ':
	case '\r':
	case '\t':
	default:
		if isDigit(c) {
			lexer.number()
		} else if isAlpha(c) {
			lexer.identifier()
		} else {
			error.ReportError(lexer.line, "Unexpected character")
		}
	}
}

func (lexer *Lexer) match(expected byte) bool {
	if lexer.isAtEnd() {
		return false
	}
	if lexer.source[lexer.current] != expected {
		return false
	}
	lexer.current += 1
	return true
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

func (lexer *Lexer) peek() byte {
	if lexer.isAtEnd() {
		return byte(0)
	}
	return lexer.source[lexer.current]
}

func (lexer *Lexer) string() {
	for lexer.peek() != '"' && !lexer.isAtEnd() {
		if lexer.peek() == '\n' {
			lexer.line += 1
		}
		lexer.advance()
	}

	if lexer.isAtEnd() {
		error.ReportError(lexer.line, "Unterminated string.")
	}

	lexer.advance()

	value := lexer.source[lexer.start+1 : lexer.current-1]
	lexer.addTokenToList(token.STRING, value)

}

func (lexer *Lexer) number() {
	for isDigit(lexer.peek()) {
		lexer.advance()
	}

	if lexer.peek() == '.' && isDigit(lexer.peekNext()) {
		lexer.advance()
	}

	for isDigit(lexer.peek()) {
		lexer.advance()
	}

	lexer.addTokenToList(
		token.NUMBER,
		lexer.source[lexer.start:lexer.current],
	)
}

func (lexer *Lexer) peekNext() byte {
	if lexer.current+1 >= len(lexer.source) {
		return byte(0)
	}
	return lexer.source[lexer.current+1]
}

func (lexer *Lexer) identifier() {
	for isAlphaNumberic(lexer.peek()) {
		lexer.advance()
	}
	text := lexer.source[lexer.start:lexer.current]
	if text == token.AND {
		lexer.addToken(token.IDENTIFIER)
	} else if text == token.CLASS {
		lexer.addToken(token.CLASS)
	} else if text == token.ELSE {
		lexer.addToken(token.ELSE)
	} else if text == token.FALSE {
		lexer.addToken(token.ELSE)
	} else if text == token.FOR {
		lexer.addToken(token.FOR)
	} else if text == token.FUN {
		lexer.addToken(token.FUN)
	} else if text == token.IF {
		lexer.addToken(token.FUN)
	} else if text == token.NIL {
		lexer.addToken(token.NIL)
	} else if text == token.OR {
		lexer.addToken(token.OR)
	} else if text == token.PRINT {
		lexer.addToken(token.PRINT)
	} else if text == token.RETURN {
		lexer.addToken(token.RETURN)
	} else if text == token.SUPER {
		lexer.addToken(token.SUPER)
	} else if text == token.THIS {
		lexer.addToken(token.THIS)
	} else if text == token.TRUE {
		lexer.addToken(token.TRUE)
	} else if text == token.VAR {
		lexer.addToken(token.VAR)
	} else if text == token.WHILE {
		lexer.addToken(token.VAR)
	} else {
		lexer.addToken(token.IDENTIFIER)
	}
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c == '-')
}

func isAlphaNumberic(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
