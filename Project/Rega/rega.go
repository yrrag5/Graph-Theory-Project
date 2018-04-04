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

		// Adding ? state	
		case '?':
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

/////// Added on from nfa.go //////
	// Checks for one element in the stacks
	if len(nfastack) != 1 {
		fmt.Println("Error, the nfa isn't correct:", len(nfastack), nfastack)
	}

	return nfastack[0]
}

// Adds states that can be accepted using empty strings
func addState(l []*state, s *state, a *state) []*state {
	l = append(l,s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l 
}

// Checks to see if the pomatch reg expression matches with the string 
func pomatch(po string, s string) bool{
	ismatch := false
	ponfa := poregtonfa(po)
	// States in the current nfa
	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	// Goes through the hardcoded strings and the current state
	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)

			}
		}
		// Moves to the next state
		current, next = next, []*state{}
	}

	for _, c := range current {
		// if current state is the final state then accepts
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}

func main() {
	// Testing a serious a postmatch regurlar expressions 

	// Returns true
	fmt.Println(pomatch("ab.c*|", "ccc"))

	// Returns false
	fmt.Println(pomatch("ab.c*|", "abc"))

	// Returns true 
	fmt.Println(pomatch("ab.c*|", "c"))
}