package main

import (
	"flag"
	"fmt"

	"github.com/ruoskija/automata/internal/automata"
)

func main() {

	rule, steps, initialState := parseFlags()

	system := automata.NewSystem(rule)
	if err := system.SetState(initialState); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	for i := 0; i < steps; i++ {
		fmt.Println(system.String())
		system.Step()
	}

	return
}

func parseFlags() (int, int, string) {
	rulePtr := flag.Int("rule", 30, "rule aka Wolfram code")
	stepsPtr := flag.Int("step", 16, "number of steps printed")
	statePtr := flag.String(
		"init",
		"0000000000000001000000000000000",
		"initial state of the automata")
	flag.Parse()
	return *rulePtr, *stepsPtr, *statePtr
}
