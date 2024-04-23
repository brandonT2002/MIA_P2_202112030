package structs

import (
	"bytes"
	"fmt"
	"strings"
)

type BlockFile struct {
	Content string // 64 bytes
}

func (b *BlockFile) Encode() []byte {
	var buf bytes.Buffer
	// Content
	if b.Content != "" {
		if len(b.Content) < 64 {
			buf.Write([]byte(b.Content)[:len(b.Content)])
			buf.Write(bytes.Repeat([]byte{0x00}, 64-len(b.Content)))
		} else {
			buf.Write([]byte(b.Content)[:64])
		}
	} else {
		buf.Write(make([]byte, 64))
	}
	return buf.Bytes()
}

func DecodeBlockFile(data []byte) *BlockFile {
	return &BlockFile{string(data[:64])}
}

func (b *BlockFile) GetDot(i int) string {
	content := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(b.Content, "\n", "\\n"), "\"", "\\\""), "'", "\\'")

	return fmt.Sprintf(`block%v[label=<
	<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0" CELLPADDING="4">
		<TR><TD BGCOLOR="#FFECA9" PORT="B%v">Bloque %v</TD></TR>
		<TR><TD>%v</TD></TR>
	</TABLE>
>];`, i, i, i, content)
}

func (b *BlockFile) GetDotB(i int) string {
	content := ""
	for r, c := range []rune(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(b.Content, "\n", "\\n"), "\"", "\\\""), "'", "\\'")) {
		content += string(c)
		if r%8 == 7 {
			content += "<BR/>"
		}
	}
	return fmt.Sprintf(`
		n%v[label = <<TABLE BORDER="0">
		<TR><TD>Bloque Archivo %v</TD></TR>
		<TR><TD><FONT FACE="Consolas">%v</FONT></TD></TR>
	</TABLE>>];`, i, i, content)
}

func (b *BlockFile) ToString() string {
	return fmt.Sprintf("content: %s\n", b.Content)
}
