package structs

import (
	"bytes"
	"fmt"
	"strings"
)

type BlockFolder struct {
	Content []*Content // 64 bytes
}

func NewBlockFolder() *BlockFolder {
	return &BlockFolder{Content: []*Content{{Inode: -1}, {Inode: -1}, {Inode: -1}, {Inode: -1}}}
}

func (b *BlockFolder) Encode() []byte {
	var buf bytes.Buffer
	// Content
	for _, c := range b.Content {
		buf.Write(c.Encode())
	}
	return buf.Bytes()
}

func DecodeBlockFolder(data []byte) *BlockFolder {
	content := []*Content{}
	for i := 0; i < 4; i++ {
		content = append(content, DecodeContent(data[i*16:16+i*16]))
	}
	return &BlockFolder{content}
}

func (b *BlockFolder) GetDot(i int) string {
	pointers := ""
	for p := 0; p < len(b.Content); p++ {
		pointers += fmt.Sprintf(`
		<TR><TD>%v</TD><TD PORT="A%v">%v</TD></TR>`, strings.TrimSpace(strings.ReplaceAll(b.Content[p].Name, "\x00", "")), p, b.Content[p].Inode)
	}
	return fmt.Sprintf(`block%v[label=<
	<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0" CELLPADDING="4">
		<TR><TD COLSPAN="2" BGCOLOR="#D1BCD2" PORT="B%v">Bloque %v</TD></TR>%v
	</TABLE>
>];`, i, i, i, pointers)
}

func (b *BlockFolder) GetDotB(i int) string {
	pointers := ""
	for p := 0; p < len(b.Content); p++ {
		pointers += fmt.Sprintf("\n\t\t<TR><TD ALIGN=\"LEFT\">%v</TD><TD ALIGN=\"LEFT\">%v</TD></TR>", strings.TrimSpace(strings.ReplaceAll(b.Content[p].Name, "\x00", "")), b.Content[p].Inode)
	}
	return fmt.Sprintf(`
		n%v[label = <<TABLE BORDER="0">
		<TR><TD COLSPAN="2">Bloque Carpeta %v</TD></TR>%v
	</TABLE>>];`, i, i, pointers)
}

func (b *BlockFolder) ToString() string {
	contents := ""
	for _, c := range b.Content {
		if contents != "" {
			contents += "\n"
		}
		contents += c.ToString()
	}
	return contents
}
