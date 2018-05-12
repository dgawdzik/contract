package contract

import (
	"fmt"
	"testing"
)

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
	defer func() {
		got := recover()
		if got != nil {
			t.Errorf("expected no panic but got [%v]", got)
		}
	}()

	obj := new(interface{})
	exec(obj != nil, msg)
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
		{func(cond bool, msg string) { Fail(msg) }, create(fail, fmt.Sprintf(failMsg, "path must not have been executed")), "path must not have been executed"},
	}

	for _, test := range tests {
		testPanic(t, test.exec, test.want, test.msg)
	}
}
