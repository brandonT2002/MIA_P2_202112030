package commands

import (
	"fmt"
	env "mia/Classes/Env"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
	"path/filepath"
	"strings"
)

type Cat struct {
	Result string
	Type   utils.Type
	Files  []string
	Line   int
	Column int
}

func NewCat(line, column int, files []string) *Cat {
	return &Cat{Type: utils.CAT, Line: line, Column: column, Files: files}
}

func (c *Cat) GetLine() int {
	return c.Line
}

func (c *Cat) GetColumn() int {
	return c.Column
}

func (c *Cat) GetType() utils.Type {
	return c.Type
}

func (c *Cat) Exec() {
	if env.CurrentLogged.User != nil {
		if len(c.Files) > 0 {
			var absolutePath, _ = filepath.Abs(env.GetPath(*env.CurrentLogged.Driveletter))
			var file, err = os.Open(absolutePath)
			if err != nil {
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()
			var readedBytes = make([]byte, 153)
			_, err = file.Read(readedBytes)
			if err != nil {
				fmt.Println("-> Error al leer el archivo:", err)
				return
			}
			var mbr = structs.DecodeMBR(readedBytes)
			for _, p := range mbr.Partitions {
				if p.Size != 0 && utils.CompareStrings(p.Name, *env.CurrentLogged.Partition) {
					_, err = file.Seek(int64(p.Start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					var superBlockData = make([]byte, structs.SizeOfSuperBlock())
					_, err = file.Read(superBlockData)
					if err != nil {
						fmt.Println("Error al leer el super block:", err)
						return
					}
					var superBlock = structs.DecodeSuperBlock(superBlockData)
					var tree = structs.NewTree(*superBlock, *file)
					fmt.Println("\033[33m-> cat: \033[0m")
					for i, f := range c.Files {
						var content, _ = tree.ReadFile(f)
						fmt.Println(fmt.Sprintf("\033[53m\t-> (%v)\n\t%-20v\033[0m", i+1, strings.ReplaceAll(content, "\n", "\n\t")))
					}
					return
				}
			}
		} else {
			c.printError("-> Error cat: No se incluyeron archivos.")
		}
	} else {
		c.printError("-> Error cat: No hay ningún usuario loggeado actualmente.")
	}
}

func (m *Cat) printError(text string) {
	fmt.Printf("\033[31m%v [%v:%v]\033[0m\n", text, m.Line, m.Column)
}

func (c *Cat) GetResult() string { return "" }
