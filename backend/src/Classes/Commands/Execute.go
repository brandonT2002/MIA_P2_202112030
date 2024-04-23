package commands

import (
	"fmt"
	"io/ioutil"
	utils "mia/Classes/Utils"
)

type CallParser interface {
	ExecutionParser(input string)
}

type Execute struct {
	Params map[string]string
	Result string
	Type   utils.Type
	Line   int
	Column int
	Parser CallParser
}

func NewExecute(line, column int, params map[string]string) *Execute {
	return &Execute{Params: params, Type: utils.EXECUTE, Line: line, Column: column}
}

func (exe *Execute) GetLine() int {
	return exe.Line
}

func (exe *Execute) GetColumn() int {
	return exe.Column
}

func (exe *Execute) GetType() utils.Type {
	return exe.Type
}

func (exe *Execute) SetParser(parser interface{}) {
	exe.Parser = parser.(CallParser)
}

func (exe *Execute) Exec() {
	if _, exists := exe.Params["path"]; exists {
		file, err := ioutil.ReadFile(exe.Params["path"])
		if err != nil {
			fmt.Printf("\033[31m-> Error execute: No se pudo leer el script. [%v:%v]\033[0m\n", exe.Line, exe.Column+1)
			return
		}
		exe.Parser.ExecutionParser(string(file))
		return
	}
	fmt.Printf("\033[31m-> Error execute: No se encontrón el parámetro obligatorio \"path\". [%v:%v]\033[0m\n", exe.Line, exe.Column+1)
}

func (exe *Execute) GetResult() string { return "" }
