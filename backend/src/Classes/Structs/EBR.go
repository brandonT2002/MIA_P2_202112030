package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	utils "mia/Classes/Utils"
	"unicode/utf8"
)

type EBR struct {
	Mount rune   // 1 byte
	Fit   rune   // 1 byte
	Start int    // 4 bytes
	Size  int    // 4 bytes
	Next  int    // 4 bytes
	Name  string // 16 bytes
} // 30 bytes

func NewEBR(fit rune, start, size int, name string) *EBR {
	return &EBR{
		Mount: '0',
		Fit:   fit,
		Start: start,
		Size:  size,
		Next:  -1,
		Name:  name,
	}
}

func (e *EBR) Encode() []byte {
	var buf bytes.Buffer
	// Mount
	buf.WriteByte(byte(e.Mount))
	// Fit
	buf.WriteByte(byte(e.Fit))
	// Start
	startBytes := make([]byte, 4)
	if e.Start != 0 {
		binary.BigEndian.PutUint32(startBytes, uint32(e.Start))
	}
	buf.Write(startBytes)
	// Size
	sizeBytes := make([]byte, 4)
	if e.Size != 0 {
		binary.BigEndian.PutUint32(sizeBytes, uint32(e.Size))
	}
	buf.Write(sizeBytes)
	// Next
	nextBytes := make([]byte, 4)
	if e.Next != 0 {
		binary.BigEndian.PutUint32(nextBytes, uint32(e.Next))
	}
	buf.Write(nextBytes)
	// Name
	if e.Name != "" {
		buf.Write([]byte(e.Name)[:16])
	} else {
		buf.Write(make([]byte, 16))
	}
	return buf.Bytes()
}

func DecodeEBR(data []byte) *EBR {
	// Mount
	mount, _ := utf8.DecodeRune(data[:1])
	// Fit
	fit, _ := utf8.DecodeRune(data[1:2])
	// Start
	start := int(binary.BigEndian.Uint32(data[2:6]))
	// Size
	size := int(binary.BigEndian.Uint32(data[6:10]))
	// Next
	next := int(binary.BigEndian.Uint32(data[10:14]))
	if next == 0 {
		next = -1
	}
	// Name
	name := string(data[14:])
	return &EBR{mount, fit, start, size, next, utils.ClearString(name)}
}

func (e *EBR) ToString() string {
	return fmt.Sprintf("Mount: %-5v Start: %-10d Size: %-10d Name: %-16s Next: %d", e.Mount, e.Start, e.Size, e.Name, e.Next)
}
