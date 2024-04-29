package commands

import (
	"fmt"
	env "mia/Classes/Env"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Mkdir struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMkdir(line, column int, params map[string]string) *Mkdir {
	return &Mkdir{Type: utils.MKDIR, Line: line, Column: column, Params: params}
}

func (m *Mkdir) GetLine() int {
	return m.Line
}

func (m *Mkdir) GetColumn() int {
	return m.Column
}

func (m *Mkdir) GetType() utils.Type {
	return m.Type
}

func (m *Mkdir) Exec() {
	if env.CurrentLogged.User != nil {
		if m.validateParams() {
			var absolutePath, _ = filepath.Abs(env.GetPath(*env.CurrentLogged.Driveletter))
			var file, err = os.OpenFile(absolutePath, os.O_WRONLY|os.O_CREATE, 0644)
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
					var disk = env.Disks[*env.CurrentLogged.Driveletter]
					if m.isIn(m.Params["path"], disk.Ids[*env.CurrentLogged.IDPart].Mkdirs) {
						m.printError(fmt.Sprintf("-> Error mkdir: No se creó la carpeta '%v' porque ya existe.", m.Params["path"]))
						return
					}
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
					var path = strings.Split(m.Params["path"], "/")
					var dir = []string{}
					for _, d := range path {
						if d != "" {
							dir = append(dir, d)
						}
					}
					if m.Params["r"] == "1" {
						var c = 0
						for c < len(dir) {
							var tmpDir = []string{}
							for i := 0; i < c+1; i++ {
								tmpDir = append(tmpDir, dir[i])
							}
							if !tree.SearchDir("/" + strings.Join(tmpDir, "/")) {
								tree.Mkdir("/"+strings.Join(tmpDir, "/"), env.GetPath(*env.CurrentLogged.Driveletter))
							}
							c++
						}
					} else {
						if len(dir) > 1 {
							var tmpDir = []string{}
							for i := 0; i < len(dir)-1; i++ {
								tmpDir = append(tmpDir, dir[i])
							}
							if !tree.SearchDir("/"+strings.Join(tmpDir, "/")) && len(dir) > 1 {
								m.printError(fmt.Sprintf("-> Error mkdir: No se creó la carpeta '%v', no existe la ruta donde intentó crearse.", m.Params["path"]))
							}
						}
						tree.Mkdir(m.Params["path"], env.GetPath(*env.CurrentLogged.Driveletter))
					}
					if superBlock.Filesystem_type == 3 {
						_, err = file.Seek(int64(p.Start+structs.SizeOfSuperBlock()), 0)
						if err != nil {
							fmt.Println("Error al mover el puntero del archivo:", err)
							return
						}
						for r := 0; r < superBlock.Inodes_count; r++ {
							journalData := make([]byte, structs.SizeOfJournal())
							_, err = file.Read(journalData)
							if err != nil {
								fmt.Println("Error al leer el super block:", err)
								return
							}
							if string(journalData) == strings.Repeat("\x00", structs.SizeOfJournal()) {
								var file1, err = os.OpenFile(absolutePath, os.O_WRONLY|os.O_CREATE, 0644)
								_, err = file1.Seek(int64(p.Start+structs.SizeOfSuperBlock()+r*structs.SizeOfJournal()), 0)
								if err != nil {
									fmt.Println("Error al mover el puntero del archivo:", err)
									return
								}
								defer file1.Close()
								_, err = file1.Write(structs.NewJournal("mkdir", m.Params["path"], "", time.Now()).Encode())
								if err != nil {
									fmt.Println("Error al escribir Journal:", err)
									return
								}
								break
							}
						}
					}
					tree.WriteInDisk(env.GetPath(*env.CurrentLogged.Driveletter), int32(p.Start), superBlock.Encode())
					disk.Ids[*env.CurrentLogged.IDPart].Mkdirs = append(disk.Ids[*env.CurrentLogged.IDPart].Mkdirs, m.Params["path"])
					m.printSuccess(fmt.Sprintf("-> mkdir: Nueva carpeta creada exitosamente '%v'", m.Params["path"]))
				}
			}
		} else {
			m.printError("-> Error mkdir: Faltan parámetros obligatorios para crear un directorio.")
		}
	} else {
		m.printError("-> Error mkdir: No hay ningún usuario loggeado actualmente.")
	}
}

func (m *Mkdir) isIn(value string, arr []string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func (m *Mkdir) validateParams() bool {
	_, pathExist := m.Params["path"]
	if pathExist {
		m.Params["path"] = strings.ReplaceAll(m.Params["path"], "\"", "")
	}
	return pathExist
}

func (m *Mkdir) printError(text string) {
	fmt.Printf("\033[31m%v [%v:%v]\033[0m\n", text, m.Line, m.Column)
}

func (m *Mkdir) printSuccess(text string) {
	fmt.Printf("\033[32m%v [%v:%v]\033[0m\n", text, m.Line, m.Column)
	m.Result += fmt.Sprintf("%v.\n", text)
}

func (m *Mkdir) GetResult() string { return m.Result }
