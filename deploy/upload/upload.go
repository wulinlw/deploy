package upload

import (
	//"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Upload(fileAbsolutePath string, fileContent []byte, storagePath string) {
	fileAbsolutePath = strings.Replace(fileAbsolutePath, "\\", "/", -1)
	_, fileName := filepath.Split(fileAbsolutePath)
	mkdirErr := os.MkdirAll(storagePath, 0777)
	checkErr(mkdirErr)
	storagePath = filepath.ToSlash(filepath.Dir(storagePath)) + "/"
	newFile := storagePath + fileName
	//fmt.Println(newFile, len(fileContent))
	fileErr := ioutil.WriteFile(newFile, fileContent, 0777)
	checkErr(fileErr)
}

//func main() {
//	file := "E:\\soft_package\\Office_Visio_Pro_2007_Win32_ChnSimp_Disk_Kit_MVL_CD.iso"
//	fileContent, err := ioutil.ReadFile(file)
//	checkErr(err)
//	//fmt.Println(string(fileContent))
//	Upload(file, fileContent, "D:/gotest/grpc_my/v/")
//}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
