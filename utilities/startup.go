package main

import (
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"os/exec"
	"syscall"
	"regexp"
)

func main() {
	var py_user, uid, gid string = os.Getenv("PYTHON_USER"), os.Args[1], os.Args[2]
	var home string = "/home/" + py_user

	pattern := py_user + ":x:[0-9]+:[0-9]+:Developer"
	re := regexp.MustCompile(pattern)

	pass_byte, _ := ioutil.ReadFile("/etc/passwd")
	pass_str := string(pass_byte)
	pass_str = strings.Replace(
		pass_str, re.FindString(pass_str), py_user + ":x:" + uid + ":" + gid + ":Developer", -1)
	pass_file, _ := os.Create("/etc/passwd")
	pass_file.WriteString(pass_str)
	pass_file.Close()


	uid_int, _ := strconv.Atoi(uid)
	gid_int, _ := strconv.Atoi(gid)
	os.Chown(home, uid_int, gid_int)
	os.Chown(home + "/.local", uid_int, gid_int)
	os.Chown(home + "/.local/bin", uid_int, gid_int)
	os.Chown(home + "/.cache", uid_int, gid_int)
	os.Chown(home + "/.cache/pip", uid_int, gid_int)
	os.Chown(home + "/.cache/pip/wheels", uid_int, gid_int)
	os.Chown(home + "/.conda", uid_int, gid_int)
	os.Chown(home + "/.conda/environments.txt", uid_int, gid_int)
	os.Chown(home + "/.condarc", uid_int, gid_int)

	if _, err := os.Stat(os.Getenv("APK_REQUIREMENTS")); err == nil {
		cmd1 := exec.Command("paste", "-s", "-d", ",", os.Getenv("APK_REQUIREMENTS"))
		output1, err1 := cmd1.CombinedOutput()
		if err1 != nil {
			println("System requirements read error!")
			println(err1.Error())
			return
		}
		args := []string{"add"}
		pkgs := strings.Split(string(output1), ",")
		args = append(args,pkgs...)
		cmd2 := exec.Command("apk", args...)
		cmd2.SysProcAttr = &syscall.SysProcAttr{}
		cmd2.SysProcAttr.Credential = &syscall.Credential{Uid: 0, Gid: 0}
		_, err2 := cmd2.CombinedOutput()
		if err2 != nil {
			println("System requirements install error!")
			println(err2.Error())
			return
		}
	}
	if _, err := os.Stat(os.Getenv("APT_REQUIREMENTS")); err == nil {
		cmd1 := exec.Command("paste", "-s", "-d", ",", os.Getenv("APT_REQUIREMENTS"))
		output1, err1 := cmd1.CombinedOutput()
		if err1 != nil {
			println("System requirements read error!")
			println(err1.Error())
			return
		}
		args := []string{"install"}
		pkgs := strings.Split(string(output1)[:len(string(output1))-1] + ",-y", ",")
		args = append(args,pkgs...)
		cmd2 := exec.Command("apt", args...)
		cmd2.SysProcAttr = &syscall.SysProcAttr{}
		cmd2.SysProcAttr.Credential = &syscall.Credential{Uid: 0, Gid: 0}
		_, err2 := cmd2.CombinedOutput()
		if err2 != nil {
			println("System requirements install error!")
			println(err2.Error())
			return
		}
	}
}
