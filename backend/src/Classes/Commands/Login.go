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

type Login struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewLogin(line, column int, params map[string]string) *Login {
	return &Login{Type: utils.LOGIN, Line: line, Column: column, Params: params}
}

func (l *Login) GetLine() int {
	return l.Line
}

func (l *Login) GetColumn() int {
	return l.Column
}

func (l *Login) GetType() utils.Type {
	return l.Type
}

func (l *Login) Exec() {
	if l.validateParams() {
		l.login()
	} else {
		l.printError("-> Error login: Faltan parámetros obligatorios para iniciar sesión.")
	}
}

func (l *Login) login() {
	if env.CurrentLogged.User == nil {
		var driveletter = string(l.Params["id"][0])
		if disk, ok := env.Disks[driveletter]; ok {
			if part, ok := disk.Ids[l.Params["id"]]; ok {
				var diskPath, _ = filepath.Abs(env.GetPath(driveletter))
				var absolutePath, _ = filepath.Abs(diskPath)
				var namePartition = part.Name
				var file, err = os.Open(absolutePath)
				if err != nil {
					fmt.Println("Error al abrir el archivo:", err)
					return
				}
				defer file.Close()
				var readedBytes = make([]byte, 153)
				_, err = file.Read(readedBytes)
				if err != nil {
					fmt.Println("Error al leer el archivo:", err)
					return
				}
				var mbr = structs.DecodeMBR(readedBytes)
				for _, p := range mbr.Partitions {
					if p.Size != 0 && utils.CompareStrings(p.Name, namePartition) {
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
						var content, founded = tree.ReadFile("/users.txt")
						if founded {
							var user = l.isValidUser(content, l.Params["user"], l.Params["pass"])
							if user != nil {
								env.CurrentLogged.User = structs.NewUser(user.Id, user.Group, user.Name, user.Password)
								env.CurrentLogged.Partition = &namePartition
								env.CurrentLogged.Driveletter = &driveletter
								var id = l.Params["id"]
								env.CurrentLogged.IDPart = &id
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
											_, err = file1.Write(structs.NewJournal("login", "", fmt.Sprintf("%v,%v,%v", user.Name, user.Password, l.Params["id"]), time.Now()).Encode())
											if err != nil {
												fmt.Println("Error al escribir Journal:", err)
												return
											}
											break
										}
									}
								}
								l.printSuccess(fmt.Sprintf("-> login: Sesión iniciada exitosamente. (%v)", user.Name))
							} else {
								l.printError(fmt.Sprintf("-> Error login: El usuario %v no existe en el sistema.", l.Params["user"]))
							}
							return
						}
						l.printError("-> Error login: No existe el archivo /users.txt.")
						return
					}
				}
			} else {
				l.printError(fmt.Sprintf("-> Error login: No existe el código de partición %v en el disco %v para iniciar sesión.", l.Params["id"], driveletter))
			}
		} else {
			l.printError(fmt.Sprintf("-> Error login: No existe el disco %v.", driveletter))
		}
	} else {
		l.printError("-> Error login: Hay un usuario loggeado actualmente.")
	}
}

func (l *Login) isValidUser(content, user, password string) *structs.User {
	var users = l.getUsers(content)
	for _, u := range users {
		if u.Name == user && u.Password == password {
			return u
		}
	}
	return nil
}

func (l *Login) getUsers(content string) []*structs.User {
	var users = []*structs.User{}
	var registers = [][]string{}
	for _, line := range strings.Split(content, "\n") {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			var register = []string{}
			for _, field := range strings.Split(trimmedLine, ",") {
				register = append(register, strings.TrimSpace(field))
			}
			registers = append(registers, register)
		}
	}
	for _, reg := range registers {
		if reg[1] == "U" && reg[0] != "0" {
			users = append(users, structs.NewUser(reg[0], reg[2], reg[3], reg[4]))
		}
	}
	return users
}

func (l *Login) validateParams() bool {
	_, userExist := l.Params["user"]
	_, passExist := l.Params["pass"]
	_, idExist := l.Params["id"]
	return userExist && passExist && idExist
}

func (l *Login) printError(text string) {
	fmt.Printf("\033[31m%s [%d:%d]\033[0m\n", text, l.Line, l.Column)
}

func (l *Login) printSuccess(text string) {
	fmt.Printf("\033[32m%s [%d:%d]\033[0m\n", text, l.Line, l.Column)
}

func (l *Login) GetResult() string { return "" }
