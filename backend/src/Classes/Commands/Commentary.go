package commands

import (
	"fmt"
	utils "mia/Classes/Utils"
	"regexp"
)

type Commentary struct {
	Result string
	Type   utils.Type
	Line   int
	Column int
}

func NewCommentary(line, column int, text string) *Commentary {
	return &Commentary{Type: utils.COMMENTARY, Line: line, Column: column, Result: text}
}

func (c *Commentary) GetLine() int {
	return c.Line
}

func (c *Commentary) GetColumn() int {
	return c.Column
}

func (c *Commentary) GetType() utils.Type {
	return c.Type
}

func (c *Commentary) Exec() {
	re := regexp.MustCompile(`#\s*`)
	c.Result = re.ReplaceAllString(c.Result, "")
	fmt.Printf("-> %v. [%v:%v]\n", c.Result, c.Line, c.Column+1)
	c.Result = fmt.Sprintf("%v.\n", c.Result)
}

func (c *Commentary) GetResult() string { return c.Result }
