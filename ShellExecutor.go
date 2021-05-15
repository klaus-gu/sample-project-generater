package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

/**
执行shell文件
*/
func Execute(script_path string, args ...string) error {
	log.Printf("开始执行Shell脚本... %s\n", script_path)
	var err error
	var cmd *exec.Cmd
	path, _ := GetCurrentPath()
	cmd = exec.Command("sh", path+script_path, args[0], args[1])
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), err
}
