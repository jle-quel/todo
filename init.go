package main

import "fmt"
import "os"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func parse_initialize(argv []string) error {
	var err error

	if len(argv) != 0 {
		err = fmt.Errorf("usage: todo init")
	}

	return err
}

func create_repository(path string) error {
	var err error

	if err = os.Mkdir(path, 0700); err != nil {
		if os.IsExist(err) {
			err = nil
			fmt.Println("todo: reinitialized existing repository")
		} else {
			err = fmt.Errorf("todo: failed to create repository")
		}
	} else {
		fmt.Println("todo: initialized repository")
	}

	return err
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func initialize(argv []string, pwd string) error {
	var todo Todo
	var err error

	if err = parse_initialize(argv); err != nil {
		return err
	}
	if err = create_repository(pwd + "/.todo"); err != nil {
		return err
	}

	todo = make(Todo)
	todo["master"] = []Task{}

	if err = write_file(pwd+"/.todo/tasks", todo); err != nil {
		return err
	}
	if err = write_file(pwd+"/.todo/info", Info{"master", 0}); err != nil {
		return err
	}

	return nil

}
