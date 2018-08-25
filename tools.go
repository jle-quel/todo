package main

import "io/ioutil"
import "os"
import "fmt"
import "encoding/json"
import "strings"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func repository_exist(pwd string) bool {
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

func get_pwd(pwd string) (string, error) {
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

func get_branch(pwd string) (string, error) {
	var data []byte
	var buf string
	var err error

	if data, err = ioutil.ReadFile(pwd); err != nil {
		if os.IsExist(err) {
			return "", fmt.Errorf("todo: failed to read repository")
		} else {
			return "", fmt.Errorf("todo: repository has not been initialized")
		}
	}
	if err = json.Unmarshal(data, &buf); err != nil {
		return "", fmt.Errorf("todo: failed to read repository")
	}

	return buf, nil
}

func get_task(pwd string) (Todo, error) {
	var data []byte
	var buf Todo
	var err error

	if data, err = ioutil.ReadFile(pwd); err != nil {
		if os.IsExist(err) {
			return nil, fmt.Errorf("todo: failed to read repository")
		} else {
			return nil, fmt.Errorf("todo: repository has not been initialized")
		}
	}
	if err = json.Unmarshal(data, &buf); err != nil {
		return nil, fmt.Errorf("todo: failed to read repository")
	}

	return buf, nil
}

func write_file(pwd string, value interface{}) error {
	var fd *os.File
	var array []byte
	var err error

	if fd, err = os.OpenFile(pwd, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600); err != nil {
		return fmt.Errorf("todo: failed to write in repository")
	}
	if array, err = json.Marshal(value); err != nil {
		return fmt.Errorf("todo: failed to write in repository")
	}
	if _, err = fd.Write(array); err != nil {
		return fmt.Errorf("todo: failed to write in repository")
	}
	if err = fd.Close(); err != nil {
		return fmt.Errorf("todo: failed to write in repository")
	}

	return nil
}
