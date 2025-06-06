package blog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func clearDirectory(dirPath string) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry.Name())
		if err := os.RemoveAll(fullPath); err != nil {
			return fmt.Errorf("failed to remove %s: %w", fullPath, err)
		}
	}

	return nil
}

func truncateWithMinMax(s string, minLen, maxLen int) string {
	runes := []rune(s)

	if len(runes) <= maxLen || len(runes) < minLen {
		return s
	}

	truncated := string(runes[:maxLen])
	if lastSpace := strings.LastIndex(truncated, " "); lastSpace >= minLen {
		return string(runes[:lastSpace])
	}

	return string(runes[:maxLen])
}
