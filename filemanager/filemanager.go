package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

/*
After creating filemanager package - using without pointer due to only reading the data - if you want to write
better we can use pointer to get address and chage the value by using derefence using '&variableName' symbol.
*/
func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println("Could not open file")
		fmt.Print(err)
		return nil, errors.New("fail to open file")
	}

	scanner := bufio.NewScanner(file) // read line by line from file
	var lines []string
	for scanner.Scan() {
		// step by step read by scan form file
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("reading the file content failed")
	}
	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file ")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}
	file.Close()
	return nil
}

func New(inputPath string, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

/*
// Befor creating or introduce struct in this filemanger package

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Could not open file")
		fmt.Print(err)
		return nil, errors.New("fail to open file")
	}

	scanner := bufio.NewScanner(file) // read line by line from file
	var lines []string
	for scanner.Scan() {
		// step by step read by scan form file
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("reading the file content failed")
	}
	file.Close()
	return lines, nil
}

func WritJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file ")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}
	file.Close()
	return nil
}
*/
