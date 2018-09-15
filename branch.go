package main

import "fmt"
import "strings"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	print_branchs(todo Todo, info Info) error {
	var keys	[]string

	keys = get_keys(todo)

	for index, _ := range keys {
		if keys[index] == info.Branch {
			fmt.Printf("* \033[0;32m%s\033[0m\n", strings.ToUpper(keys[index]))
		} else {
			fmt.Printf("  %s\n", strings.ToUpper(keys[index]))
		}
	}

	return nil
}

func	append_branch(argv []string, pwd string, todo Todo, info Info) error {
	var branch	string
	var err		error

	for index, _ := range argv {
		branch = strings.ToLower(argv[index])
		if branch_exist(branch, todo) == false {
			todo[branch] = []Task{}
		}
		info.Branch = branch
	}

	if err = write_file(pwd + "/.todo/tasks", todo); err != nil {
		return err
	}

	return write_file(pwd + "/.todo/info", info)
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	branch(argv []string, pwd string) error {
	var todo	Todo
	var info	Info
	var err		error

	if pwd, err = get_pwd(pwd); err != nil {
		return err
	}
	if todo, err = get_task(pwd + "/.todo/tasks"); err != nil {
		return err
	}
	if info, err = get_info(pwd + "/.todo/info"); err != nil {
		return err
	}

	if len(argv) > 0 {
		return append_branch(argv, pwd, todo, info)
	}

	return print_branchs(todo, info)
}
