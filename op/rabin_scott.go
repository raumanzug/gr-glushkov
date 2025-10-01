package op

import (
	"iter"

	"github.com/raumanzug/gr-glushkov/glushkov"

	"github.com/raumanzug/gr-set/set"
	"github.com/raumanzug/gr-set/simple"
)

type rabinScottConstruction[TState comparable, TAction any] struct {
	nfa glushkov.INFA[TState, TAction]
}

// RabinScott generates an object which can parse strings against the
// regular language related to it.
func RabinScott[TState comparable, TAction any](
	nfa glushkov.INFA[TState, TAction],
) glushkov.IAutomaton[TAction] {
	retval := rabinScottConstruction[TState, TAction]{
		nfa: nfa,
	}

	return &retval
}

type powerState[TState any] struct {
	body set.ISet[TState]
}

func (pAutomaton *rabinScottConstruction[TState, TAction]) Match(
	in iter.Seq[TAction],
) bool {
	currentPowerstate := powerState[TState]{
		body: simple.NewSet[TState](),
	}
	currentPowerstate.body.Add(pAutomaton.nfa.Start())

	for action := range in {
		_currentStateBody := simple.NewSet[TState]()
		for leftState := range currentPowerstate.body.Generator() {

			_currentStateBody.AddSet(
				pAutomaton.nfa.Next(
					leftState,
					action,
				),
			)

		}
		currentPowerstate.body = _currentStateBody
	}

	return !currentPowerstate.body.IsDisjoint(
		pAutomaton.nfa.Finals(),
	)
}
