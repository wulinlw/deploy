package command

import (
	"bytes"
	//"fmt"
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

//func main() {
//	//  re1 := svnCheckout("https://svn.td.gamebar.com/svn/private/liwai/test", "/root/goproject/svntest")
//	//	re1 := svnUp("/root/goproject/svntest/test")
//	//	fmt.Println(re1)
//	//	re2 := svnUpToRevision("/root/goproject/svntest/test", 502)
//	//	fmt.Println(re2)
//	//	re3 := specifiedCommand("wc -l /root/goproject/svntest/test/1.txt", "")
//	//	fmt.Println(re3)
//	//re4 := specifiedCommand("ps aux", "")
//	re4 := complexCommand("ps aux |grep apache", "")
//	fmt.Println(re4)
//}

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
func SvnUp(dir string) int {
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
func SvnCheckout(svnUrl string, dir string) int {
	cmd := exec.Command("svn", "checkout", svnUrl)
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}
}

//获得svn info信息
func SvnInfo(dir string) string {
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
func SvnUpToRevision(dir string, revision int) int {
	cmd := exec.Command("svn", "up", "-r", strconv.Itoa(revision))
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	checkErr(err)
	//fmt.Println(out.String())
	return praseCommandOut(out.String(), regexpStr["svnUp"])
}

//执行指定命令，简单命令，不包含管道操作、&&等
func SpecifiedCommand(command string, dir string) string {
	if dir == "" {
		dir = "/root/"
	}
	command = strings.TrimSpace(command)
	if command == "" {
		log.Println("lost command")
	}
	parts := strings.Fields(command)
	var cmd *exec.Cmd
	if len(parts) == 1 {
		cmd = exec.Command(command)
	} else {
		commandFile := parts[0]
		parts = parts[1:len(parts)]
		commandSlice := []string{}
		for _, param := range parts {
			commandSlice = append(commandSlice, param)
		}
		cmd = exec.Command(commandFile, commandSlice...)
		///usr/bin/wc -l /root/goproject/svntest/test/1.txt
		//cmd = exec.Command("/usr/bin/wc", "-l", "/root/goproject/svntest/test/1.txt")
	}
	cmd.Dir = dir
	env := "/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin:/usr/local/go/bin:/root/bin"
	cmd.Env = strings.Split(env, ":")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	checkErr(err)
	//fmt.Println(out.String())
	return out.String()
}

//复杂命令 管道操作等
func ComplexCommand(command string, dir string) string {
	if dir == "" {
		dir = "/root/"
	}
	command = strings.TrimSpace(command)
	if command == "" {
		log.Println("lost command")
	}
	parts := strings.Fields(command)
	var cmd *exec.Cmd
	if len(parts) == 1 {
		cmd = exec.Command(command)
	} else {
		cmd = exec.Command("/bin/bash", "-c", command)
	}
	cmd.Dir = dir
	cmd.Env = []string{
		"SHELL=/bin/bash",
		"PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin:/usr/local/go/bin:/root/bin",
		"LC_ALL=en_US.UTF-8",
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	checkErr(err)
	//fmt.Println(out.String())
	return out.String()
}
