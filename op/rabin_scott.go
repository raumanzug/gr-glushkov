package op

import (
	"github.com/raumanzug/gr-glushkov/glushkov"

	"github.com/raumanzug/gr-set/set"
	"github.com/raumanzug/gr-set/simple"
)

type rabinScottConstruction[TState comparable, TAction comparable] struct {
	nfa   glushkov.INFA[TState, TAction]
	state set.ISet[TState]
}

// RabinScott generates an object which can parse strings against the
// regular language related to it.
func RabinScott[TState comparable, TAction comparable](
	nfa glushkov.INFA[TState, TAction],
) glushkov.IAutomaton[TAction] {
	retval := rabinScottConstruction[TState, TAction]{
		nfa:   nfa,
		state: simple.NewSet[TState](),
	}
	retval.state.Add(nfa.Start())

	return &retval
}

func (pAutomaton *rabinScottConstruction[TState, TAction]) IsFinal() bool {
	return !pAutomaton.state.IsDisjoint(pAutomaton.nfa.Finals())
}

func (pAutomaton *rabinScottConstruction[TState, TAction]) Next(
	action TAction,
) {
	nextPowerState := simple.NewSet[TState]()
	for leftState := range pAutomaton.state.Generator() {
		nextPowerState.AddSet(
			pAutomaton.nfa.Next(leftState, action),
		)
	}

	pAutomaton.state = nextPowerState
}

func (pAutomaton *rabinScottConstruction[TState, TAction]) PermittedActions() set.ISet[TAction] {
	retval := simple.NewSet[TAction]()
	for leftState := range pAutomaton.state.Generator() {
		for action := range pAutomaton.nfa.Actions().Generator() {
			if !pAutomaton.nfa.Next(leftState, action).IsEmpty() {
				retval.Add(action)
			}
		}
	}

	return retval
}
