package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	utils "mia/Classes/Utils"
	"unicode/utf8"
)

type Partition struct {
	Status      rune   // 1 byte
	Type        rune   // 1 byte
	Fit         rune   // 1 byte
	Start       int    // 4 bytes
	Size        int    // 4 bytes
	Name        string // 16 bytes
	Correlative int    // 4 bytes
	Id          string // 4 bytes
} // 35 bytes

func NewPartition(status, type_, fit rune, start, size int, name string, correlative int) *Partition {
	return &Partition{Status: status, Type: type_, Fit: fit, Start: start, Size: size, Name: name, Correlative: correlative}
}

func (p *Partition) Encode(flag bool) []byte {
	if flag {
		var buf bytes.Buffer
		// Status
		buf.WriteByte(byte(p.Status))
		// Type
		buf.WriteByte(byte(p.Type))
		// Fit
		buf.WriteByte(byte(p.Fit))
		// Start
		startBytes := make([]byte, 4)
		if p.Start != 0 {
			binary.BigEndian.PutUint32(startBytes, uint32(p.Start))
		}
		buf.Write(startBytes)
		// Size
		sizeBytes := make([]byte, 4)
		if p.Size != 0 {
			binary.BigEndian.PutUint32(sizeBytes, uint32(p.Size))
		}
		buf.Write(sizeBytes)
		// Name
		if p.Name != "" {
			buf.Write([]byte(p.Name)[:16])
		} else {
			buf.Write(make([]byte, 16))
		}
		// Correlative
		correlativeBytes := make([]byte, 4)
		if p.Correlative != 0 {
			binary.BigEndian.PutUint32(correlativeBytes, uint32(p.Correlative))
		}
		buf.Write(correlativeBytes)
		// Id
		if p.Id != "" {
			buf.Write([]byte(p.Id)[:4])
		} else {
			buf.Write(make([]byte, 4))
		}
		return buf.Bytes()
	}
	return bytes.Repeat([]byte{0x00}, 35)
}

func DecodePartition(data []byte) *Partition {
	// Status
	status, _ := utf8.DecodeRune(data[:1])
	// Type
	type_, _ := utf8.DecodeRune(data[1:2])
	// Fit
	fit, _ := utf8.DecodeRune(data[2:3])
	// Start
	start := int(binary.BigEndian.Uint32(data[3:7]))
	// Size
	size := int(binary.BigEndian.Uint32(data[7:11]))
	// Name
	name := string(data[11:27])
	// Correlative
	correlative := int(binary.BigEndian.Uint32(data[27:31]))
	// Id
	id := string(data[31:])
	return &Partition{status, type_, fit, start, size, utils.ClearString(name), correlative, utils.ClearString(id)}
}

func (p *Partition) ToString() string {
	return fmt.Sprintf("\n\tStart: %-10d Status: %c Size: %-10d Name: %s", p.Start, p.Status, p.Size, p.Name)
}
