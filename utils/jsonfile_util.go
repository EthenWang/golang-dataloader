package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadJson(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteJson(data interface{}, filePath string) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err == nil {
		err = ioutil.WriteFile(filePath, file, 0644)
	}
	return err
}

func GetFileList(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	fileList := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileList = append(fileList, file.Name())
	}
	return fileList, err
}
