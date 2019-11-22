package automata

import (
	"errors"
	"strconv"
	"strings"
)

func getNeighbors(s []uint, me int) [3]uint {
	sLen := len(s)
	if sLen < 3 {
		return [3]uint{0, 0, 0}
	}

	if me == 0 {
		return [3]uint{s[sLen-1], s[me], s[1]}
	}
	if me == sLen-1 {
		return [3]uint{s[me-1], s[me], s[0]}
	}
	return [3]uint{s[me-1], s[me], s[me+1]}
}

func updateCell(s []uint, r []uint, i int) uint {
	neighbors := getNeighbors(s, i)
	ruleIndex := neighbors[2] * 1
	ruleIndex += neighbors[1] * 2
	ruleIndex += neighbors[0] * 4
	return r[ruleIndex]
}

func updateState(s []uint, r []uint) []uint {
	ns := make([]uint, len(s))
	for i := 0; i < len(s); i++ {
		ns[i] = updateCell(s, r, i)
	}
	return ns
}

func ruleFromInt(n int) []uint {
	rule := make([]uint, 8)
	i := 0
	for i < 8 {
		rule[i] = uint(n % 2)
		n = n >> 1
		i++
	}
	return rule
}

// NewSystem creates an elementary cellular automata system with the provided rule number
// the default state is "00001000"
func NewSystem(ruleNumber int) System {
	state := []uint{0, 0, 0, 0, 1, 0, 0, 0}
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
