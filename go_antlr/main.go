package main

import (
	"bufio"
	"fmt"
	"go_antlr/parser"
	"os"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type calcListener struct {
	*parser.BaseGrammarListener

	stack []Term
}

func (l *calcListener) push(i Term) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() Term {
	if len(l.stack) < 1 {
		panic("stack is empty")
	}

	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

func (l *calcListener) ExitBOp(c *parser.BOpContext) {
	right, left := l.pop(), l.pop()
	var op Op
	switch c.GetOp().GetText() {
	case "*":
		op = TIMES
	case "+":
		op = PLUS
	case "-":
		op = MINUS
	case "/":
		op = DIVIDES
	}
	l.push(Bop{left, op, right})
}

func (l *calcListener) ExitLit(c *parser.LitContext) {
	num, err := strconv.Atoi(c.GetText())

	if err != nil {
		panic("Unexpected number")
	}
	l.push(Lit{num})
}

func (l *calcListener) ExitIfZ(c *parser.IfZContext) {
	elseE, thenE, condE := l.pop(), l.pop(), l.pop()
	l.push(Ifz{condE, thenE, elseE})
}

func (l *calcListener) ExitLet(c *parser.LetContext) {
	var name string = c.GetVar_().GetText()
	assign := l.pop()
	l.push(Let{name, assign})
}

func (l *calcListener) ExitVar(c *parser.VarContext) {
	l.push(Var{c.GetText()})
}

func calc(input string) Term {
	is := antlr.NewInputStream(input)
	lexer := parser.NewGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewGrammarParser(stream)

	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Statement())

	return listener.pop()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Basic Interpreter")
	fmt.Println("---------------------")

	var env map[string]int = make(map[string]int)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		term := calc(text)
		val, _ := term.value(env)
		fmt.Println(val)
	}
}
