package proc

import (
	"log"
	"os"
	"strconv"
	"syscall"
)

const (
	PidFile = "/var/run/go_service.pid"
)

var (
	pidFile string
	sigs    chan os.Signal
)

func SetPid(pidfile string) int {
	pidFile = pidfile
	if pidFile == "" {
		pidFile = PidFile
	}

	// 判断pid是否存在，如果存在且有效退出程序
	if Existed() {
		log.Printf("pid file  %s  exist \n", pidFile)
		os.Exit(1)
		return 0
	}

	pid := os.Getpid()
	_ = os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0666)
	return pid
}

// 判断pid文件是否存在，以及pid进程是否存在
func Existed() bool {
	if !IsExist(pidFile) {
		return false
	}

	res, err := os.ReadFile(pidFile)
	if err != nil {
		return false
	}
	pid, _ := strconv.Atoi(string(res))
	if pid <= 0 {
		return false
	}

	if err = syscall.Kill(pid, 0); err == nil {
		return true
	}

	return false
}

// 删除pid文件
func DelPid() {
	if IsExist(pidFile) {
		return
	}
	os.Remove(pidFile)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	// Check if error is "no such file or directory"
	if _, ok := err.(*os.PathError); ok {
		return false
	}
	return false
}
