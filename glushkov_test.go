package gr

import (
	"slices"
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
	automaton := op.RabinScott(nfa)

	if !automaton.Match(slices.Values([]rune("ab"))) {
		t.Fail()
	}

	if !automaton.Match(slices.Values([]rune("abab"))) {
		t.Fail()
	}

	if !automaton.Match(slices.Values([]rune("ac"))) {
		t.Fail()
	}

	if !automaton.Match(slices.Values([]rune(""))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("acab"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("abac"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("aba"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("aca"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("a"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("aa"))) {
		t.Fail()
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
	automaton := op.RabinScott(nfa)

	if automaton.Match(slices.Values([]rune("ab"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("abab"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("ac"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune(""))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("acab"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("abac"))) {
		t.Fail()
	}

	if !automaton.Match(slices.Values([]rune("aba"))) {
		t.Fail()
	}

	if !automaton.Match(slices.Values([]rune("aca"))) {
		t.Fail()
	}

	if !automaton.Match(slices.Values([]rune("a"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("aa"))) {
		t.Fail()
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
	automaton := op.RabinScott(nfa)

	if !automaton.Match(slices.Values([]rune(""))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("acab"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("abac"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("aba"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("aca"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("a"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("aa"))) {
		t.Fail()
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
	automaton := op.RabinScott(nfa)

	if automaton.Match(slices.Values([]rune(""))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("acab"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("abac"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("aba"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("aca"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("a"))) {
		t.Fail()
	}

	if automaton.Match(slices.Values([]rune("aa"))) {
		t.Fail()
	}
}
