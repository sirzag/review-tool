# Code Review Tool

A command-line tool that automatically reviews code changes in your Git repository using large language models (LLMs).

## Overview

This tool analyzes git diffs in your repository and provides intelligent code review suggestions. It works by:

1. Collecting changed files from git
2. Analyzing the diffs and full file context using LLMs
3. Providing actionable feedback on code quality, consistency, and potential issues

## Features

- **Automated Code Reviews**: Get instant feedback on your code changes
- **Multiple LLM Support**: Works with Codellama, Llama3, and Qwen2.5
- **Actionable Suggestions**: Provides specific code snippets to improve your code
- **Local Processing**: Uses Ollama to run models locally, keeping your code private

## Prerequisites

- Git
- Go (1.18+)
- [Ollama](https://ollama.ai/) installed and running locally

## Installation

```bash
# Clone the repository
git clone https://github.com/sirzag/review-tool.git
cd review-tool

# Build the tool
go build -o review-tool ./cmd/main
```

## Usage

1. Make sure Ollama is running with your preferred model (default is Qwen2.5. It works best and consumes the least memory)
2. Navigate to your Git repository
3. Run the review tool:

```bash
/path/to/review-tool
```

The tool will:
- Collect all changed files in the current repository
- Send the diffs to the language model for analysis
- Print suggested improvements to the console

## Sample Output

```
📁 path/to/file.go
  🛑 [15-17]: Missing error handling after database query
    * Suggestion: if err != nil {
        return nil, fmt.Errorf("failed to query database: %w", err)
    }

  ⚠️ [42]: Inconsistent variable naming convention
    * Suggestion: var userID int // instead of var userId int
```

## Configuration

The tool uses Ollama with the Qwen2.5 model by default. You can modify the `main.go` file to change:

- The base URL for the Ollama service
- The language model (Codellama, Llama3, or Qwen)

## Roadmap

[] - Add interface for openai
[] - Add interface for claude
[] - Create more review flavours (new prompt templates)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).
