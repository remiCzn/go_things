package listener

import (
	"go_antlr/parser"
	"go_antlr/utils"
)

func (l *CalcListener) ExitStatement(c *parser.StatementContext) {
	utils.Log(l.stack)
}

func (l *CalcListener) ExitReturn(c *parser.ReturnContext) {
	utils.Log(l.stack)
}
