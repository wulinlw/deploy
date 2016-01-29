package main

import (
	//"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func Upload(fileAbsolutePath string, fileContent []byte, storagePath string) {
	_, fileName := filepath.Split(fileAbsolutePath)
	newFile := storagePath + fileName
	err := ioutil.WriteFile(newFile, fileContent, 0777)
	checkErr(err)
}

func main() {
	file := "E:\\soft_package\\Office_Visio_Pro_2007_Win32_ChnSimp_Disk_Kit_MVL_CD.iso"
	fileContent, err := ioutil.ReadFile(file)
	checkErr(err)
	//fmt.Println(string(fileContent))
	Upload(file, fileContent, "D:/gotest/grpc_my/v/")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
