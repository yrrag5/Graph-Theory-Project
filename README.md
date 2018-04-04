# Graph-Theory-Project

Name - Garry Cummins

Student ID - G00335806

# NFA regurlar expressions

This program is based on the project for the module graph theory as part of the course Software Development in GMIT.It focuses on the go programming language to build NFA (Non-deterministic Finite Automata) using regurlar expressions (regex) and graph algorithms to match with a given string.

# Running the program

Before cloning the repository you fill first need to install go onto your machine. Once thats been achieved clone the repo using the git clone command in your cmd with the url. You can run each class using the (go run) command and class name (shunt.go, nfa.go, and rega.go) so long as your in the correct directory of each class in your cmd or terminal. Visual studio code is the most optimal way of viewing the code of the repo.

# How it works

Using two graph algorithms based on theory from the module - The Shunting Yard Algorithm, and The Thompson's Construction. We were given the task to code up these algorithmas to create NFAs with regurlar expressions.

The shunting yard translates the regex from infix to postfix to allow the program to understand the regex due to the lack of parentheses.

Thompsons construction creates the regex and uses structs (nfa and state) that when the regex are being read in, break up the NFA, then recreates it with extended parameters. Once thats complete, the string will then be read in with each character being tested. If it ends within a accept state, the string will be accepted and will return true. Otherwise it will return false. All strings used have been hardcoded for testing within the main function of regex.go.

