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

type Mkfs struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMkfs(line, column int, params map[string]string) *Mkfs {
	return &Mkfs{Type: utils.MKFS, Line: line, Column: column, Params: params}
}

func (m *Mkfs) GetLine() int {
	return m.Line
}

func (m *Mkfs) GetColumn() int {
	return m.Column
}

func (m *Mkfs) GetType() utils.Type {
	return m.Type
}

func (m *Mkfs) Exec() {
	if m.validateParams() {
		m.mkfs()
	} else {
		m.printError("Error mkfs: Faltan parámetros obligatorios para formatear la partición")
		m.Result += "Error mkfs: Faltan parámetros obligatorios para formatear la partición"
	}
}

func (m *Mkfs) mkfs() {
	driveletter := string(m.Params["id"][0])
	if disk, ok := env.Disks[driveletter]; ok {
		absolutePath, _ := filepath.Abs(env.GetPath(driveletter))
		if partitionData, ok := disk.Ids[m.Params["id"]]; ok {
			namePartition := partitionData.Name
			file, err := os.Open(absolutePath)
			if err != nil {
				fmt.Println("Error al abrir el archivo:", err)
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
				if mbr.Partitions[i].Size != 0 && m.compareStrings(mbr.Partitions[i].Name, namePartition) {
					if m.Params["fs"] == "2fs" {
						m.ext2(absolutePath, mbr.Partitions[i])
					} else {
						m.ext3(absolutePath, mbr.Partitions[i])
					}
					m.printSuccess(driveletter, namePartition, m.Params["id"], string(mbr.Partitions[i].Type), m.Params["fs"])
					disk.Ids[m.Params["id"]].Mkdirs = append(disk.Ids[m.Params["id"]].Mkdirs, "/users.txt")
					return
				}
			}
		} else {
			m.printError(fmt.Sprintf("Error mkfs: No existe el código de partición \"%s\" para formatear en el disco \"%s\"", disk.Ids[m.Params["id"]], driveletter))
		}
	} else {
		m.printError(fmt.Sprintf("Error mkfs: No existe el disco \"%s\" para formatear la partición", driveletter))
	}
}

func (m *Mkfs) compareStrings(s1, s2 string) bool {
	return strings.TrimSpace(strings.ReplaceAll(s1, "\x00", "")) == strings.TrimSpace(strings.ReplaceAll(s2, "\x00", ""))
}

func (m *Mkfs) ext2(absolutePath string, partition *structs.Partition) {
	n := int((partition.Size - structs.SizeOfSuperBlock()) / (4 + structs.SizeOfInode() + 3*structs.SizeOfBlockFolder()))
	superBlock := structs.SuperBlock{}
	superBlock.Filesystem_type = 2
	superBlock.Inodes_count = n
	superBlock.Blocks_count = 3 * n
	superBlock.Free_inodes_count = n - 2
	superBlock.Free_blocks_count = 3*n - 2
	superBlock.Mtime = time.Time{}
	superBlock.Umtime = time.Time{}
	superBlock.Mnt_count = 0
	superBlock.Magic = 0xEF53
	superBlock.Inode_size = structs.SizeOfInode()
	superBlock.Block_size = structs.SizeOfBlockFolder()
	superBlock.First_ino = 2
	superBlock.First_blo = 2
	superBlock.Bm_inode_start = partition.Start + structs.SizeOfSuperBlock()
	superBlock.Bm_block_start = superBlock.Bm_inode_start + n
	superBlock.Inode_start = superBlock.Bm_block_start + 3*n
	superBlock.Block_start = superBlock.Inode_start + n*structs.SizeOfInode()

	inode0 := structs.NewInode('0', 0)
	inode0.Block[0] = 0

	blockFolder := structs.NewBlockFolder()
	blockFolder.Content[0] = &structs.Content{Name: m.padRight(".", 12), Inode: 0}
	blockFolder.Content[1] = &structs.Content{Name: m.padRight("..", 12), Inode: 0}
	blockFolder.Content[2] = &structs.Content{Name: m.padRight("users.txt", 12), Inode: 1}

	userstxt := "1,G,root      \n1,U,root      ,root      ,123       \n"
	inode1 := structs.NewInode('1', len(userstxt))
	inode1.Block[0] = 1

	blockFile := structs.BlockFile{Content: userstxt + strings.Repeat(" ", 64-len(userstxt))}

	file, err := os.OpenFile(absolutePath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()
	_, err = file.Seek(int64(partition.Start), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return
	}
	_, err = file.Write(superBlock.Encode())
	if err != nil {
		fmt.Println("Error al escribir el superBloque en el archivo:", err)
		return
	}
	_, err = file.Seek(int64(superBlock.Bm_inode_start), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return
	}
	bmInodes_bmBlocks := make([]byte, n+3*n)
	for i := 0; i < 2; i++ {
		bmInodes_bmBlocks[i] = '1'
	}
	for i := 2; i < n; i++ {
		bmInodes_bmBlocks[i] = '0'
	}
	for i := n; i < n+2; i++ {
		bmInodes_bmBlocks[i] = '1'
	}
	for i := n + 2; i < len(bmInodes_bmBlocks); i++ {
		bmInodes_bmBlocks[i] = '0'
	}
	_, err = file.Write(bmInodes_bmBlocks)
	if err != nil {
		fmt.Println("Error al escribir el mapa de bits de los inodos en el archivo:", err)
		return
	}
	_, err = file.Seek(int64(superBlock.Inode_start), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return
	}
	_, err = file.Write(inode0.Encode())
	if err != nil {
		fmt.Println("Error al escribir el primer inodo en el archivo", err)
		return
	}
	_, err = file.Write(inode1.Encode())
	if err != nil {
		fmt.Println("Error al escribir en el segundo inodo en el archivo:", err)
		return
	}
	_, err = file.Seek(int64(superBlock.Block_start), 0)
	if err != nil {
		fmt.Println("Error al escribir el bloqueFolder en el archivo", err)
		return
	}
	_, err = file.Write(blockFolder.Encode())
	if err != nil {
		fmt.Println("Error al escrbir el bloqueFolder en el archivo", err)
		return
	}
	_, err = file.Write(blockFile.Encode())
	if err != nil {
		fmt.Println("Error al escribir el blockFile en el archivo:", err)
		return
	}
}

func (m *Mkfs) ext3(absolutePath string, partition *structs.Partition) {
	n := int((partition.Size - structs.SizeOfSuperBlock()) / (4 + structs.SizeOfJournal() + structs.SizeOfInode() + 3*structs.SizeOfBlockFolder()))
	superBlock := structs.SuperBlock{}
	superBlock.Filesystem_type = 3
	superBlock.Inodes_count = n
	superBlock.Blocks_count = 3 * n
	superBlock.Free_inodes_count = n - 2
	superBlock.Free_blocks_count = 3*n - 2
	superBlock.Mtime = time.Time{}
	superBlock.Umtime = time.Time{}
	superBlock.Mnt_count = 0
	superBlock.Magic = 0xEF53
	superBlock.Inode_size = structs.SizeOfInode()
	superBlock.Block_size = structs.SizeOfBlockFolder()
	superBlock.First_ino = 2
	superBlock.First_blo = 2
	superBlock.Bm_inode_start = partition.Start + structs.SizeOfSuperBlock() + n*structs.SizeOfJournal()
	superBlock.Bm_block_start = superBlock.Bm_inode_start + n
	superBlock.Inode_start = superBlock.Bm_block_start + 3*n
	superBlock.Block_start = superBlock.Inode_start + n*structs.SizeOfInode()

	inode0 := structs.NewInode('0', 0)
	inode0.Block[0] = 0

	blockFolder := structs.NewBlockFolder()
	blockFolder.Content[0] = &structs.Content{Name: m.padRight(".", 12), Inode: 0}
	blockFolder.Content[1] = &structs.Content{Name: m.padRight("..", 12), Inode: 0}
	blockFolder.Content[2] = &structs.Content{Name: m.padRight("users.txt", 12), Inode: 1}

	userstxt := "1,G,root      \n1,U,root      ,root      ,123       \n"
	inode1 := structs.NewInode('1', len(userstxt))
	inode1.Block[0] = 1

	blockFile := structs.BlockFile{Content: strings.Repeat(" ", 64)}
	for c, char := range userstxt {
		if c >= len(blockFile.Content) {
			break
		}
		blockFile.Content = blockFile.Content[:c] + string(char) + blockFile.Content[c+1:]
	}

	journal1 := structs.NewJournal("mkdir", "/", "", time.Now())
	journal2 := structs.NewJournal("mkfile", "users.txt", userstxt, time.Now())

	file, err := os.OpenFile(absolutePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()
	_, err = file.Seek(int64(partition.Start), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero:", err)
		return
	}
	_, err = file.Write(superBlock.Encode())
	if err != nil {
		fmt.Println("Error al escribir SuperBlock:", err)
		return
	}
	_, err = file.Seek(int64(partition.Start+structs.SizeOfSuperBlock()), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero:", err)
		return
	}
	_, err = file.Write(append(journal1.Encode(), journal2.Encode()...))
	if err != nil {
		fmt.Println("Error al escribir los Journals:", err)
		return
	}
	_, err = file.Seek(int64(superBlock.Bm_inode_start), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero:", err)
		return
	}
	bmInodes_bmBlocks := make([]byte, n+3*n)
	for i := 0; i < 2; i++ {
		bmInodes_bmBlocks[i] = '1'
	}
	for i := 2; i < n; i++ {
		bmInodes_bmBlocks[i] = '0'
	}
	for i := n; i < n+2; i++ {
		bmInodes_bmBlocks[i] = '1'
	}
	for i := n + 2; i < len(bmInodes_bmBlocks); i++ {
		bmInodes_bmBlocks[i] = '0'
	}
	_, err = file.Write(bmInodes_bmBlocks)
	if err != nil {
		fmt.Println("Error al escribir la lista de inodos:", err)
		return
	}
	_, err = file.Seek(int64(superBlock.Inode_start), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return
	}
	_, err = file.Write(inode0.Encode())
	if err != nil {
		fmt.Println("Error al escribir el primer inodo en el archivo", err)
		return
	}
	_, err = file.Write(inode1.Encode())
	if err != nil {
		fmt.Println("Error al escribir en el segundo inodo en el archivo:", err)
		return
	}
	_, err = file.Seek(int64(superBlock.Block_start), 0)
	if err != nil {
		fmt.Println("Error al escribir el bloqueFolder en el archivo", err)
		return
	}
	_, err = file.Write(blockFolder.Encode())
	if err != nil {
		fmt.Println("Error al escrbir el bloqueFolder en el archivo", err)
		return
	}
	_, err = file.Write(blockFile.Encode())
	if err != nil {
		fmt.Println("Error al escribir el blockFile en el archivo:", err)
		return
	}
}

func (m *Mkfs) padRight(s string, length int) string {
	return s + strings.Repeat(" ", length-len(s))
}

func (m *Mkfs) validateParams() bool {
	_, idExists := m.Params["id"]
	return idExists
}

func (m *Mkfs) printError(text string) {
	fmt.Printf("\033[31m-> %s. [%v:%v]\033[0m\n", text, m.Line, m.Column+1)
	m.Result += fmt.Sprintf("%s.\n", text)
}

func (m *Mkfs) printSuccess(diskName, name, ID, typePart, typeFs string) {
	if typePart == "P" {
		typePart = "PRIMARIA "
	} else if typePart == "E" {
		typePart = "EXTENDIDA"
	} else {
		typePart = "LOGICA "
	}
	if typeFs == "2fs" {
		typeFs = "EXT2"
	} else {
		typeFs = "EXT3"
	}
	fmt.Printf("\033[32m-> mkfs: Partición formateada \"%s\" exitosamente en el disco \"%s\". %s (%s: %s) [%v:%v]\033[0m\n", typeFs, diskName, typePart, name, ID, m.Line, m.Column)
	m.Result += fmt.Sprintf("mkfs: Partición formateada \"%s\" exitosamente en el disco \"%s\". %s (%s: %s)\n", typeFs, diskName, typePart, name, ID)
}

func (m *Mkfs) GetResult() string { return m.Result }
