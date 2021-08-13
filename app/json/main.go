package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

type person struct { // 表面上json包没法new一个person值，但是，反射可以new这个结构，这里的私有结构体的私有只是提示给编译器看的
	Name string
	Age  int
}

type Team struct {
	Leader   *person // 指针也能被填充
	Follower person
}

func main() {
	testCase()
}

type StringOld string

func (s *StringOld) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		s1, ok := unquote(b)
		if !ok {
			return errors.New("parse string fail: " + string(b))
		}
		*s = StringOld(s1)
		return nil
	}
	*s = StringOld(b)
	return nil
}

func (s *StringOld) UnmarshalText(text []byte) error {
	*s = StringOld(text)
	return nil
}

type String string

func (s *String) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		s1 := ""
		if err := json.Unmarshal(b, &s1); err != nil {
			return err
		}
		*s = String(s1)
		return nil
	}
	*s = String(b)
	return nil
}

type StringSlice []string

func (s *StringSlice) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		s1 := ""
		if err := json.Unmarshal(b, &s1); err != nil {
			return err
		}
		*s = []string{s1}
		return nil
	} else if b[0] == '[' {
		var s1 []string
		if err := json.Unmarshal(b, &s1); err != nil {
			return err
		}
		*s = s1
		return nil
	}
	return errors.New("parse string fail: " + string(b))
}

func testCase() {
	s := &struct {
		Os      String      `json:"os"`
		Version String      `json:"version"`
		Cpu1    StringSlice `json:"Cpu1"`
		Cpu2    StringSlice `json:"Cpu2"`
	}{}
	if err := json.Unmarshal([]byte(`{"os":"aa", "version":1, "cpu1":"arm1", "cpu2":["arm1", "arm2"]}`), s); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", s)
	v, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v))
}

func testMap() {
	//s := []byte(`{"a":"A", "b":2}`)
	s := []byte(`{"mfp":"{\"1\":\"13.7\",\"11\":\"E99163D8-5A2B-47ED-A9F3-55555AE5E5E2\",\"18\":\"iPhone\",\"2\":\"796997f731e5681f6e9e5ea298866da43e1cf2f0\",\"23\":\"zh-Hans-CN\",\"24\":\"460,02,中国移动\",\"33\":\"1614525706\",\"37\":\"45.276970\",\"38\":\"3117580288\",\6002\",\"69\":\"iPhone10,2\",\"70\":\"255989469184\",\"73\":\"1597084210.236371\",\"75\":\"CN\",\"80\":\"L3Zhci9tb2JpbGUvTGlicmFyeS9Vc2VyQ29uZmlndXJhdGlvblByb2ZpbGVzL1B1YmxpY0luZm8vTUNNZXRhLnBsaXN0\",\"83\":\"none\",\"85\":\"D21AP\"}","src":"getaid"}`)
	m := map[string]interface{}{}
	err := json.Unmarshal(s, &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", m)
}

func testStruct() {
	str := `{"leader":{"Name":"wang", "age":22},"Follower":{"Name":"sun", "age":20}}`
	t := &Team{}
	if err := json.Unmarshal([]byte(str), t); err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("leader: %#v, follower: %#v", t.Leader, t.Follower)

}

func testEscape() {
	a := struct {
		A string
	}{`{"1":"13.6","11":"E99163D8-5A2B-47ED-A9F3-55555AE5E5E2","18":"张小花的 iPhone","2":"696997f731e5681f6e9e5ea298866da43e1cf2f0","23":"zh-Hans-CN","24":"460,02,中国移动","33":"1614525706","37":"45.276970","38":"3117580288","4":"46002","69":"iPhone10,2","70":"255989469":"1597084210.236371","75":"CN","80":"L3Zhci9tb2JpbGUvTGlicmFyeS9Vc2VyQ29uZmlndXJhdGlvblByb2ZpbGVzL1B1YmxpY0luZm8vTUNNZXRhLnBsaXN0","83":"none","85":"D21AP"}`}
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(b))
}

// unquote converts a quoted JSON string literal s into an actual string t.
// The rules are different than for Go, so cannot use strconv.Unquote.
func unquote(s []byte) (t string, ok bool) {
	s, ok = unquoteBytes(s)
	t = string(s)
	return
}

func unquoteBytes(s []byte) (t []byte, ok bool) {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return
	}
	s = s[1 : len(s)-1]

	// Check for unusual characters. If there are none,
	// then no unquoting is needed, so return a slice of the
	// original bytes.
	r := 0
	for r < len(s) {
		c := s[r]
		if c == '\\' || c == '"' || c < ' ' {
			break
		}
		if c < utf8.RuneSelf {
			r++
			continue
		}
		rr, size := utf8.DecodeRune(s[r:])
		if rr == utf8.RuneError && size == 1 {
			break
		}
		r += size
	}
	if r == len(s) {
		return s, true
	}

	b := make([]byte, len(s)+2*utf8.UTFMax)
	w := copy(b, s[0:r])
	for r < len(s) {
		// Out of room? Can only happen if s is full of
		// malformed UTF-8 and we're replacing each
		// byte with RuneError.
		if w >= len(b)-2*utf8.UTFMax {
			nb := make([]byte, (len(b)+utf8.UTFMax)*2)
			copy(nb, b[0:w])
			b = nb
		}
		switch c := s[r]; {
		case c == '\\':
			r++
			if r >= len(s) {
				return
			}
			switch s[r] {
			default:
				return
			case '"', '\\', '/', '\'':
				b[w] = s[r]
				r++
				w++
			case 'b':
				b[w] = '\b'
				r++
				w++
			case 'f':
				b[w] = '\f'
				r++
				w++
			case 'n':
				b[w] = '\n'
				r++
				w++
			case 'r':
				b[w] = '\r'
				r++
				w++
			case 't':
				b[w] = '\t'
				r++
				w++
			case 'u':
				r--
				rr := getu4(s[r:])
				if rr < 0 {
					return
				}
				r += 6
				if utf16.IsSurrogate(rr) {
					rr1 := getu4(s[r:])
					if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
						// A valid pair; consume.
						r += 6
						w += utf8.EncodeRune(b[w:], dec)
						break
					}
					// Invalid surrogate; fall back to replacement rune.
					rr = unicode.ReplacementChar
				}
				w += utf8.EncodeRune(b[w:], rr)
			}

		// Quote, control characters are invalid.
		case c == '"', c < ' ':
			return

		// ASCII
		case c < utf8.RuneSelf:
			b[w] = c
			r++
			w++

		// Coerce to well-formed UTF-8.
		default:
			rr, size := utf8.DecodeRune(s[r:])
			r += size
			w += utf8.EncodeRune(b[w:], rr)
		}
	}
	return b[0:w], true
}

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
// or it returns -1.
func getu4(s []byte) rune {
	if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
		return -1
	}
	var r rune
	for _, c := range s[2:6] {
		switch {
		case '0' <= c && c <= '9':
			c = c - '0'
		case 'a' <= c && c <= 'f':
			c = c - 'a' + 10
		case 'A' <= c && c <= 'F':
			c = c - 'A' + 10
		default:
			return -1
		}
		r = r*16 + rune(c)
	}
	return r
}
