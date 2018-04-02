// Graph Theory Project 
// Author: Garry Cummins

package main

// Imports
import(
	"fmt"
)
/* The shunting algorithim implements a expression and converts it from a 
   infix notion to a postfix 
*/   

func intopost(infix string) string {
	// The sequence order of characters 
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	
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
	// Checks for infix and postfix inputs
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: " + intopost("a.b.c*"))

	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: " + intopost("(a(b|d))*"))

	fmt.Println("Infix: ", "a.(b|d)).c*")
	fmt.Println("Postfix: " + intopost("a(b|d)).c*"))

	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: " + intopost("a.(b.b)+.c"))

	
}