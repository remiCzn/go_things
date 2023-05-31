package listener

import (
	"go_antlr/ast"
	"go_antlr/parser"
)

func (l *CalcListener) ExitLet(c *parser.LetContext) {
	var name string = c.GetVar_().GetText()
	assign := l.Pop()
	l.Push(ast.Let{VarName: name, AssignTerm: assign})
}
