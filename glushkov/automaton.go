package glushkov

import (
	"iter"
)

// IAutomaton is an interface for a type that parses strings.
// Match replies true iff the argument is sentence of the
// language related to this automaton.
type IAutomaton[TAction any] interface {
	Match(iter.Seq[TAction]) bool
}
