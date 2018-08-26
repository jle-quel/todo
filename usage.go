package main

import "fmt"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func usage() {
	fmt.Println("Todo is a tool for managing your project. 🐜 ")
	fmt.Println()

	fmt.Println("Usage:\n\n\ttodo command [arguments]")
	fmt.Println()

	fmt.Println("The commands are:\n")
	fmt.Println("\tinit\tcreate an empty todo repository or reinitialize an existing one")
	fmt.Println("\tadd\tcreate and add task in the current branch")
	fmt.Println("\tbranch\tcreate, list or switch branch")
	fmt.Println("\tls\tlist tasks of the current branch")
	fmt.Println("\tswitch\tmodify the status of tasks in the current branch")
	fmt.Println("\trm\tdelete current branch and/or tasks of the current branch")
	fmt.Println()

	fmt.Println("Use \"todo help [command]\" for more information about a command.")
}
