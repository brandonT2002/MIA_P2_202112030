package commands

import (
	"bytes"
	"fmt"
	"math"
	env "mia/Classes/Env"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Rep struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewRep(line, column int, params map[string]string) *Rep {
	return &Rep{Type: utils.REP, Line: line, Column: column, Params: params}
}

func (r *Rep) GetLine() int {
	return r.Line
}

func (r *Rep) GetColumn() int {
	return r.Column
}

func (r *Rep) GetType() utils.Type {
	return r.Type
}

func (r *Rep) Exec() {
	name, nameExists := r.Params["name"]
	path, pathExists := r.Params["path"]
	_, idExists := r.Params["id"]

	if !nameExists || !pathExists || !idExists {
		r.printError("-> Error rep: Faltan parámetros obligatorios para generar el reporte")
		return
	}
	r.Params["path"] = strings.ReplaceAll(path, "\"", "")

	switch strings.ToLower(name) {
	case "mbr":
		r.reportMBR()
	case "disk":
		r.reportDisk()
	case "inode":
		r.reportInode()
	case "block":
		r.reportBlock()
	case "journaling":
		r.reportJournaling()
	case "bm_inode":
		r.reportBMInode()
	case "bm_block":
		r.reportBMBlock()
	case "tree":
		r.reportTree()
	case "sb":
		r.reportSb()
	default:
		_, rutaExists := r.Params["ruta"]
		if !rutaExists {
			r.printError("-> Error rep: Faltan parámetros obligatorios para generar el reporte")
			return
		}
		switch strings.ToLower(name) {
		case "file":
			r.reportFile()
		case "ls":
			r.reportLs()
		default:
			r.printError("-> Nombre de reporte desconocido")
		}
	}
}

func (r *Rep) reportMBR() {
	driveletter := string(r.Params["id"][0])
	if _, ok := env.Disks[driveletter]; ok {
		diskPath, _ := filepath.Abs(env.GetPath(driveletter))
		absolutePath, _ := filepath.Abs(diskPath)
		file, err := os.Open(absolutePath)
		if err != nil {
			if os.IsNotExist(err) {
				r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
				return
			}
			fmt.Println("Error al abrir el archivo:", err)
			return
		}
		defer file.Close()
		readedBytes := make([]byte, 153)
		_, err = file.Read(readedBytes)
		if err != nil {
			fmt.Println("Error al leer el archivo:", err)
			return
		}
		mbr := structs.DecodeMBR(readedBytes)
		dot := "digraph MBR{\n\tnode [shape=plaintext];\n"
		dot += "\n\ttabla[label=<\n\t\t<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\">"
		dot += "\n\t\t\t<TR>\n\t\t\t\t<TD BORDER=\"1\">\n\t\t\t\t\t<TABLE BORDER=\"1\" CELLBORDER=\"0\" CELLSPACING=\"0\" CELLPADDING=\"4\">"
		dot += fmt.Sprintf("\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD COLSPAN=\"2\" BGCOLOR=\"#4A235A\"><FONT COLOR=\"white\">%v</FONT></TD>\n\t\t\t\t\t\t</TR>", r.Params["id"])
		dot += "n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR=\"#4A235A\"><FONT COLOR=\"white\">MBR</FONT></TD>\n\t\t\t\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#4A235A\"></TD>\n\t\t\t\t\t\t</TR>"
		dot += fmt.Sprintf("\t\t\t<TR>\n\t\t\t\t<TD BGCOLOR=\"#FFFFFF\">mbr_tamano</TD>\n\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#FFFFFF\">%v</TD>\n\t\t\t</TR>\n", mbr.Size)
		dot += fmt.Sprintf("\t\t\t<TR>\n\t\t\t\t<TD BGCOLOR=\"#E8DAEF\">mbr_fecha_creacion</TD>\n\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#E8DAEF\">%v</TD>\n\t\t\t</TR>\n", mbr.Date)
		dot += fmt.Sprintf("\t\t\t<TR>\n\t\t\t\t<TD BGCOLOR=\"#FFFFFF\">mbr_disk_signature</TD>\n\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#FFFFFF\">%v</TD>\n\t\t\t</TR>\n", mbr.Dsk)
		for i := 0; i < len(mbr.Partitions); i++ {
			if mbr.Partitions[i].Size != 0 {
				dot += "\t\t\t<TR>\n\t\t\t\t<TD BGCOLOR=\"#4A235A\"><FONT COLOR=\"white\">Particion</FONT></TD>\n\t\t\t\t<TD COLSPAN=\"2\" BGCOLOR=\"#4A235A\"></TD>\n\t\t\t</TR>"
				dot += fmt.Sprintf("\t\t\t<TR>\n\t\t\t\t<TD BGCOLOR=\"#FFFFFF\">part_status</TD>\n\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#FFFFFF\">%c</TD>\n\t\t\t</TR>\n", mbr.Partitions[i].Status)
				dot += fmt.Sprintf("\t\t\t<TR>\n\t\t\t\t<TD BGCOLOR=\"#E8DAEF\">part_type</TD>\n\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#E8DAEF\">%c</TD>\n\t\t\t</TR>\n", mbr.Partitions[i].Type)
				dot += fmt.Sprintf("\t\t\t<TR>\n\t\t\t\t<TD BGCOLOR=\"#FFFFFF\">part_fit</TD>\n\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#FFFFFF\">%c</TD>\n\t\t\t</TR>\n", mbr.Partitions[i].Fit)
				dot += fmt.Sprintf("\t\t\t<TR>\n\t\t\t\t<TD BGCOLOR=\"#E8DAEF\">part_start</TD>\n\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#E8DAEF\">%v</TD>\n\t\t\t</TR>\n", mbr.Partitions[i].Start)
				dot += fmt.Sprintf("\t\t\t<TR>\n\t\t\t\t<TD BGCOLOR=\"#FFFFFF\">part_size</TD>\n\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#FFFFFF\">%v</TD>\n\t\t\t</TR>\n", mbr.Partitions[i].Size)
				dot += fmt.Sprintf("\t\t\t<TR>\n\t\t\t\t<TD BGCOLOR=\"#E8DAEF\">part_name</TD>\n\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#E8DAEF\">%v</TD>\n\t\t\t</TR>\n", strings.TrimSpace(strings.ReplaceAll(mbr.Partitions[i].Name, "\x00", "")))
				if mbr.Partitions[i].Type == 'E' {
					iterEBR := r.getListEBR(mbr.Partitions[i].Start, mbr.Partitions[i].Size, file).GetIterable()
					for _, ebr := range iterEBR {
						if ebr.Size != 0 {
							dot += "\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR=\"#F08080\"><FONT COLOR=\"white\">Particion Logica</FONT></TD>\n\t\t\t\t\t\t\t<TD COLSPAN=\"2\" BGCOLOR=\"#F08080\"></TD>\n\t\t\t\t\t\t</TR>"
							dot += fmt.Sprintf("\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR=\"#FFFFFF\">part_status</TD>\n\t\t\t\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#FFFFFF\">%c</TD>\n\t\t\t\t\t\t</TR>", ebr.Mount)
							dot += fmt.Sprintf("\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR=\"#F5B7B1\">part_next</TD>\n\t\t\t\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#F5B7B1\">%v</TD>\n\t\t\t\t\t\t</TR>", ebr.Next)
							dot += fmt.Sprintf("\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR=\"#FFFFFF\">part_fit</TD>\n\t\t\t\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#FFFFFF\">%c</TD>\n\t\t\t\t\t\t</TR>", ebr.Fit)
							dot += fmt.Sprintf("\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR=\"#F5B7B1\">part_start</TD>\n\t\t\t\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#F5B7B1\">%v</TD>\n\t\t\t\t\t\t</TR>", ebr.Start)
							dot += fmt.Sprintf("\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR=\"#FFFFFF\">part_size</TD>\n\t\t\t\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#FFFFFF\">%v</TD>\n\t\t\t\t\t\t</TR>", ebr.Size)
							dot += fmt.Sprintf("\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR=\"#F5B7B1\">part_name</TD>\n\t\t\t\t\t\t\t<TD COLSPAN=\"1\" BGCOLOR=\"#F5B7B1\">%v</TD>\n\t\t\t\t\t\t</TR>", strings.TrimSpace(strings.ReplaceAll(ebr.Name, "\x00", "")))
						}
					}
				}
			}
		}
		dot += "\n\t\t\t\t\t</TABLE>\n\t\t\t\t</TD>\n\t\t\t</TR>"
		dot += "\n\t\t</TABLE>\n\t>];"
		dot += "\n}"
		r.generateFile(dot, driveletter)
	} else {
		r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
	}
}

func (r *Rep) reportDisk() {
	driveletter := string(r.Params["id"][0])
	if _, ok := env.Disks[driveletter]; ok {
		diskPath, _ := filepath.Abs(env.GetPath(driveletter))
		absolutePath, _ := filepath.Abs(diskPath)
		file, err := os.Open(absolutePath)
		if err != nil {
			if os.IsNotExist(err) {
				r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
				return
			}
			fmt.Println("Error al abrir el archivo:", err)
			return
		}
		defer file.Close()
		readedBytes := make([]byte, 153)
		_, err = file.Read(readedBytes)
		if err != nil {
			fmt.Println("Error al leer el archivo:", err)
			return
		}
		mbr := structs.DecodeMBR(readedBytes)
		lastNoEmptyByte := 152
		dotParts := ""
		occupiedCells := 10
		extendedParts := ""
		for i := 0; i < len(mbr.Partitions); i++ {
			if mbr.Partitions[i].Size != 0 {
				if mbr.Partitions[i].Start-lastNoEmptyByte > 1 {
					space := r.calculateSpace(float64(mbr.Partitions[i].Start), float64(lastNoEmptyByte), float64(mbr.Size))
					occupiedCells += int(space)
					dotParts += fmt.Sprintf("\n\t\t\t\t<TD COLSPAN=\"%v\" ROWSPAN=\"6\">Libre<BR/>%v %%</TD>", int(space), r.percentage(float64(mbr.Partitions[i].Start), float64(lastNoEmptyByte+1), float64(mbr.Size)))
				}
				space := r.calculateSpace(float64(mbr.Partitions[i].Size), 0, float64(mbr.Size))
				if mbr.Partitions[i].Type == 'P' {
					occupiedCells += int(space)
					dotParts += fmt.Sprintf("\n\t\t\t\t<TD COLSPAN=\"%v\" ROWSPAN=\"6\">%v<BR/>Primaria<BR/>%v %%</TD>", int(space), strings.TrimSpace(strings.ReplaceAll(mbr.Partitions[i].Name, "\x00", "")), r.percentage(float64(mbr.Partitions[i].Size), 0, float64(mbr.Size)))
				} else if mbr.Partitions[i].Type == 'E' {
					extendedParts = "\n\t\t\t<TR>"
					lastNoEmptyByteExt := mbr.Partitions[i].Start - 1
					occupiedExtend := 0
					iterEBR := r.getListEBR(mbr.Partitions[i].Start, mbr.Partitions[i].Size, file).GetIterable()
					if iterEBR[0].Size != 0 {
						for _, ebr := range iterEBR {
							if ebr.Start-lastNoEmptyByteExt > 1 {
								space = r.calculateSpace(float64(ebr.Start), float64(lastNoEmptyByteExt), float64(mbr.Size))
								occupiedExtend += int(space)
								extendedParts += fmt.Sprintf("\n\t\t\t\t<TD COLSPAN=\"%v\" ROWSPAN=\"5\">Libre<BR/>%v %%</TD>", int(space), r.percentage(float64(ebr.Start), float64(lastNoEmptyByteExt+1), float64(mbr.Size)))
							}
							extendedParts += "\n\t\t\t\t<TD COLSPAN=\"10\" ROWSPAN=\"5\">EBR</TD>"
							space = r.calculateSpace(float64(ebr.Size), 0, float64(mbr.Size))
							extendedParts += fmt.Sprintf("\n\t\t\t\t<TD COLSPAN=\"%v\" ROWSPAN=\"5\">%v<BR/>Logica<BR/>%v %%</TD>", int(space), strings.TrimSpace(strings.ReplaceAll(ebr.Name, "\x00", "")), r.percentage(float64(ebr.Size), 0, float64(mbr.Size)))
							lastNoEmptyByteExt = ebr.Start + ebr.Size - 1
							occupiedExtend += 10 + int(space)
						}
					} else if iterEBR[0].Next != -1 {
						occupiedExtend += 10
						extendedParts += "\n\t\t\t\t<TD COLSPAN=\"10\" ROWSPAN=\"5\">EBR</TD>"
						for e := 1; i < len(iterEBR); e++ {
							if iterEBR[i].Start-lastNoEmptyByteExt > 1 {
								space = r.calculateSpace(float64(iterEBR[e].Start), float64(lastNoEmptyByteExt)+1, float64(mbr.Size))
								occupiedExtend += int(space)
								extendedParts += fmt.Sprintf("\n\t\t\t\t<TD COLSPAN=\"%v\" ROWSPAN=\"5\">Libre<BR/>%v %%</TD>", int(space), r.percentage(float64(iterEBR[e].Start), float64(lastNoEmptyByteExt+1), float64(mbr.Size)))
							}
							extendedParts += "\n\t\t\t\t<TD COLSPAN=\"10\" ROWSPAN=\"5\">EBR</TD>"
							space = r.calculateSpace(float64(iterEBR[e].Size), 0, float64(mbr.Size))
							extendedParts += fmt.Sprintf("\n\t\t\t\t<TD COLSPAN=\"%v\" ROWSPAN=\"5\">%v<BR/>Logica<BR/>%v %%</TD>", int(space), strings.TrimSpace(strings.ReplaceAll(iterEBR[e].Name, "\x00", "")), r.percentage(float64(iterEBR[e].Size), 0, float64(mbr.Size)))
							lastNoEmptyByteExt = iterEBR[e].Start + iterEBR[e].Size - 1
							occupiedExtend += 10 + int(space)
						}
					} else {
						occupiedExtend += 10
						extendedParts += "\n\t\t\t\t<TD COLSPAN=\"10\" ROWSPAN=\"5\">EBR</TD>"
					}
					if mbr.Partitions[i].Start+mbr.Partitions[i].Size-lastNoEmptyByteExt > 1 {
						space = r.calculateSpace(float64(mbr.Partitions[i].Start)+float64(mbr.Partitions[i].Size), float64(lastNoEmptyByteExt)+1, float64(mbr.Size))
						occupiedExtend += int(space)
						extendedParts += fmt.Sprintf("\n\t\t\t\t<TD COLSPAN=\"%v\" ROWSPAN=\"5\">Libre<BR/>%v %%</TD>", int(space), r.percentage(float64(mbr.Partitions[i].Start+mbr.Partitions[i].Size), float64(lastNoEmptyByteExt+1), float64(mbr.Size)))
					}
					extendedParts += "\n\t\t\t</TR>"
					occupiedCells += occupiedExtend
					dotParts += fmt.Sprintf("\n\t\t\t\t<TD COLSPAN=\"%v\" ROWSPAN=\"1\">%v<BR/>Extendida</TD>", occupiedExtend, strings.TrimSpace(strings.ReplaceAll(mbr.Partitions[i].Name, "\x00", "")))
				}
				lastNoEmptyByte = mbr.Partitions[i].Start + mbr.Partitions[i].Size - 1
			}
		}
		if mbr.Size-lastNoEmptyByte > 1 {
			space := r.calculateSpace(float64(mbr.Size), float64(lastNoEmptyByte)+1, float64(mbr.Size))
			dotParts += fmt.Sprintf("\n\t\t\t\t<TD COLSPAN=\"%v\" ROWSPAN=\"6\">Libre<BR/>%v %%</TD>", int(space), r.percentage(float64(mbr.Size), float64(lastNoEmptyByte+1), float64(mbr.Size)))
			occupiedCells += int(space)
		}
		dot := "digraph Disk{\n\tnode [shape=plaintext];"
		dot += "\n\ttabla[label=<\n\t\t<TABLE BORDER=\"1\" CELLBORDER=\"1\" CELLSPACING=\"2\" CELLPADDING=\"4\">"
		dot += fmt.Sprintf("\n\t\t\t<TR>\n\t\t\t\t<TD COLSPAN=\"%v\">Disco: %v</TD>\n\t\t\t</TR>", occupiedCells, driveletter)
		dot += "\n\t\t\t<TR>\n\t\t\t\t<TD COLSPAN=\"10\" ROWSPAN=\"6\">MBR</TD>"
		dot += dotParts
		dot += "\n\t\t\t</TR>"
		dot += extendedParts
		dot += "\n\t\t</TABLE>\n\t>];"
		dot += "\n}"
		r.generateFile(dot, driveletter)
	} else {
		r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%v\" para reportar.", driveletter))
	}
}

func (r *Rep) reportInode() {
	driveletter := string(r.Params["id"][0])
	if disk, ok := env.Disks[driveletter]; ok {
		if partitionData, ok := disk.Ids[r.Params["id"]]; ok {
			namePartition := partitionData.Name
			diskPath, _ := filepath.Abs(env.GetPath(driveletter))
			absolutePath, _ := filepath.Abs(diskPath)
			file, err := os.Open(absolutePath)
			if err != nil {
				if os.IsNotExist(err) {
					r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
					return
				}
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()
			readedBytes := make([]byte, 153)
			_, err = file.Read(readedBytes)
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return
			}
			mbr := structs.DecodeMBR(readedBytes)
			for i := 0; i < len(mbr.Partitions); i++ {
				if mbr.Partitions[i].Size != 0 && utils.CompareStrings(mbr.Partitions[i].Name, namePartition) {
					_, err = file.Seek(int64(mbr.Partitions[i].Start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					superBlockData := make([]byte, structs.SizeOfSuperBlock())
					_, err = file.Read(superBlockData)
					if err != nil {
						fmt.Println("Error al leer el super block:", err)
						return
					}
					superBlock := structs.DecodeSuperBlock(superBlockData)
					_, err = file.Seek(int64(superBlock.Bm_inode_start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					bmInodes := make([]byte, superBlock.Inodes_count)
					_, err = file.Read(bmInodes)
					if err != nil {
						fmt.Println("Error al leer el Inode", err)
						return
					}
					dot := "digraph Inodes{\n\tnode [shape=box];\n\trankdir=LR;"
					for i := 0; i < len(bmInodes); i++ {
						if bmInodes[i] == '1' {
							_, err := file.Seek(int64(superBlock.Inode_start+i*structs.SizeOfInode()), 0)
							if err != nil {
								fmt.Println("Error al mover el puntero en el archivo:", err)
								return
							}
							readedBytes := make([]byte, structs.SizeOfInode())
							_, err = file.Read(readedBytes)
							if err != nil && err.Error() != "EOF" {
								fmt.Println("-> Error al leer el archivo:", err)
								return
							}
							inode := structs.DecodeInode(readedBytes)

							dot += fmt.Sprintf(`
	n%v[label = <<TABLE BORDER="0">
        <TR><TD colspan="2">Inodo %v</TD></TR>
        <TR><TD ALIGN="LEFT">uid:</TD><TD ALIGN="LEFT">%v</TD></TR>
        <TR><TD ALIGN="LEFT">gid:</TD><TD ALIGN="LEFT">%v</TD></TR>
        <TR><TD ALIGN="LEFT">size:</TD><TD ALIGN="LEFT">%v</TD></TR>
        <TR><TD ALIGN="LEFT">atime:</TD><TD ALIGN="LEFT">%v</TD></TR>
        <TR><TD ALIGN="LEFT">ctime:</TD><TD ALIGN="LEFT">%v</TD></TR>
        <TR><TD ALIGN="LEFT">mtime:</TD><TD ALIGN="LEFT">%v</TD></TR>`, i, i, inode.Uid, inode.Gid, inode.Size, inode.Atime, inode.Ctime, inode.Mtime)
							for r := 0; r < len(inode.Block); r++ {
								dot += fmt.Sprintf("\n\t\t<TR><TD ALIGN=\"LEFT\">apt%v:</TD><TD ALIGN=\"LEFT\">%v</TD></TR>", r+1, inode.Block[r])
							}
							dot += fmt.Sprintf(`
		<TR><TD ALIGN="LEFT">type:</TD><TD ALIGN="LEFT">%v</TD></TR>
		<TR><TD ALIGN="LEFT">perm:</TD><TD ALIGN="LEFT">%v</TD></TR>
	</TABLE>>];`, inode.Type, string(inode.Perm))
							if i > 0 {
								dot += fmt.Sprintf("\n\tn%v -> n%v;", i-1, i)
							}
						}
					}
					dot += "\n}"
					r.generateFile(dot, fmt.Sprintf("%v: %v", driveletter, r.Params["id"]))
					return
				}
			}
		} else {
			r.printError(fmt.Sprintf("-> Error rep: No existe el código de partición \"%s\" para reportar el disco \"%v\"", r.Params["id"], driveletter))
		}
	} else {
		r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%v\" para reportar", driveletter))
	}
}

func (r *Rep) reportBlock() {
	driveletter := string(r.Params["id"][0])
	if disk, ok := env.Disks[driveletter]; ok {
		if partitionData, ok := disk.Ids[r.Params["id"]]; ok {
			namePartition := partitionData.Name
			diskPath, _ := filepath.Abs(env.GetPath(driveletter))
			absolutePath, _ := filepath.Abs(diskPath)
			file, err := os.Open(absolutePath)
			if err != nil {
				if os.IsNotExist(err) {
					r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
					return
				}
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()
			readedBytes := make([]byte, 153)
			_, err = file.Read(readedBytes)
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return
			}
			mbr := structs.DecodeMBR(readedBytes)
			for i := 0; i < len(mbr.Partitions); i++ {
				if mbr.Partitions[i].Size != 0 && utils.CompareStrings(mbr.Partitions[i].Name, namePartition) {
					_, err = file.Seek(int64(mbr.Partitions[i].Start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					superBlockData := make([]byte, structs.SizeOfSuperBlock())
					_, err = file.Read(superBlockData)
					if err != nil {
						fmt.Println("Error al leer el super block:", err)
						return
					}
					superBlock := structs.DecodeSuperBlock(superBlockData)
					tree := structs.Tree{SuperBlock: superBlock, File: *file}
					blocks := tree.GetBlocks()
					dot := "digraph Blocks{\n\tnode [shape=box];\n\trankdir=LR;"
					for i := 0; i < len(blocks); i++ {
						if blocks[i].BlockFile != nil {
							dot += fmt.Sprintf("%v", blocks[i].BlockFile.GetDotB(int(blocks[i].I)))
						} else if blocks[i].BlockFolder != nil {
							dot += fmt.Sprintf("%v", blocks[i].BlockFolder.GetDotB(int(blocks[i].I)))
						} else if blocks[i].BlockPointers != nil {
							dot += fmt.Sprintf("%v", blocks[i].BlockPointers.GetDotB(int(blocks[i].I)))
						}
						if i > 0 {
							dot += fmt.Sprintf("\n\tn%v -> n%v;", blocks[i-1].I, blocks[i].I)
						}
					}
					dot += "\n}"
					r.generateFile(dot, fmt.Sprintf("%v: %v", namePartition, driveletter))
					return
				}
			}
		} else {
			r.printError(fmt.Sprintf("-> Error rep: No existe el código de partición \"%s\" para reportar el disco \"%v\"", r.Params["id"], driveletter))
		}
	} else {
		r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%v\" para reportar", driveletter))
	}
}

func (r *Rep) reportJournaling() {
	driveletter := string(r.Params["id"][0])
	if disk, ok := env.Disks[driveletter]; ok {
		if partitionData, ok := disk.Ids[r.Params["id"]]; ok {
			namePartition := partitionData.Name
			diskPath, _ := filepath.Abs(env.GetPath(driveletter))
			absolutePath, _ := filepath.Abs(diskPath)
			file, err := os.Open(absolutePath)
			if err != nil {
				if os.IsNotExist(err) {
					r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
					return
				}
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()
			readedBytes := make([]byte, 153)
			_, err = file.Read(readedBytes)
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return
			}
			mbr := structs.DecodeMBR(readedBytes)
			for i := 0; i < len(mbr.Partitions); i++ {
				if mbr.Partitions[i].Size != 0 && utils.CompareStrings(mbr.Partitions[i].Name, namePartition) {
					_, err = file.Seek(int64(mbr.Partitions[i].Start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					superBlockData := make([]byte, 68)
					_, err = file.Read(superBlockData)
					if err != nil {
						fmt.Println("Error al leer el super block:", err)
						return
					}
					superBlock := structs.DecodeSuperBlock(superBlockData)
					if superBlock.Filesystem_type == 3 {
						_, err = file.Seek(int64(mbr.Partitions[i].Start+68), 0)
						if err != nil {
							fmt.Println("Error al mover el puntero del archivo:", err)
							return
						}
						dot := "digraph Inodes{\n\tnode [shape=plaintext];\n\trankdir=LR;"
						dot += fmt.Sprintf("\n\tn%v[label = <<TABLE BORDER=\"1\" >", i)
						dot += fmt.Sprintf("\n\t\t<TR><TD COLSPAN=\"4\">%v: %v</TD></TR>", driveletter, namePartition)
						dot += "\n\t\t<TR><TD>Operacion</TD><TD>Path</TD><TD>Contenido</TD><TD>Fecha</TD></TR>"
						for r := 0; r < superBlock.Inodes_count; r++ {
							readedBytes := make([]byte, 212)
							_, err = file.Read(readedBytes)
							if err != nil {
								fmt.Println("Error al leer el archivo:", err)
								return
							}
							if !bytes.Equal(readedBytes, bytes.Repeat([]byte{0}, 212)) {
								dot += structs.DecodeJournal(readedBytes).GetDot()
							}
						}
						dot += "\n\t</TABLE>>];"
						dot += "\n}"
						r.generateFile(dot, fmt.Sprintf("%v: %v", namePartition, driveletter))
					} else {
						r.printError(fmt.Sprintf("-> Error rep: No puede generarse el reporte para \"%v\" en el disco \"%v\". Journaling no disponible en EXT2", r.Params["id"], driveletter))
					}
					return
				}
			}
		} else {
			r.printError(fmt.Sprintf("-> Error rep: No existe el código de partición \"%s\" para reportar el disco \"%v\"", r.Params["id"], driveletter))
		}
	} else {
		r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%v\" para reportar", driveletter))
	}
}

func (r *Rep) reportBMInode() {
	driveletter := string(r.Params["id"][0])
	if disk, ok := env.Disks[driveletter]; ok {
		if partitionData, ok := disk.Ids[r.Params["id"]]; ok {
			namePartition := partitionData.Name
			diskPath, _ := filepath.Abs(env.GetPath(driveletter))
			absolutePath, _ := filepath.Abs(diskPath)
			file, err := os.Open(absolutePath)
			if err != nil {
				if os.IsNotExist(err) {
					r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
					return
				}
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()
			readedBytes := make([]byte, 153)
			_, err = file.Read(readedBytes)
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return
			}
			mbr := structs.DecodeMBR(readedBytes)
			for i := 0; i < len(mbr.Partitions); i++ {
				if mbr.Partitions[i].Size != 0 && utils.CompareStrings(mbr.Partitions[i].Name, namePartition) {
					_, err = file.Seek(int64(mbr.Partitions[i].Start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					superBlockData := make([]byte, 68)
					_, err = file.Read(superBlockData)
					if err != nil {
						fmt.Println("Error al leer el super block:", err)
						return
					}
					superBlock := structs.DecodeSuperBlock(superBlockData)
					_, err = file.Seek(int64(superBlock.Bm_inode_start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					readedBytes := make([]byte, superBlock.Inodes_count)
					_, err = file.Read(readedBytes)
					if err != nil && err.Error() != "EOF" {
						fmt.Println("-> Error al leer el archivo:", err)
					}
					bmInodes := string(readedBytes)
					matriz := ""
					i := 0
					for i < len(bmInodes) {
						matriz += string(bmInodes[i]) + " "
						if (i+1)%20 == 0 {
							matriz += "\n"
						}
						i += 1
					}
					file, err := os.Create(r.Params["path"])
					if err != nil {
						fmt.Println("Error al crear al archivo:", err)
						return
					}
					defer file.Close()
					_, err = file.WriteString(matriz)
					if err != nil {
						fmt.Println("Error al escribir en el archivo:", err)
						return
					}
					r.printSuccess(strings.ToLower(r.Params["name"]), fmt.Sprintf("%v: %v", namePartition, driveletter))
					return
				}
			}
		} else {
			r.printError(fmt.Sprintf("-> Error rep: No existe el código de partición \"%s\" para reportar el disco \"%v\"", r.Params["id"], driveletter))
		}
	} else {
		r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%v\" para reportar", driveletter))
	}
}

func (r *Rep) reportBMBlock() {
	driveletter := string(r.Params["id"][0])
	if disk, ok := env.Disks[driveletter]; ok {
		if partitionData, ok := disk.Ids[r.Params["id"]]; ok {
			namePartition := partitionData.Name
			diskPath, _ := filepath.Abs(env.GetPath(driveletter))
			absolutePath, _ := filepath.Abs(diskPath)
			file, err := os.Open(absolutePath)
			if err != nil {
				if os.IsNotExist(err) {
					r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
					return
				}
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()
			readedBytes := make([]byte, 153)
			_, err = file.Read(readedBytes)
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return
			}
			mbr := structs.DecodeMBR(readedBytes)
			for i := 0; i < len(mbr.Partitions); i++ {
				if mbr.Partitions[i].Size != 0 && utils.CompareStrings(mbr.Partitions[i].Name, namePartition) {
					_, err = file.Seek(int64(mbr.Partitions[i].Start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					superBlockData := make([]byte, 68)
					_, err = file.Read(superBlockData)
					if err != nil {
						fmt.Println("Error al leer el super block:", err)
						return
					}
					superBlock := structs.DecodeSuperBlock(superBlockData)
					_, err = file.Seek(int64(superBlock.Bm_block_start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					readedBytes := make([]byte, superBlock.Blocks_count)
					_, err = file.Read(readedBytes)
					if err != nil && err.Error() != "EOF" {
						fmt.Println("-> Error al leer el archivo:", err)
					}
					bmBlocks := string(readedBytes)
					matriz := ""
					i := 0
					for i < len(bmBlocks) {
						matriz += string(bmBlocks[i]) + " "
						if (i+1)%20 == 0 {
							matriz += "\n"
						}
						i += 1
					}
					file, err := os.Create(r.Params["path"])
					if err != nil {
						fmt.Println("Error al crear al archivo:", err)
						return
					}
					defer file.Close()
					_, err = file.WriteString(matriz)
					if err != nil {
						fmt.Println("Error al escribir en el archivo:", err)
						return
					}
					r.printSuccess(strings.ToLower(r.Params["name"]), fmt.Sprintf("%v: %v", namePartition, driveletter))
					return
				}
			}
		} else {
			r.printError(fmt.Sprintf("-> Error rep: No existe el código de partición \"%s\" para reportar el disco \"%v\"", r.Params["id"], driveletter))
		}
	} else {
		r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%v\" para reportar", driveletter))
	}
}

func (r *Rep) reportTree() {
	driveletter := string(r.Params["id"][0])
	if disk, ok := env.Disks[driveletter]; ok {
		if partitionData, ok := disk.Ids[r.Params["id"]]; ok {
			namePartition := partitionData.Name
			diskPath, _ := filepath.Abs(env.GetPath(driveletter))
			absolutePath, _ := filepath.Abs(diskPath)
			file, err := os.Open(absolutePath)
			if err != nil {
				if os.IsNotExist(err) {
					r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
					return
				}
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()
			readedBytes := make([]byte, 153)
			_, err = file.Read(readedBytes)
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return
			}
			mbr := structs.DecodeMBR(readedBytes)
			for i := 0; i < len(mbr.Partitions); i++ {
				if mbr.Partitions[i].Size != 0 && utils.CompareStrings(mbr.Partitions[i].Name, namePartition) {
					_, err = file.Seek(int64(mbr.Partitions[i].Start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					superBlockData := make([]byte, 68)
					_, err = file.Read(superBlockData)
					if err != nil {
						fmt.Println("Error al leer el super block:", err)
						return
					}
					superBlock := structs.DecodeSuperBlock(superBlockData)
					tree := structs.Tree{SuperBlock: superBlock, File: *file}
					r.generateFile(tree.GetDot(driveletter, namePartition), fmt.Sprintf("%v: %v", namePartition, driveletter))
					return
				}
			}
		} else {
			r.printError(fmt.Sprintf("-> Error rep: No existe el código de partición \"%s\" para reportar el disco \"%v\"", r.Params["id"], driveletter))
		}
	} else {
		r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%v\" para reportar", driveletter))
	}
}

func (r *Rep) reportSb() {
	driveletter := string(r.Params["id"][0])
	if disk, ok := env.Disks[driveletter]; ok {
		if partitionData, ok := disk.Ids[r.Params["id"]]; ok {
			namePartition := partitionData.Name
			diskPath, _ := filepath.Abs(env.GetPath(driveletter))
			absolutePath, _ := filepath.Abs(diskPath)
			file, err := os.Open(absolutePath)
			if err != nil {
				if os.IsNotExist(err) {
					r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
					return
				}
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()
			readedBytes := make([]byte, 153)
			_, err = file.Read(readedBytes)
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return
			}
			mbr := structs.DecodeMBR(readedBytes)
			for i := 0; i < len(mbr.Partitions); i++ {
				if mbr.Partitions[i].Size != 0 && utils.CompareStrings(mbr.Partitions[i].Name, namePartition) {
					_, err = file.Seek(int64(mbr.Partitions[i].Start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					superBlockData := make([]byte, 68)
					_, err = file.Read(superBlockData)
					if err != nil {
						fmt.Println("Error al leer el super block:", err)
						return
					}
					superBlock := structs.DecodeSuperBlock(superBlockData)
					dot := "digraph SuperBlock{\n\tnode [shape=plaintext];"
					dot += "\n\ttabla[label=<\n\t\t<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\">"
					dot += "\n\t\t\t<TR>\n\t\t\t\t<TD BORDER=\"1\">\n\t\t\t\t\t<TABLE BORDER=\"1\" CELLBORDER=\"0\" CELLSPACING=\"0\" CELLPADDING=\"4\">"
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD COLSPAN="2" BGCOLOR="#145A32"><FONT COLOR="white">%v</FONT></TD>\n\t\t\t\t\t\t</TR>`, driveletter)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#145A32"><FONT COLOR="white">SuperBlock</FONT></TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#145A32"><FONT COLOR="white">%v</FONT></TD>\n\t\t\t\t\t\t</TR>`, namePartition)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#FFFFFF">filesystem_type</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#FFFFFF">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Filesystem_type)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#27AE60">inodes_count</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#27AE60">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Inodes_count)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#FFFFFF">blocks_count</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#FFFFFF">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Blocks_count)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#27AE60">free_inodes_count</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#27AE60">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Free_inodes_count)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#FFFFFF">free_blocks_count</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#FFFFFF">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Free_blocks_count)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#27AE60">mtime</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#27AE60">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Mtime)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#FFFFFF">umtime</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#FFFFFF">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Umtime)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#27AE60">mnt_count</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#27AE60">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Mnt_count)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#FFFFFF">magic</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#FFFFFF">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Magic)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#27AE60">inode_size</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#27AE60">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Inode_size)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#FFFFFF">block_size</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#FFFFFF">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Block_size)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#27AE60">first_ino</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#27AE60">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.First_ino)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#FFFFFF">first_blo</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#FFFFFF">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.First_blo)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#27AE60">bm_inode_start</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#27AE60">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Bm_inode_start)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#FFFFFF">bm_block_start</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#FFFFFF">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Bm_block_start)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#27AE60">inode_start</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#27AE60">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Inode_start)
					dot += fmt.Sprintf(`\n\t\t\t\t\t\t<TR>\n\t\t\t\t\t\t\t<TD BGCOLOR="#FFFFFF">block_start</TD>\n\t\t\t\t\t\t\t<TD COLSPAN="1" BGCOLOR="#FFFFFF">%v</TD>\n\t\t\t\t\t\t</TR>`, superBlock.Block_start)
					dot += "\n\t\t\t\t\t</TABLE>\n\t\t\t\t</TD>\n\t\t\t</TR>"
					dot += "\n\t\t</TABLE>\n\t>];"
					dot += "\n}"
					r.generateFile(dot, fmt.Sprintf("%v: %v", namePartition, driveletter))
					return
				}
			}
		} else {
			r.printError(fmt.Sprintf("-> Error rep: No existe el código de partición \"%s\" para reportar el disco \"%v\"", r.Params["id"], driveletter))
		}
	} else {
		r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%v\" para reportar", driveletter))
	}
}

func (r *Rep) reportFile() {
	driveletter := string(r.Params["id"][0])
	if disk, ok := env.Disks[driveletter]; ok {
		if partitionData, ok := disk.Ids[r.Params["id"]]; ok {
			namePartition := partitionData.Name
			diskPath, _ := filepath.Abs(env.GetPath(driveletter))
			absolutePath, _ := filepath.Abs(diskPath)
			file, err := os.Open(absolutePath)
			if err != nil {
				if os.IsNotExist(err) {
					r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%s\" para reportar.", driveletter))
					return
				}
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()
			readedBytes := make([]byte, 153)
			_, err = file.Read(readedBytes)
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return
			}
			mbr := structs.DecodeMBR(readedBytes)
			for i := 0; i < len(mbr.Partitions); i++ {
				if mbr.Partitions[i].Size != 0 && utils.CompareStrings(mbr.Partitions[i].Name, namePartition) {
					_, err = file.Seek(int64(mbr.Partitions[i].Start), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					superBlockData := make([]byte, 68)
					_, err = file.Read(superBlockData)
					if err != nil {
						fmt.Println("Error al leer el super block:", err)
						return
					}
					// superBlock := structs.DecodeSuperBlock(superBlockData)
					// tree := structs.Tree{SuperBlock: superBlock, File: *file}
				}
			}
		} else {
			r.printError(fmt.Sprintf("-> Error rep: No existe el código de partición \"%s\" para reportar el disco \"%v\"", r.Params["id"], driveletter))
		}
	} else {
		r.printError(fmt.Sprintf("-> Error rep: No existe el disco \"%v\" para reportar", driveletter))
	}
}

func (r *Rep) reportLs() {}

func (r *Rep) getListEBR(start, size int, file *os.File) *structs.ListEBR {
	listEBR := structs.NewListEBR(start, size)
	_, err := file.Seek(int64(start), 0)
	if err != nil {
		return nil
	}

	readedBytes := make([]byte, 30)
	_, err = file.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
		return listEBR
	}

	ebr := structs.DecodeEBR(readedBytes)
	listEBR.Insert(ebr)
	for ebr.Next != -1 {
		file.Seek(int64(ebr.Next), 0)
		readedBytes := make([]byte, 30)
		_, err = file.Read(readedBytes)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("-> Error al leer el archivo:", err)
			return listEBR
		}
		ebr = structs.DecodeEBR(readedBytes)
		listEBR.Insert(ebr)
	}
	return listEBR
}

func (r *Rep) generateFile(dot, diskname string) {
	absolutePath := filepath.Join(r.Params["path"])
	destdir := filepath.Dir(absolutePath)
	extension := filepath.Ext(absolutePath)
	absolutePathDot := strings.ReplaceAll(absolutePath, extension, ".dot")
	if _, err := os.Stat(destdir); os.IsNotExist(err) {
		if err := os.MkdirAll(destdir, os.ModePerm); err != nil {
			fmt.Println("Error al crear el directorio:", err)
			return
		}
	}
	file, err := os.Create(absolutePathDot)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(dot)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
	file.Close()
	cmd := exec.Command("dot", "-T"+extension[1:], absolutePathDot, "-o", absolutePath)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al ejecutar el comando dot:", err)
		return
	}
	err = os.Remove(strings.Replace(absolutePath, extension, ".dot", 1))
	if err != nil {
		fmt.Println("Error al eliminar el archivo .dot:", err)
		return
	}
	r.printSuccess(strings.ToLower(r.Params["name"]), diskname)
}

func (r *Rep) percentage(start, firstEmptyByte, size float64) float64 {
	number := math.Round(((start - firstEmptyByte) / size) * 100)
	num := number - math.Floor(number)
	if math.Round(num*100)/100 > 0 {
		return number
	}
	return math.Floor(number)
}

func (r *Rep) calculateSpace(start, firstEmptyByte, size float64) float64 {
	num := math.Round(((start - firstEmptyByte) / size) * 200)
	if num >= 1 {
		return num
	}
	return 1
}

func (r *Rep) printError(text string) {
	fmt.Printf("\033[31m%v [%v:%v]\033[0m\n", text, r.Line, r.Column)
}

func (r *Rep) printSuccess(type_, diskName string) {
	fmt.Printf("\033[35m-> rep: Reporte generado exitosamente. \"%s\" %s [%d:%d]\033[0m\n", type_, diskName, r.Line, r.Column)
}

func (r *Rep) GetResult() string { return "" }
