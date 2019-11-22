package main

import (
	"flag"
	"fmt"

	"github.com/ruoskija/automata/internal/automata"
)

func main() {

	rulePtr := flag.Int("rule", 30, "rule aka Wolfram code")
	statePtr := flag.String("init", "0000000000000001000000000000000", "initial state of the automata")
	stepsPtr := flag.Int("step", 16, "number of steps printed")
	flag.Parse()

	system := automata.NewSystem(*rulePtr)
	if err := system.SetState(*statePtr); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	for i := 0; i < *stepsPtr; i++ {
		fmt.Println(system.String())
		system.Step()
	}

	return
}
