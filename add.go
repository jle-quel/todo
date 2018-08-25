package main

import "fmt"
import "os"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func parse_add(argv []string) error {
	var err error

	if len(argv) < 1 {
		err = fmt.Errorf("usage: todo add [arguments...]")
	}

	return err
}

func (self Todo) CheckTask(branch string, task string) bool {
	for index, _ := range self[branch] {
		if task == self[branch][index].Name {
			return false
		}
	}

	return true
}

func append_task(branch string, todo Todo, pwd string, argv []string) error {
	var err error

	for index, _ := range argv {
		if todo.CheckTask(branch, argv[index]) {
			todo[branch] = append(todo[branch], Task{argv[index], 0})

			if err = write_file(pwd, todo); err != nil {
				return err
			}
		} else {
			fmt.Fprintf(os.Stderr, "todo: task \"%s\" already exist\n", argv[index])
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func add(argv []string, pwd string) error {
	var branch string
	var todo Todo
	var err error

	if err = parse_add(argv); err != nil {
		return err
	}
	if pwd, err = get_pwd(pwd); err != nil {
		return err
	}

	if branch, err = get_branch(pwd + "/.todo/branch"); err != nil {
		return err
	}
	if todo, err = get_task(pwd + "/.todo/tasks"); err != nil {
		return err
	}
	if err = append_task(branch, todo, pwd+"/.todo/tasks", argv); err != nil {
		return err
	}

	return nil
}
