package diffs_collector

import "fmt"

func Collect() ([]GitDiff, error) {
	files, err := getGitChangedFiles()
	if err != nil {
		return []GitDiff{}, fmt.Errorf("unable to get changed files: %w", err)
	}

	diffs := make([]GitDiff, len(files))

	for i, file := range files {
		diff, err := getFileDiff(file)
		if err != nil {
			return diffs, fmt.Errorf("unable to get diff for file %s: %w", file, err)
		}
		diffs[i] = diff
	}

	return diffs, nil
}
