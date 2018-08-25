package	main

import "os"
import "fmt"
import "strconv"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	id_exist(argv int, task []Task) bool {
	var length	int
	var result	bool

	length = len(task)
	result = false

	if argv >= 0 && argv < length {
		result = true
	}

	return result
}

func	get_id(argv int, task []Task) int {
	var id	int

	id = 0

	for id, _ = range task {
		if id == argv {
			break ;
		}
		id += 1
	}

	return id
}

func	parse_switch(argv []string, todo Todo, info Info) ([]int, error) {
	var tmp		int
	var result	[]int
	var err		error

	if len(argv) < 1 {
		return nil, fmt.Errorf("usage: todo switch [arguments...]")
	}

	for index, _ := range argv {
		if tmp, err = strconv.Atoi(argv[index]); err != nil {
			fmt.Fprintf(os.Stderr, "todo: id \"%s\" is not supported\n", argv[index])
		} else if id_exist(tmp, todo[info.Branch]) {
			result = append(result, tmp)
		} else {
			fmt.Fprintf(os.Stderr, "todo: id \"%s\" does not exist\n", argv[index])
		}
	}

	return result, nil
}

func	switch_task(argv []int, pwd string, todo Todo, info Info) error {
	var id		int
	var err	error

	for index, _ := range argv {
		id = get_id(argv[index], todo[info.Branch])

		if todo[info.Branch][id].Status == true {
			todo[info.Branch][id].Status = false
		} else {
			todo[info.Branch][id].Status = true
		}
	}

	if err = write_file(pwd + "/.todo/tasks", todo); err != nil {
		return err
	}

	return nil
}


////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	_switch(argv []string, pwd string) error {
	var todo	Todo
	var info	Info
	var array	[]int
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
	
	if array, err = parse_switch(argv, todo, info); err != nil {
		return err
	}

	return switch_task(array, pwd, todo, info)
}
