package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	sync.RWMutex
	reader io.Reader
	bytes  int64
	ops    int
}

type writeCounter struct {
	sync.RWMutex
	writer io.Writer
	bytes  int64
	ops    int
}

type readWriteCounter struct {
	rc ReadCounter
	wc WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{NewReadCounter(readwriter), NewWriteCounter(readwriter)}
}

func (rvc *readWriteCounter) Read(p []byte) (int, error) {
	bts, err := rvc.rc.Read(p)
	if err != nil {
		return 0, err
	}
	return bts, nil
}

func (rvc *readWriteCounter) Write(p []byte) (int, error) {
	bts, err := rvc.wc.Write(p)
	if err != nil {
		return 0, err
	}
	return bts, nil
}

func (rvc *readWriteCounter) ReadCount() (int64, int) {
	return rvc.rc.ReadCount()
}

func (rvc *readWriteCounter) WriteCount() (int64, int) {
	return rvc.wc.WriteCount()
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.Lock()
	defer rc.Unlock()
	bts, err := rc.reader.Read(p)
	if err != nil {
		return 0, err
	}
	rc.bytes += int64(bts)
	rc.ops++
	return bts, nil
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.Lock()
	defer rc.Unlock()
	return rc.bytes, rc.ops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.Lock()
	defer wc.Unlock()
	bts, err := wc.writer.Write(p)
	if err != nil {
		return 0, err
	}
	wc.bytes += int64(bts)
	wc.ops++
	return bts, nil
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.Lock()
	defer wc.Unlock()
	return wc.bytes, wc.ops
}
