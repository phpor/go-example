package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/savaki/jq"
)

func main() {
	filter := flag.String("filter", ".", "filter")
	raw := flag.Bool("r", false, "raw string")
	flag.Parse()

	op, _ := jq.Parse(*filter)           // create an Op

	stdin := bufio.NewReader(os.Stdin)
	for {
		line, err := stdin.ReadBytes('\n')
		if len(line) == 0 && err != nil {
			break
		}
		line = bytes.TrimRight(line, "\n")
		v, err := op.Apply(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
		} else {
			if *raw {
				v = bytes.TrimPrefix(v, []byte{'"'})
				v = bytes.TrimSuffix(v, []byte{'"'})
			}
			fmt.Printf("%s\n", string(v))
		}

	}
}