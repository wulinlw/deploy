package upload

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Upload(relativePath string, fileContent []byte) {
	storagePath, _ := filepath.Abs("/usr/local/gbops/scripts/" + relativePath)
	dir, _ := filepath.Split(storagePath)
	if !isDirExists(dir) {
		os.MkdirAll(dir, 0777)
	}
	err := ioutil.WriteFile(storagePath, fileContent, 0777)
	checkErr(err)
}

func GetFileList(path string) string {
	var folder = filepath.ToSlash(path) + string(os.PathSeparator)
	var fileList []string
	readFolders(folder, &fileList)
	for i, fullPath := range fileList {
		fileList[i] = strings.Replace(fullPath, folder+string(os.PathSeparator), "", -1)
	}
	b, _ := json.Marshal(fileList)
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
			*fileList = append((*fileList)[:], newPath)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
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
