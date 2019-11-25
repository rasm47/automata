package automata

import "testing"

func TestNewSystem(t *testing.T) {
	rule := 123
	initialState := "01010101"
	system, err := NewSystem(rule, initialState)
	if err != nil {
		t.Error(err)
	}

	expectedRule := []uint{1, 1, 0, 1, 1, 1, 1, 0}
	if len(system.rule) != len(expectedRule) {
		t.Error("the length of the rule of the new system is not correct")
	}
	for i := range system.rule {
		if expectedRule[i] != system.rule[i] {
			t.Errorf("At rule index %v got %v, expected to get %v", i, expectedRule[i], system.rule[i])
		}
	}

	expectedState := []uint{0, 1, 0, 1, 0, 1, 0, 1}
	if len(system.rule) != len(expectedState) {
		t.Error("the length of the state of the new system is not correct")
	}
	for i := range system.state {
		if expectedState[i] != system.state[i] {
			t.Errorf("At state index %v got %v, expected to get %v", i, expectedState[i], system.state[i])
		}
	}

	_, err = NewSystem(-1, initialState)
	if err == nil {
		t.Error("Expected rule -1 to result in an error")
	}

	_, err = NewSystem(1234567, initialState)
	if err == nil {
		t.Error("Expected rule 1234567 to result in an error")
	}

	_, err = NewSystem(123, "teststring123")
	if err == nil {
		t.Error("Expected initialstate \"teststring123\" to result in an error")
	}

	_, err = NewSystem(123, "")
	if err == nil {
		t.Error("Expected empty string as initial state to result in an error")
	}
}

func TestSystem_String(t *testing.T) {
	system := System{}
	system.state = []uint{0, 0, 1, 1, 1, 0}
	expectedString := "001110"
	recievedString := system.String()
	if expectedString != recievedString {
		t.Errorf("expected %v; got %v\n", expectedString, recievedString)
	}

}

func TestSystem_Step(t *testing.T) {
	system := System{}
	system.state = []uint{0, 1, 0}
	system.rule = []uint{0, 0, 0, 0, 0, 0, 0, 0}

	system.Step()

	recievedState := system.state
	expectedState := []uint{0, 0, 0}

	if len(recievedState) != len(expectedState) {
		t.Error("Step changed state length")
	}
	for i := range recievedState {
		if expectedState[i] != recievedState[i] {
			t.Errorf("At state index %v got %v, expected to get %v", i, expectedState[i], recievedState[i])
		}
	}
}
