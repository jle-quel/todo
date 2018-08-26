package	main

import "os"
import "fmt"
import "strconv"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	(self *Task) GetNewStatus() {
	if (*self).Status == true {
		(*self).Status = false
	} else {
		(*self).Status = true
	}
}

func	switch_task(argv []string, pwd string, todo Todo, info Info) error {
	var id	int
	var err	error

	for index, _ := range argv {
		if id, err = strconv.Atoi(argv[index]); err != nil {
			fmt.Fprintf(os.Stderr, "todo: id \"%s\" is not supported\n", argv[index])
		} else if id_exist(id, todo[info.Branch]) {
			todo[info.Branch][id].GetNewStatus()
		} else {
			fmt.Fprintf(os.Stderr, "todo: id \"%s\" does not exist\n", argv[index])
		}
	}

	return write_file(pwd + "/.todo/tasks", todo)
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	_switch(argv []string, pwd string) error {
	var todo	Todo
	var info	Info
	var err		error

	if len(argv) == 0 {
		return fmt.Errorf("usage: todo switch [arguments]")
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
	
	return switch_task(argv, pwd, todo, info)
}
