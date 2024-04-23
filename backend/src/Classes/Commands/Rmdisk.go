package commands

import (
	"fmt"
	env "mia/Classes/Env"
	utils "mia/Classes/Utils"
	"os"
)

type Rmdisk struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewRmdisk(line, column int, params map[string]string) *Rmdisk {
	return &Rmdisk{Type: utils.RMDISK, Line: line, Column: column, Params: params}
}

func (r *Rmdisk) GetLine() int {
	return r.Line
}

func (r *Rmdisk) GetColumn() int {
	return r.Column
}

func (r *Rmdisk) GetType() utils.Type {
	return r.Type
}

func (r *Rmdisk) Exec() {
	if r.ValidateParams() {
		file := os.Remove(env.GetPath(r.Params["driveletter"]))
		if file != nil {
			fmt.Printf("\033[31m-> Error rmdisk: No se puede eliminar disco %s. [%v:%v]\033[0m\n", r.Params["driveletter"], r.Line, r.Column+1)
			return
		}
		fmt.Printf("\033[96m-> rmdisk: Disco %s eliminado exitosamente. [%v:%v]\033[0m\n", r.Params["driveletter"], r.Line, r.Column+1)
	} else {
		fmt.Printf("\033[31m-> Error rmdisk: No se encontrón el parámetro obligatorio \"driveletter\". [%v:%v]\033[0m\n", r.Line, r.Column+1)
	}
}

func (r *Rmdisk) ValidateParams() bool {
	if _, exist := r.Params["driveletter"]; exist {
		return true
	}
	return false
}

func (r *Rmdisk) GetResult() string { return "" }
