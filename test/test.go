package main

import (
	"fmt"
	//"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Upload(relativePath string, fileContent []byte) {
	//storagePath, _ := filepath.Abs("/usr/local/gbops///" + relativePath)
	storagePath, _ := filepath.Abs("D:\\gotest\\grpc_my\\test\\as\\" + relativePath)
	dir, _ := filepath.Split(storagePath)
	if !isDirExists(dir) {
		os.MkdirAll(dir, 0777)
	}
	fmt.Println(storagePath)
	err := ioutil.WriteFile(storagePath, fileContent, 0777)
	checkErr(err)
}

func main() {
	dir := "D:\\gotest\\grpc_my\\"
	//file := "D:\\gotest\\grpc_my\\gbops.rar"
	file := "gbops.rar"
	fileContent, err := ioutil.ReadFile(dir + file)
	checkErr(err)
	//fmt.Println(string(fileContent))
	Upload(file, fileContent)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func isDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}
