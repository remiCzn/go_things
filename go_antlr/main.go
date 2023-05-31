package main

import (
	"fmt"
	"go_antlr/api"
	"go_antlr/listener"
	"go_antlr/parser"
	"io/ioutil"
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func parse(input string) listener.CalcListener {
	is := antlr.NewInputStream(input)
	lexer := parser.NewGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewGrammarParser(stream)

	var listener listener.CalcListener = *listener.New()
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
	return listener
}

func main() {
	args := os.Args
	if len(args) > 1 {
		filename := args[1]
		buf, e := ioutil.ReadFile(filename)
		if e != nil {
			log.Fatal("Unable to read file:", e)
		}
		text := string(buf)
		l := parse(text)
		fmt.Println(l.Execute())
	} else {
		log.Println("No File Input")
		api.Run(":3000")
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Println("Basic Interpreter")
		// fmt.Println("---------------------")
		// for {
		// 	fmt.Print("-> ")
		// 	text, _ := reader.ReadString('\n')
		// 	// convert CRLF to LF
		// 	text = strings.Replace(text, "\n", "", -1)
		// 	term := parse(text)
		// 	val, _ := term.Value(env)
		// 	fmt.Println(val)
		// }
	}

}
