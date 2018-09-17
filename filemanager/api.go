package filemanager

import (
	"errors"
	"os"
)

func WriteStringsToFile(input []string, fileName string) error {
	file, err := getOrCreateFile(fileName)
	if err != nil {
		return errors.New("writeStringToFileError: " + err.Error())
	}
	defer file.Close()
	for _, str := range input {
		_, err := file.WriteString(str)
		if err != nil {
			return errors.New("writeStringToFileError: " + err.Error())
		}
	}
	file.Sync()
	return nil
}

func getOrCreateFile(fileName string) (file *os.File, err error) {
	if _, err = os.Stat("files/" + fileName); os.IsNotExist(err) {
		file, err = os.Create("files/" + fileName)
	} else {
		file, err = os.Open("files/" + fileName)
	}
	if err != nil {
		return file, errors.New("GetOrCreateFileError: " + err.Error())
	}
	return file, nil
}
