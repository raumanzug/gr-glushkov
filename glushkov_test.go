package gr

import (
	"testing"

	"github.com/raumanzug/gr-glushkov/op"

	"github.com/raumanzug/gr-set/simple"
)

func Test_A(t *testing.T) {
	actions := simple.NewSet[rune]()
	actions.Add('a')
	actions.Add('b')
	actions.Add('c')

	// (ab)*|ac
	a1 := op.ActionGlushkov('a')
	b1 := op.ActionGlushkov('b')
	a2 := op.ActionGlushkov('a')
	c2 := op.ActionGlushkov('c')
	op.ConcatGlushkov(&a1, b1)
	op.ConcatGlushkov(&a2, c2)
	op.KleeneStarGlushkov(&a1)
	glushkovData := op.UnionGlushkov(a1, a2)

	nfa := op.Glushkov(actions, glushkovData)

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "ab")

		if !isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('a')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "abab")

		if !isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('a')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "ac")

		if !isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "")

		if !isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('a')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "acab")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "abac")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aba")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('b')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aca")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "a")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('b')
		targetPermittedActions.Add('c')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aa")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}
}

func Test_B(t *testing.T) {
	actions := simple.NewSet[rune]()
	actions.Add('a')
	actions.Add('b')
	actions.Add('c')

	// ((a | b)(b | c))* a
	a1 := op.ActionGlushkov('a')
	b1 := op.ActionGlushkov('b')
	b2 := op.ActionGlushkov('b')
	c2 := op.ActionGlushkov('c')
	a3 := op.ActionGlushkov('a')
	glushkovData := op.UnionGlushkov(a1, b1)
	bc2 := op.UnionGlushkov(b2, c2)
	op.ConcatGlushkov(&glushkovData, bc2)
	op.KleeneStarGlushkov(&glushkovData)
	op.ConcatGlushkov(&glushkovData, a3)

	nfa := op.Glushkov(actions, glushkovData)

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "ab")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('a')
		targetPermittedActions.Add('b')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "abab")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('a')
		targetPermittedActions.Add('b')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "ac")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('a')
		targetPermittedActions.Add('b')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('a')
		targetPermittedActions.Add('b')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "acab")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('a')
		targetPermittedActions.Add('b')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "abac")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('a')
		targetPermittedActions.Add('b')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aba")

		if !isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('b')
		targetPermittedActions.Add('c')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aca")

		if !isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('b')
		targetPermittedActions.Add('c')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "a")

		if !isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()
		targetPermittedActions.Add('b')
		targetPermittedActions.Add('c')

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aa")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}
}

func Test_Epsilon(t *testing.T) {
	actions := simple.NewSet[rune]()
	actions.Add('a')
	actions.Add('b')
	actions.Add('c')

	// ()
	glushkovData := op.EpsilonGlushkov[rune]()

	nfa := op.Glushkov(actions, glushkovData)

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "")

		if !isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "acab")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "abac")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aba")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aca")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "a")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aa")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}
}

func Test_Zero(t *testing.T) {
	actions := simple.NewSet[rune]()
	actions.Add('a')
	actions.Add('b')
	actions.Add('c')

	// ()
	glushkovData := op.UnionGlushkov[rune]()

	nfa := op.Glushkov(actions, glushkovData)

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "acab")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "abac")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aba")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aca")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "a")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}

	{
		automaton := op.RabinScott(nfa)
		isFinal, permittedActions := parse(automaton, "aa")

		if isFinal {
			t.Fail()
		}

		targetPermittedActions := simple.NewSet[rune]()

		if !permittedActions.Eq(targetPermittedActions) {
			t.Fail()
		}
	}
}
