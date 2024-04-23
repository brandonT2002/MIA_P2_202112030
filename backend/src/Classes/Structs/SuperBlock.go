package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

type SuperBlock struct {
	Filesystem_type   int       // 4 bytes
	Inodes_count      int       // 4 bytes
	Blocks_count      int       // 4 bytes
	Free_inodes_count int       // 4 bytes
	Free_blocks_count int       // 4 bytes
	Mtime             time.Time // 4 bytes
	Umtime            time.Time // 4 bytes
	Mnt_count         int       // 4 bytes
	Magic             int       // 4 bytes
	Inode_size        int       // 4 bytes
	Block_size        int       // 4 bytes
	First_ino         int       // 4 bytes
	First_blo         int       // 4 bytes
	Bm_inode_start    int       // 4 bytes
	Bm_block_start    int       // 4 bytes
	Inode_start       int       // 4 bytes
	Block_start       int       // 4 bytes
} // 68 bytes

func (s *SuperBlock) Encode() []byte {
	var buf bytes.Buffer
	// Filesystem_type
	filesystem_typeBytes := make([]byte, 4)
	if s.Filesystem_type != 0 {
		binary.BigEndian.PutUint32(filesystem_typeBytes, uint32(s.Filesystem_type))
	}
	buf.Write(filesystem_typeBytes)
	// Inodes_count
	inodes_countBytes := make([]byte, 4)
	if s.Inodes_count != 0 {
		binary.BigEndian.PutUint32(inodes_countBytes, uint32(s.Inodes_count))
	}
	buf.Write(inodes_countBytes)
	// Blocks_count
	blocks_countBytes := make([]byte, 4)
	if s.Blocks_count != 0 {
		binary.BigEndian.PutUint32(blocks_countBytes, uint32(s.Blocks_count))
	}
	buf.Write(blocks_countBytes)
	// Free_inodes_count
	free_inodes_countBytes := make([]byte, 4)
	if s.Free_inodes_count != 0 {
		binary.BigEndian.PutUint32(free_inodes_countBytes, uint32(s.Free_inodes_count))
	}
	buf.Write(free_inodes_countBytes)
	// Free_blocks_count
	free_blocks_countBytes := make([]byte, 4)
	if s.Free_blocks_count != 0 {
		binary.BigEndian.PutUint32(free_blocks_countBytes, uint32(s.Free_blocks_count))
	}
	buf.Write(free_blocks_countBytes)
	// Mtime
	mtimeBytes := make([]byte, 4)
	if !s.Mtime.IsZero() {
		binary.BigEndian.PutUint32(mtimeBytes, uint32(s.Mtime.Unix()))
	}
	buf.Write(mtimeBytes)
	// Umtime
	umtimeBytes := make([]byte, 4)
	if !s.Umtime.IsZero() {
		binary.BigEndian.PutUint32(umtimeBytes, uint32(s.Umtime.Unix()))
	}
	buf.Write(umtimeBytes)
	// Mnt_count
	mnt_countBytes := make([]byte, 4)
	if s.Mnt_count != 0 {
		binary.BigEndian.PutUint32(mnt_countBytes, uint32(s.Mnt_count))
	}
	buf.Write(mnt_countBytes)
	// Magic
	magicBytes := make([]byte, 4)
	if s.Magic != 0 {
		binary.BigEndian.PutUint32(magicBytes, uint32(s.Magic))
	}
	buf.Write(magicBytes)
	// Inode_size
	inode_sizeBytes := make([]byte, 4)
	if s.Inode_size != 0 {
		binary.BigEndian.PutUint32(inode_sizeBytes, uint32(s.Inode_size))
	}
	buf.Write(inode_sizeBytes)
	// Block_size
	block_sizeBytes := make([]byte, 4)
	if s.Block_size != 0 {
		binary.BigEndian.PutUint32(block_sizeBytes, uint32(s.Block_size))
	}
	buf.Write(block_sizeBytes)
	// First_ino
	first_inoBytes := make([]byte, 4)
	if s.First_ino != 0 {
		binary.BigEndian.PutUint32(first_inoBytes, uint32(s.First_ino))
	}
	buf.Write(first_inoBytes)
	// First_blo
	first_bloBytes := make([]byte, 4)
	if s.First_blo != 0 {
		binary.BigEndian.PutUint32(first_bloBytes, uint32(s.First_blo))
	}
	buf.Write(first_bloBytes)
	// Bm_inode_start
	bm_inode_startBytes := make([]byte, 4)
	if s.Bm_inode_start != 0 {
		binary.BigEndian.PutUint32(bm_inode_startBytes, uint32(s.Bm_inode_start))
	}
	buf.Write(bm_inode_startBytes)
	// Bm_block_start
	bm_block_startBytes := make([]byte, 4)
	if s.Bm_block_start != 0 {
		binary.BigEndian.PutUint32(bm_block_startBytes, uint32(s.Bm_block_start))
	}
	buf.Write(bm_block_startBytes)
	// Inode_start
	inode_startBytes := make([]byte, 4)
	if s.Inode_start != 0 {
		binary.BigEndian.PutUint32(inode_startBytes, uint32(s.Inode_start))
	}
	buf.Write(inode_startBytes)
	// Block_start
	block_startBytes := make([]byte, 4)
	if s.Block_start != 0 {
		binary.BigEndian.PutUint32(block_startBytes, uint32(s.Block_start))
	}
	buf.Write(block_startBytes)
	return buf.Bytes()
}

func DecodeSuperBlock(data []byte) *SuperBlock {
	// Filesystem_type
	filesystem_type := int(binary.BigEndian.Uint32(data[:4]))
	// Inodes_count
	inodes_count := int(binary.BigEndian.Uint32(data[4:8]))
	// Blocks_count
	blocks_count := int(binary.BigEndian.Uint32(data[8:12]))
	// Free_inodes_count
	free_inodes_count := int(binary.BigEndian.Uint32(data[12:16]))
	// Free_blocks_count
	free_blocks_count := int(binary.BigEndian.Uint32(data[16:20]))
	// Mtime
	unixTime := binary.BigEndian.Uint32(data[20:24])
	mtime := time.Unix(int64(unixTime), 0)
	// Umtime
	unixTime = binary.BigEndian.Uint32(data[24:28])
	umtime := time.Unix(int64(unixTime), 0)
	// Mnt_count
	mnt_count := int(binary.BigEndian.Uint32(data[28:32]))
	// Magic
	magic := int(binary.BigEndian.Uint32(data[32:36]))
	// Inode_size
	inode_size := int(binary.BigEndian.Uint32(data[36:40]))
	// Block_size
	block_size := int(binary.BigEndian.Uint32(data[40:44]))
	// First_ino
	first_ino := int(binary.BigEndian.Uint32(data[44:48]))
	// First_blo
	first_blo := int(binary.BigEndian.Uint32(data[48:52]))
	// Bm_inode_start
	bm_inode_start := int(binary.BigEndian.Uint32(data[52:56]))
	// Bm_block_start
	bm_block_start := int(binary.BigEndian.Uint32(data[56:60]))
	// Inode_start
	inode_start := int(binary.BigEndian.Uint32(data[60:64]))
	// Block_start
	block_start := int(binary.BigEndian.Uint32(data[64:]))
	return &SuperBlock{
		filesystem_type,
		inodes_count,
		blocks_count,
		free_inodes_count,
		free_blocks_count,
		mtime,
		umtime,
		mnt_count,
		magic,
		inode_size,
		block_size,
		first_ino,
		first_blo,
		bm_inode_start,
		bm_block_start,
		inode_start,
		block_start,
	}
}

func (s *SuperBlock) ToString() string {
	return fmt.Sprintf(`filesystem_type %v
inodes_count %v
blocks_count %v
free_inodes_count %v
free_blocks_count %v
mtime %v
umtime %v
mnt_count %v
magic %v
inode_size %v
block_size %v
first_ino %v
first_blo %v
bm_inode_start %v
bm_block_start %v
inode_start %v
block_start %v`,
		s.Filesystem_type,
		s.Inodes_count,
		s.Blocks_count,
		s.Free_inodes_count,
		s.Free_blocks_count,
		s.Mtime,
		s.Umtime,
		s.Mnt_count,
		s.Magic,
		s.Inode_size,
		s.Block_size,
		s.First_ino,
		s.First_blo,
		s.Bm_inode_start,
		s.Bm_block_start,
		s.Inode_start,
		s.Block_start,
	)
}
