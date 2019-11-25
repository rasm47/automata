package automata

import (
	"errors"
	"strconv"
	"strings"
)

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

func updateState(state []uint, rule []uint) []uint {
	ns := make([]uint, len(state))
	for i := len(state) - 1; i >= 0; i-- {
		ns[i] = updateCell(state, rule, i)
	}
	return ns
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

// NewSystem creates an elementary cellular automata system
// with the provided rule number and initial state
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
func (sys *System) String() string {

	strs := []string{}
	for _, n := range sys.state {
		s := strconv.Itoa(int(n))
		strs = append(strs, s)
	}

	return strings.Join(strs, "")
}

// Step updates the state by one increment
func (sys *System) Step() {
	sys.state = updateState(sys.state, sys.rule)
	return
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
