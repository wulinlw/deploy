package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var regexpStr map[string]string = map[string]string{
	"svnCheckout": "Checked out revision ([\\d]{1,10})",
	"svnUp":       "[Updated to|At] revision ([\\d]{1,10})",
}

func main() {
	//  re1 := svnCheckout("https://svn.td.gamebar.com/svn/private/liwai/test", "/root/goproject/svntest")
	//	re1 := svnUp("/root/goproject/svntest/test")
	//	fmt.Println(re1)
	//	re2 := svnUpToRevision("/root/goproject/svntest/test", 502)
	//	fmt.Println(re2)
	//	re3 := specifiedCommand("wc -l /root/goproject/svntest/test/1.txt", "")
	//	fmt.Println(re3)
	//specifiedCommand("wc -l /root/goproject/svntest/test/1.txt", "")
	specifiedCommand("echo $path", "")
}

/**
svn up
删除文件后，执行svn up返回
Restored '1.txt'
At revision 501.

--分割线--
更新文件后，执行svn up返回
U    1.txt
Updated to revision 502.

--分割线--
增加文件后，执行svn up返回
A    2.txt
Updated to revision 503.

*/
func svnUp(dir string) int {
	cmd := exec.Command("svn", "up")
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	checkErr(err)
	//fmt.Println(out.String())
	return praseCommandOut(out.String(), regexpStr["svnUp"])
}

//svn checkout版本到指定的目录
func svnCheckout(svnUrl string, dir string) int {
	cmd := exec.Command("svn", "checkout", svnUrl)
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return praseCommandOut(out.String(), regexpStr["svnCheckout"])
}

//检测是否checkout成功，成功返回版本号，失败返回0
func praseCommandOut(str string, regexpString string) int {
	re, err := regexp.Compile(regexpString)
	checkErr(err)
	res := re.FindStringSubmatch(str)
	if len(res) >= 2 {
		version, _ := strconv.Atoi(res[len(res)-1])
		return version
	}
	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//获得svn info信息
func svnInfo(dir string) string {
	cmd := exec.Command("svn", "info")
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	checkErr(err)
	return out.String()
}

//svn 升级到指定版本
//http://www.cnblogs.com/mfryf/p/4654110.html
func svnUpToRevision(dir string, revision int) int {
	cmd := exec.Command("svn", "up", "-r", strconv.Itoa(revision))
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	checkErr(err)
	//fmt.Println(out.String())
	return praseCommandOut(out.String(), regexpStr["svnUp"])
}

//执行指定命令
func specifiedCommand(command string, dir string) {
	if dir == "" {
		dir = "/root/"
	}
	command = strings.TrimSpace(command)
	if command == "" {
		log.Fatal("lost command")
	}
	parts := strings.Fields(command)
	var cmd *exec.Cmd
	if len(parts) == 1 {
		cmd = exec.Command(command)
	} else {
		parts = parts[1:len(parts)]
		//commandSlice := []string{"-c", command}
		//1 i,v in commandSlice
		//commandSlice = append(commandSlice, parts[1:len(parts)])
		//cmd = exec.Command("/bin/bash", "-c", command, parts...)
		cmd = exec.Command("/usr/bin/wc ", "-l", "/root/goproject/svntest/test/1.txt")
	}
	fmt.Println(exec.LookPath("wc"))
	cmd.Dir = dir
	env := "/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin:/usr/local/go/bin:/root/bin"
	cmd.Env = strings.Split(env, ":")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	checkErr(err)
	fmt.Println(out.String())
}
