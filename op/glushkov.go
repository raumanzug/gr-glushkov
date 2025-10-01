package op

import (
	"github.com/raumanzug/gr-glushkov/glushkov"

	"github.com/raumanzug/gr-set/set"
	"github.com/raumanzug/gr-set/simple"
)

// Each action in a regular expression must be _linearized_, i.e. this action
// must be renamed, e.g. by adding an index, so that the action is unique in
// the regular expression. Use this type to do this. You can have multiple
// instances of this type for the same action in the body field. The
// state field does not contain any data on an instance of this type,
// only the identity.
type linearizedAction_t[TAction any] struct {
	state struct{}
	body  TAction
}

// glushkovEdge_t represents an edge in the non-deterministic finite state
// machine. For each linearized action, there is a state in the NFA generated
// by the Glushkov construction, and each state is either the initial state
// or a linearized action.
type glushkovEdge_t[TAction any] struct {
	from *linearizedAction_t[TAction]
	to   *linearizedAction_t[TAction]
}

// GlushkovData provides four parameters for the Glushkov construction.
// The Glushkov construction expects four parameters, which are obtained
// recursively from the regular expression for which the Glushkov
// construction generates the nfa.  This type contains these four parameters.
type GlushkovData[TAction any] struct {
	d set.ISet[*linearizedAction_t[TAction]]
	f set.ISet[*glushkovEdge_t[TAction]]
	l bool
	p set.ISet[*linearizedAction_t[TAction]]
}

// ActionGlushkov produces the four parameters from which the Glushkov
// construction generates a nfa that accepts a regular language
// consisting of one action.  This action is passed to this function
// as an argument.
func ActionGlushkov[TAction any](action TAction) (out GlushkovData[TAction]) {
	linearizedAction := linearizedAction_t[TAction]{
		body: action,
	}
	out.d = simple.NewSet[*linearizedAction_t[TAction]]()
	out.d.Add(&linearizedAction)
	out.f = simple.NewSet[*glushkovEdge_t[TAction]]()
	out.p = simple.NewSet[*linearizedAction_t[TAction]]()
	out.p.Add(&linearizedAction)

	return
}

// ConcatGlushkov produces the four parameters from which the Glushkov
// construction generates an nfa that accepts a regular language formed
// from the concatenation of two regular languages.  This function is
// called with the quadruples for these two regular languages from which the
// Glushkov construction would generate the corresponding nfa.  The function
// ConcatGlushkov overwrites the first parameter with the result.
func ConcatGlushkov[TAction any](
	pData *GlushkovData[TAction],
	other GlushkovData[TAction],
) {

	pData.f.AddSet(other.f)
	for d := range pData.d.Generator() {
		for p := range other.p.Generator() {
			glushkovEdge := glushkovEdge_t[TAction]{
				from: d,
				to:   p,
			}
			pData.f.Add(&glushkovEdge)
		}
	}

	if other.l {
		pData.d.AddSet(other.d)
	} else {
		pData.d = other.d
	}

	if pData.l {
		pData.p.AddSet(other.p)
	}

	pData.l = pData.l && other.l

	return
}

// EpsilonGlushkov produces the four parameters from which the Glushkov
// construction generates an nfa that accepts a regular language consisting
// only of an empty string.
func EpsilonGlushkov[TAction any]() (out GlushkovData[TAction]) {
	out.d = simple.NewSet[*linearizedAction_t[TAction]]()
	out.f = simple.NewSet[*glushkovEdge_t[TAction]]()
	out.l = true
	out.p = simple.NewSet[*linearizedAction_t[TAction]]()

	return
}

// KleeneStarGlushkov produces the four parameters from which the Glushkov
// construction generates an nfa that accepts the sentences of a regular
// language formed from the Kleene closure of another regular language.
// The four parameters from which the Glushkov construction would
// generate this other language are passed to this function as arguments.
// This argument is overwritten by the result of this function.
func KleeneStarGlushkov[TAction any](pData *GlushkovData[TAction]) {
	pData.l = true

	for d := range pData.d.Generator() {
		for p := range pData.p.Generator() {
			glushkovEdge := glushkovEdge_t[TAction]{
				from: d,
				to:   p,
			}
			pData.f.Add(&glushkovEdge)
		}
	}
}

// UnionGlushkov produces the four parameters from which the Glushkov
// construction generates an nfa that accepts the sentences of the
// regular language that forms the union of several regular languages.
// The quadruples from which the Glushkov construction would form these
// regular languages are passed to this function as arguments.
func UnionGlushkov[TAction any](
	dataList ...GlushkovData[TAction],
) (out GlushkovData[TAction]) {
	out.d = simple.NewSet[*linearizedAction_t[TAction]]()
	out.f = simple.NewSet[*glushkovEdge_t[TAction]]()
	out.l = false
	out.p = simple.NewSet[*linearizedAction_t[TAction]]()

	for _, data := range dataList {
		out.d.AddSet(data.d)
		out.f.AddSet(data.f)
		out.l = out.l || data.l
		out.p.AddSet(data.p)
	}

	return
}

// extractStates extract the state field from linearizedActions_t
// instances.
func extractStates[TAction any](
	linearizedActions set.ISet[*linearizedAction_t[TAction]],
) set.ISet[*struct{}] {

	retval := simple.NewSet[*struct{}]()

	for pLinearizedAction := range linearizedActions.Generator() {
		retval.Add(&pLinearizedAction.state)
	}

	return retval
}

// gatherStates gathers the states included in the quadrupels from which
// the Glushkov construction generates a nfa.
func gatherStates[TAction any](
	data GlushkovData[TAction],
) (states set.ISet[*struct{}]) {
	states = extractStates(data.d)

	for pEdge := range data.f.Generator() {
		states.Add(&pEdge.from.state)
		states.Add(&pEdge.to.state)
	}

	states.AddSet(extractStates(data.p))

	return
}

// Glushkov performs the Glushkov construction.  Their arguments are
// the alphabet, i.e. the set of actions, and the quadrupel recursively
// obtained by the regular expression.
func Glushkov[TAction comparable](
	actions set.ISet[TAction],
	data GlushkovData[TAction],
) glushkov.INFA[*struct{}, TAction] {

	retval := NewNFA[*struct{}, TAction](
		actions,
		extractStates(data.d),
		make(map[*struct{}]map[TAction]set.ISet[*struct{}]),
		&struct{}{},
		gatherStates(data),
	)
	retval.States().Add(retval.Start())

	if data.l {
		retval.Finals().Add(retval.Start())
	}

	for p := range data.p.Generator() {
		AddTransition(
			retval,
			retval.Start(),
			p.body,
			&p.state,
		)
	}

	for f := range data.f.Generator() {
		AddTransition(
			retval,
			&f.from.state,
			f.to.body,
			&f.to.state,
		)
	}

	return retval
}
