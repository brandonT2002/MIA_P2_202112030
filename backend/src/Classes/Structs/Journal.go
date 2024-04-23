package structs

import (
	"bytes"
	"encoding/binary"
	"strings"
	"time"
)

type Journal struct {
	Operation string    // 8 bytes
	Path      string    // 100 bytes
	Content   string    // 100 bytes
	Date      time.Time // 4 bytes
} // 212 bytes

func NewJournal(operation, path, content string, date time.Time) *Journal {
	return &Journal{operation, path, content, date}
}

func (j *Journal) Encode() []byte {
	var result []byte
	// Codificar la operación
	if j.Operation != "" {
		result = append(result, []byte(j.Operation)...)
		result = append(result, bytes.Repeat([]byte{0}, 8-len(j.Operation))...)
	} else {
		result = append(result, bytes.Repeat([]byte{0}, 8)...)
	}
	// Codificar la ruta
	if j.Path != "" {
		result = append(result, []byte(j.Path)...)
		result = append(result, bytes.Repeat([]byte{0}, 100-len(j.Path))...)
	} else {
		result = append(result, bytes.Repeat([]byte{0}, 100)...)
	}
	// Codificar el contenido
	if j.Content != "" {
		result = append(result, []byte(j.Content)...)
		result = append(result, bytes.Repeat([]byte{0}, 100-len(j.Content))...)
	} else {
		result = append(result, bytes.Repeat([]byte{0}, 100)...)
	}
	// Codificar la fecha
	dateBytes := make([]byte, 4)
	timestamp := j.Date.Unix()
	for i := 0; i < 4; i++ {
		dateBytes[i] = byte(timestamp >> uint(8*(3-i)))
	}
	result = append(result, dateBytes...)
	return result
}

func DecodeJournal(data []byte) *Journal {
	journal := &Journal{}
	// Decodificar la operación
	journal.Operation = decodeString(data[:8])
	// Decodificar la ruta
	journal.Path = decodeString(data[8:108])
	// Decodificar el contenido
	journal.Content = decodeString(data[108:208])
	// Decodificar la fecha
	timestamp := binary.BigEndian.Uint32(data[208:])
	if timestamp != 0 {
		journal.Date = time.Unix(int64(timestamp), 0)
	}
	return journal
}

func (j *Journal) GetDot() string {
	// Reemplaza los caracteres especiales en el contenido
	content := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(j.Content, "\n", "\\n"), "\"", "\\\""), "'", "\\'")

	// Formatea la fecha en una cadena legible
	dateStr := j.Date.Format("2006-01-02") // El formato "2006-01-02" es para "año-mes-día"

	// Forma la cadena de salida
	return "\n\t\t<TR><TD>" + strings.TrimSpace(j.Operation) + "</TD><TD>" + j.Path + "</TD><TD>" + content + "</TD><TD>" + dateStr + "</TD></TR>"
}

func decodeString(data []byte) string {
	return string(bytes.TrimRight(data, "\x00"))
}
