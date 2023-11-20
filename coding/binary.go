package coding

import (
	"encoding/binary"
	"errors"
	"fmt"
	"sync"
)

const (
	// MaxHeaderSize = 2 + 10 + 10 + 10 + 4 (10 refer to binary.MaxVarintLen64)
	MaxHeaderSize = 36
	Uint32Size    = 4
	Uint16Size    = 2
)

var UnmarshalError = errors.New("an error occurred in Unmarshal")

// RequestHeader request header structure looks like:
// +--------------+----------------+----------+------------+----------+
// | CompressType |      Method    |    ID    | RequestLen | Checksum |
// +--------------+----------------+----------+------------+----------+
// |    uint16    | uvarint+string |  uvarint |   uvarint  |  uint32  |
// +--------------+----------------+----------+------------+----------+
type RequestHeader struct {
	sync.RWMutex
	CompressType uint16
	Method       string
	ID           uint64
	RequestLen   uint32
	Checksum     uint32
}

// HeaderToBytes  Transfer RequestHeader to byte slice
func (r *RequestHeader) HeaderToBytes() []byte {
	// lock
	r.RLock()
	defer r.RUnlock()
	header := make([]byte, MaxHeaderSize)
	idx := 0
	// 1. CompressType to byte slice
	binary.LittleEndian.PutUint16(header[idx:], r.CompressType)
	idx += Uint16Size
	// 2. Method to byte slice
	idx += writeString(header[idx:], r.Method)
	// 3. ID to byte slice
	idx += binary.PutUvarint(header[idx:], uint64(r.ID))
	// 4. RequestLen

	return nil
}

func writeString(header []byte, s string) int {
	// String write to header, two parameters: len + str
	idx := 0
	idx += binary.PutUvarint(header[idx:], uint64(len(s))) // len's len
	copy(header[idx:], s)
	idx += len(s) // s's len
	return idx
}
func Encoding() {
	num := 1212
	header := make([]byte, 120)
	header1 := make([]byte, 120)
	binary.LittleEndian.PutUint16(header, uint16(num))
	binary.BigEndian.PutUint16(header1, uint16(num))
	fmt.Println("Encoding:", header)
	fmt.Println("Encoding:", header1)

}
