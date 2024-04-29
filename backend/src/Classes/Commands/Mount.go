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

type Mount struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMount(line, column int, params map[string]string) *Mount {
	return &Mount{Type: utils.MOUNT, Line: line, Column: column, Params: params}
}

func (m *Mount) GetLine() int {
	return m.Line
}

func (m *Mount) GetColumn() int {
	return m.Column
}

func (m *Mount) GetType() utils.Type {
	return m.Type
}

func (m *Mount) Exec() {
	if m.validateParams() {
		m.mount()
		return
	} else if m.validateEmptyParams() {
		m.viewMounteds()
		return
	} else {
		m.printError("Error mount: Faltan parámetros obligatorios para montar la partición")
		m.Result += "Error mount: Faltan parámetros obligatorios para montar la partición"
	}
}

func (m *Mount) mount() {
	m.Params["driveletter"] = strings.ReplaceAll(m.Params["driveletter"], `"`, "")
	absolutePath, _ := filepath.Abs(env.GetPath(m.Params["driveletter"]))
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		m.printError(fmt.Sprintf("Error mount: No existe el disco %s para montar la partición", m.Params["driveletter"]))
		m.Result += fmt.Sprintf("Error mount: No existe el disco %s para montar la partición", m.Params["driveletter"])
		return
	}
	file, err := os.OpenFile(absolutePath, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("-> Error al abrir el archivo:", err)
		return
	}
	defer file.Close()
	readedBytes := make([]byte, 153)
	_, err = file.Read(readedBytes)
	if err != nil {
		fmt.Println("-> Error al leer el archivo:", err)
		return
	}
	mbr := structs.DecodeMBR(readedBytes)
	for i := 0; i < len(mbr.Partitions); i++ {
		if mbr.Partitions[i].Size != 0 && m.compareStrings(mbr.Partitions[i].Name, m.Params["name"]) {
			if mbr.Partitions[i].Status == '0' {
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
				_, err = file.Write([]byte("1"))
				if err != nil {
					fmt.Println("Error al escribir en el archivo:", err)
					return
				}
				thisDisk := env.Disks[m.Params["driveletter"]]
				newID := fmt.Sprintf("%s%d30", m.Params["driveletter"], thisDisk.NextId)
				thisDisk.Ids[newID] = &env.PartData{
					Name:   m.Params["name"],
					Mkdirs: []string{},
				}
				thisDisk.NextId++
				m.printSuccess(m.Params["driveletter"], m.Params["name"], newID, string(mbr.Partitions[i].Type))
				return
			}
			m.printError(fmt.Sprintf("Error mount: Intenta montar la partición en el disco \"%s\" que ya estpa montada", strings.Split(filepath.Base(absolutePath), ".")[0]))
			return
		}
	}
	i := m.getExtended(mbr.Partitions)
	if i != -1 {
		listEBR := m.getListEBR(mbr.Partitions[i].Start, mbr.Partitions[i].Size, file).GetIterable()
		for _, ebr := range listEBR {
			if ebr.Name != "" && strings.TrimSpace(ebr.Name) == m.Params["name"] {
				if ebr.Mount == '0' {
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
					_, err = file.Write([]byte("1"))
					if err != nil {
						fmt.Println("Error al escribir en el archivo:", err)
						return
					}
					thisDisk := env.Disks[m.Params["driveletter"]]
					newID := fmt.Sprintf("%d%s30", thisDisk.NextId, m.Params["driveletter"])
					thisDisk.Ids[newID] = &env.PartData{
						Name:   m.Params["name"],
						Mkdirs: []string{},
					}
					thisDisk.NextId++
					m.printSuccess(m.Params["driveletter"], m.Params["name"], newID, "L")
					return
				}
				m.printError(fmt.Sprintf("Error mount: Intenta montar la partición en el disco \"%s\" que ya estpa montada", strings.Split(filepath.Base(absolutePath), ".")[0]))
				return
			}
		}
	}
	m.printError(fmt.Sprintf("Error mount: Intenta montar una partición inexistente en el disco \"%s\"", strings.Split(filepath.Base(absolutePath), ".")[0]))
}

func (m *Mount) compareStrings(s1, s2 string) bool {
	return strings.TrimSpace(strings.ReplaceAll(s1, "\x00", "")) == strings.TrimSpace(strings.ReplaceAll(s2, "\x00", ""))
}

func (m *Mount) getListEBR(start, size int, file *os.File) *structs.ListEBR {
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

func (m *Mount) getExtended(partitions []*structs.Partition) int {
	for i := 0; i < len(partitions); i++ {
		if partitions[i].Type == 'E' {
			return i
		}
	}
	return -1
}

func (m *Mount) validateParams() bool {
	_, pathExists := m.Params["driveletter"]
	_, nameExists := m.Params["name"]
	return pathExists && nameExists
}

func (m *Mount) validateEmptyParams() bool {
	return len(m.Params) == 0
}

func (m *Mount) viewMounteds() {
	if len(env.Disks) > 0 {
		fmt.Println("\033[33m -> mount: \033[0m")
		m.Result += "mount: "
		fmt.Println("\033[33m\t -> Particiones Montadas\033[0m")
		m.Result += "Particiones Montadas"
		for diskName, diskInfo := range env.Disks {
			for id, value := range diskInfo.Ids {
				fmt.Printf("\033[33m\t -> %-20s %-20s %-20s\033[0m\n", id, value.Name, diskName)
				m.Result += fmt.Sprintf("%-20s %-20s %-20s\n", id, value.Name, diskName)
			}
		}
		println()
	} else {
		fmt.Println("\033[33m -> mount: No hay particiones montadas: \033[0m")
		m.Result += "mount: No hay particiones montadas: "
	}
}

func (m *Mount) printError(text string) {
	fmt.Printf("\033[31m-> %s. [%v:%v]\033[0m\n", text, m.Line, m.Column+1)
	m.Result += fmt.Sprintf("%s.\n", text)
}

func (m *Mount) printSuccess(diskName, name, newID, type_ string) {
	if type_ == "P" {
		type_ = "PRIMARIA "
	} else if type_ == "E" {
		type_ = "EXTENDIDA"
	} else {
		type_ = "LOGICA "
	}
	fmt.Printf("\033[32m-> mount: Partición montada exitosamente en %s. %s (%s: %s) [%d:%d]\033[0m\n", diskName, type_, name, newID, m.Line, m.Column)
	m.Result += fmt.Sprintf("mount: Partición montada exitosamente en %s. %s (%s: %s)\n", diskName, type_, name, newID)
}

func (m *Mount) GetResult() string { return m.Result }
