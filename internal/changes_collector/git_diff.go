package changes_collector

import (
	"fmt"
	"os/exec"
)

type GitDiff struct {
	Filepath string
	Diffs    string
}

func getGitDiff() ([]GitDiff, error) {

	res := make([]GitDiff, 1)
	gdiff := exec.Command("git", "diff")

	out, err := gdiff.Output()
	if (err != nil) {
		return res, fmt.Errorf("unable to get diff: %w", err)
	}

	fmt.Println(out)


	return res, nil
}
