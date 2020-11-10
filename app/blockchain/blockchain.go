package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
)

type Block struct {
	pre    *Block
	next   *Block
	data   fmt.Stringer
	rand   [8]byte
	digest [32]byte
}

type BlockChain struct {
	head     *Block
	tail     *Block
	initData fmt.Stringer
}

func New(data fmt.Stringer) *BlockChain {
	r := makeRand()
	digest := sha256.Sum256(append([]byte(data.String()), r[:]...))
	firstBlock := &Block{
		pre:    nil,
		next:   nil,
		data:   data,
		rand:   r,
		digest: digest,
	}
	return &BlockChain{
		head:     firstBlock,
		tail:     firstBlock,
		initData: data,
	}
}

func makeRand() [8]byte {
	r := [8]byte{}
	rand.Read(r[:])
	return r
}

func makeDigest(preDigest [32]byte, data []byte, rand [8]byte) [32]byte {
	s := append(preDigest[:], data...)
	s = append(s, rand[:]...)
	return sha256.Sum256(s)
}

func (bc *BlockChain) Add(data fmt.Stringer) *BlockChain {
	if bc == nil {
		return nil
	}
	r := makeRand()
	b := &Block{
		pre:    bc.tail,
		next:   nil,
		data:   data,
		rand:   r,
		digest: makeDigest(bc.tail.digest, []byte(data.String()), r),
	}
	bc.tail.next = b
	bc.tail = b
	return bc
}

func (b *Block) String() string {
	return fmt.Sprintf("%s", b.data)
}

func (bc *BlockChain) Verify() bool {
	ret := true
	bc.Walk(func(b *Block) bool {
		if b.pre == nil { // 第一个块儿先不校验，逻辑上
			if b.data.String() == bc.initData.String() {
				return true
			}
			ret = false
			return false
		}
		d := makeDigest(b.pre.digest, []byte(b.data.String()), b.rand)
		if !bytes.Equal(b.digest[:], d[:]) {
			ret = false
			return false
		}
		return true
	})
	return ret
}

func (bc *BlockChain) Walk(f func(b *Block) bool) {
	b := bc.head
	for {
		if !f(b) {
			break
		}
		if b.next == nil {
			break
		}
		b = b.next
	}
}

func (bc *BlockChain) Delete(n int) {
	i := 0
	bc.Walk(func(b *Block) bool {
		if i == n {
			b.pre.next = b.next
			b.next.pre = b.pre
			return false
		}
		i++
		return true
	})
}

func (bc *BlockChain) Dump() {
	bc.Walk(func(b *Block) bool {
		fmt.Printf("%s\n", b)
		return true
	})
}
