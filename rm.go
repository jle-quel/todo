package	main

import "os"
import "fmt"
import "strconv"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	remove_this_task(id int, task []Task) []Task {
	var result	[]Task

	for index, _ := range task {
		if id != index {
			result = append(result, task[index])
		}
	}

	return result
}

func	rm_task(argv []string, pwd string, todo Todo, info Info) error {
	var id	int
	var err	error

	for index, _ := range argv {
		if id, err = strconv.Atoi(argv[index]); err != nil {
			fmt.Fprintf(os.Stderr, "todo: id \"%s\" is not supported\n", argv[index])
		} else if id_exist(id, todo[info.Branch]) {
			todo[info.Branch] = remove_this_task(id, todo[info.Branch])
		} else {
			fmt.Fprintf(os.Stderr, "todo: id \"%s\" does not exist\n", argv[index])
		}
	}

	return write_file(pwd + "/.todo/tasks", todo)
}

func	rm_branch(pwd string, todo Todo, info Info) error {
	var err	error

	if info.Branch == "master" {
		return fmt.Errorf("todo: master branch cannot be erased")
	}

	delete(todo, info.Branch);

	if err = write_file(pwd + "/.todo/tasks", todo); err != nil {
		return err
	}

	info.Branch = "master"
	return write_file(pwd + "/.todo/info", info)
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	rm(argv []string, pwd string) error {
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
	
	if len(argv) == 0 {
		return rm_branch(pwd, todo, info)
	}

	return rm_task(argv, pwd, todo, info)
}
