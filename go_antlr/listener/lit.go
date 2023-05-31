package listener

import (
	"go_antlr/ast"
	"go_antlr/parser"
	"strconv"
)

func (l *CalcListener) ExitLit(c *parser.LitContext) {

	num, err := strconv.Atoi(c.GetText())

	if err != nil {
		panic("Unexpected number")
	}
	l.Push(ast.Lit{Val: num})
}
