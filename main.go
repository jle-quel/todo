package main

import "fmt"
import "os"

var state_machine = []StateMachine{
	{"init", _init},
	{"add", add},
	{"branch", branch},
	{"ls", list},
	{"switch", _switch},
}

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	core() error {
	var pwd	string

	if pwd = os.Getenv("PWD"); pwd == "" {
		return fmt.Errorf("todo: $PWD must be initialized")
	}

	for index, _ := range state_machine {
		if state_machine[index].Name == os.Args[1] {
			return state_machine[index].Function(os.Args[2:], pwd)
		}
	}

	return fmt.Errorf("todo: unknown subcommand \"%s\"", os.Args[1])
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	main() {
	var err	error

	if len(os.Args) < 2 {
		usage()
	} else if err = core(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}
