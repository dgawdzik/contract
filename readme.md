# Contracts

This is a GO library that adds support for design by contract. See [DbC](https://en.wikipedia.org/wiki/Design_by_contract) for introductory description.

## Theory

At the heart of Design by Contract is the idea that we choose to engage in economic activity because
of mutual benefits that we derive when interacting with each other. Relevant to this notion is the relationship between a client and a supplier. We can view each GO package or function as either
supplying a service or being a client of a package or function that supplies a service. A contract
defines obligations and establishing benefits when interactions between packages and functions take
place. This is similar, in many ways, to a business contract. As an example, a contract for a
cellular service can be established between you and a cellular provider. Under this contract, terms
clearly define obligations and benefits of each party. A cellular provider gets a benefit of your
money but must provide you with a cellular service. On the other hand, you get a benefit of a
service but must provide money in exchange. A contract clearly spells out benefits and obligations.
To put it in another way, a pre-condition, a condition that must be true, or is required to be true
before you can obtain a service, is the promise of money that you must pay for the service. The post-condition, the benefit that you get, or a condition that needs to be ensured by the supplier,
is the service that you obtain when you make a phone call.

Design by contract benefits not just a purely object oriented language but also a language that in
many ways is a hybrid between object oriented, procedural, and functional. One could argue that
functions are one of the major abstractions present in GO. We know that some functions are partial
and we need to know what is expected from us before calling them. More formally, we need to know
what is a pre-condition that we need to satisfy before making a call. With the help of this library,
one can clearly define such a requirement using `contract.Requires`. When we see that `contract.Requires` condition failed, a bug can quickly be identified as being in the calling code. In summary, a
function can only be called if pre-condition is satisfied; all bets are off if this is not the case. Pre-condition i.e. `contract.Requires` provides a benefit to the person implementing a function. It
makes the code simpler to implement as some possibilities are eliminated by `contract.Requires`. The implementation must only concern itself with the possibilities that are still open as defined by pre-condition.

After calling GO function, we need to know what is guaranteed by a function we just called.
This defines a benefit to the calling code. More formally we want to know what is guaranteed or
ensured by a function we just called. In this library, `contract.Ensures` expresses the
benefit we obtain from calling a function. If for some reason, there is a failure of the
`contract.Ensures`, we know the implementation of the function is incorrect as the code
does not live to its expectation. When this is the case, we can focus our effort on fixing the
function that promised but did not deliver.

`contract.Check` allows us to clearly define assumptions about our code that we believe to
be true at certain point of function execution. Were such assumption turn out to be incorrect, as
manifested by failure of `contract.Check`, we should go back and correct the code that was written
claiming these assumptions were true.

`contract.Fail` is useful when it is our understanding that certain portion of code should never be
executed or reached. If this proves not to be the case, we should re-examine the code and make
necessary corrections.

`contract.Invariant` is specified against private/package members of a structure and it defines one
or more conditions that are always true between executions of public functions/methods. If an
invariant turns out to be false at any point of program execution, this would indicate that package
implementation needs to be corrected. The invariant ensures that we change the structure from one
valid state into another valid state as defined by the set of conditioned expressed by invariant
function.


To summarize, contracts allow us to fail fast. We can
clearly express what is required before calling a function and what benefit we obtain. Finally,
failures of different types of contracts clearly give indication of which part of the code has bugs.
This makes the exercise of correcting them simpler. Testing, including, property based testing and
design by contract are trying to address our inability to implement formal proof for code
correctness. Pre-conditions and post-conditions are useful even in functions with no side effects,
as they limit input domains and output ranges making code easier to develop and reason about.

## Future Improvements
1. Use reflection to report which function in which package failed.

## Usage and Installation

The package can be added using **go get**:

```bash
go get https://github.com/dgawdzik/contract
```

or via your favorite dependency package manager such as **dep**:

```bash
dep ensure -add github/dgawdzik/contract
```

Add to source file via import:

```go
import "github.com/dgawdzik/contract"
```

## Contract failures

When condition provided to `contract.Requires`, `contract.Ensures`, `contract.Check`, and
`contract.Invariant`, turns out to be false, or when `contract.Fail` executes, panic is invoked and
provided with structure `contract.Exception`. The structure contains code representation of the
condition that failed followed by generic message related to type of contract that failed.
The message embeds the string parameter provided to the contract.

## Pre-conditions: contract.Requires
Multiple requires contract clauses with message parameter can be defined. The first parameter
specifies a condition that we expect to be true upon call to a function.

Example:

```go
func NewStdoutSplitWriter(writer io.Writer) {
	contract.Requires(writer != nil, "writer must be provided")
	contract.Requires(writer.Out != nil, "writer.Out buffer must be set")

    ...
}
```

## Post-conditions: contract.Ensures

The library allows to define multiple ensures conditions that can have message parameter. The first parameter specifies a condition that we expect to be true when function exits.

Example:

```go
func NewStdoutSplitWriter(writer io.Writer) *StdoutSplitWriter {
    ...

	contract.Ensures(result != nil, "new splitter must have been created")
	contract.Ensures(result.writer == writer, "writer must have been set")
	return result
}
```

## Check-condition: contract.Check

This contract can appear multiple times inside of function body. First parameter to `contract.Check`
specifies a condition that we expect to be true at certain point of function execution.

Example:

```go
func NewStdoutSplitWriter(writer io.Writer) *StdoutSplitWriter {
    ...
    value := someVar * 2
    ... 
	contract.Check(value > 0, "value must be positive")
    ...
}
```

## Fail-condition: contract.Fail

`contract.Fail` can appear multiple times inside of function body. The only parameter is a message
that should describe a reason for a failure.

Example:

```go
func translateLogLevel(level Level) logrus.Level {
	var result logrus.Level

	switch byte(level) {
	case DebugLevel:
		result = logrus.DebugLevel
	case InfoLevel:
		result = logrus.InfoLevel
	case WarnLevel:
		result = logrus.WarnLevel
	case ErrorLevel:
		result = logrus.ErrorLevel
	default:
		contract.Fail(fmt.Sprintf("unexpected value [%v] for Level type", level))
	}

	return result
}
```

## State-invariant: contract.Invariant

`contract.Invariant` can appear multiple times inside of function by convention named invariant. The
condition should check the state of a structure against which invariant is defined.

Example:

```go
// define state
type state struct {
	name  string
	value int
	obj   interface{}
}

// define state invariant
func (s *state) invariant() {
	contract.Invariant(s != nil, "state must be provided")
	contract.Invariant(!IsEmpty(s.name), "name must be set")
	contract.Invariant(s.value >= 0, "value must be positive")
	contract.Invariant(s.obj != nil, "obj must be set")
}

// check state invariant upon function entry
func (s *state) SomeFunction() {
    s.invariant()
    ...
}
```

# Credits
The inventor of Design by Contract [Bertrand Meyer](https://en.wikipedia.org/wiki/Bertrand_Meyer).
