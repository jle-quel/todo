package main

import "fmt"
import "strings"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	print_branch(todo Todo, info Info) {
	var keys	[]string

	keys = get_keys(todo)

	for index, _ := range keys {
		if keys[index] == info.Branch {
			fmt.Printf("* \033[0;32m%s\033[0m\n", strings.ToUpper(keys[index]))
		} else {
			fmt.Printf("  %s\n", strings.ToUpper(keys[index]))
		}
	}
}

func	branch_exist(branch string, todo Todo) bool {
	var result	bool

	result = false

	for key, _ := range todo {
		if branch == key {
			result = true
		}
	}

	return result
}

func	append_branch(info Info, todo Todo, pwd string, argv []string) error {
	var index	int
	var err		error

	index = 0

	for index, _ = range argv {
		if branch_exist(argv[index], todo) == false {
			todo[argv[index]] = []Task{}
			if err = write_file(pwd+"/.todo/tasks", todo); err != nil {
				return err
			}
		}
	}

	info.Branch = argv[index]
	if err = write_file(pwd + "/.todo/info", info); err != nil {
		return err
	}

	return nil
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
		return append_branch(info, todo, pwd, argv)
	} else {
		print_branch(todo, info)
	}

	return nil
}
