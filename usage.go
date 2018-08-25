package main

import "fmt"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func usage() {
	fmt.Println("Todo is the missing tool of Git. 🐜 ")
	fmt.Println()

	fmt.Println("Usage:\n\n\ttodo command [arguments]")
	fmt.Println()

	fmt.Println("The commands are:\n")
	fmt.Println("\tinit\tcreate an empty todo repository or reinitialize an existing one")
	fmt.Println("\tadd\tcreate and add task in the current branch")
	fmt.Println("\tbranch\tcreate, list or switch branch")
	fmt.Println("\tls\tlist tasks of current branch")
	fmt.Println()

	fmt.Println("Use \"todo help [command]\" for more information about a command.")
}
