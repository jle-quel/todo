package main

import "fmt"
import "strings"
import "os"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func print_branch(active_branch string, branch string) {
	if active_branch == branch {
		fmt.Printf("%s ★\n", strings.ToUpper(branch))
	} else {
		fmt.Printf("%s \n", strings.ToUpper(branch))
	}
}

func print_tasks(task []Task) {
	for index, _ := range task {
		switch task[index].Status {
		case 0:
			fmt.Printf("\t[%d] \033[0;31m%s\033[0m\n", task[index].Id, task[index].Name)
		case 1:
			fmt.Printf("\t[%d] \033[0;33m%s\033[0m\n", task[index].Id, task[index].Name)
		case 2:
			fmt.Printf("\t[%d] \033[0;32m%s\033[0m\n", task[index].Id, task[index].Name)
		}
	}
}

func print_all(todo Todo, info Info, argv []string) {
	for key, _ := range todo {
		if len(todo[key]) != 0 {
			print_branch(info.Branch, key)
			print_tasks(todo[key])
		}
	}
}

func print_some(todo Todo, info Info, argv []string) {
	for index, _ := range argv {
		if _, exist := todo[argv[index]]; exist == true {
			print_branch(info.Branch, argv[index])
			print_tasks(todo[argv[index]])
		} else {
			fmt.Fprintf(os.Stderr, "todo: branch \"%s\" doest not exist\n", argv[index])
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func list(argv []string, pwd string) error {
	var todo Todo
	var info Info
	var err error

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
		print_all(todo, info, argv)
	} else {
		print_some(todo, info, argv)
	}

	return nil
}
