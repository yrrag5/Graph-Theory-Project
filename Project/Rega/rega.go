// Graph Theory Project
// Author: Garry Cummins
// ID: G00335806

/* References: https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
			   http://www.perlmonks.org/?node_id=805819	
*/

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

func intopost(infix string) string {
	// The sequence order of characters 
	specials := map[rune]int{'*': 10, '.': 9, '|': 8, '+': 7, '?': 6}
	
	// Rune is used as a character
	pofix, s := []rune{}, []rune{}
    
	// The range is used to convert strings to a array of runes
	for _, r:= range infix {
		switch {
			// ( characters are added to the stack 
			case r == '(':
				s = append(s,r);
			// ) is added from the character stack to the postfix	
			case r == ')':
				for s[len(s)-1] != '(' {
					pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
				}
				s = s[:len(s)-1]
			case specials[r] > 0: // Characters after () will be added to the stack
				for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
					// Finds all elements bar the last 
					pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
				}
				s = append(s, r)
			default:
				pofix = append(pofix, r)
		}// switch
	}// for
	
	// Characters still left in the stack will be appended to postfix 
	// to clear the stack
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)
}

func main() {
	// Testing a series of postmatch regurlar expressions 
	fmt.Println()
	// ab.c
	fmt.Println("Using regex ab.c*|")
	
	fmt.Println("Input = ccc")
	// Returns true
	fmt.Println(pomatch("ab.c*|", "ccc"))
	
	fmt.Println("Input = abd")
	//Returns false
	fmt.Println(pomatch("ab.c*|", "abd"))
	fmt.Println()
    /////////////////////////////////////

	// bcd|.c*
	fmt.Println("Using regex bcd|.c*")

	fmt.Println("Input = bc")
	// Returns true
	fmt.Println(pomatch("bcd|.c*|", "bc"))

	
	fmt.Println("Input = ac")
	// Returns false
	fmt.Println(pomatch("bcd|.c*|", "ac"))
	fmt.Println()
	///////////////////////////////////////
	
	// xyz|.z*|
	fmt.Println("Using regex xyz|.z*|")

	fmt.Println("Input = x")
	// Returns false 
	fmt.Println(pomatch("xyz|.z*|", "x"))

	
	fmt.Println("Input = zzz")
	// Returns true
	fmt.Println(pomatch("xyz|.z*|", "zzz"))
	fmt.Println()
	///////////////////////////////////////

	// User Input
	fmt.Println("////// User Input //////") 
	fmt.Println()

	//Postfix
	rVal := intopost("fg.h*|")
	fmt.Println("Postfix value: " , rVal)
	
	var uInput string
	fmt.Println("Enter regular expression you wish to match with fgh*.|: ")
	fmt.Scanln(&uInput)

	mVal := pomatch(rVal, uInput)
	fmt.Println("The match is:", mVal)
	
}// Main