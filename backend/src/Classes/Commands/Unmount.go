package commands

import (
	"fmt"
	"io"
	"io/ioutil"
	env "mia/Classes/Env"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
	"path/filepath"
	"strings"
)

type Unmount struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewUnmount(line, column int, params map[string]string) *Unmount {
	return &Unmount{Type: utils.UNMOUNT, Line: line, Column: column, Params: params}
}

func (u *Unmount) GetLine() int {
	return u.Line
}

func (u *Unmount) GetColumn() int {
	return u.Column
}

func (u *Unmount) GetType() utils.Type {
	return u.Type
}

func (u *Unmount) Exec() {
	if u.validateParams() {
		u.unmount()
	} else {
		u.printError("Error unmount: Faltan parámetros obligatorios para desmontar la partición")
	}
}

func (u *Unmount) unmount() {
	driveletter := string(u.Params["id"][0])
	if disk, ok := env.Disks[driveletter]; ok {
		absolutePath, _ := filepath.Abs(env.GetPath(driveletter))
		if disk.Ids[u.Params["id"]] != nil {
			namePartition := disk.Ids[u.Params["id"]].Name
			file, err := os.Open(absolutePath)
			if err != nil {
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()
			readedBytes, err := ioutil.ReadAll(io.LimitReader(file, 153))
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return
			}
			mbr := structs.DecodeMBR(readedBytes)
			for i := 0; i < len(mbr.Partitions); i++ {
				if mbr.Partitions[i].Size != 0 && u.compareStrings(mbr.Partitions[i].Name, namePartition) {
					file, err := os.OpenFile(absolutePath, os.O_RDWR, 0644)
					if err != nil {
						fmt.Println("Error al abrir el archivo:", err)
						return
					}
					defer file.Close()
					_, err = file.Seek(int64(13+i*35), 0)
					if err != nil {
						fmt.Println("Error al posicionar el puntero:", err)
						return
					}
					_, err = file.Write([]byte("0"))
					if err != nil {
						fmt.Println("Error al escribir en el archivo:", err)
						return
					}
					delete(disk.Ids, u.Params["id"])
					u.printSuccess(filepath.Base(absolutePath)[:len(filepath.Base(absolutePath))-len(filepath.Ext(absolutePath))], u.Params["id"], namePartition, string(mbr.Partitions[i].Type))
					return
				}
			}
			i := u.getExtended(mbr.Partitions)
			if i != -1 {
				listEBR := u.getListEBR(mbr.Partitions[i].Start, mbr.Partitions[i].Size, file).GetIterable()
				for _, ebr := range listEBR {
					if ebr.Name != "" && strings.TrimSpace(ebr.Name) == namePartition {
						file, err := os.OpenFile(absolutePath, os.O_RDWR, 0644)
						if err != nil {
							fmt.Println("Error al abrir el archivo:", err)
							return
						}
						defer file.Close()
						_, err = file.Seek(int64(ebr.Start), 0)
						if err != nil {
							fmt.Println("Error al buscar en el archivo:", err)
							return
						}
						_, err = file.Write([]byte("0"))
						if err != nil {
							fmt.Println("Error al escribir en el archivo:", err)
							return
						}
						delete(disk.Ids, u.Params["id"])
						u.printSuccess(filepath.Base(absolutePath)[:len(filepath.Base(absolutePath))-len(filepath.Ext(absolutePath))], u.Params["id"], namePartition, string(mbr.Partitions[i].Type))
						return
					}
				}
			}
			u.printError(fmt.Sprintf("Error unmount: Intenta desmontar una partición desmontada o inexistente en el disco \"%s\"", driveletter))
			return
		}
		u.printError(fmt.Sprintf("Error unmount: No existe el código de partición \"%s\" para desmontar en el disco \"%s\"", u.Params["id"], driveletter))
		return
	}
	u.printError(fmt.Sprintf("Error unmount: No existe el disco \"%s\" para desmontar la partición", driveletter))
}

func (u *Unmount) compareStrings(s1, s2 string) bool {
	return strings.TrimSpace(strings.ReplaceAll(s1, "\x00", "")) == strings.TrimSpace(strings.ReplaceAll(s2, "\x00", ""))
}

func (u *Unmount) getListEBR(start, size int, file *os.File) *structs.ListEBR {
	listEBR := structs.NewListEBR(start, size)
	_, err := file.Seek(int64(start), 0)
	if err != nil {
		return nil
	}
	ebr := structs.DecodeEBR(make([]byte, 30))
	listEBR.Insert(ebr)
	for ebr.Next != -1 {
		file.Seek(int64(ebr.Next), 0)
		ebr = structs.DecodeEBR(make([]byte, 30))
		listEBR.Insert(ebr)
	}
	return listEBR
}

func (u *Unmount) getExtended(partitios []*structs.Partition) int {
	for i := 0; i < len(partitios); i++ {
		if partitios[i].Type == 'E' {
			return i
		}
	}
	return -1
}

func (u *Unmount) validateParams() bool {
	_, idExists := u.Params["id"]
	return idExists
}

func (u *Unmount) printError(text string) {
	fmt.Printf("\033[31m-> %s. [%v:%v]\033[0m\n", text, u.Line, u.Column+1)
}

func (u *Unmount) printSuccess(diskName, codePart, name, type_ string) {
	if type_ == "P" {
		type_ = "PRIMARIA "
	} else if type_ == "E" {
		type_ = "EXTENDIDA"
	} else {
		type_ = "LOGICA "
	}
	fmt.Printf("\033[32m-> unmount: Partición desmontada exitosamente en %s. %s (%s: %s) [%d:%d]\033[0m\n", diskName, type_, codePart, name, u.Line, u.Column)
}

func (u *Unmount) GetResult() string { return "" }
