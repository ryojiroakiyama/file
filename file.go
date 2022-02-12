// Package file provides basic functions for file operations.
package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//ToFile creates a fileName file and writes contents to the file.
//If successful, ToFile returns nil error.
//Else returns any error encountered.
func ToFile(fileName string, contents []byte) (err error) {
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
	return
}

//ToTmpFile creates a tmpprary file and write contents to the file.
//ToFile returns string which is name of the file created.
// any error encountered.
func ToTmpFile(src io.Reader) (fileName string, err error) {
	tmpfile, err := os.CreateTemp("", "")
	if err != nil {
		return "", fmt.Errorf("ToTmpFile: %v", err)
	}
	fileName = tmpfile.Name()
	//fmt.Println("create:", fileName)
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
	return
}

func FileBytes(fileName string) ([]byte, error) {
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
	return err
}
