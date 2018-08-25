package main

import "fmt"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func parse_branch(argv []string) error {
	var err error

	if len(argv) < 1 {
		err = fmt.Errorf("usage: todo branch [arguments...]")
	}

	return err
}

func check_branch_exist(branch string, todo Todo) bool {
	for key, _ := range todo {
		if branch == key {
			return true
		}
	}

	return false
}

func append_branch(todo Todo, pwd string, argv []string) error {
	var index int
	var err error

	index = 0

	for index, _ = range argv {
		if check_branch_exist(argv[index], todo) == false {
			todo[argv[index]] = []Task{}

			if err = write_file(pwd+"/.todo/tasks", todo); err != nil {
				return err
			}
		}
	}
	if err = write_file(pwd+"/.todo/branch", argv[index]); err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func branch(argv []string, pwd string) error {
	var todo Todo
	var err error

	if err = parse_branch(argv); err != nil {
		return err
	}
	if pwd, err = get_pwd(pwd); err != nil {
		return err
	}

	if todo, err = get_task(pwd + "/.todo/tasks"); err != nil {
		return err
	}
	if err = append_branch(todo, pwd, argv); err != nil {
		return err
	}

	return nil
}
