package automata

import (
	"errors"
	"strconv"
)

// NewSystem creates an elementary cellular automata system
// with the provided rule number and initial state.
// Invalid inputs generate errors.
func NewSystem(ruleNumber int, initialState string) (System, error) {
	state, err := newState(initialState)
	if err != nil {
		return System{}, err
	}
	rule, err := newRule(ruleNumber)
	if err != nil {
		return System{}, err
	}
	return System{state, rule}, nil
}

// String gives the state of the system as a string
func (system *System) String() string {
	output := ""
	for _, n := range system.state {
		output += strconv.Itoa(int(n))

	}

	return output
}

// Step updates the state by one increment
func (system *System) Step() {

	updatedState := make([]uint, len(system.state))
	for i := len(system.state) - 1; i >= 0; i-- {
		updatedState[i] = updateCell(system.state, system.rule, i)
	}

	system.state = updatedState
	return
}

func newRule(n int) ([]uint, error) {
	if n < 0 || n > 255 {
		return []uint{}, errors.New("rule number out of bounds (0-255)")
	}

	const ruleSize int = 8
	rule := make([]uint, ruleSize)

	for i := 0; i < ruleSize; i++ {
		rule[i] = uint(n % 2)
		n = n >> 1
	}

	return rule, nil
}

func newState(str string) ([]uint, error) {
	state := []uint{}
	for _, char := range str {
		if char == '1' {
			state = append(state, 1)
		} else if char == '0' {
			state = append(state, 0)
		}
	}
	if len(state) < 3 {
		return []uint{}, errors.New("invalid state")
	}
	return state, nil
}

func getNeighbors(state []uint, myPos int) [3]uint {
	stateLen := len(state)

	if myPos == 0 {
		return [3]uint{state[stateLen-1], state[myPos], state[1]}
	}
	if myPos == stateLen-1 {
		return [3]uint{state[myPos-1], state[myPos], state[0]}
	}
	return [3]uint{state[myPos-1], state[myPos], state[myPos+1]}
}

func updateCell(state []uint, rule []uint, index int) uint {
	neighbors := getNeighbors(state, index)
	ruleIndex := neighbors[2] * 1
	ruleIndex += neighbors[1] * 2
	ruleIndex += neighbors[0] * 4
	return rule[ruleIndex]
}
