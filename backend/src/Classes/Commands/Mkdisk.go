package commands

import (
	"fmt"
	env "mia/Classes/Env"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
	"strconv"
	"strings"
)

type Mkdisk struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMkdisk(line, column int, params map[string]string) *Mkdisk {
	return &Mkdisk{Type: utils.MKDISK, Line: line, Column: column, Params: params}
}

func (m *Mkdisk) GetLine() int {
	return m.Line
}

func (m *Mkdisk) GetColumn() int {
	return m.Column
}

func (m *Mkdisk) GetType() utils.Type {
	return m.Type
}

func (m *Mkdisk) Exec() {
	if m.ValidateParams() {
		file, _ := os.Create(env.GetPath(fmt.Sprintf("%c", env.Asciiletter)))
		units := m.RecalculateUnits()
		size, _ := strconv.Atoi(m.Params["size"])
		for i := 0; i < size; i++ {
			byte := make([]byte, units)
			file.Write(byte)
		}
		file.Close()
		file, _ = os.OpenFile(env.GetPath(fmt.Sprintf("%c", env.Asciiletter)), os.O_RDWR, 0644)
		file.Seek(0, 0)
		mbr := structs.NewMBR(units*size, m.GetFit())
		env.Disks[fmt.Sprintf("%c", env.Asciiletter)] = &env.DiscoData{
			Ids:    map[string]*env.PartData{},
			NextId: 1,
		}
		file.Write(mbr.Encode())
		file.Close()
		fmt.Printf("\033[96m-> mkdisk: Disco %c creado exitosamente. (%s %sB) [%v:%v]\033[0m\n", env.Asciiletter, m.Params["size"], m.Params["unit"], m.Line, m.Column+1)
		env.Asciiletter++
	} else {
		fmt.Printf("\033[31m-> Error mkdisk: Faltan parámetros obligatorios. [%v:%v]\033[0m\n", m.Line, m.Column+1)
	}
}

func (m *Mkdisk) ValidateParams() bool {
	if _, exist := m.Params["size"]; exist {
		return true
	}
	return false
}

func (m *Mkdisk) RecalculateUnits() int {
	m.Params["unit"] = strings.ToUpper(m.Params["unit"])
	if _, exist := m.Params["unit"]; exist {
		if m.Params["unit"] == "K" {
			return 1024
		}
	}
	return 1024 * 1024
}

func (m *Mkdisk) GetFit() rune {
	m.Params["fit"] = strings.ToUpper(m.Params["fit"])
	if m.Params["fit"] == "FF" {
		return 'F'
	}
	if m.Params["fit"] == "BF" {
		return 'B'
	}
	return 'W'
}

func (m *Mkdisk) GetResult() string { return "" }
