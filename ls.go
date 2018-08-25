package main

import "fmt"
//import "os"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	parse_list(argv []string) error {
	var err	error

	if len(argv) != 0 {
		err = fmt.Errorf("usage: todo ls")
	}

	return err
}

func	list_task(todo Todo, info Info) error {
	var id	int

	id = 0

	for _, task := range todo[info.Branch] {
		if task.Status == false {
			fmt.Printf("%d → \033[0;31m%s\033[0m\n", id, task.Name)
		} else {
			fmt.Printf("%d → \033[0;32m%s\033[0m\n", id, task.Name)
		}
		
		id += 1
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

	if pwd, err = get_pwd(pwd); err != nil {
		return err
	}
	if todo, err = get_task(pwd + "/.todo/tasks"); err != nil {
		return err
	}
	if info, err = get_info(pwd + "/.todo/info"); err != nil {
		return err
	}

	if err = parse_list(argv); err != nil {
		return err
	}

	return list_task(todo, info)
}
