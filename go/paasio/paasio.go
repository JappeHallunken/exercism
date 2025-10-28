package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	sync.RWMutex
	reader io.Reader
	n      int64
	nops   int
}

type writeCounter struct {
	sync.RWMutex
	writer io.Writer
	n      int64
	nops   int
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		NewReadCounter(readwriter),
		NewWriteCounter(readwriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.Lock()
	defer rc.Unlock()

	n, err := rc.reader.Read(p)
	rc.n += int64(n)
	rc.nops++
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.RLock()
	defer rc.RUnlock()

	return rc.n, rc.nops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.Lock()
	defer wc.Unlock()

	n, err := wc.writer.Write(p)
	wc.n += int64(n)
	wc.nops++
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.RLock()
	defer wc.RUnlock()

	return wc.n, wc.nops
}
