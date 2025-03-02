package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/sirzag/review-tool/internal/diffs_collector"
	"github.com/sirzag/review-tool/internal/llm"
	"github.com/sirzag/review-tool/internal/ollama_service"
	"github.com/sirzag/review-tool/internal/prompts"
	"github.com/sirzag/review-tool/internal/writer"
)

func main() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "Reviewing files "
	s.Color("blue")
	s.Start()

	diffs, err := diffs_collector.Collect()
	if err != nil {
		panic(err)
	}

	promptTempl, err := prompts.GetPromptTemplate()
	if err != nil {
		panic(err)
	}

	cwd, _ := os.Getwd()
	llmService := ollama_service.New(llm.LLMConfig{
		BaseURL: "http://localhost:11434",
		Model:   llm.Qwen,
	})

	for i, diff := range diffs {
		if diff.IsDeleted {
			continue
		}

		fileContent, err := os.ReadFile(filepath.Join(cwd, diff.Filepath))
		if err != nil {
			fmt.Println("unable to read file content", err)
			continue
		}

		var prompt bytes.Buffer
		err = promptTempl.Execute(&prompt, prompts.PromptData{
			FileContent: string(fileContent),
			DiffContent: diff.Diffs,
			Language:    diff.Language,
			FilePath:    diff.Filepath,
		})
		if err != nil {
			panic(err)
		}

		os.WriteFile(fmt.Sprintf("debug/%v.md", i), prompt.Bytes(), 0644)

		res, err := llmService.Prompt(context.Background(), prompt.String())
		if err != nil {
			fmt.Println("unable prompt llm", err)
			continue
		}

		val, err := json.MarshalIndent(res, "", "  ")
		os.WriteFile(fmt.Sprintf("debug/%v.json", i), val, 0644)
		writer.WriteToStd(diff.Filepath, res.Observations)
	}
	s.Stop()
}
