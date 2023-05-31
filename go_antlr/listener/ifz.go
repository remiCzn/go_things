package listener

import (
	"go_antlr/ast"
	"go_antlr/parser"
)

func (l *CalcListener) ExitIfZ(c *parser.IfZContext) {

	elseE, thenE, condE := l.Pop(), l.Pop(), l.Pop()
	l.Push(ast.Ifz{CondTerm: condE, ThenTerm: thenE, ElseTerm: elseE})
}
