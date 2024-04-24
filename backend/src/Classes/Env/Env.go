package env

import (
	"fmt"
	structs "mia/Classes/Structs"
)

type DiscoData struct {
	Ids    map[string]*PartData
	NextId int
}

type PartData struct {
	Name   string
	Mkdirs []string
}

var Disks = map[string]*DiscoData{}
var Asciiletter = 65
var CurrentLogged = struct {
	User        *structs.User
	Partition   *string
	Driveletter *string
	IDPart      *string
}{
	User:        nil,
	Partition:   nil,
	Driveletter: nil,
	IDPart:      nil,
}

func GetPath(driveletter string) string {
	return fmt.Sprintf("/home/jefferson/Escritorio/SO/%s.dsk", driveletter)
}

/*

	disks = {"A": {"ids": {"A130": {"name": "Particion1"}}}}

*/
