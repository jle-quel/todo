package	main

import "fmt"

var help_function = []HelpFunction {
	{"init", init_help},
	{"add", add_help},
	{"branch", branch_help},
	{"ls", ls_help},
	{"switch", switch_help},
	{"rm", rm_help},
}


////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	init_help() error {
	fmt.Println("usage: todo init")
	fmt.Println()

	fmt.Println("This command creates an empty todo repository")
	fmt.Println("It will contain all the information about your current project")
	fmt.Println("If the directory has already been initialized, todo init will be reinitialized")
	fmt.Println()
	fmt.Println("After initializing your directory, you will be able to add task/branches ...")
	fmt.Println()

	fmt.Println("See also: todo help add")
	return nil
}

func	add_help() error {
	fmt.Println("usage: todo add [arguments]")
	fmt.Println()

	fmt.Println("This command creates and adds as many tasks given to the repository")
	fmt.Println()
	fmt.Println("After adding some tasks, you will be able to list/modify them ...")
	fmt.Println()

	fmt.Println("See also: todo help ls")
	return nil
}

func	branch_help() error {
	fmt.Println("usage: todo branch [arguments]")
	fmt.Println()

	fmt.Println("This command, given with arguments will add branches to the repository")
	fmt.Println("If the branch already exists it will change the current branch to the latter")
	fmt.Println()
	fmt.Println("This same command, given without arguments will print all the branches of the repository")
	fmt.Println("The current branch will be started by a ★")
	fmt.Println()
	fmt.Println("After adding some branches you will be able to create tasks on specific branches")
	fmt.Println()

	fmt.Println("See also: todo help add")
	return nil
}

func	ls_help() error {
	fmt.Println("usage: todo ls")
	fmt.Println()

	fmt.Println("This command will print all the tasks on the actual branch")
	fmt.Println("If the task is completed it will be printed in \033[0;32mgreen\033[0m")
	fmt.Println()
	fmt.Println("The output is like this \"0 -> task\"")
	fmt.Println("The first positive integer is the id of the task")
	fmt.Println("You will use this id to switch the status of tasks")
	fmt.Println()

	fmt.Println("See also: todo help switch")
	return nil
}

func	switch_help() error {
	fmt.Println("usage: todo switch [arguments]")
	fmt.Println()

	fmt.Println("This command will modify the status of tasks")
	fmt.Println("A Task can be switch between todo and done")
	fmt.Println("After added some tasks and completed it, you will be able to list them and see the status")
	fmt.Println()

	fmt.Println("See also: todo help ls")
	return nil
}

func	rm_help() error {
	fmt.Println("usage: todo rm [arguments]")
	fmt.Println()

	fmt.Println("This command, given with arguments will remove tasks from the current branch")
	fmt.Println("This same command, given without arguments will delete the current branch")
	fmt.Println("You will be now switched by default on the master one")
	fmt.Println("The master branch is a const branch, you cannot erase it")
	fmt.Println()

	fmt.Println("See also: todo help init")
	return nil
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	help(argv []string, pwd string) error {
	if len(argv) != 1 {
		return fmt.Errorf("usage: todo help [argument]")
	}

	for index, _ := range help_function {
		if help_function[index].Name == argv[0] {
			return help_function[index].Call()
		}
	}

	return fmt.Errorf("todo: unknown help topic \"%s\"", argv[0])
}
