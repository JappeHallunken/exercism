package circular

import (
	"errors"
	"io"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

type Buffer struct {
	data  []byte
	start int
	count int
}

var ErrBufferFull = errors.New("buffer full")

func NewBuffer(size int) *Buffer {
	data := make([]byte, size)
	return &Buffer{data: data}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, io.EOF
	}
	value := b.data[b.start]
	b.start = (b.start + 1) % len(b.data)
	b.count--
	return value, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.data) == b.count {
		return ErrBufferFull
	}
	idx := (b.start + b.count) % len(b.data)
	b.data[idx] = c
	b.count++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	err := b.WriteByte(c)
	if err != nil {
		b.data[b.start] = c
		b.start = (b.start + 1) % len(b.data)
	}
}

func (b *Buffer) Reset() {
	b.start, b.count = 0, 0
}
