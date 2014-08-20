// 参考： https://groups.google.com/forum/#!topic/golang-nuts/mOvX0bmJoeI
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type SeekingBuffer struct {
	b      []byte
	buffer *bytes.Buffer
	offset int64
	size   int64
}

func NewSeekingBuffer(b []byte) *SeekingBuffer {
	if b == nil {
		return nil
	}
	return &SeekingBuffer{
		b:      b,
		buffer: bytes.NewBuffer(b),
		offset: 0,
	}
}

func (fb *SeekingBuffer) Read(p []byte) (n int, err error) {
	n, err = fb.buffer.Read(p)
	fb.offset += int64(n)
	return n, err
}

func (fb *SeekingBuffer) Seek(offset int64, whence int) (ret int64, err error) {
	var newoffset int64
	switch whence {
	case 0:
		newoffset = offset
	case 1:
		newoffset = fb.offset+offset
	case 2:
		newoffset = int64(len(fb.b))-offset
	}
	if newoffset == fb.offset {
		return newoffset, nil
	}
	fb.buffer = bytes.NewBuffer(fb.b[newoffset:])
	fb.offset = newoffset
	return fb.offset, nil
}

func main() {
	fb := NewSeekingBuffer([]byte("Hello, playground"))
	pos, err := fb.Seek(0, 0)
	fmt.Println("pos = ", pos, ", error = ", err)
	pos, err = fb.Seek(7, 0)
	fmt.Println("pos = ", pos, ", error = ", err)
	io.Copy(os.Stdout, fb)

}
