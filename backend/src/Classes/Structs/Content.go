package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Content struct {
	Name  string // 12 bytes
	Inode int32  // 4 bytes
} // 16 bytes

func (c *Content) Encode() []byte {
	var buf bytes.Buffer
	// Name
	if c.Name != "" {
		buf.Write([]byte(c.Name)[:12])
	} else {
		buf.Write(make([]byte, 12))
	}
	// Inode
	inodeBytes := make([]byte, 4)
	if c.Inode != 0 {
		binary.BigEndian.PutUint32(inodeBytes, uint32(c.Inode))
	}
	buf.Write(inodeBytes)
	return buf.Bytes()
}

func DecodeContent(data []byte) *Content {
	// Name
	name := string(data[:12])
	// Inode
	inode := int32(binary.BigEndian.Uint32(data[12:16]))
	return &Content{name, inode}
}

func (c *Content) ToString() string {
	return fmt.Sprintf("name: %s inode: %d", c.Name, c.Inode)
}
