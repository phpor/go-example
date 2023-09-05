package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"io"
	"os"
	"strings"
)

type relation struct {
	from string
	to   string
}

// 根据  go mod graph 的结果生成mermaid图
func main() {
	var outType string
	flag.StringVar(&outType, "out-type", "no-version | with-version", "out type")
	flag.Parse()
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("flowchart TB\n")

	if outType == "with-version" {
		withVersion(b)
	} else {
		noVersion(b)
	}

}

func noVersion(b []byte) {
	arr := bytes.Split(b, []byte{'\n'})
	relations := hashset.New()
	group := map[string]*hashset.Set{}
	class := func(module string) {
		module = strings.Split(module, "@")[0]
		p1Arr := strings.Split(module, "/")
		g := strings.Join(p1Arr[0:2], "/")
		if v, ok := group[g]; !ok {
			group[g] = hashset.New(module)
		} else {
			v.Add(module)
		}
	}
	for _, l := range arr {
		a := strings.Split(string(l), " ")
		if len(a) != 2 {
			continue
		}
		class(a[0])
		class(a[1])
		r := strings.Split(a[0], "@")[0] + "-->" + strings.Split(a[1], "@")[0]
		relations.Add(r)
	}
	for _, r := range relations.Values() {
		fmt.Println(r)
	}

	buildSubGraph(group)
}

func withVersion(b []byte) {
	arr := bytes.Split(bytes.ReplaceAll(b, []byte{'@'}, []byte{'#'}), []byte{'\n'})
	group := map[string]*hashset.Set{}
	var relations []*relation
	class := func(module string) {
		p1Arr := strings.Split(module, "#")
		g := p1Arr[0]
		if v, ok := group[g]; !ok {
			group[g] = hashset.New(module)
		} else {
			v.Add(module)
		}
	}
	for _, l := range arr {
		a := strings.Split(string(l), " ")
		if len(a) != 2 {
			continue
		}
		class(a[0])
		class(a[1])
		relations = append(relations, &relation{a[0], a[1]})
	}
	for _, v := range relations {
		fmt.Printf("\t%s --> %s\n", v.from, v.to)
	}
	buildSubGraph(group)
}

func buildSubGraph(group map[string]*hashset.Set) {
	for g := range group {
		if group[g].Size() == 1 {
			continue
		}
		fmt.Printf("\tsubgraph %s\n", g)
		for _, v := range group[g].Values() {
			fmt.Printf("\t\t%s\n", v.(string))
		}
		fmt.Printf("\tend\n")
	}
}
