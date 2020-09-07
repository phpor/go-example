package main

import (
	"crypto/md5"
	"encoding/binary"
	"flag"
	"fmt"
)

func main() {
	id := flag.String("id", "01AgVJOC1b3U3iLivGQkVQIogk", "id")
	flag.Parse()
	sid, pid := shard(*id)
	fmt.Printf("%d, %d\n", sid, pid)
}

func shard(id string) (uint16, uint16) {
	hash := md5.Sum([]byte(id))
	sid := binary.BigEndian.Uint16(hash[:2]) % 32768
	pid := (sid / 256) % 128
	return sid, pid
}
