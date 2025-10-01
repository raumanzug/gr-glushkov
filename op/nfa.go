package op

import (
	"github.com/raumanzug/gr-glushkov/glushkov"

	"github.com/raumanzug/gr-set/set"
	"github.com/raumanzug/gr-set/simple"
)

// NFA implements the INFA interface
type NFA[TState comparable, TAction comparable] struct {
	actions set.ISet[TAction]
	finals  set.ISet[TState]
	next    map[TState]map[TAction]set.ISet[TState]
	states  set.ISet[TState]
	start   TState
}

// NewNFA generates a NFA object.
func NewNFA[TState comparable, TAction comparable](
	actions set.ISet[TAction],
	finals set.ISet[TState],
	next map[TState]map[TAction]set.ISet[TState],
	start TState,
	states set.ISet[TState],
) glushkov.INFA[TState, TAction] {
	retval := NFA[TState, TAction]{
		actions: actions,
		finals:  finals,
		next:    next,
		states:  states,
		start:   start,
	}
	return &retval
}

func (nfa NFA[TState, TAction]) Actions() set.ISet[TAction] {
	return nfa.actions
}

func (nfa NFA[TState, TAction]) Finals() set.ISet[TState] {
	return nfa.finals
}

func (pNfa *NFA[TState, TAction]) Next(left TState, action TAction) set.ISet[TState] {
	var transitions map[TAction]set.ISet[TState]
	{
		var ok bool
		transitions, ok = pNfa.next[left]
		if !ok {
			transitions = make(map[TAction]set.ISet[TState])
			pNfa.next[left] = transitions
		}
	}

	var nextStates set.ISet[TState]
	{
		var ok bool
		nextStates, ok = transitions[action]
		if !ok {
			nextStates = simple.NewSet[TState]()
			transitions[action] = nextStates
		}
	}

	return nextStates
}

func (nfa NFA[TState, TAction]) States() set.ISet[TState] {
	return nfa.states
}

func (nfa NFA[TState, TAction]) Start() TState {
	return nfa.start
}

// AddTransition adds a transition consisting in two states and an action
// taken from the alphabet.
func AddTransition[TState any, TAction any](
	nfa glushkov.INFA[TState, TAction],
	from TState,
	action TAction,
	to TState,
) {
	nextStates := nfa.Next(from, action)
	nextStates.Add(to)
}
