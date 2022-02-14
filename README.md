# file
ファイル操作をまとめたパッケージです。
Package file provides basic functions for file operations.

FUNCTIONS

func BindFiles(srcNames []string, dstName string) (err error)
    BindFiles binds contents of files stocked in srcNames to the file named
    dstName or any error encountered. The order to bind is the order of the
    slice index.

func FileBytes(fileName string) ([]byte, error)
    FileBytes returns a byte array readed from the file named fileName or any
    error encountered.

func ToFile(fileName string, contents []byte) (err error)
    ToFile creates a fileName file and writes contents to the file. If
    successful, ToFile returns nil error. Else if faulse, ToFile returns any
    error encountered. Maybe same as os.WriteFile(fileName, contents, 066).

func ToTmpFile(src io.Reader) (fileName string, err error)
    ToTmpFile creates a tmpprary file and write contents to the file. If
    successful, ToTmpFile returns a string which is the name of the created file
    and nil error. Else if faulse, ToTmpFile returns a empty string and any
    error encountered.
