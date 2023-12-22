package ast

import "horse-lang/src/token"

type Expr interface {
	assign(visitorExpr ExprVisitor) interface{}
}

type ExprVisitor interface {
	VisitBinaryExpr(expr *Binary) interface{}
	VisitGroupingExpr(expr *Grouping) interface{}
	VisitLiteralExpr(expr *Literal) interface{}
	VisitUnaryExpr(expr *Unary) interface{}
}

type Binary struct {
	left Expr
	operator token.Token
	right Expr
}

func (expr *Binary) accept(visitor ExprVisitor) interface{} {
	return visitor.VisitBinaryExpr(expr)
}

type Grouping struct {
	expression Expr
}

func (expr *Grouping) accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(expr)
}

type Literal struct {
	value interface{}
}

func (expr *Literal) accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(expr)
}

type Unary struct {
	operator token.Token
	right Expr
}

func (expr *Unary) accept(visitor ExprVisitor) interface{} {
	return visitor.VisitUnaryExpr(expr)
}

