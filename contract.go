package contract

import (
	"fmt"
)

//
// Constants
//

// Defines constants for Type of contract.
const (
	requires exType = iota
	ensures
	assert
	fail
	invariant
)

const (
	requiresMsg  = "Pre-condition violated. Invalid implementation of calling code given method pre-condition [%s]."
	ensuresMsg   = "Post-condition violated. Invalid implementation of method given post-condition [%s]."
	assertMsg    = "Assertion violated. Invalid assumption about state of computation given assert condition [%s]."
	failMsg      = "Fail condition triggered. Invalid program path executed with failed condition [%s]. "
	invariantMsg = "Invariant violated. Invalid state given invariant [%s]."
)

//
// Types
//

// Type specifies one of requries, ensures, assert, or fail contract types.
type exType int

// Exception contains info related to the type of contract failure.
type Exception struct {
	failed exType
	msg    string
}

//
// Package Functions
//

func create(failed exType, msg string) (result Exception) {
	result.failed = failed
	result.msg = msg

	return result
}

//
// API Functions
//

// Requires defines a requirement that needs to be satisfied by a caller to a method or function in
// terms of pre-condition. A call to a method that defines a pre-condition is only valid if the
// condition upon method entry is true. Condition being false upon method entry indicates a bug in
// the caller.
func Requires(condition bool, msg string) {
	if !condition {
		panic(create(requires, fmt.Sprintf(requiresMsg, msg)))
	}
}

// Ensures defines a benefit that a method or function that is being called guarantees. It
// establishes condition that is true upon exit that a client code can benefit from. Condition being
// false upon method, or function exit, indicates a bug in the implementation of the method or
// function.
func Ensures(condition bool, msg string) {
	if !condition {
		panic(create(ensures, fmt.Sprintf(ensuresMsg, msg)))
	}
}

// Assert defines a condition that is belived to be true at certain point of program execution.
// Assert condition being invalid, indicates that assumptions about program computation need to be
// revised.
func Assert(condition bool, msg string) {
	if !condition {
		panic(create(assert, fmt.Sprintf(assertMsg, msg)))
	}
}

// Fail statement indicates that it is belived that a program excecution should never reach a
// portion of the code that contains the Fail statement. Excecution of Fail statement indicates that
// a particular path in program execution is possible and understanding of program control flow
// needs to be revised.
func Fail(msg string) {
	panic(create(fail, fmt.Sprintf(failMsg, msg)))
}

// Invariant defines an invariant i.e. a condition that is true before every call to a public
// function or method and which is also true after every call to a public function or method
// completes. The invariant is most often check against state which is contained in Go using structs.
func Invariant(condition bool, msg string) {
	if !condition {
		panic(create(invariant, fmt.Sprintf(invariantMsg, msg)))
	}
}

//
// API Methods
//

func (ex Exception) Error() string {
	return ex.msg
}

// IsRequires returns true when exception was raised because of Requires contract failure, returns
// false otherwise.
func (ex Exception) IsRequires() bool {
	return ex.failed == requires
}

// IsEnsures returns true when exception was raised because of Ensures contract failure, returns
// false otherwise.
func (ex Exception) IsEnsures() bool {
	return ex.failed == ensures
}

// IsAssert returns true when exception was raised because of Assert contract failure, returns
// false otherwise.
func (ex Exception) IsAssert() bool {
	return ex.failed == assert
}

// IsFail returns true when exception was raised because of Fail contract failure, returns
// false otherwise.
func (ex Exception) IsFail() bool {
	return ex.failed == fail
}

// IsInvariant returns true when exception was raised because of Invariant contract failure, returns
// false otherwise.
func (ex Exception) IsInvariant() bool {
	return ex.failed == invariant
}
