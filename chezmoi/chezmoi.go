package chezmoi

import (
	"os/exec"
	"regexp"
)

func GetDiffs() ([]Diff, error) {
	diffs := []Diff{}
	output, err := exec.Command("chezmoi", "diff").Output()
	if err != nil {
		panic(err)
	}
	
	expression := regexp.MustCompile(`diff --git a\/([^\s]*)`)
	matches := expression.FindAllSubmatch(output, -1)
	if len(matches) == 0 {
		return diffs, nil
	}
	for _, match := range matches {
		path := match[1]
		diffs = append(diffs, Diff{FilePath: string(path)})
	}
	return diffs, nil
}

func Add(path string) error {
	cmd := exec.Command("chezmoi", "add", path)
	return cmd.Run()
}

func ChangeDirectory() error {
	cmd := exec.Command("chezmoi", "cd")
	return cmd.Run()
}

type Diff struct {
	FilePath string
}
