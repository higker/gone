package utils

import (
	"bytes"
	"os/exec"
)

/**
 *
 * Create BY YooDing
 *
 * Des: bash or shell or cmd command util
 *
 * Time: 2019/7/9 7:42 PM.
 *
 * <a href="https://github.com/YooDing/gone">Github</a>
 */
const (
	bash = "bash"
	cmd = "cmd"
)

func ShellOut(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(bash, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func ExeShell(command string) error{
	return exec.Command(bash, "-c", command).Run()
}

func CmdOut(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command(bash, "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	}
}

func ExeCmd(command string) error {
	return exec.Command(cmd,"-c",command).Run()
}