// Graph Theory Project
// Author: Garry Cummins

package main

// Imports 
import(
	"fmt"
)

// State within the nondeterministic finite automata
type state struct {
	symbol rune
	edge1 *state
	edge2 *state
}

// NFA stuct is used to keep track of the intial and accept states
type nfa struct {
	initial *state
	accept *state
}

// Creates a nfa using pofix regurlar expression
func poregtonfa(pofix string) *nfa {
	// Pointers set to nfa
	nfastack := []*nfa{}

	// for loop goes through the pofix to remove a item from a stack
	// and push them 
	for _, r := range pofix {
		switch r {
		case '.':
			// Removes two values from the stack
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// Gives frag1 the intial state of frag2	
			frag1.accept.edge1 = frag2.initial
			// Appends the pointer
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			// Removes two values from the stack
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			// nfa's are both pointed to the intial states
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})


		case '*':
			// Only removes one frag due to * character only using one state 
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			// Initial  will point at the new intial and accept states
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		
		// Adding plus state
		case '+':	
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			// Initial  will point at the new intial and accept states
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			// Pushes the accept state to the new initial state
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}// Switch
	}// for

	return nfastack[0]
}

func main() {
	// Pofix expression ab.c*|
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}