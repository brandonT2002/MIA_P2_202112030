package commands

import (
	"fmt"
	"io/ioutil"
	env "mia/Classes/Env"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Mkfile struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMkfile(line, column int, params map[string]string) *Mkfile {
	return &Mkfile{Type: utils.MKDIR, Line: line, Column: column, Params: params}
}

func (m *Mkfile) GetLine() int {
	return m.Line
}

func (m *Mkfile) GetColumn() int {
	return m.Column
}

func (m *Mkfile) GetType() utils.Type {
	return m.Type
}

func (m *Mkfile) Exec() {
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
					var sizeStr, sizeExist = m.Params["size"]
					var contStr, contExist = m.Params["cont"]
					var size = 0
					var content = ""
					if sizeExist {
						size, _ = strconv.Atoi(sizeStr)
						if size > 0 {
							for len(content) < size {
								content += "0123456789"
							}
							content = content[:size]
						} else {
							m.printError(fmt.Sprintf("-> Error mkfile: No se creó el archivo '%v', el atributo size debe ser un número positivo.", m.Params["path"]))
							return
						}
					} else if contExist {
						contStr = strings.ReplaceAll(contStr, "\"", "")
						var absolutePath1, _ = filepath.Abs(contStr)
						var fileData, err = ioutil.ReadFile(absolutePath1)
						if err != nil {
							fmt.Println("Error al abrir el archivo:", err)
							return
						}
						content = string(fileData)
					}
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
						tree.Mkfile(m.Params["path"], env.GetPath(*env.CurrentLogged.Driveletter))
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
						tree.Mkfile(m.Params["path"], env.GetPath(*env.CurrentLogged.Driveletter))
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
								_, err = file1.Write(structs.NewJournal("mkfile", m.Params["path"], "", time.Now()).Encode())
								if err != nil {
									fmt.Println("Error al escribir Journal:", err)
									return
								}
								break
							}
						}
					}
					if content != "" {
						tree.WriteFile(m.Params["path"], env.GetPath(*env.CurrentLogged.Driveletter), p.Start, content)
					}
					tree.WriteInDisk(env.GetPath(*env.CurrentLogged.Driveletter), int32(p.Start), superBlock.Encode())
					disk.Ids[*env.CurrentLogged.IDPart].Mkdirs = append(disk.Ids[*env.CurrentLogged.IDPart].Mkdirs, m.Params["path"])
					m.printSuccess(fmt.Sprintf("-> mkfile: Nueva carpeta creada exitosamente '%v'", m.Params["path"]))
				}
			}
		} else {
			m.printError("-> Error mkfile: Faltan parámetros obligatorios para crear un directorio.")
		}
	} else {
		m.printError("-> Error mkfile: No hay ningún usuario loggeado actualmente.")
	}
}

func (m *Mkfile) isIn(value string, arr []string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func (m *Mkfile) validateParams() bool {
	_, pathExist := m.Params["path"]
	if pathExist {
		m.Params["path"] = strings.ReplaceAll(m.Params["path"], "\"", "")
	}
	return pathExist
}

func (m *Mkfile) printError(text string) {
	fmt.Printf("\033[31m-> %v [%v:%v]\033[0m\n", text, m.Line, m.Column)
	m.Result += fmt.Sprintf("%v.\n", text)
}

func (m *Mkfile) printSuccess(text string) {
	fmt.Printf("\033[32m-> %v [%v:%v]\033[0m\n", text, m.Line, m.Column)
	m.Result += fmt.Sprintf("%v.\n", text)
}

func (m *Mkfile) GetResult() string { return m.Result }
