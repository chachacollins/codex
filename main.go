package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ExtractCodeBlocks(md string) []string {
	re := regexp.MustCompile("```(?s)(.*?)```")
	matches := re.FindAllStringSubmatch(md, -1)
	var results []string
	for _, match := range matches {
		if len(match) > 1 {
			content := strings.TrimSpace(match[1])
			if idx := strings.Index(content, "\n"); idx != -1 {
				content = strings.TrimSpace(content[idx:])
			}
			results = append(results, content)
		}
	}
	return results
}

func main() {
	fmt.Println("Hello world")
}
