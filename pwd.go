package main

import "os"
import "fmt"
import "strings"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	repository_exist(pwd string) bool {
	var fd	*os.File
	var err	error

	if fd, err = os.OpenFile(pwd + "/.todo", os.O_RDONLY, 0600); err != nil {
		return false
	}
	if err = fd.Close(); err != nil {
		return false
	}

	return true
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	get_pwd(pwd string) (string, error) {
	var array	[]string
	var length	int

	array = strings.Split(pwd, "/")
	length = len(array) - 1

	for length > 0 {
		if repository_exist(pwd) {
			return pwd, nil
		}

		pwd = strings.TrimRight(pwd, "/" + array[length])
		length -= 1
	}

	return "", fmt.Errorf("todo: repository not found (or any of the parent directories)")
}
