package ast

import "horse-lang/token"

type Node interface {
	TokenLiteral() string // return the literal value of the node
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program is the root node
// It is a collection of statement nodes
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) <= 0 {
		return ""
	}
	return p.Statements[0].TokenLiteral() // huh?
}

// Let statement
type LetStatement struct {
	Token token.Token // let token
	Name  *Identifier // identifer of the name of the binding val (eg. in let a = 5, this will hold "a")
	Value Expression  // value of the evaluated expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
