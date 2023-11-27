package token

const (
	// Single character tokens
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	COMMA       = ","
	DOT         = ","
	MINUS       = "-"
	PLUS        = "+"
	SEMICOLON   = ";"
	SLASH       = "/"
	STAR        = "*"

	// One or two character tokens
	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL         = "="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	// Literals
	IDENTIFIER = "IDENTIFIER"
	STRING     = "string"
	NUMBER     = "number"

	// Keywords
	AND    = "and"
	CLASS  = "class"
	ELSE   = "else"
	FALSE  = "false"
	FUN    = "fun"
	FOR    = "for"
	IF     = "if"
	NIL    = "nil"
	OR     = "or"
	PRINT  = "print"
	RETURN = "return"
	SUPER  = "super"
	THIS   = "this"
	TRUE   = "true"
	VAR    = "var"
	WHILE  = "while"

	EOF = "EOF"
)

type TokenType string

type Token struct {
	Token   TokenType
	Lexeme  string
	Literal string
	Line    int
}

func (token Token) ToString() string {
	return string(token.Token) + "" + token.Lexeme + "" + token.Literal
}
