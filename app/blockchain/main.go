package main

import "fmt"

type Data string

func (d Data) String() string {
	return string(d)
}

func main() {
	bc := New(Data("init data"))
	bc.Add(Data("A -> B : 100"))
	bc.Add(Data("B -> C : 100"))
	bc.Add(Data("C -> D : 100"))
	bc.Add(Data("D -> E : 100"))
	verify := func() {
		if bc.Verify() {
			fmt.Printf("Verify ok\n")
		} else {
			fmt.Printf("Verify fail\n")
		}
	}

	bc.Dump()
	verify()

	// ==== 搞点儿破坏
	//bc.Delete(2)  // 删除一个区块，校验就得失败
	bc.Walk(func(b *Block) bool { // 修改某个区块的内容，校验就得失败
		if b.data.String() == "B -> C : 100" {
			b.data = Data("B -> C : 200")
			return false
		}
		return true
	})

	bc.Dump()
	verify()

}
