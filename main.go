package main

import "fmt"
import "os"

var state_machine = []StateMachine{
	{"init", initialize},
	{"add", add},
	{"branch", branch},
}

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func usage() {
	fmt.Println("Todo is the missing tool of Git.")
	fmt.Println()

	fmt.Println("Usage:\n\n\ttodo command [arguments]")
	fmt.Println()

	fmt.Println("The commands are:\n")
	fmt.Println("\tinit\tcreate an empty todo repository or reinitialize an existing one")
	fmt.Println("\tadd\tcreate and add task in the current branch")
	fmt.Println("\tbranch\tcreate or switch to specific branch")
	fmt.Println()

	fmt.Println("Use \"todo help [command]\" for more information about a command.")
}

func core() error {
	var pwd string

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

func main() {
	var err error

	if len(os.Args) < 2 {
		usage()
	} else if err = core(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}
