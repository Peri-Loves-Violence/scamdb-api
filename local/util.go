package local

import (
	"io/ioutil"
)

func listDir(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return make([]string, 0), err
	}
	var names []string
	for _, f := range files {
		if f.IsDir() {
			names = append(names, f.Name())
		}
	}
	return names, nil
}

func listFiles(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return make([]string, 0), err
	}
	var names []string
	for _, f := range files {
		if !f.IsDir() {
			names = append(names, f.Name())
		}
	}
	return names, nil
}
