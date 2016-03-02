package upload

import (
	//"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"encoding/json"
)

func Upload(fileAbsolutePath string, fileContent []byte, storagePath string) {
	fileAbsolutePath = strings.Replace(fileAbsolutePath, "\\", "/", -1)
	_, fileName := filepath.Split(fileAbsolutePath)
	//storagePath = filepath.ToSlash(filepath.Dir(storagePath)) + "/"
	if storagePath[len(storagePath)-1:len(storagePath)] != "/" {
		storagePath = storagePath + "/"
	}
	mkdirErr := os.MkdirAll(storagePath, 0777)
	checkErr(mkdirErr)
	//fmt.Println(storagePath)
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

func GetFileList(path string) string {
	//var folder = filepath.ToSlash("C:\\go_test\\deploy")
	var folder = filepath.ToSlash(path)
	var fileList []string
	readFolders(folder, &fileList)
	//fmt.Println(fileList)
	for i, fullPath := range fileList {
		fileList[i] = strings.Replace(fullPath, folder, "", -1)
	}
	//fmt.Println(fileList)
	b, _ := json.Marshal(fileList)
	//fmt.Println(string(b))
	return string(b)
}

func readFolders(path string, fileList *[]string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		checkErr(err)
	}
	for _, f := range files {
		if f.Name() == ".git" {
			continue
		}
		var newPath string
		if f.IsDir() {
			newPath = filepath.ToSlash(path + string(os.PathSeparator) + f.Name())
			readFolders(newPath, fileList)
		} else {
			newPath = filepath.ToSlash(path + string(os.PathSeparator) + f.Name())
			//fmt.Println(newPath)
			*fileList = append((*fileList)[:], newPath)
			//fmt.Println(fileList)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
