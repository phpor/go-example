package main

import (
	"bufio"
	"bytes"
	"io"
)

// 基于io.Reader的分类器。 用bufio的NewScanner实现
type classifier2 struct {
	reader  *bufio.Scanner
	key     func([]byte) []byte
	done    bool
	lines   [][]byte
	tmpLine []byte
}

func NewClassicer2(reader io.Reader, key func([]byte) []byte) *classifier2 {
	return &classifier2{
		reader: bufio.NewScanner(reader),
		key:    key,
	}
}

func (c *classifier2) Scan() bool {
	if c.done {
		return false
	}
	c.lines = nil
	var key []byte
	if len(c.tmpLine) > 0 {
		c.lines = append(c.lines, c.tmpLine)
		c.tmpLine = nil
	}
	if len(c.lines) > 0 {
		key = c.key(c.lines[0])
	}
	for c.reader.Scan() {
		buf := c.reader.Bytes() // 注意：这里拿出来的东西，如果不是理解用完的话，需要copy走的，否则，下次scan就会被覆盖
		line := make([]byte, len(buf))
		copy(line, buf)
		// 这里判断 c.reader.Err() 的意义也不大
		k := c.key(line)
		if len(key) == 0 {
			key = k
			c.lines = append(c.lines, line)
			continue
		}
		if bytes.Equal(key, k) {
			c.lines = append(c.lines, line)
			continue
		}
		c.tmpLine = line
		break
	}
	if len(c.lines) > 0 {
		return true
	}
	c.done = true
	return false
}

func (c *classifier2) Lines() [][]byte {
	return c.lines
}
