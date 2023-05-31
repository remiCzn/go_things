package listener

import (
	"go_antlr/ast"
	"go_antlr/parser"
)

func (l *CalcListener) ExitBOp(c *parser.BOpContext) {

	right, left := l.Pop(), l.Pop()
	var op ast.Op
	switch c.GetOp().GetText() {
	case "*":
		op = ast.TIMES
	case "+":
		op = ast.PLUS
	case "-":
		op = ast.MINUS
	case "/":
		op = ast.DIVIDES
	}
	l.Push(ast.Bop{T1: left, Op: op, T2: right})
}
