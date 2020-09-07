package main

import (
	"bufio"
	"bytes"
	"io"
)

// 基于io.Reader的分类器
type classifier struct {
	reader  *bufio.Reader
	key     func([]byte) []byte
	done    bool
	lines   [][]byte
	tmpLine []byte
}

func NewClassifier(reader io.Reader, key func([]byte) []byte) *classifier {
	return &classifier{
		reader: bufio.NewReader(reader),
		key:    key,
	}
}

func (c *classifier) Scan() bool {
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

	for {
		line, err := c.reader.ReadBytes('\n')
		if len(line) > 0 {
			line = bytes.TrimSpace(line)

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
		if err != nil {
			break
		}
	}

	if len(c.lines) > 0 {
		return true
	}
	c.done = true
	return false
}

func (c *classifier) Lines() [][]byte {
	return c.lines
}
