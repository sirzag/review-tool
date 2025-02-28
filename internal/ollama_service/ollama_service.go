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

	allowedFilesToRead, ok := ctx.Value("allowedFilesToRead").([]string)
	if !ok {
		fmt.Println("allowedFilesToRead is not set in context")
		allowedFilesToRead = []string{}
	} else {
		fmt.Printf("allowedFilesToRead are restricted to %v \n", allowedFilesToRead)
	}

	reqBody := ollamaRequest{
		Model:  string(s.model),
		Prompt: prompt,
		Stream: false,
		Format: formatSpec{
			Type: "object",
			Properties: map[string]property{
				"review": {
					Type: "string",
				},
			},
			Required: []string{"review"},
		},
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

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	var ollamaResp ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	llmResp := &llm.LLMResponse{}
	json.Unmarshal([]byte(ollamaResp.Response), &llmResp)

	// Here we'll need to parse the response text into our LLMResponse structure
	// For now just returning the raw response
	return llmResp, nil
}
