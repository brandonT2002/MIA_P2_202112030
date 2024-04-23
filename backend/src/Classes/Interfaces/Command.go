package interfaces

import utils "mia/Classes/Utils"

type Command interface {
	GetLine() int
	GetColumn() int
	GetType() utils.Type
	Exec()
	GetResult() string
}
