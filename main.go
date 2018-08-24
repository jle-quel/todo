package main

import "fmt"
import "os"

var state_machine = []StateMachine{
	{"init", initialize},
}

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func (self *Environ) Get() {
	if (*self).Pwd = os.Getenv("PWD"); (*self).Pwd == "" {
		fmt.Fprintf(os.Stderr, "todo: $PWD must be set")
		os.Exit(2)
	}
	if (*self).Home = os.Getenv("HOME"); (*self).Home == "" {
		fmt.Fprintf(os.Stderr, "todo: $HOME must be set")
		os.Exit(2)
	}

	(*self).Argv = os.Args[1:]
}

func usage() {
	fmt.Println("Todo is the missing tool of Git.")
	fmt.Println()
	
	fmt.Println("Usage:\n\n\ttodo command [arguments]")
	fmt.Println()

	fmt.Println("The commands are:\n")
	fmt.Println("\tinit\tcreate an empty todo repository or reinitialize an existing one")
	fmt.Println()

	fmt.Println("Use \"todo help [command]\" for more information about a command.")
}

func core() error {
	var environ Environ

	environ.Get()

	for index, _ := range state_machine {
		if state_machine[index].Name == environ.Argv[0] {
			return state_machine[index].Function(environ)
		}
	}

	return fmt.Errorf("todo: unknown subcommand \"%s\"", environ.Argv[0])
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
