package main

import "os"
import "fmt"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	task_exist(argv string, task []Task) bool {
	var result	bool

	result = false

	for index, _ := range task {
		if task[index].Name == argv {
			result = true
		}
	}

	return result
}

func	parse_add(argv []string, todo Todo, info Info) ([]string, error) {
	var result	[]string

	if len(argv) < 1 {
		return nil, fmt.Errorf("usage: todo add [arguments...]")
	}
	if len(todo[info.Branch]) == 0 {
		return argv, nil
	}
	
	for index, _ := range argv {
		if task_exist(argv[index], todo[info.Branch]) {
			fmt.Fprintf(os.Stderr, "todo: task \"%s\" already exist\n", argv[index])
		} else {
			result = append(result, argv[index])
		}
	}

	return result, nil
}

func	append_task(argv []string, pwd string, todo Todo, info Info) error {
	var err	error

	for index, _ := range argv {
		todo[info.Branch] = append(todo[info.Branch], Task{argv[index], false})
	}
	
	if err = write_file(pwd + "/.todo/tasks", todo); err != nil {
		return err
	}
	if err = write_file(pwd + "/.todo/info", info); err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	add(argv []string, pwd string) error {
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
	
	if argv, err = parse_add(argv, todo, info); err != nil {
		return err
	}

	return append_task(argv, pwd, todo, info)
}
