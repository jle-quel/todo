package	main

import "sort"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	get_keys(todo Todo) []string {
	var result	[]string

	for key, _ := range todo {
		result = append(result, key)
	}

	sort.Strings(result)

	return result 
}

func	id_exist(argv int, task []Task) bool {
	var result	bool

	result = false

	for index, _ := range task {
		if index == argv {
			result = true
		}
	}
	
	return result
}

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

func	branch_exist(branch string, todo Todo) bool {
	var result	bool

	result = false

	for key, _ := range todo {
		if branch == key {
			result = true
		}
	}

	return result
}

