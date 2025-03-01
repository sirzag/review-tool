package writer

import (
	"fmt"
	"slices"
	"strings"

	"github.com/sirzag/review-tool/internal/llm"
)

func WriteToStd(filename string, observation []llm.Observation) {
	notEmptySuggestionIdx := slices.IndexFunc(observation, func(s llm.Observation) bool {
		return s.Suggestion != ""
	})

	if notEmptySuggestionIdx == -1 {
		return
	}

	fmt.Print("\n\n")
	fmt.Println(fmt.Sprintf("📁 %s", filename))

	for _, obs := range observation {
		if obs.Suggestion == "" {
			continue
		}

		icon := getTypeIcon(obs.Type)
		fmt.Println(fmt.Sprintf("  %s [%s]: %s", icon, obs.Lines, obs.Description))
		fmt.Println(fmt.Sprintf("    * Suggestion: %s", strings.ReplaceAll(obs.Suggestion, "\n", "\n    ")))
	}
	fmt.Print("\n\n")
}

func getTypeIcon(t string) string {
	switch t {
	case "ISSUE":
		return "🛑"
	case "CONSISTENCY", "STYLE":
		return "⚠️"
	case "IMPROVEMENT":
		return "ℹ️"
	default:
		return ""
	}
}
