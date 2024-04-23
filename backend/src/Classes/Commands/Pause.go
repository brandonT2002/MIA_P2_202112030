package commands

import (
	"fmt"
	utils "mia/Classes/Utils"
)

type Pause struct {
	Result string
	Type   utils.Type
	Line   int
	Column int
}

func NewPause(line, column int) *Pause {
	return &Pause{Type: utils.PAUSE, Line: line, Column: column}
}

func (p *Pause) GetLine() int {
	return p.Line
}

func (p *Pause) GetColumn() int {
	return p.Column
}

func (p *Pause) GetType() utils.Type {
	return p.Type
}

func (p *Pause) Exec() {
	var input string
	fmt.Printf("\033[33m-> pause \033[0m")
	fmt.Scanln(&input)
}

func (p *Pause) GetResult() string { return "" }
