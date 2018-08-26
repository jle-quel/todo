package main

import "os"
import "fmt"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	append_task(argv []string, pwd string, todo Todo, info Info) error {
	for index, _ := range argv {
		if task_exist(argv[index], todo[info.Branch]) {
			fmt.Fprintf(os.Stderr, "todo: task \"%s\" already exist\n", argv[index])
		} else {
			todo[info.Branch] = append(todo[info.Branch], Task{argv[index], false})
		}
	}

	return write_file(pwd + "/.todo/tasks", todo)
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	add(argv []string, pwd string) error {
	var todo	Todo
	var info	Info
	var err		error

	if len(argv) == 0 {
		return fmt.Errorf("usage: todo add [arguments...]")
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
	
	return append_task(argv, pwd, todo, info)
}
