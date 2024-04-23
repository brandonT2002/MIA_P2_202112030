package callparser

import (
	"fmt"
	commands "mia/Classes/Commands"
	utils "mia/Classes/Utils"
	listener "mia/Language"
	parser "mia/Language/Parser"

	"github.com/antlr4-go/antlr/v4"
)

type CallParser struct {
}

func (c *CallParser) ExecutionParser(input string) string {
	Action := ""
	inputStream := antlr.NewInputStream(input)
	scanner := parser.NewScanner(inputStream)
	tokens := antlr.NewCommonTokenStream(scanner, antlr.TokenDefaultChannel)
	parser := parser.NewParserParser(tokens)
	parser.RemoveErrorListeners()
	parserErrors := listener.NewMIAErrorListener()
	parser.AddErrorListener(parserErrors)

	parser.BuildParseTrees = true
	tree := parser.Init()
	var listener *listener.MIAListener = listener.NewMIAListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("\033[31m-> Ha ocurrido un error en la línea de comandos.\033[0m")
				Action += fmt.Sprintf("\033[31mHa ocurrido un error en la línea de comandos.\033[0m")
			}
		}()
	}()

	for _, fail := range parserErrors.Errors {
		fmt.Printf("\033[31m-> %s. [%v:%v]\033[0m\n", fail.Msg, fail.Line, fail.Column+1)
		// Action += fmt.Sprintf("\033[31m%s. [%v:%v]\033[0m\n", fail.Msg, fail.Line, fail.Column+1)
		Action += fmt.Sprintf("%s. [%v:%v]\n", "Comando no reconocido", fail.Line, fail.Column+1)
	}

	for _, cmd := range listener.Execute {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(cmd.GetType())
					fmt.Printf("\033[31m-> %s. [%v:%v]\033[0m\n", fmt.Sprintf("%v", r), cmd.GetLine(), cmd.GetColumn())
					Action += fmt.Sprintf("\033[31m%s. [%v:%v]\033[0m\n", fmt.Sprintf("%v", r), cmd.GetLine(), cmd.GetColumn())
				}
			}()
			if cmd.GetType() == utils.EXECUTE {
				cmd.(*commands.Execute).SetParser(c)
			}
			cmd.Exec()
			Action += cmd.GetResult()
		}()
	}
	return Action
}
