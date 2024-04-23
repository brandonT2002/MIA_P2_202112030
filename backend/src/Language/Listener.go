package language

import (
	interfaces "mia/Classes/Interfaces"
	parser "mia/Language/Parser"
)

type MIAListener struct {
	*parser.BaseParserListener
	Execute []interfaces.Command
}

func NewMIAListener() *MIAListener {
	return new(MIAListener)
}

func (l *MIAListener) ExitInit(ctx *parser.InitContext) {
	l.Execute = ctx.GetResult()
}
