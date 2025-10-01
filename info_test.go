package gr

import (
	"fmt"
	"slices"

	"github.com/raumanzug/gr-glushkov/glushkov"
	"github.com/raumanzug/gr-glushkov/op"

	"github.com/raumanzug/gr-set/simple"
)

// compile an automaton which we can use to parse against
// regular expression (ab)*|ac
func compileRegex() glushkov.IAutomaton[rune] {

	// defines the alphabet consisting in three letters a, b and c.
	actions := simple.NewSet[rune]()
	actions.Add('a')
	actions.Add('b')
	actions.Add('c')
	// actions now contains the alphabet.

	// determine the parameters for feeding Glushkov construction
	// procedure in order to generate a nfa which can parse against
	// regular expression (ab)*|ac
	a1 := op.ActionGlushkov('a')             // regex: a
	b1 := op.ActionGlushkov('b')             // regex: b
	a2 := op.ActionGlushkov('a')             // do not recycle a1 here!
	c2 := op.ActionGlushkov('c')             // regex: c
	op.ConcatGlushkov(&a1, b1)               // regex: ab
	op.ConcatGlushkov(&a2, c2)               // regex: ac
	op.KleeneStarGlushkov(&a1)               // regex: (ab)*
	glushkovData := op.UnionGlushkov(a1, a2) // regex: (ab)*|ac
	// Now, we acquired food for Glushkov construction.

	// perform Glushkov construction
	nfa := op.Glushkov(actions, glushkovData)
	automaton := op.RabinScott(nfa)
	// Now, we acquired an automaton which can parse against
	// regular expression (ab)*|ac.

	return automaton
}

// Parses the sentence string using the automaton specified in the
// automaton argument.
func parse(automaton glushkov.IAutomaton[rune], sentence string) bool {
	return automaton.Match(slices.Values([]rune(sentence)))
}

func Example() {

	// compile regular expression (ab)*|ac
	automaton := compileRegex()

	// parse a sentence using this automaton:
	fmt.Println(parse(automaton, "abab"))
	// Output: true
}
