package commands

import (
	"fmt"
	env "mia/Classes/Env"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Mkusr struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMkusr(line, column int, params map[string]string) *Mkusr {
	return &Mkusr{Type: utils.MKUSR, Line: line, Column: column, Params: params}
}

func (m *Mkusr) GetLine() int {
	return m.Line
}

func (m *Mkusr) GetColumn() int {
	return m.Column
}

func (m *Mkusr) GetType() utils.Type {
	return m.Type
}

func (m *Mkusr) Exec() {
	if env.CurrentLogged.User != nil {
		if env.CurrentLogged.User.Name == "root" {
			if m.validateParams() {
				if len(m.Params["user"]) <= 10 {
					if len(m.Params["pass"]) <= 10 {
						if len(m.Params["grp"]) <= 10 {
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
									var content, exists = tree.ReadFile("/users.txt")
									if exists {
										var users = tree.GetUsers(content)
										var idInt, _ = strconv.Atoi(users[len(users)-1].Id)
										var newUser = fmt.Sprintf("%v,U,%-10v,%-10v,%-10v", idInt+1, m.Params["grp"], m.Params["user"], m.Params["pass"])
										tree.WriteFile("/users.txt", env.GetPath(*env.CurrentLogged.Driveletter), p.Start, newUser)
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
													_, err = file1.Write(structs.NewJournal("mkusr", "", fmt.Sprintf("%v,%v,%v", m.Params["user"], m.Params["pass"], m.Params["grp"]), time.Now()).Encode())
													if err != nil {
														fmt.Println("Error al escribir Journal:", err)
														return
													}
													break
												}
											}
										}
										m.printSuccess(fmt.Sprintf("-> mkusr: Usuario %-10v creado exitosamente. (%v: %v)", m.Params["user"], env.CurrentLogged.Partition, env.CurrentLogged.Driveletter))
									} else {
										m.printError("-> Error mkusr: No existe el archivo /users.txt.")
									}
									return
								}
							}
						} else {
							m.printError("-> Error mkusr: El nombre de grupo no puede contener más de 10 caracteres.")
						}
					} else {
						m.printError("-> Error mkusr: La contraseña no puede contener más de 10 caracteres.")
					}
				} else {
					m.printError("-> Error mkusr: El nombre de un usuario no puede contener más de 10 caracteres.")
				}
			} else {
				m.printError("-> Error mkusr: Faltan parámetros obligatorios para crear un usuario.")
			}
		} else {
			m.printError("-> Error mkusr: Solo el usuario 'root' puede crear usuarios.")
		}
	} else {
		m.printError("-> Error mkusr: No hay ningún usuario loggeado actualmente.")
	}
}

func (m *Mkusr) validateParams() bool {
	_, userExist := m.Params["user"]
	_, passExist := m.Params["pass"]
	_, grpExist := m.Params["grp"]
	return userExist && passExist && grpExist
}

func (m *Mkusr) printError(text string) {
	fmt.Printf("\033[31m%v [%v:%v]\033[0m\n", text, m.Line, m.Column)
}

func (m *Mkusr) printSuccess(text string) {
	fmt.Printf("\033[32m%v [%v:%v]\033[0m\n", text, m.Line, m.Column)
}

func (m *Mkusr) GetResult() string { return "" }
