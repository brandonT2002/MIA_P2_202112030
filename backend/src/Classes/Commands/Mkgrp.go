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

type Mkgrp struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMkgrp(line, column int, params map[string]string) *Mkgrp {
	return &Mkgrp{Type: utils.MKGRP, Line: line, Column: column, Params: params}
}

func (m *Mkgrp) GetLine() int {
	return m.Line
}

func (m *Mkgrp) GetColumn() int {
	return m.Column
}

func (m *Mkgrp) GetType() utils.Type {
	return m.Type
}

func (m *Mkgrp) Exec() {
	if env.CurrentLogged.User != nil {
		if env.CurrentLogged.User.Name == "root" {
			if m.validateParams() {
				if len(m.Params["name"]) <= 10 {
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
								var groups = tree.GetGroups(content)
								var idInt, _ = strconv.Atoi(groups[len(groups)-1].Id)
								var newGroup = fmt.Sprintf("%v,G,%-10v", idInt+1, m.Params["name"])
								tree.WriteFile("/users.txt", env.GetPath(*env.CurrentLogged.Driveletter), p.Start, newGroup)
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
											_, err = file1.Write(structs.NewJournal("mkgrp", "", fmt.Sprintf("%v", m.Params["name"]), time.Now()).Encode())
											if err != nil {
												fmt.Println("Error al escribir Journal:", err)
												return
											}
											break
										}
									}
								}
								m.printSuccess(fmt.Sprintf("-> mkgrp: Grupo %-10v creado exitosamente. (%v: %v)", m.Params["name"], env.CurrentLogged.Partition, env.CurrentLogged.Driveletter))
							} else {
								m.printError("-> Error mkgrp: No existe el archivo /users.txt.")
							}
							return
						}
					}
				} else {
					m.printError("-> Error mkgrp: El nombre de un grupo no puede contener más de 10 caracteres.")
				}
			} else {
				m.printError("-> Error mkgrp: Faltan parámetros obligatorios para crear un grupo.")
			}
		} else {
			m.printError("-> Error mkgrp: Solo el usuario 'root' puede crear grupos.")
		}
	} else {
		m.printError("-> Error mkgrp: No hay ningún usuario loggeado actualmente.")
	}
}

func (m *Mkgrp) validateParams() bool {
	_, nameExist := m.Params["name"]
	return nameExist
}

func (m *Mkgrp) printError(text string) {
	fmt.Printf("\033[31m%v [%v:%v]\033[0m\n", text, m.Line, m.Column)
}

func (m *Mkgrp) printSuccess(text string) {
	fmt.Printf("\033[32m%v [%v:%v]\033[0m\n", text, m.Line, m.Column)
}

func (m *Mkgrp) GetResult() string { return "" }
