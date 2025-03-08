package llm

import "context"

type AnalysisStatus string
type LLMModel string

const (
	Codellama  LLMModel = "codellama:13b"
	Llama3     LLMModel = "llama3.1"
	Qwen       LLMModel = "qwen2.5"
)

const (
	ContextRequired AnalysisStatus = "context_required"
	Done            AnalysisStatus = "done"
)

type LLMResponse struct {
	Observations []Observation `json:"observations,omitempty"`
}

type Observation struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Suggestion  string `json:"suggestion"`
	Lines       string `json:"lines"`
}

type LLM interface {
	Prompt(ctx context.Context, prompt string) (*LLMResponse, error)
}

type LLMConfig struct {
	BaseURL string   `json:"baseURL"`
	Model   LLMModel `json:"model"`
}
