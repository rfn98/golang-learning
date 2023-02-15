package main

import (
    "path/filepath"
    "runtime"
    "fmt"
    "os"
    "io"
)

var (
    _, b, _, _ = runtime.Caller(0)
    basePath   = filepath.Dir(b)
    filePath = basePath + "/test.txt"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	var _, err = os.Stat(filePath)

	if err != nil {
		fmt.Println(err.Error())
	}

	if os.IsNotExist(err) {
		var file, err = os.Create(filePath)
		if isError(err) { return }
		defer file.Close()
	}

	fmt.Println("FILE HAS CREATED", filePath)
}

func writeFile() {
	var file, err = os.OpenFile(filePath, os.O_RDWR, 0644) //file dibuka dengan level akses read dan write dengan kode permission 0644
	if isError(err) { return }
	defer file.Close()

	// WRITE FILE
	_, err = file.WriteString("HALO\n")
	if isError(err) { return }
	_, err = file.WriteString("TEST WRITING FILE IN GOLANG.\n")
	if isError(err) { return }

	//SAVE CHANGE
	err = file.Sync()
	if isError(err) { return }

	fmt.Println("FILE HAS WRITTEN")
}

func readFile() {
	var file, err = os.OpenFile(filePath, os.O_RDONLY, 0644) //file dibuka dengan level akses readonly dengan kode permission 0644
	if isError(err) { return }
	defer file.Close()

	//READ FILE
	var text = make([]byte, 34) // ALOKASI ELEMEN 34
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) { break }
		}
		fmt.Println(n)
		if n == 0 {
			break
		}
	}
	if isError(err) { return }

	fmt.Println("FILE HAS READ")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(filePath)
	if isError(err) { return }

	fmt.Println("FILE HAS DELETED")
}

func main() {
	deleteFile()
    // readFile()
    // writeFile()
    // createFile()
}