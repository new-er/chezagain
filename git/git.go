package git

import (
	"os/exec"
)


func AddAll() error {
	cmd := exec.Command("git", "add", ".")
	return cmd.Run()
}

func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	return cmd.Run()
}


func Push() error {
	cmd := exec.Command("git", "push")
	return cmd.Run()
}
