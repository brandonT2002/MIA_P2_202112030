package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

type MBR struct {
	Size       int          // 4 bytes
	Date       time.Time    // 4 bytes
	Dsk        int          // 4 bytes
	Fit        rune         // 1 byte
	Partitions []*Partition // 140 bytes
} // 153 bytes

func NewMBR(size int, fit rune) *MBR {
	rand.Seed(time.Now().UnixNano())
	return &MBR{size, time.Now(), rand.Intn(10001), fit, []*Partition{{Status: '-'}, {Status: '-'}, {Status: '-'}, {Status: '-'}}}
}

func (m *MBR) Encode() []byte {
	var buf bytes.Buffer
	// Size
	sizeBytes := make([]byte, 4)
	if m.Size != 0 {
		binary.BigEndian.PutUint32(sizeBytes, uint32(m.Size))
	}
	buf.Write(sizeBytes)
	// Date
	dateBytes := make([]byte, 4)
	if !m.Date.IsZero() {
		binary.BigEndian.PutUint32(dateBytes, uint32(m.Date.Unix()))
	}
	buf.Write(dateBytes)
	// Dsk
	dskBytes := make([]byte, 4)
	if m.Dsk != 0 {
		binary.BigEndian.PutUint32(dskBytes, uint32(m.Dsk))
	}
	buf.Write(dskBytes)
	// Fit
	buf.WriteByte(byte(m.Fit))
	// Parts
	for _, p := range m.Partitions {
		if p != nil {
			buf.Write(p.Encode(true))
		} else {
			buf.Write(p.Encode(false))
		}
	}
	return buf.Bytes()
}

func DecodeMBR(data []byte) *MBR {
	// Size
	size := int(binary.BigEndian.Uint32(data[:4]))
	// Date
	unixTime := binary.BigEndian.Uint32(data[4:8])
	date := time.Unix(int64(unixTime), 0)
	// Dsk
	dsk := int(binary.BigEndian.Uint32(data[8:12]))
	// Fit
	fit, _ := utf8.DecodeRune(data[12:13])
	// Partitions
	partitions := []*Partition{
		DecodePartition(data[13:48]),
		DecodePartition(data[48:83]),
		DecodePartition(data[83:118]),
		DecodePartition(data[118:]),
	}
	return &MBR{size, date, dsk, fit, partitions}
}

func (m *MBR) ToString() string {
	prts := ""
	for _, p := range m.Partitions {
		prts += p.ToString()
	}
	return fmt.Sprintf("Size: %-10d Date: %-20v DSK: %v", m.Size, m.Date.Format("02/01/2006 15:04:05"), m.Dsk) + prts
}
