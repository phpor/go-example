package main

import "github.com/nakabonne/gosivy/agent"

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		panic(err)
	}
	defer agent.Close()
}
