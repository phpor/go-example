package  main

import "fmt"

type parent struct {

}
func (p *parent) hello() {
	fmt.Println("parent hello")
}
func (p *parent) say() {
	p.hello()
}
type child struct {
	parent
}
func (c *child) hello() {
	c.parent.say()
}

func main() {
	c := &child{}
	c.say()
}
