package cli

import (
	"errors"
	"reflect"
	"strings"
)

var ErrActionNotFound = errors.New("action not exists")

type Cli struct {
	param map[string]interface{}
}

func (c *Cli) AddParam(key string, val interface{}) {
	if c.param == nil {
		c.param = map[string]interface{}{}
	}
	c.param[key] = val
}

func (c *Cli) GetParam(key string) interface{} {
	if c.param == nil {
		return nil
	}
	return c.param[key]
}

func Dispatch(c interface{}, action string) error {
	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)
	method := "Action" + hump(action)

	_, ok := t.MethodByName(method)
	if !ok {
		return ErrActionNotFound
	}

	ret := v.MethodByName(method).Call(nil)
	if err, ok := ret[0].Interface().(error); ok {
		return err
	}
	return nil
}

func ActionList(c interface{}) (ret []string) {
	t := reflect.TypeOf(c)
	if t == nil {
		return
	}
	num := t.NumMethod()
	for i := 0; i < num; i++ {
		m := t.Method(i)
		if strings.HasPrefix(m.Name, "Action") {
			name := strings.TrimPrefix(m.Name, "Action")
			ret = append(ret, toDashSplit(name))
		}
	}
	return
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func toLower(c byte) byte {
	if isUpper(c) {
		c = c + 32
	}
	return c
}

func toUpper(c byte) byte {
	if isLower(c) {
		c = c - 32
	}
	return c
}

// 将驼峰式转换为 - 分隔
func toDashSplit(str string) string {
	length := len(str)
	if length == 0 {
		return ""
	}
	ret := make([]byte, 2*length)
	ret[0] = toLower(str[0])
	j := 1
	for i := 1; i < length; i++ {
		if isUpper(str[i]) {
			ret[j] = '-'
			j++
			ret[j] = toLower(str[i])
		} else {
			ret[j] = str[i]
		}
		j++
	}
	return string(ret[:j])
}

// 将 - 分隔转换为驼峰表示
func hump(str string) string {
	b := []byte(str)
	length := len(b)
	if length == 0 {
		return ""
	}
	b[0] = toUpper(b[0])
	j := 1
	for i := 1; i < length; i++ {
		if b[i] == '-' {
			i++
			if i == length { // 如果 - 结尾，则直接忽略
				break
			}
			b[i] = toUpper(b[i])
		}
		b[j] = b[i]
		j++
	}
	return string(b[:j])
}
