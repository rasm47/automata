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

func ruleFromInt(n int) []uint {
	const ruleSize int = 8
	rule := make([]uint, ruleSize)

	for i := 0; i < ruleSize; i++ {
		rule[i] = uint(n % 2)
		n = n >> 1
	}

	return rule
}

// NewSystem creates an elementary cellular automata system
// with the provided rule number
func NewSystem(ruleNumber int) System {
	state := []uint{0, 0, 0, 0, 1, 0, 0, 0} // dummy initial state
	rule := ruleFromInt(ruleNumber)
	return System{state, rule}
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

// SetState set the system state manually
func (sys *System) SetState(str string) error {
	state := []uint{}
	for _, char := range str {
		if char == '1' {
			state = append(state, 1)
		} else if char == '0' {
			state = append(state, 0)
		}
	}
	if len(state) < 3 {
		return errors.New("invalid state")
	}
	sys.state = state
	return nil
}
