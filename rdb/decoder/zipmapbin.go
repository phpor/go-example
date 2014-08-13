package decoder

import (
	"fmt"
	"os"

	"github.com/phpor/go/rdb"
	"encoding/binary"
)

type zipmapbin struct {
	db int
	i  int
	rdb.NopDecoder
}

func (p *zipmapbin) StartDatabase(n int) {
	p.db = n
}

func (p *zipmapbin) Hset(key, field, value []byte) {
	lenByte := make([]byte, 4)
	binary.BigEndian.PutUint32(lenByte, uint32(len(key)))
	os.Stdout.Write(lenByte)
	os.Stdout.Write(key)

	binary.BigEndian.PutUint32(lenByte, uint32(len(field)))
	os.Stdout.Write(lenByte)
	os.Stdout.Write(field)

	binary.BigEndian.PutUint32(lenByte, uint32(len(value)))
	os.Stdout.Write(lenByte)
	os.Stdout.Write(value)
}
