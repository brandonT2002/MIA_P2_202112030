package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type BlockPointers struct {
	Pointers []int32 // 64 bytes
}

func NewBlockPointers() *BlockPointers {
	return &BlockPointers{
		Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	}
}

func (b *BlockPointers) Encode() []byte {
	var buf bytes.Buffer
	// Pointers
	for _, pointer := range b.Pointers {
		pointerBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(pointerBytes, uint32(pointer))
		buf.Write(pointerBytes)
	}
	return buf.Bytes()
}

func DecodeBlockPointers(data []byte) *BlockPointers {
	// Pointers
	pointers := []int32{}
	for i := 0; i < 16; i++ {
		pointers = append(pointers, int32(binary.BigEndian.Uint32(data[i*4:4+i*4])))
	}
	return &BlockPointers{pointers}
}

func (b *BlockPointers) GetDot(i int) string {
	pointers := ""
	for p := 0; p < len(b.Pointers); p++ {
		pointers += fmt.Sprintf(`
		<TR><TD PORT="A%v">%v</TD></TR>`, p, b.Pointers[p])
	}
	return fmt.Sprintf(`block%v[label=<
	<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0" CELLPADDING="4">
		<TR><TD BGCOLOR="#FFCB97" PORT="B%v">Bloque %v</TD></TR>%v
	</TABLE>
>];`, i, i, i, pointers)
}

func (b *BlockPointers) GetDotB(i int) string {
	content := ""
	for r := 0; r < len(b.Pointers); r++ {
		content += string(b.Pointers[r])
		if r < len(b.Pointers)-1 {
			content += ","
		}
		if r%4 == 3 {
			content += "<BR/>"
		} else {
			content += " "
		}
	}
	return fmt.Sprintf(`
	n%v[label = <<TABLE BORDER="0">
		<TR><TD>Bloque Apuntadores %v</TD></TR>
		<TR><TD><FONT FACE="Consolas">%v</FONT></TD></TR>
	</TABLE>>];`, i, i, content)
}

func (b *BlockPointers) ToString() string {
	return fmt.Sprintf("pointers: %v\n", b.Pointers)
}
