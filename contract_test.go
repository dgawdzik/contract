package contract

import (
	"fmt"
	"testing"
)

//
// Package Types
//

type state struct {
	name  string
	value int
	obj   interface{}
}

//
// Package Functions
//

func check(t *testing.T, want Exception) {
	got := recover()
	if got != want {
		t.Errorf("expected [%v] but got [%v]", want, got)
	}
}

func testPanic(t *testing.T, exec func(bool, string), want Exception, msg string) {
	defer check(t, want)

	obj := new(interface{})
	obj = nil

	exec(obj != nil, msg)
	t.Errorf("%v panic should have been triggered", want)
}

func testNoPanic(t *testing.T, exec func(bool, string), msg string) {
	defer checkNoPanic(t)

	obj := new(interface{})
	exec(obj != nil, msg)
}

func checkNoPanic(t *testing.T) {
	got := recover()

	if got != nil {
		t.Errorf("expected no panic but got [%v]", got)
	}
}

func checkInvariantPanic(t *testing.T, s *state, want Exception) {
	defer check(t, want)

	s.invariant()
}

//
// Tests
//

func TestContractSucceeds(t *testing.T) {
	tests := []struct {
		exec func(bool, string)
		msg  string
	}{
		{Requires, "object must be provided"},
		{Ensures, "object must have been created"},
		{Assert, "object must have been set"},
		{Invariant, "object must have been set"},
	}

	for _, test := range tests {
		testNoPanic(t, test.exec, test.msg)
	}
}

func TestContractFails(t *testing.T) {

	tests := []struct {
		exec func(bool, string)
		want Exception
		msg  string
	}{
		{Requires, create(requires, fmt.Sprintf(requiresMsg, "object must be provided")), "object must be provided"},
		{Ensures, create(ensures, fmt.Sprintf(ensuresMsg, "object must have been created")), "object must have been created"},
		{Assert, create(assert, fmt.Sprintf(assertMsg, "object must have been set")), "object must have been set"},
		{Invariant, create(invariant, fmt.Sprintf(invariantMsg, "object must have been set")), "object must have been set"},
		{func(cond bool, msg string) { Fail(msg) }, create(fail, fmt.Sprintf(failMsg, "path must not have been executed")), "path must not have been executed"},
	}

	for _, test := range tests {
		testPanic(t, test.exec, test.want, test.msg)
	}
}

func TestInvariantSucceeds(t *testing.T) {
	defer checkNoPanic(t)

	tests := []*state{
		&state{name: "name1", value: 3, obj: "a"},
		&state{name: "name2", value: 3, obj: new(interface{})},
	}

	for _, test := range tests {
		test.invariant()
	}
}

func TestInvariantFails(t *testing.T) {
	tests := []struct {
		s   *state
		msg string
	}{
		{nil, "state must be provided"},
		{&state{name: "", value: 3, obj: "a"}, "name must be set"},
		{&state{name: "name1", value: -1, obj: new(interface{})}, "value must be positive"},
		{&state{name: "name2", value: 1, obj: nil}, "obj must be set"},
	}

	for _, test := range tests {
		checkInvariantPanic(t, test.s, create(invariant, fmt.Sprintf(invariantMsg, test.msg)))
	}
}

/*
	Invariant
*/

func (s *state) invariant() {
	Invariant(s != nil, "state must be provided")
	Invariant(!IsEmpty(s.name), "name must be set")
	Invariant(s.value >= 0, "value must be positive")
	Invariant(s.obj != nil, "obj must be set")
}
