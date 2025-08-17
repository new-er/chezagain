package git

import (
	"os/exec"
)


func AddAll() error {
	println("Adding all changes to git...")
	cmd := exec.Command("git", "add", ".")
	return cmd.Run()
}

func Commit(message string) error {
	println("Committing changes with message:", message)
	cmd := exec.Command("git", "commit", "-m", message)
	return cmd.Run()
}


func Push() error {
	println("Pushing changes to remote repository...")
	cmd := exec.Command("git", "push")
	return cmd.Run()
}
