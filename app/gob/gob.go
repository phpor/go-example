package gob

import (
	"encoding/gob"
	"os"
	"fmt"
)

func main() {
	en2 := gob.NewDecoder(os.Stdin)
	var e interface{}
	en2.Decode(e)
	fmt.Print(e)

}
