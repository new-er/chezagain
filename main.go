package main

import (
	"fmt"

	"github.com/new-er/chezagain/chezmoi"
	"github.com/new-er/chezagain/git"
)

const (
	chezmoiPath = "~/local/share/chezmoi"
)

func main() {
	diffs, err := chezmoi.GetDiffs()
	if err != nil {
		panic(err)
	}
	if len(diffs) == 0 {
		println("No diffs found.")
		return
	}
	println("Diffs:")
	for _, diff := range diffs {
		println(diff.FilePath)
	}
	println("Add Diffs? (y/n)")
	var input string
	fmt.Scanln(&input)
	if input != "y" {
		println("Exiting...")
		return
	}
	println("Adding Diffs...")

	for _, diff := range diffs {
		chezmoi.Add(fmt.Sprintf("%s/%s", "~/", diff.FilePath))
		println(diff.FilePath)
	}

	chezmoi.ChangeDirectory()
	git.AddAll()
	git.Commit("add diffs")
	git.Push()
}
