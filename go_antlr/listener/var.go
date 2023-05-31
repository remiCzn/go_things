package listener

import (
	"go_antlr/ast"
	"go_antlr/parser"
)

func (l *CalcListener) ExitVar(c *parser.VarContext) {
	l.Push(ast.Var{Name: c.GetText()})
}
