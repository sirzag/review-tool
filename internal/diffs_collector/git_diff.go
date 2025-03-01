package diffs_collector

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type GitDiff struct {
	Filepath  string
	Language  string
	Diffs     string
	IsDeleted bool
}

func getFileDiff(file string) (GitDiff, error) {
	out, err := exec.Command("git", "diff", "--no-color", "--no-prefix", "--", file).Output()
	if err != nil {
		return GitDiff{}, fmt.Errorf("unable to get diff: %w", err)
	}
	outText := string(out)
	isDeleted := strings.Contains(outText, "deleted file mode")
	ext := strings.TrimPrefix(filepath.Ext(file), ".")
	return GitDiff{file, outText, ext, isDeleted}, nil
}

func getGitChangedFiles() ([]string, error) {
	out, err := exec.Command("git", "diff", "--name-only").Output()
	if err != nil {
		return []string{}, fmt.Errorf("unable to get diff: %w", err)
	}

	res := []string{}
	for _, str := range strings.Split(string(out), "\n") {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		res = append(res, str)
	}
	return res, nil
}
