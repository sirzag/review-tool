package ollama_service

type ollamaRequest struct {
	Model  string     `json:"model"`
	Prompt string     `json:"prompt"`
	Stream bool       `json:"stream"`
	Format formatSpec `json:"format"`
}

type formatSpec struct {
	Type       string              `json:"type"`
	Properties map[string]property `json:"properties"`
	Required   []string            `json:"required"`
}

type property struct {
	Type  string   `json:"type"`
	Enum  []string `json:"enum,omitempty"`
	Items *items   `json:"items,omitempty"`
}

type items struct {
	Type string `json:"type"`
	Enum  []string `json:"enum,omitempty"`
}

type ollamaResponse struct {
	Response string `json:"response"`
}
