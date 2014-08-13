package decoder


import (
	"fmt"

	"github.com/phpor/go/rdb"
)

type diff struct {
	db int
	i  int
	rdb.Nopdiff
}

func (p *diff) StartDatabase(n int) {
	p.db = n
}

func (p *diff) Set(key, value []byte, expiry int64) {
	fmt.Printf("db=%d %q -> %q\n", p.db, key, value)
}

func (p *diff) Hset(key, field, value []byte) {
	fmt.Printf("db=%d %q . %q -> %q\n", p.db, key, field, value)
}

func (p *diff) Sadd(key, member []byte) {
	fmt.Printf("db=%d %q { %q }\n", p.db, key, member)
}

func (p *diff) StartList(key []byte, length, expiry int64) {
	p.i = 0
}

func (p *diff) Rpush(key, value []byte) {
	fmt.Printf("db=%d %q[%d] -> %q\n", p.db, key, p.i, value)
	p.i++
}

func (p *diff) Zadd(key []byte, score float64, member []byte) {
	fmt.Printf("db=%d %q[%d] -> {%q, score=%g}\n", p.db, key, p.i, member, score)
	p.i++
}

func (p *diff) StartZSet(key []byte, cardinality, expiry int64) {
	p.i = 0
}
