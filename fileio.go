// Package fileio provides basic functions for file operations.
package fileio

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//GenFile creates a fileName file and writes contents to the file.
//If successful, GenFile returns nil error.
//Else if faulse, GenFile returns any error encountered.
//Maybe same as os.WriteFile(fileName, contents, 066).
func GenFile(fileName string, contents []byte) (err error) {
	dstFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("ToFile: %v", err)
	}
	defer func() {
		if cerr := dstFile.Close(); cerr != nil {
			err = fmt.Errorf("ToFile: %v", cerr)
		}
		if err != nil && dstFile != nil {
			os.Remove(dstFile.Name())
		}
	}()
	_, werr := dstFile.Write(contents)
	if werr != nil {
		return fmt.Errorf("ToFile: %v", werr)
	}
	err = dstFile.Sync()
	return
}

//GenTmpFile creates a tmpprary file and write contents to the file.
//If successful, GenTmpFile returns a string
//which is the name of the created file and nil error.
//Else if faulse, GenTmpFile returns a empty string and any error encountered.
func GenTmpFile(src io.Reader) (fileName string, err error) {
	tmpfile, err := os.CreateTemp("", "")
	if err != nil {
		return "", fmt.Errorf("ToTmpFile: %v", err)
	}
	fileName = tmpfile.Name()
	defer func() {
		if cerr := tmpfile.Close(); cerr != nil {
			err = fmt.Errorf("ToTmpFile: %v", cerr)
		}
		if err != nil && fileName != "" {
			os.Remove(fileName)
		}
	}()
	_, err = io.Copy(tmpfile, src)
	if err != nil {
		return "", fmt.Errorf("ToTmpFile: %v", err)
	}
	err = tmpfile.Sync()
	return
}

//FileContents returns a byte array readed
//from the file named fileName or any error encountered.
func FileContents(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("FileBytes: %v", err)
	}
	defer file.Close()
	srcBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("FileBytes: %v", err)
	}
	return srcBytes, nil
}

//BindFiles binds contents of files stocked
//in srcNames to the file named dstName or any error encountered.
//The order to bind is the order of the slice index.
func BindFiles(srcNames []string, dstName string) (err error) {
	dstfile, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("BindFiles: %v", err)
	}
	defer func() {
		if cerr := dstfile.Close(); cerr != nil {
			err = fmt.Errorf("BindFiles: %v", cerr)
		}
		if err != nil {
			os.Remove(dstName)
		}
	}()
	for _, srcName := range srcNames {
		srcfile, err := os.Open(srcName)
		if err != nil {
			return fmt.Errorf("BindFiles: %v", err)
		}
		defer srcfile.Close()
		_, err = io.Copy(dstfile, srcfile)
		if err != nil {
			return fmt.Errorf("BindFiles: %v", err)
		}
	}
	err = dstfile.Sync()
	return
}
