// The _Glushkov construction_ generates the non-deterministic finite
// automaton (nfa) from a regular expression, which can decide whether a string
// is a sentence of the language defined by this regular expression or not.
//
// This module implements these constructions.  In addition, it provides
// a parser for strings that uses the constructed non-deterministic finite
// automaton to check whether strings are sentences in the language defined
// by a regular expression.
package gr
