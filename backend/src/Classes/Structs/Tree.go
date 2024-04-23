package structs

import (
	"fmt"
	utils "mia/Classes/Utils"
	"os"
	"strings"
)

type DataBlock struct {
	I             int32
	BlockFile     *BlockFile
	BlockFolder   *BlockFolder
	BlockPointers *BlockPointers
}

type Tree struct {
	SuperBlock *SuperBlock
	File       os.File
	Blocks     []*DataBlock
	FileBlocks []*SuperBlock
	current    int32
	prev       int32
	dir        bool
}

type BlockFileTuple struct {
	Block int
	File  *BlockFile
}

type PointerBlockData struct {
	NumBlock   int32
	I          int32
	NameBlock  string
	Simplicity int
}

func NewTree(SuperBlock SuperBlock, file os.File) *Tree {
	return &Tree{SuperBlock: &SuperBlock, File: file, dir: false}
}

// ============================== GRAPH TREE ================================

func (t *Tree) GetDot(diskname, partName string) string {
	dot := "digraph Tree{\n\tnode [shape=plaintext];\n\trankdir=LR;\n\t"
	dot += fmt.Sprintf("label=\"%v: %v\";\n\tlabelloc=t;\n\t", diskname, partName)
	dot += t.getDotInode(0)
	dot += "\n}"
	return dot
}

func (t *Tree) getDotInode(i int32) string {
	_, err := t.File.Seek(int64(t.SuperBlock.Inode_start)+int64(i*88), 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero:", err)
	}
	readedBytes := make([]byte, 88)
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
	}
	inode := DecodeInode(readedBytes)
	dot := inode.getDot(i)
	for p := 0; p < len(inode.Block); p++ {
		if inode.Block[p] != -1 {
			if p < 12 {
				if inode.Type == '0' {
					dot += "\n\t" + t.getDotBlockFolder(inode.Block[p])
				} else {
					dot += "\n\t" + t.getDotBlockFile(inode.Block[p])
				}
			} else if p == 12 {
				dot += "\n\t" + t.getDotBlockPointers(inode.Block[p], inode.Type, 1)
			} else if p == 13 {
				dot += "\n\t" + t.getDotBlockPointers(inode.Block[p], inode.Type, 2)
			} else if p == 14 {
				dot += "\n\t" + t.getDotBlockPointers(inode.Block[p], inode.Type, 3)
			}
			dot += fmt.Sprintf("\n\tinode%v:A%v -> block%v:B%v;", i, p, inode.Block[p], inode.Block[p])
		}
	}
	return dot
}

func (t *Tree) getDotBlockPointers(i int32, inodeType rune, simplicity int) string {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(64), 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero seek:", err)
	}
	readedBytes := make([]byte, 64)
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
	}
	blockPointers := DecodeBlockPointers(readedBytes)
	dot := blockPointers.GetDot(int(i))
	for p := 0; p < len(blockPointers.Pointers); p++ {
		if blockPointers.Pointers[p] != -1 {
			if simplicity == 1 {
				if inodeType == '0' {
					dot += "\n\t" + t.getDotBlockFolder(int32(blockPointers.Pointers[p]))
				} else {
					dot += "\n\t" + t.getDotBlockFile(int32(blockPointers.Pointers[p]))
				}
			} else {
				dot += "\n\t" + t.getDotBlockPointers(int32(blockPointers.Pointers[p]), inodeType, simplicity-1)
			}
			dot += fmt.Sprintf("\n\tblock%v:A%v -> block%v:B%v;", i, p, blockPointers.Pointers[p], blockPointers.Pointers[p])
		}
	}
	return dot
}

func (t *Tree) getDotBlockFolder(i int32) string {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(64), 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero seek:", err)
	}
	readedBytes := make([]byte, 64)
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
	}
	blockFolder := DecodeBlockFolder(readedBytes)
	dot := blockFolder.GetDot(int(i))
	for p := 0; p < len(blockFolder.Content); p++ {
		if strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) != "." && strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) != ".." && blockFolder.Content[p].Inode != -1 {
			dot += "\n\t" + t.getDotInode(blockFolder.Content[p].Inode)
			dot += fmt.Sprintf("\n\tblock%v:A%v -> inode%v:I%v;", i, p, blockFolder.Content[p].Inode, blockFolder.Content[p].Inode)
		}
	}
	return dot
}

func (t *Tree) getDotBlockFile(i int32) string {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*64, 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero seek:", err)
	}
	readedBytes := make([]byte, 64)
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
	}
	blockFile := DecodeBlockFile(readedBytes)
	return blockFile.GetDot(int(i))
}

// =================================== GET BLOCKS ==================================

func (t *Tree) GetBlocks() []*DataBlock {
	t.searchInInodes(0)
	if len(t.Blocks) > 1 {
		for i := 1; i < len(t.Blocks); i++ {
			for j := i; j > 0; j-- {
				if t.Blocks[j].I < t.Blocks[j-1].I {
					t.Blocks[j], t.Blocks[j-1] = t.Blocks[j-1], t.Blocks[j]
					continue
				}
				break
			}
		}
	}
	return t.Blocks
}

func (t *Tree) searchInInodes(i int32) {
	t.File.Seek(int64(t.SuperBlock.Inode_start)+int64(i*88), 0)
	inodeData := make([]byte, 88)
	_, err := t.File.Read(inodeData)
	if err != nil {
		fmt.Println("Error al leer el inode:", err)
		return
	}
	inode := DecodeInode(inodeData)
	for p := 0; p < len(inode.Block); p++ {
		if inode.Block[p] != -1 {
			if p < 12 {
				if inode.Type == '0' {
					t.searchInBlockFolder(inode.Block[p])
				} else {
					t.searchInBlockFile(inode.Block[p])
				}
			} else if p == 12 {
				t.searchInBlockPointers(inode.Block[p], inode.Type, 1)
			} else if p == 13 {
				t.searchInBlockPointers(inode.Block[p], inode.Type, 2)
			} else if p == 14 {
				t.searchInBlockPointers(inode.Block[p], inode.Type, 3)
			}
		}
	}
}

func (t *Tree) searchInBlockPointers(i int32, inodeType rune, simplicity int) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(64), 0)
	blockPointersData := make([]byte, 64)
	_, err := t.File.Read(blockPointersData)
	if err != nil {
		fmt.Println("Error al leer el block pointer:", err)
		return
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	t.Blocks = append(t.Blocks, &DataBlock{I: i, BlockPointers: blockPointers})
	for p := 0; p < len(blockPointers.Pointers); p++ {
		if blockPointers.Pointers[p] != -1 {
			if simplicity == 1 {
				if inodeType == '0' {
					t.searchInBlockFolder(int32(blockPointers.Pointers[p]))
				} else {
					t.searchInBlockFile(int32(blockPointers.Pointers[p]))
				}
			} else {
				t.searchInBlockPointers(int32(blockPointers.Pointers[p]), inodeType, simplicity-1)
			}
		}
	}
}

func (t *Tree) searchInBlockFolder(i int32) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i*64), 0)
	blockFolderData := make([]byte, 64)
	_, err := t.File.Read(blockFolderData)
	if err != nil {
		fmt.Println("Error al leer el block folder:", err)
		return
	}
	blockFolder := DecodeBlockFolder(blockFolderData)
	t.Blocks = append(t.Blocks, &DataBlock{I: i, BlockFolder: blockFolder})
	for p := 0; p < len(blockFolder.Content); p++ {
		if strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) != "." && strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) != ".." && blockFolder.Content[p].Inode != -1 {
			t.searchInInodes(int32(blockFolder.Content[p].Inode))
		}
	}
}

func (t *Tree) searchInBlockFile(i int32) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(64), 0)
	blockFileData := make([]byte, 64)
	_, err := t.File.Read(blockFileData)
	if err != nil {
		fmt.Println("Error al leer el block file:", err)
		return
	}
	blockFile := DecodeBlockFile(blockFileData)
	t.Blocks = append(t.Blocks, &DataBlock{I: i, BlockFile: blockFile})
}

// ================================== READ CONTENT ==================================

func (t *Tree) ReadFile(path string) (string, bool) {
	dir := strings.Split(path, "/")
	var cleanDir []string
	for _, d := range dir {
		if d != "" {
			cleanDir = append(cleanDir, d)
		}
	}
	return t.readFileInInodes(0, cleanDir)
}

func (t *Tree) readFileInInodes(i int32, path []string) (string, bool) {
	t.File.Seek(int64(t.SuperBlock.Inode_start)+int64(i)*int64(SizeOfInode()), 0)
	inodeData := make([]byte, SizeOfInode())
	_, err := t.File.Read(inodeData)
	if err != nil {
		if err.Error() == "EOF" {
			return "", false
		}
		fmt.Println("Error al leer el inode (READ FILE):", err)
		return "", false
	}
	inode := DecodeInode(inodeData)
	content := ""
	founded := false
	for p := 0; p < len(inode.Block); p++ {
		if inode.Block[p] == -1 {
			break
		}
		if inode.Block[p] != -1 {
			if p < 12 {
				if inode.Type == '0' {
					content, founded = t.readFileInBlockFolder(inode.Block[p], path)
					if founded {
						return content, founded
					}
				} else {
					cont := ""
					cont, founded = t.readFileInBlockFile(inode.Block[p])
					content += cont
				}
			}
		} else if p == 12 {
			if inode.Type == '0' {
				content, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 1)
			} else {
				cont := ""
				cont, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 1)
				content += cont
			}
		} else if p == 13 {
			if inode.Type == '0' {
				content, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 2)
			} else {
				cont := ""
				cont, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 2)
				content += cont
			}
		} else if p == 14 {
			if inode.Type == '0' {
				content, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 3)
			} else {
				cont := ""
				cont, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 3)
				content += cont
			}
		}
	}
	return content, founded
}

func (t *Tree) readFileInBlockPointers(i int32, path []string, inodeType rune, simplicity int) (string, bool) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(SizeOfBlockPointers()), 0)
	blockPointersData := make([]byte, SizeOfBlockPointers())
	_, err := t.File.Read(blockPointersData)
	if err != nil {
		if err.Error() == "EOF" {
			return "", false
		}
		fmt.Println("Error al leer el block pointer (READ FILE):", err)
		return "", false
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	content := ""
	founded := false
	for p := 0; p < len(blockPointers.Pointers); p++ {
		if blockPointers.Pointers[p] != -1 {
			if simplicity == 1 {
				if inodeType == '0' {
					content, founded = t.readFileInBlockFolder(blockPointers.Pointers[p], path)
					if founded {
						return content, founded
					}
				} else {
					cont := ""
					cont, founded = t.readFileInBlockFile(blockPointers.Pointers[p])
					content += cont
				}
			} else {
				if inodeType == '0' {
					content, founded = t.readFileInBlockPointers(blockPointers.Pointers[p], path, inodeType, simplicity-1)
				} else {
					count := ""
					count, founded = t.readFileInBlockPointers(blockPointers.Pointers[p], path, inodeType, simplicity-1)
					content += count
				}
			}
		} else {
			break
		}
	}
	return content, founded
}

func (t *Tree) readFileInBlockFolder(i int32, path []string) (string, bool) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i*64), 0)
	blockFolderData := make([]byte, 64)
	_, err := t.File.Read(blockFolderData)
	if err != nil {
		if err.Error() == "EOF" {
			return "", false
		}
		fmt.Println("Error al leer el block folder (READ FILE):", err)
		return "", false
	}
	blockFolder := DecodeBlockFolder(blockFolderData)
	for p := 0; p < len(blockFolder.Content); p++ {
		if !(utils.CompareStrings(blockFolder.Content[p].Name, ".") || utils.CompareStrings(blockFolder.Content[p].Name, "..")) && blockFolder.Content[p].Inode != -1 && utils.CompareStrings(blockFolder.Content[p].Name, path[0]) {
			path = path[1:]
			return t.readFileInInodes(blockFolder.Content[p].Inode, path)
		}
	}
	return "", false
}

func (t *Tree) readFileInBlockFile(i int32) (string, bool) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(64), 0)
	blockFileData := make([]byte, 64)
	_, err := t.File.Read(blockFileData)
	if err != nil {
		if err.Error() == "EOF" {
			return "", false
		}
		fmt.Println("Error al leer el block file (READ FILE):", err)
		return "", false
	}
	blockFile := DecodeBlockFile(blockFileData)
	return blockFile.Content, true
}

// ================================= WRITE FILE ===================================

func (t *Tree) WriteFile(path, diskpath string, partstart int, newContent string) {
	dir := strings.Split(path, "/")
	var cleanDir []string
	for _, d := range dir {
		if d != "" {
			cleanDir = append(cleanDir, d)
		}
	}
	t.writeFileInInodes(0, cleanDir, diskpath, newContent, partstart)
	t.SuperBlock.First_blo = t.findNextFreeBlock(1)[0]
	t.SuperBlock.First_ino = t.findNextFreeInode(1)[0]
}

func (t *Tree) writeFileInInodes(i int32, path []string, pathdsk, newContent string, partstart int) {
	t.File.Seek(int64(t.SuperBlock.Inode_start)+int64(i)*int64(88), 0)
	inodeData := make([]byte, 88)
	_, err := t.File.Read(inodeData)
	if err != nil {
		fmt.Println("Error al leer el inode:", err)
	}
	inode := DecodeInode(inodeData)
	if inode.Type == '0' {
		for p := 0; p < len(inode.Block); p++ {
			if inode.Block[p] != -1 {
				writed := t.writeFileInBlockFolder(inode.Block[p], path, pathdsk, newContent, int32(partstart))
				if writed {
					return
				}
			}
		}
	} else {
		blocksFile := &BlockFileTuple{Block: -1, File: nil}
		for p := 0; p < len(inode.Block); p++ {
			if inode.Block[p] != -1 {
				if p < 12 {
					blocksFile = t.writeFileInBlockFile(inode.Block[p])
				} else if p == 12 {
					blocksFile = t.writeFileInBlockPointers3(pathdsk, inode.Block[p], 1)
				} else if p == 13 {
					blocksFile = t.writeFileInBlockPointers3(pathdsk, inode.Block[p], 2)
				} else if p == 14 {
					blocksFile = t.writeFileInBlockPointers3(pathdsk, inode.Block[p], 3)
				}
			}
		}
		num, block := blocksFile.Block, blocksFile.File
		if block != nil {
			block = &BlockFile{Content: ""}
			num = t.findNextFreeBlock(1)[0]
			inode.Block[0] = int32(num)
			t.rewriteInode(pathdsk, i, inode)
		}
		contents := [][]rune{{}}
		for _, z := range newContent {
			if len(contents[len(contents)-1]) < 64 {
				contents[len(contents)-1] = append(contents[len(contents)-1], rune(z))
			} else {
				contents = append(contents, []rune{z})
			}
		}
		blockContent := contents[0]
		t.WriteInDisk(pathdsk, int32(t.SuperBlock.Block_start+num*64), []byte(string(blockContent)))
		newSizeInode := (num-1)*64 + len(blockContent)
		for len(contents) > 0 {
			newBlock := BlockFile{Content: ""}
			contenew := contents[0]
			newSizeInode = (num-1)*64 + len(contenew)
			if len(newBlock.Content) == len(contenew) {
				newBlock.Content = string(contenew)
			} else {
				for z := 0; z < len(contenew); z++ {
					contentRunes := []rune(newBlock.Content)
					contentRunes[z] = contenew[z]
					newBlock.Content = string(contentRunes)
				}
			}
			nextFreeBitBlock := t.findNextFreeBlock(1)
			for h := 0; h < len(inode.Block); h++ {
				if inode.Block[h] == -1 {
					if h < 12 {
						inode.Block[h] = int32(nextFreeBitBlock[0])
						t.WriteInDisk(pathdsk, int32(t.SuperBlock.Bm_block_start+nextFreeBitBlock[0]), []byte{1})
						t.WriteInDisk(pathdsk, int32(t.SuperBlock.Block_start+nextFreeBitBlock[0]*64), newBlock.Encode())
						t.SuperBlock.First_blo = t.findNextFreeBlock(1)[0]
						t.SuperBlock.Free_blocks_count -= 1
					} else if h == 12 {
						inode.Block[h] = int32(nextFreeBitBlock[0])
						t.writeFileInBlockPointers(pathdsk, newBlock, int32(nextFreeBitBlock[0]), 1)
						t.SuperBlock.Free_blocks_count -= 1
						t.WriteInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					} else if h == 13 {
						inode.Block[h] = int32(nextFreeBitBlock[0])
						t.writeFileInBlockPointers(pathdsk, newBlock, int32(nextFreeBitBlock[0]), 2)
						t.WriteInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					} else if h == 14 {
						inode.Block[h] = int32(nextFreeBitBlock[0])
						t.writeFileInBlockPointers(pathdsk, newBlock, int32(nextFreeBitBlock[0]), 3)
						t.WriteInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					}
					break
				} else if h == 12 && inode.Block[h] != -1 && t.validateSpacePointers(inode.Block[h], 1) {
					t.writeFileInBlockPointers(pathdsk, newBlock, inode.Block[h], 1)
					t.SuperBlock.Free_blocks_count -= 1
					t.WriteInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					break
				} else if h == 13 && inode.Block[h] != -1 && t.validateSpacePointers(inode.Block[h], 2) {
					posiblePointer := t.searchPointer(inode.Block[h], 2)
					if posiblePointer[0] != -1 {
						t.writeNewBlockInIndirect(pathdsk, int32(posiblePointer[0]), newBlock.Encode(), int32(posiblePointer[1]))
						break
					}
					t.writeFileInBlockPointers(pathdsk, newBlock, inode.Block[h], 2)
					t.SuperBlock.Free_blocks_count -= 1
					t.WriteInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					break
				} else if h == 14 && inode.Block[h] != -1 && t.validateSpacePointers(inode.Block[h], 3) {
					posiblePointer := t.searchPointer(inode.Block[h], 3)
					if posiblePointer[0] != -1 {
						t.writeNewBlockInIndirect(pathdsk, int32(posiblePointer[0]), newBlock.Encode(), int32(posiblePointer[1]))
						break
					}
					posiblePointer = t.searchPointer(inode.Block[h], 2)
					if posiblePointer[0] != -1 {
						t.writeNewBlockInIndirect(pathdsk, int32(posiblePointer[0]), (&BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}).Encode(), int32(posiblePointer[1]))
						posiblePointer = t.searchPointer(inode.Block[h], 3)
						t.writeNewBlockInIndirect(pathdsk, int32(posiblePointer[0]), newBlock.Encode(), int32(posiblePointer[1]))
						break
					}
					t.writeFileInBlockPointers(pathdsk, newBlock, inode.Block[h], 3)
					t.SuperBlock.Free_blocks_count -= 1
					t.WriteInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					break
				}
			}
			inode.Size = newSizeInode
			t.WriteInDisk(pathdsk, int32(t.SuperBlock.Inode_start)+i*64, inode.Encode())
		}
	}
}

func (t *Tree) validateSpacePointers(numBlock, simplicity int32) bool {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(numBlock)*int64(64), 0)
	blockPointersData := make([]byte, 64)
	_, err := t.File.Read(blockPointersData)
	if err != nil {
		fmt.Println("Error al leer el block pointer:", err)
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	resultado := false
	for i := 0; i < len(blockPointers.Pointers); i++ {
		if simplicity == 1 {
			if blockPointers.Pointers[i] == -1 {
				return true
			} else {
				if blockPointers.Pointers[i] != -1 {
					resultado = t.validateSpacePointers(blockPointers.Pointers[i], simplicity-1)
				} else {
					return true
				}
			}
		}
	}
	return resultado
}

func (t *Tree) writeFileInBlockPointers(pathdsk string, newBlockFile BlockFile, nextFreeBitBlock, simplicity int32) {
	file, err := os.OpenFile(pathdsk, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error al leer el archivo", err)
		return
	}
	file.Close()
	_, err = file.Seek(int64(t.SuperBlock.Block_start)+int64(nextFreeBitBlock)*int64(SizeOfBlockPointers()), 0)
	if err != nil {
		fmt.Println("Error al posicionar el seek", err)
		return
	}
	blockPointersData := make([]byte, 64)
	_, err = file.Read(blockPointersData)
	if err != nil {
		fmt.Println("Error al leer el block pointer:", err)
		return
	}
	if string(blockPointersData) != strings.Repeat("\x00", SizeOfBlockPointers()) {
		t.writeFileInBlockPointers2(pathdsk, newBlockFile, nextFreeBitBlock, simplicity)
	} else {
		blockPointers := BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
		t.WriteInDisk(pathdsk, int32(t.SuperBlock.Bm_block_start)+nextFreeBitBlock, []byte{1})
		t.WriteInDisk(pathdsk, int32(t.SuperBlock.Block_start)+nextFreeBitBlock*int32(SizeOfBlockPointers()), blockPointers.Encode())
		t.SuperBlock.Free_blocks_count -= 1
		t.writeFileInBlockPointers2(pathdsk, newBlockFile, nextFreeBitBlock, simplicity)
	}
}

func (t *Tree) writeFileInBlockPointers2(pathdsk string, newBlockFile BlockFile, nextFreeBitBlock, simplicity int32) {
	file, err := os.OpenFile(pathdsk, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error al leer el archivo", err)
		return
	}
	file.Close()
	_, err = file.Seek(int64(t.SuperBlock.Block_start)+int64(nextFreeBitBlock)*int64(SizeOfBlockPointers()), 0)
	if err != nil {
		fmt.Println("Error al posicionar el seek", err)
		return
	}
	blockPointersData := make([]byte, 64)
	_, err = file.Read(blockPointersData)
	if err != nil {
		fmt.Println("Error al leer el block pointer:", err)
		return
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	newNextBit := t.findNextFreeBlock(1)[0]
	for p := 0; p < len(blockPointers.Pointers); p++ {
		if blockPointers.Pointers[p] == -1 {
			if simplicity == 1 {
				blockPointers.Pointers[p] = int32(newNextBit)
				t.WriteInDisk(pathdsk, int32(t.SuperBlock.Block_start)+nextFreeBitBlock*int32(SizeOfBlockPointers()), blockPointers.Encode())
				t.WriteInDisk(pathdsk, int32(t.SuperBlock.Block_start)+int32(newNextBit)*int32(SizeOfBlockFile()), newBlockFile.Encode())
				t.WriteInDisk(pathdsk, int32(t.SuperBlock.Bm_block_start)+int32(newNextBit), []byte{1})
				t.SuperBlock.Free_blocks_count -= 1
				return
			} else {
				blockPointers.Pointers[p] = int32(newNextBit)
				t.WriteInDisk(pathdsk, int32(t.SuperBlock.Block_start)+nextFreeBitBlock*int32(SizeOfBlockPointers()), blockPointers.Encode())
				t.writeFileInBlockPointers(pathdsk, newBlockFile, int32(t.findNextFreeBlock(1)[0]), simplicity-1)
			}
		}
	}
}

func (t *Tree) writeNewBlockInIndirect(pathdsk string, numBlock int32, newBlock []byte, numPtr int32) {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(numBlock)*int64(SizeOfBlockPointers()), 0)
	if err != nil {
		fmt.Println("Error al posicionar el seek", err)
		return
	}
	blockPointersData := make([]byte, 64)
	_, err = t.File.Read(blockPointersData)
	if err != nil {
		fmt.Println("Error al leer el block pointer:", err)
		return
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	nextBitBlock := t.findNextFreeBlock(1)[0]
	blockPointers.Pointers[numPtr] = int32(nextBitBlock)
	t.SuperBlock.Free_blocks_count -= 1
	t.SuperBlock.First_blo = nextBitBlock
	t.WriteInDisk(pathdsk, int32(t.SuperBlock.Block_start)+numBlock*int32(SizeOfBlockPointers()), blockPointers.Encode())
	t.WriteInDisk(pathdsk, int32(t.SuperBlock.Block_start)+int32(nextBitBlock)*int32(SizeOfBlockFile()), newBlock)
	t.WriteInDisk(pathdsk, int32(t.SuperBlock.Bm_block_start)+int32(nextBitBlock), []byte{1})
}

func (t *Tree) searchPointer(numBlock, simplicity int32) []int {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(numBlock)*int64(SizeOfBlockPointers()), 0)
	blockPointersData := make([]byte, 64)
	_, err = t.File.Read(blockPointersData)
	if err != nil {
		fmt.Println("Error al leer el block pointer:", err)
		return []int{}
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	resultado := []int{-1, -1}
	for i := 0; i < len(blockPointers.Pointers); i++ {
		if simplicity == 1 {
			if blockPointers.Pointers[i] == -1 {
				return []int{int(numBlock), i}
			}
		} else {
			if blockPointers.Pointers[i] != -1 {
				resultado = t.searchPointer(blockPointers.Pointers[i], simplicity-1)
			}
		}
	}
	return resultado
}

func (t *Tree) writeFileInBlockPointers3(pathdsk string, i, simplicity int32) *BlockFileTuple {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(SizeOfBlockPointers()), 0)
	if err != nil {
		fmt.Println("Error al posicionar el seek", err)
		return nil
	}
	blockPointersData := make([]byte, 64)
	_, err = t.File.Read(blockPointersData)
	if err != nil {
		fmt.Println("Error al leer el block pointer:", err)
		return nil
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	var blockFile *BlockFileTuple
	blockFile.Block = -1
	blockFile.File = nil
	for i := 0; i < len(blockPointers.Pointers); i++ {
		if blockPointers.Pointers[i] != -1 {
			if simplicity == 1 {
				blockFile.Block = t.writeFileInBlockFile(blockPointers.Pointers[i]).Block
				blockFile.File = t.writeFileInBlockFile(blockPointers.Pointers[i]).File
			} else {
				blockFile.Block = t.writeFileInBlockPointers3(pathdsk, blockPointers.Pointers[i], simplicity-1).Block
				blockFile.File = t.writeFileInBlockPointers3(pathdsk, blockPointers.Pointers[i], simplicity-1).File
			}
		}
	}
	return blockFile
}

func (t *Tree) writeFileInBlockFolder(i int32, path []string, pathdsk, content string, partstart int32) bool {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i*int32(SizeOfBlockFolder())), 0)
	if err != nil {
		fmt.Println("Error al posicionar el seek:", err)
	}
	blockFolderData := make([]byte, 64)
	_, err = t.File.Read(blockFolderData)
	if err != nil {
		fmt.Println("Error al leer el block folder:", err)
	}
	blockFolder := DecodeBlockFolder(blockFolderData)
	for p := 0; p < len(blockFolder.Content); p++ {
		if strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) != "." && strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) != ".." && blockFolder.Content[p].Inode != -1 && strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) == path[0] {
			path = path[1:]
			t.writeFileInInodes(blockFolder.Content[p].Inode, path, pathdsk, content, int(partstart))
			return true
		}
	}
	return false
}

func (t *Tree) writeFileInBlockFile(i int32) *BlockFileTuple {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i*int32(SizeOfBlockFile())), 0)
	if err != nil {
		fmt.Println("Error al posicionar el seek:", err)
		return nil
	}
	blockFileData := make([]byte, 64)
	_, err = t.File.Read(blockFileData)
	if err != nil {
		fmt.Println("Error al leer el block file:", err)
	}
	blockFile := DecodeBlockFile(blockFileData)
	var blockFileTuple BlockFileTuple
	blockFileTuple.Block = int(i)
	blockFileTuple.File = blockFile
	return &blockFileTuple
}

// =========================================== MKDIR AND MKFILE =============================================

func (t *Tree) Mkdir(path, diskPath string) bool {
	dirs := strings.Split(path, "/")
	var cleanedDirs []string
	for _, dir := range dirs {
		if dir != "" {
			cleanedDirs = append(cleanedDirs, dir)
		}
	}
	t.current = 0
	t.prev = 0
	t.dir = true
	result := t.mkdirInInode(0, dirs, diskPath)
	t.SuperBlock.First_blo = t.findNextFreeBlock(1)[0]
	t.SuperBlock.First_ino = t.findNextFreeInode(1)[0]
	return result
}

func (t *Tree) Mkfile(path, diskPath string) bool {
	dirs := strings.Split(path, "/")
	var cleanedDirs []string
	for _, dir := range dirs {
		if dir != "" {
			cleanedDirs = append(cleanedDirs, dir)
		}
	}
	t.current = 0
	t.prev = 0
	t.dir = true
	result := t.mkdirInInode(0, dirs, diskPath)
	t.SuperBlock.First_blo = t.findNextFreeBlock(1)[0]
	t.SuperBlock.First_ino = t.findNextFreeInode(1)[0]
	return result
}

func (t *Tree) mkdirInInode(i int32, path []string, pathdsk string) bool {
	_, err := t.File.Seek(int64(t.SuperBlock.Inode_start)+int64(i)*int64(SizeOfInode()), 0)
	inodeData := make([]byte, 88)
	_, err = t.File.Read(inodeData)
	if err != nil {
		fmt.Println("Error al leer el inode:", err)
	}
	inode := DecodeInode(inodeData)
	t.prev = t.current
	t.current = i
	for p := 0; p < len(inode.Block); p++ {
		if p < 12 {
			if inode.Block[p] != -1 {
				mkdirBlockFolder := t.mkdirInBlockFolder(inode.Block[p], path, pathdsk)
				if mkdirBlockFolder {
					return mkdirBlockFolder
				}
			} else {
				inode.Block[p] = int32(t.findNextFreeBlock(1)[0])
				t.rewriteInode(pathdsk, i, inode)
				newBlockFolder := &BlockFolder{}
				for i := 0; i < 4; i++ {
					newBlockFolder.Content = append(newBlockFolder.Content, &Content{})
				}
				t.writeNewBlock(pathdsk, inode.Block[p], newBlockFolder.Encode())
				return t.mkdirInBlockFolder(inode.Block[p], path, pathdsk)
			}
		} else if p == 12 {
			if inode.Block[p] != -1 {
				copyPath := make([]string, len(path))
				for i, v := range path {
					copyPath[i] = v
				}
				posiblePtr := t.searchPointerFolder(inode.Block[p], 1, copyPath)
				if posiblePtr.NumBlock != -1 {
					switch posiblePtr.NameBlock {
					case "BLOCKFOLDER":
						return t.mkdirInBlockFolder(posiblePtr.NumBlock, []string{path[len(path)-1]}, pathdsk)
					case "INODE":
						return t.mkdirInInode(posiblePtr.NumBlock, []string{path[len(path)-1]}, pathdsk)
					case "BLOCKPOINTERS":
						_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(posiblePtr.NumBlock)*int64(SizeOfBlockPointers()), 0)
						if err != nil {
							fmt.Println("Error al posicionar el seek:", err)
						}
						blockPointersData := make([]byte, 64)
						_, err = t.File.Read(blockPointersData)
						if err != nil {
							fmt.Println("Error al leer el block pointer:", err)
						}
						blockPointers := DecodeBlockPointers(blockPointersData)
						blockPointers.Pointers[posiblePtr.I] = int32(t.findNextFreeBlock(1)[0])
						t.rewriteBlock(pathdsk, posiblePtr.NumBlock, blockPointers.Encode())

						newBlockFolder := &BlockFolder{}
						for i := 0; i < 4; i++ {
							newBlockFolder.Content = append(newBlockFolder.Content, &Content{})
						}
						t.writeNewBlock(pathdsk, blockPointers.Pointers[posiblePtr.I], newBlockFolder.Encode())
						return t.mkdirInBlockFolder(blockPointers.Pointers[posiblePtr.I], []string{path[len(path)-1]}, pathdsk)
					}
				}
				continue
			} else {
				inode.Block[p] = int32(t.findNextFreeBlock(1)[0])
				t.rewriteInode(pathdsk, i, inode)

				newBlockPointers := BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
				t.writeNewBlock(pathdsk, inode.Block[p], newBlockPointers.Encode())
				newBlockPointers.Pointers[0] = int32(t.findNextFreeBlock(1)[0])
				t.rewriteBlock(pathdsk, inode.Block[p], newBlockPointers.Encode())

				newBlockFolder := &BlockFolder{}
				for i := 0; i < 4; i++ {
					newBlockFolder.Content = append(newBlockFolder.Content, &Content{})
				}
				t.writeNewBlock(pathdsk, newBlockPointers.Pointers[0], newBlockFolder.Encode())

				return t.mkdirInBlockFolder(newBlockPointers.Pointers[0], path, pathdsk)
			}
		} else if p == 13 {
			if inode.Block[p] != -1 {
				copyPath := make([]string, len(path))
				for i, v := range path {
					copyPath[i] = v
				}
				posiblePtr := t.searchPointerFolder(inode.Block[p], 1, copyPath)
				if posiblePtr.NumBlock != -1 {
					switch posiblePtr.Simplicity {
					case 1:
						_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(posiblePtr.NumBlock)*int64(SizeOfBlockPointers()), 0)
						if err != nil {
							fmt.Println("Error al posicionar el seek:", err)
						}
						blockPointersData := make([]byte, 64)
						_, err = t.File.Read(blockPointersData)
						if err != nil {
							fmt.Println("Error al leer el block pointer:", err)
						}
						blockPointers := DecodeBlockPointers(blockPointersData)
						blockPointers.Pointers[posiblePtr.I] = int32(t.findNextFreeBlock(1)[0])
						t.rewriteBlock(pathdsk, posiblePtr.NumBlock, blockPointers.Encode())

						newBlockFolder := &BlockFolder{}
						for i := 0; i < 4; i++ {
							newBlockFolder.Content = append(newBlockFolder.Content, &Content{})
						}
						t.writeNewBlock(pathdsk, blockPointers.Pointers[0], newBlockFolder.Encode())
						return t.mkdirInBlockFolder(blockPointers.Pointers[posiblePtr.I], []string{path[len(path)-1]}, pathdsk)
					case 2:
						_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(posiblePtr.NumBlock)*int64(SizeOfBlockPointers()), 0)
						if err != nil {
							fmt.Println("Error al posicionar el seek:", err)
						}
						blockPointersData := make([]byte, 64)
						_, err = t.File.Read(blockPointersData)
						if err != nil {
							fmt.Println("Error al leer el block pointer:", err)
						}
						blockPointers := DecodeBlockPointers(blockPointersData)
						blockPointers.Pointers[posiblePtr.I] = int32(t.findNextFreeBlock(1)[0])
						t.rewriteBlock(pathdsk, posiblePtr.NumBlock, blockPointers.Encode())

						newBlockPointers := &BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
						t.writeNewBlock(pathdsk, blockPointers.Pointers[posiblePtr.I], newBlockPointers.Encode())
						newBlockPointers.Pointers[0] = int32(t.findNextFreeBlock(1)[0])
						t.rewriteBlock(pathdsk, blockPointers.Pointers[posiblePtr.I], newBlockPointers.Encode())

						newBlockFolder := &BlockFolder{}
						for i := 0; i < 4; i++ {
							newBlockFolder.Content = append(newBlockFolder.Content, &Content{})
						}
						t.writeNewBlock(pathdsk, newBlockPointers.Pointers[0], newBlockFolder.Encode())
						return t.mkdirInBlockFolder(newBlockPointers.Pointers[0], []string{path[len(path)-1]}, pathdsk)
					}
				}
				continue
			} else {
				inode.Block[p] = int32(t.findNextFreeBlock(1)[0])
				t.rewriteInode(pathdsk, i, inode)

				newBlockPointers1 := &BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
				t.writeNewBlock(pathdsk, inode.Block[p], newBlockPointers1.Encode())
				newBlockPointers1.Pointers[0] = int32(t.findNextFreeBlock(1)[0])
				t.rewriteBlock(pathdsk, inode.Block[p], newBlockPointers1.Encode())

				newBlockPointers2 := &BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
				t.writeNewBlock(pathdsk, newBlockPointers1.Pointers[0], newBlockPointers2.Encode())
				newBlockPointers2.Pointers[0] = int32(t.findNextFreeBlock(1)[0])
				t.rewriteBlock(pathdsk, newBlockPointers1.Pointers[0], newBlockPointers2.Encode())

				newBlockFolder := &BlockFolder{}
				for i := 0; i < 4; i++ {
					newBlockFolder.Content = append(newBlockFolder.Content, &Content{})
				}
				t.writeNewBlock(pathdsk, newBlockPointers2.Pointers[0], newBlockFolder.Encode())

				return t.mkdirInBlockFolder(newBlockPointers2.Pointers[0], path, pathdsk)
			}
		} else if p == 14 {
			if inode.Block[p] != -1 {
				copyPath := make([]string, len(path))
				for i, v := range path {
					copyPath[i] = v
				}
				posiblePtr := t.searchPointerFolder(inode.Block[p], 3, copyPath)
				if posiblePtr.NumBlock != -1 {
					switch posiblePtr.NameBlock {
					case "BLOCKFOLDER":
						return t.mkdirInBlockFolder(posiblePtr.NumBlock, []string{path[len(path)-1]}, pathdsk)
					case "INODE":
						return t.mkdirInInode(posiblePtr.NumBlock, []string{path[len(path)-1]}, pathdsk)
					case "BLOCKPOINTERS":
						switch posiblePtr.Simplicity {
						case 1:
							t.File.Seek(int64(t.SuperBlock.Block_start)+int64(posiblePtr.NumBlock)*int64(SizeOfBlockPointers()), 0)
							blockPointersData := make([]byte, 64)
							_, err = t.File.Read(blockPointersData)
							if err != nil {
								fmt.Println("Error al leer el block pointer:", err)
							}
							blockPointers := DecodeBlockPointers(blockPointersData)
							blockPointers.Pointers[posiblePtr.I] = int32(t.findNextFreeBlock(1)[0])
							t.rewriteBlock(pathdsk, posiblePtr.NumBlock, blockPointers.Encode())

							newBlockFolder := &BlockFolder{}
							for i := 0; i < 4; i++ {
								newBlockFolder.Content = append(newBlockFolder.Content, &Content{})
							}
							t.writeNewBlock(pathdsk, blockPointers.Pointers[posiblePtr.NumBlock], NewBlockFolder().Encode())
							return t.mkdirInBlockFolder(blockPointers.Pointers[posiblePtr.NumBlock], []string{path[len(path)-1]}, pathdsk)

						case 2:
							t.File.Seek(int64(t.SuperBlock.Block_start)+int64(posiblePtr.NumBlock)*int64(SizeOfBlockPointers()), 0)
							blockPointersData := make([]byte, 64)
							_, err = t.File.Read(blockPointersData)
							if err != nil {
								fmt.Println("Error al leer el block pointer:", err)
							}
							blockPointers := DecodeBlockPointers(blockPointersData)
							blockPointers.Pointers[posiblePtr.NumBlock] = int32(t.findNextFreeBlock(1)[0])
							t.rewriteBlock(pathdsk, posiblePtr.NumBlock, blockPointers.Encode())

							newBlockPointers := &BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
							t.writeNewBlock(pathdsk, blockPointers.Pointers[posiblePtr.I], newBlockPointers.Encode())
							newBlockPointers.Pointers[0] = int32(t.findNextFreeBlock(1)[0])
							t.rewriteBlock(pathdsk, blockPointers.Pointers[posiblePtr.I], newBlockPointers.Encode())

							newBlockFolder := &BlockFolder{}
							for i := 0; i < 4; i++ {
								newBlockFolder.Content = append(newBlockFolder.Content, &Content{})
							}
							t.writeNewBlock(pathdsk, blockPointers.Pointers[posiblePtr.NumBlock], newBlockFolder.Encode())
							return t.mkdirInBlockFolder(blockPointers.Pointers[posiblePtr.I], []string{path[len(path)-1]}, pathdsk)

						case 3:
							t.File.Seek(int64(t.SuperBlock.Block_start)+int64(posiblePtr.NumBlock)*int64(SizeOfBlockPointers()), 0)
							blockPointersData := make([]byte, 64)
							_, err = t.File.Read(blockPointersData)
							if err != nil {
								fmt.Println("Error al leer el block pointer:", err)
							}
							blockPointers := DecodeBlockPointers(blockPointersData)
							blockPointers.Pointers[posiblePtr.I] = int32(t.findNextFreeBlock(1)[0])
							t.rewriteBlock(pathdsk, posiblePtr.NumBlock, blockPointers.Encode())

							newBlockPointers1 := &BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
							t.writeNewBlock(pathdsk, blockPointers.Pointers[posiblePtr.I], newBlockPointers1.Encode())
							newBlockPointers1.Pointers[0] = int32(t.findNextFreeBlock(1)[0])
							t.rewriteBlock(pathdsk, blockPointers.Pointers[posiblePtr.I], newBlockPointers1.Encode())

							newBlockPointers2 := &BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
							t.writeNewBlock(pathdsk, newBlockPointers1.Pointers[0], newBlockPointers2.Encode())
							newBlockPointers2.Pointers[0] = int32(t.findNextFreeBlock(1)[0])
							t.rewriteBlock(pathdsk, newBlockPointers1.Pointers[0], newBlockPointers2.Encode())

							newBlockFolder := &BlockFolder{}
							for i := 0; i < 4; i++ {
								newBlockFolder.Content = append(newBlockFolder.Content, &Content{})
							}
							t.writeNewBlock(pathdsk, newBlockPointers2.Pointers[0], newBlockFolder.Encode())
							return t.mkdirInBlockFolder(newBlockPointers2.Pointers[0], []string{path[len(path)-1]}, pathdsk)
						}
					}
				}
				continue
			} else {
				inode.Block[p] = int32(t.findNextFreeBlock(1)[0])
				t.rewriteInode(pathdsk, i, inode)

				newBlockPointers1 := &BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
				t.writeNewBlock(pathdsk, inode.Block[p], newBlockPointers1.Encode())
				newBlockPointers1.Pointers[0] = int32(t.findNextFreeBlock(1)[0])
				t.rewriteBlock(pathdsk, inode.Block[p], newBlockPointers1.Encode())

				newBlockPointers2 := &BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
				t.writeNewBlock(pathdsk, newBlockPointers1.Pointers[0], newBlockPointers2.Encode())
				newBlockPointers2.Pointers[0] = int32(t.findNextFreeBlock(1)[0])
				t.rewriteBlock(pathdsk, newBlockPointers1.Pointers[0], newBlockPointers2.Encode())

				newBlockPointers3 := &BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
				t.writeNewBlock(pathdsk, newBlockPointers2.Pointers[0], newBlockPointers3.Encode())
				newBlockPointers3.Pointers[0] = int32(t.findNextFreeBlock(1)[0])
				t.rewriteBlock(pathdsk, newBlockPointers2.Pointers[0], newBlockPointers3.Encode())

				newBlockFolder := &BlockFolder{}
				for i := 0; i < 4; i++ {
					newBlockFolder.Content = append(newBlockFolder.Content, &Content{})
				}
				t.writeNewBlock(pathdsk, newBlockPointers3.Pointers[0], newBlockFolder.Encode())

				return t.mkdirInBlockFolder(newBlockPointers3.Pointers[0], path, pathdsk)
			}
		}
	}
	return false
}

func (t *Tree) mkdirInBlockFolder(i int32, path []string, pathdsk string) bool {
	var _, err = t.File.Seek(int64(t.SuperBlock.Block_start+int(i)*SizeOfBlockFolder()), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return false
	}
	var readedBytes = make([]byte, SizeOfBlockFolder())
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
		return false
	}
	var blockFolder = DecodeBlockFolder(readedBytes)
	if len(path) > 1 {
		for _, c := range blockFolder.Content {
			if c.Inode != -1 && utils.CompareStrings(c.Name, path[0]) {
				path = path[1:]
				return t.mkdirInInode(c.Inode, path, pathdsk)
			}
		}
	} else {
		for p, c := range blockFolder.Content {
			if c.Inode == -1 {
				blockFolder.Content[p].Name = t.padRight(path[0], 12)
				blockFolder.Content[p].Inode = int32(t.findNextFreeInode(1)[0])
				t.rewriteBlock(pathdsk, i, blockFolder.Encode())
				if t.dir {
					var newInode = NewInode('0', 0)
					newInode.Block[0] = int32(t.findNextFreeBlock(1)[0])
					t.writeNewInode(pathdsk, blockFolder.Content[p].Inode, newInode)

					t.prev = t.current
					t.current = c.Inode

					var firstBlockFolder = NewBlockFolder()
					firstBlockFolder.Content[0] = &Content{Name: t.padRight(".", 12), Inode: t.current}
					firstBlockFolder.Content[1] = &Content{Name: t.padRight("..", 12), Inode: t.prev}
					t.writeNewBlock(pathdsk, newInode.Block[0], firstBlockFolder.Encode())
				} else {
					var newInode = NewInode('1', 0)
					t.writeNewInode(pathdsk, blockFolder.Content[p].Inode, newInode)
				}
				return true
			}
		}
	}
	return false
}

func (m *Tree) padRight(s string, length int) string {
	return s + strings.Repeat(" ", length-len(s))
}

func (t *Tree) searchPointerInodeFolder(numInode int32, path []string) *PointerBlockData {
	var result = &PointerBlockData{-1, -1, "", -1}
	var _, err = t.File.Seek(int64(t.SuperBlock.Block_start+int(numInode)*SizeOfInode()), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return result
	}
	var readedBytes = make([]byte, SizeOfInode())
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
		return result
	}
	var inode = DecodeInode(readedBytes)
	for i, b := range inode.Block {
		if b != -1 {
			if i < 12 {
				result = t.searchPointerBlockFolder(b, path)
				if result.NumBlock != -1 {
					return result
				}
			} else if i == 12 {
				result = t.searchPointerFolder(b, 1, path)
				if result.NumBlock != -1 {
					return result
				}
			} else if i == 13 {
				result = t.searchPointerFolder(b, 2, path)
				if result.NumBlock != -1 {
					return result
				}
			} else if i == 14 {
				result = t.searchPointerFolder(b, 3, path)
				if result.NumBlock != -1 {
					return result
				}
			}
		} else {
			return &PointerBlockData{numInode, int32(i), "INODE", -1}
		}
	}
	return result
}

func (t *Tree) searchPointerFolder(numBlock int32, simplicity int, path []string) *PointerBlockData {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(numBlock)*int64(SizeOfBlockPointers()), 0)
	if err != nil {
		fmt.Println("Error al posicionar el seek:", err)
		return nil
	}
	blockPointersData := make([]byte, 64)
	_, err = t.File.Read(blockPointersData)
	if err != nil {
		fmt.Println("Error al leer el block pointer:", err)
		return nil
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	result := &PointerBlockData{-1, -1, "", -1}
	for i := 0; i < len(blockPointers.Pointers); i++ {
		if simplicity == 1 {
			if blockPointers.Pointers[i] == -1 {
				return &PointerBlockData{numBlock, int32(i), "BLOCKPOINTERS", 1}
			} else {
				result = t.searchPointerBlockFolder(blockPointers.Pointers[i], path)
				if result.NumBlock != -1 {
					return result
				}
			}
		} else {
			if blockPointers.Pointers[i] == -1 {
				return &PointerBlockData{numBlock, int32(i), "BLOCKPOINTERS", simplicity}
			} else {
				result = t.searchPointerFolder(blockPointers.Pointers[i], simplicity-1, path)
				if result.NumBlock != -1 {
					return result
				}
			}
		}
	}
	return result
}

func (t *Tree) searchPointerBlockFolder(numBlock int32, path []string) *PointerBlockData {
	var resultado = &PointerBlockData{-1, -1, "", -1}
	var _, err = t.File.Seek(int64(t.SuperBlock.Block_start+int(numBlock)*SizeOfBlockFolder()), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return resultado
	}
	var readedBytes = make([]byte, SizeOfBlockFolder())
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
		return resultado
	}
	var blockFolder = DecodeBlockFolder(readedBytes)
	if len(path) > 1 {
		for _, c := range blockFolder.Content {
			if c.Inode != -1 && utils.CompareStrings(c.Name, path[0]) {
				path = path[1:]
				return t.searchPointerInodeFolder(c.Inode, path)
			}
		}
	} else {
		for i, c := range blockFolder.Content {
			if c.Inode == -1 {
				path = path[1:]
				return &PointerBlockData{numBlock, int32(i), "BLOCKFOLDER", -1}
			}
		}
	}
	return resultado
}

func (t *Tree) writeNewBlock(pathdsk string, numBlock int32, block []byte) {
	t.WriteInDisk(pathdsk, int32(t.SuperBlock.Block_start)+numBlock*int32(SizeOfBlockFolder()), block)
	t.WriteInDisk(pathdsk, int32(t.SuperBlock.Bm_block_start)+numBlock, []byte{'1'})
	t.SuperBlock.Free_blocks_count--
}

func (t *Tree) writeNewInode(pathdsk string, numInode int32, inode *Inode) {
	t.WriteInDisk(pathdsk, int32(t.SuperBlock.Inode_start)+numInode*int32(SizeOfInode()), inode.Encode())
	t.WriteInDisk(pathdsk, int32(t.SuperBlock.Bm_inode_start)+numInode, []byte{'1'})
	t.SuperBlock.Free_inodes_count--
}

func (t *Tree) rewriteBlock(pathdsk string, numBlock int32, block []byte) {
	t.WriteInDisk(pathdsk, int32(t.SuperBlock.Block_start)+numBlock*int32(SizeOfBlockFolder()), block)
}

func (t *Tree) rewriteInode(pathdsk string, numInode int32, inode *Inode) {
	t.WriteInDisk(pathdsk, int32(t.SuperBlock.Inode_start)+numInode*int32(SizeOfInode()), inode.Encode())
}

// =========================================== SEARCH DIRECTORY =============================================

func (t *Tree) SearchDir(path string) bool {
	var dir = []string{}
	for _, p := range strings.Split(path, "/") {
		if p != "" {
			dir = append(dir, p)
		}
	}
	return t.searchDirInInode(0, dir)
}

func (t *Tree) searchDirInInode(i int, path []string) bool {
	var _, err = t.File.Seek(int64(t.SuperBlock.Inode_start+i*SizeOfInode()), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return false
	}
	var readedBytes = make([]byte, SizeOfInode())
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
		return false
	}
	var inode = DecodeInode(readedBytes)
	var result = false
	for p, b := range inode.Block {
		if b != -1 {
			if p < 12 {
				var searchDirBlockFolder = t.searchDirInBlockFolder(b, path)
				if searchDirBlockFolder {
					return searchDirBlockFolder
				}
			} else if p == 12 {
				result = t.searchDirInBlockPointers(p, path, 1)
			} else if p == 13 {
				result = t.searchDirInBlockPointers(p, path, 2)
			} else if p == 14 {
				result = t.searchDirInBlockPointers(p, path, 3)
			}
		}
	}
	return result
}

func (t *Tree) searchDirInBlockPointers(i int, path []string, simplicity int) bool {
	var _, err = t.File.Seek(int64(t.SuperBlock.Block_start+i*SizeOfBlockPointers()), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return false
	}
	var readedBytes = make([]byte, SizeOfBlockPointers())
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
		return false
	}
	var blockPointers = DecodeBlockPointers(readedBytes)
	var founded = false
	for _, b := range blockPointers.Pointers {
		if b != -1 {
			if simplicity == 1 {
				founded = t.searchDirInBlockFolder(b, path)
				if founded {
					return founded
				}
			} else {
				founded = t.searchDirInBlockPointers(int(b), path, simplicity-1)
			}
		}
	}
	return founded
}

func (t *Tree) searchDirInBlockFolder(i int32, path []string) bool {
	var _, err = t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(SizeOfBlockFolder()), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return false
	}
	var readedBytes = make([]byte, SizeOfBlockFolder())
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
		return false
	}
	var blockFolder = DecodeBlockFolder(readedBytes)
	if len(path) > 1 {
		for _, c := range blockFolder.Content {
			if c.Inode != -1 && utils.CompareStrings(c.Name, path[0]) {
				path = path[1:]
				return t.searchDirInInode(int(c.Inode), path)
			}
		}
	} else {
		for _, c := range blockFolder.Content {
			if c.Inode != -1 && utils.CompareStrings(c.Name, path[0]) {
				path = path[1:]
				return true
			}
		}
	}
	return false
}

// ========================================= FIND NEXT BIT =============================================

func (t *Tree) findNextFreeInode(count int) []int {
	var freeInodes = []int{}
	var _, err = t.File.Seek(int64(t.SuperBlock.Bm_inode_start), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return freeInodes
	}
	var readedBytes = make([]byte, t.SuperBlock.Inodes_count)
	_, err = t.File.Read(readedBytes)
	if err != nil {
		fmt.Println("-> Error al leer el archivo:", err)
		return freeInodes
	}
	var bmInode = string(readedBytes)
	for i, b := range bmInode {
		if len(freeInodes) == count {
			break
		}
		if b == '0' {
			freeInodes = append(freeInodes, i)
		}
	}
	return freeInodes
}

func (t *Tree) findNextFreeBlock(count int) []int {
	var freeBlocks = []int{}
	var _, err = t.File.Seek(int64(t.SuperBlock.Bm_block_start), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return freeBlocks
	}
	var readedBytes = make([]byte, t.SuperBlock.Blocks_count)
	_, err = t.File.Read(readedBytes)
	if err != nil {
		fmt.Println("-> Error al leer el archivo:", err)
		return freeBlocks
	}
	var bmBlock = string(readedBytes)
	for i, b := range bmBlock {
		if len(freeBlocks) == count {
			break
		}
		if b == '0' {
			freeBlocks = append(freeBlocks, i)
		}
	}
	return freeBlocks
}

// ========================================= USERS AND GROUPS ===============================================

func (t *Tree) GetUsers(content string) []*User {
	var users = []*User{}
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
			users = append(users, NewUser(reg[0], reg[2], reg[3], reg[4]))
		}
	}
	return users
}

func (t *Tree) GetGroups(content string) []*Group {
	var groups = []*Group{}
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
		if reg[1] == "G" && reg[0] != "0" {
			groups = append(groups, NewGroup(reg[0], reg[2]))
		}
	}
	return groups
}

// ========================================== WRITE IN DISK =================================================

func (t *Tree) WriteInDisk(path string, seek int32, content []byte) {
	var file1, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	_, err = file1.Seek(int64(seek), 0)
	if err != nil {
		fmt.Println("Error al mover el puntero del archivo:", err)
		return
	}
	defer file1.Close()
	_, err = file1.Write(content)
	if err != nil {
		fmt.Println("Error al escribir Journal:", err)
		return
	}
}
