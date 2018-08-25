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

func append_task(info Info, todo Todo, pwd string, argv []string) error {
	var err error

	for index, _ := range argv {
		if todo.CheckTask(info.Branch, argv[index]) {
			todo[info.Branch] = append(todo[info.Branch], Task{argv[index], 0, info.Id})
			info.Id += 1

			if err = write_file(pwd+"/.todo/tasks", todo); err != nil {
				return err
			}
			if err = write_file(pwd+"/.todo/info", info); err != nil {
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
	var todo Todo
	var info Info
	var err error

	if err = parse_add(argv); err != nil {
		return err
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
	if err = append_task(info, todo, pwd, argv); err != nil {
		return err
	}

	return nil
}
