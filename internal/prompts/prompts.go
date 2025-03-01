package prompts

import (
	_ "embed"
	"fmt"
	"text/template"
)

//go:embed file_review.md
var promptTemplate string

func GetPromptTemplate() (*template.Template, error) {
	t, err := template.New("fileReview").Parse(promptTemplate)
	if err != nil {
		return nil, fmt.Errorf("unable to parse template: %w", err)
	}

	return t, nil
}

type PromptData struct {
	FilePath    string
	DiffContent string
	Language    string
	FileContent string
}
