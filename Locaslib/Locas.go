package Locaslib

import (
	"fmt"
	"os"
)

func ReadWD() ([]os.DirEntry, error) { // Work dir type : []Entry
	file, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}
	return file, nil

}
func ReadWDSTR() ([]string, error) { // Work dir type : []String
	var listof_files []string
	file, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}
	for _, val := range file {
		listof_files = append(listof_files, val.Name())
	}
	return listof_files, nil
}
func TypeofFile() []string {
	var short []string
	X, _ := ReadWD()
	for _, val := range X {
		if val.IsDir() {
			short = append(short, fmt.Sprintf("|DIR : %s", val.Name()))
		} else {
			short = append(short, fmt.Sprintf("|File : %s", val.Name()))
		}

	}
	return short
}
