package ollama_service

import (
	"bytes"
	"context"
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/sirzag/review-tool/internal/llm"
)

type OllamaService struct {
	client  *http.Client
	baseURL string
	model   llm.LLMModel
}

func New(config llm.LLMConfig) llm.LLM {
	return &OllamaService{
		client:  &http.Client{},
		baseURL: config.BaseURL,
		model:   config.Model,
	}
}

func (s *OllamaService) Prompt(ctx context.Context, prompt string) (*llm.LLMResponse, error) {
	if prompt == "" {
		return nil, fmt.Errorf("prompt is empty")
	}

	reqBody := ollamaRequest{
		Model:  string(s.model),
		Prompt: prompt,
		Stream: false,
		Format: LLMResponseFormat,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("%s/api/generate", s.baseURL),
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	var ollamaResp ollamaResponse
	if err := json.NewDecoder(res.Body).Decode(&ollamaResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	llmResp := llm.LLMResponse{}
	json.Unmarshal([]byte(ollamaResp.Response), &llmResp)

	return &llmResp, nil
}

var LLMResponseFormat = formatSpec{
	Type: "object",
	Properties: map[string]property{
		"observations": {
			Type: "array",
			Items: &items{
				Type: "object",
				Properties: map[string]property{
					"type": {
						Type: "string",
						Enum: []string{"ISSUE", "STYLE", "IMPROVEMENT", "CONSISTENCY"},
					},
					"description": {
						Type: "string",
					},
					"suggestion": {
						Type: "string",
					},
					"lines": {
						Type: "string",
					},
				},
				Required: []string{"type", "description", "suggestion", "lines"},
			},
		},
	},
	Required: []string{"observations"},
}
