package glushkov

import (
	"github.com/raumanzug/gr-set/set"
)

// IAutomaton is an interface for a type that parses strings
// of TAction instances.
//
// If IsFinal method returns the value “true”, this means
// that the parsed string is sentence in the regular language.
type IAutomaton[TAction any] interface {

	// IsFinal returns true iff a final state was reached.
	IsFinal() bool

	// The Next method switches the automaton to the next state
	// by consuming an action that is passed as an argument to
	// this method.
	Next(TAction)

	// PermittedActions returns all actions that can follow the
	// current state.
	PermittedActions() set.ISet[TAction]
}
