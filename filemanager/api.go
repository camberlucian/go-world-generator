package filemanager

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

func WriteStringsToFile(input *[]string, fileName string) error {
	file, err := GetOrCreateFile(fileName, true)
	if err != nil {
		return errors.New("writeStringToFileError: " + err.Error())
	}
	defer file.Close()
	for _, str := range *input {
		_, err := file.WriteString(str)
		if err != nil {
			return errors.New("writeStringToFileError: " + err.Error())
		}
	}
	file.Sync()
	return nil
}

func GetOrCreateFile(fileName string, overwrite bool) (file *os.File, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return file, errors.New("GetOrCreateFileError: " + err.Error())
	}
	path := filepath.Join(wd, "filemanager", "files", fileName)
	fmt.Println("GETTING OR CREATING FILE AT: " + path)

	if _, err = os.Stat(path); os.IsNotExist(err) {
		fmt.Println("File Doesnt Exist")
		file, err = os.Create(path)
		if err != nil {
			return file, errors.New("GetOrCreateFileError: " + err.Error())
		}
		fmt.Println("CREATED")
	} else if overwrite {
		fmt.Println("REMOVING FILE")
		err := os.Remove(path)
		if err != nil {
			return file, errors.New("GetOrCreateFileError: " + err.Error())
		}
		file, err = os.Create(path)
		if err != nil {
			return file, errors.New("GetOrCreateFileError: " + err.Error())
		}
	} else {
		fmt.Println("File Exists")
		file, err = os.OpenFile(path, os.O_RDWR, 0644)
		if err != nil {
			return file, errors.New("GetOrCreateFileError: " + err.Error())
		}
	}

	fmt.Println("Returning File")
	return file, nil
}

func ExportImage(m image.Image, fileName string) error {
	file, err := GetOrCreateFile(fileName, true)
	if err != nil {
		fmt.Println("Could not get or create file: " + err.Error())
	}
	err = png.Encode(file, m)
	if err != nil {
		fmt.Println("COULD NOT WRITE IMAGE: " + err.Error())
	}
	return err
}
