package cache

import (
	"errors"
	"sync"
)

// sync.Map 没有长度，不方便控制，还是要使用原始的map才行

type storeGroup struct {
	size int
	db   map[string]interface{}
	mu   sync.Mutex
}

var store map[string]storeGroup

// Create a cache
func CreateGroup(group string, size int) error {
	if _, ok := store[group]; ok {
		return nil
	}
	store[group] = storeGroup{size: size, db: map[string]interface{}{}}
	return nil
}

func Set(group string, key string, val interface{}) error {
	if g, ok := store[group]; !ok {
		return errors.New("group not init")
	} else {
		g.mu.Lock()
		if _, ok := g.db[key]; !ok {
			if len(g.db) >= g.size {
				// 如何删除掉map中的一个值？
			}
		}
		g.db[key] = val
		g.mu.Unlock()
	}

	return nil
}

func init() {

}
