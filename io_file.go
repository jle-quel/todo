package main

import "os"
import "fmt"
import "io/ioutil"
import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	get_info(pwd string) (Info, error) {
	var data	[]byte
	var buf		Info
	var err		error

	if data, err = ioutil.ReadFile(pwd); err != nil {
		return Info{}, fmt.Errorf("todo: failed to read repository")
	}
	if err = json.Unmarshal(data, &buf); err != nil {
		return Info{}, fmt.Errorf("todo: failed to read repository")
	}

	return buf, nil
}

func	get_task(pwd string) (Todo, error) {
	var data	[]byte
	var buf		Todo
	var err		error

	if data, err = ioutil.ReadFile(pwd); err != nil {
		return nil, fmt.Errorf("todo: failed to read repository")
	}
	if err = json.Unmarshal(data, &buf); err != nil {
		return nil, fmt.Errorf("todo: failed to read repository")
	}

	return buf, nil
}

func	write_file(pwd string, value interface{}) error {
	var fd		*os.File
	var array	[]byte
	var err		error

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
