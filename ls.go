package main

import "fmt"
import "strings"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	print_task(index int, task Task) {
	if task.Status == false {
		fmt.Printf("\t%d | \033[0;31m✘\033[0m %s\n", index, task.Name)
	} else {
		fmt.Printf("\t%d | \033[0;32m✔︎\033[0m %s\n", index, task.Name)
	}
}

func	print_branch(branch string, info Info) {
	if branch == info.Branch {
		fmt.Printf("* \033[0;32m%s\033[0m\n", strings.ToUpper(branch))
	} else {
		fmt.Printf("  %s\n", strings.ToUpper(branch))
	}
}

func	list_task(todo Todo, info Info) error {
	var keys	[]string
	var length	int

	keys = get_keys(todo)
	length = len(keys) - 1

	for index, key := range keys {
		print_branch(key, info)
		for id, task := range todo[key] {
			print_task(id, task)
		}
		if index < length { 
			fmt.Println()
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	list(argv []string, pwd string) error {
	var todo	Todo
	var info	Info
	var err		error

	if len(argv) != 0 {
		return fmt.Errorf("usage: todo ls")
	}

	if pwd, err = get_pwd(pwd); err != nil {
		return err
	}
	if todo, err = get_task(pwd + "/.todo/tasks"); err != nil {
		return err
	}
	if info, err = get_info(pwd + "/.todo/info"); err != nil {
		return err
	}

	return list_task(todo, info)
}
