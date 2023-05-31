package listener

import (
	"go_antlr/ast"
	"go_antlr/parser"
	"go_antlr/utils"
)

type CalcListener struct {
	*parser.BaseGrammarListener

	stack []ast.Term
	env   map[string]int
}

func New() *CalcListener {
	return &CalcListener{
		stack: []ast.Term{},
		env:   make(map[string]int),
	}
}

func (l *CalcListener) Push(i ast.Term) {
	l.stack = append(l.stack, i)
}

func (l *CalcListener) Pop() ast.Term {
	if len(l.stack) < 1 {
		panic("stack is empty")
	}

	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

func (l *CalcListener) Execute() int {
	var value int
	for i := 0; i < len(l.stack); i++ {
		t := l.stack[i]
		utils.Log(t)
		value, l.env = t.Value(l.env)
	}
	return value
}
