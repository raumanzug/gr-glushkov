package glushkov

import (
	"github.com/raumanzug/gr-set/set"
)

// INFA is interface for non-deterministic finite automatons (nfa).
// A nfa is defined by an _alphabet_ of actions, a set of
// states, a subset of this set of _final states_, a start state
// and a _transition relation_ between two states and an action.
type INFA[TState any, TAction any] interface {
	Actions() set.ISet[TAction]
	Finals() set.ISet[TState]
	Next(TState, TAction) set.ISet[TState]
	States() set.ISet[TState]
	Start() TState
}
