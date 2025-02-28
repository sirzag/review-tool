package llm

import "context"

type AnalysisStatus string
type LLMModel string

const (
	Codellama  LLMModel = "codellama"
	Llama3     LLMModel = "llama3"
	DeepseekR1 LLMModel = "deepseek-r1"
)

const (
	ContextRequired AnalysisStatus = "context_required"
	Done            AnalysisStatus = "done"
)

type LLMResponse struct {
	Review string `json:"review,omitempty"`
}

type LLM interface {
	Prompt(ctx context.Context, prompt string) (*LLMResponse, error)
}

type LLMConfig struct {
	BaseURL string   `json:"baseURL"`
	Model   LLMModel `json:"model"`
}
