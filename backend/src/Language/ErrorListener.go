package language

import (
	"github.com/antlr4-go/antlr/v4"
)

type CustomSyntaxError struct {
	Line   int
	Column int
	Msg    string
}

type MIAErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []*CustomSyntaxError
}

func NewMIAErrorListener() *MIAErrorListener {
	return new(MIAErrorListener)
}

func (c *MIAErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &CustomSyntaxError{
		Line:   line,
		Column: column + 1,
		Msg:    msg,
	})
}
