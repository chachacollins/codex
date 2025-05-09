package main

import (
	"fmt"
	"os"
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

func usage() {
	fmt.Println("[Usage]: codex <input> <output>")
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "[Error]: not enough options provided\n")
		usage()
		os.Exit(1)
	}
	input := args[1]
	output := args[2]
	inputFile, err := os.ReadFile(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error]: %s", err)
		os.Exit(1)
	}
	codeBlock := ExtractCodeBlocks(string(inputFile))
	outputFile, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error]: %s", err)
		os.Exit(1)
	}
	defer outputFile.Close()
	for _, code := range codeBlock {
		outputFile.WriteString(code + "\n")
	}
	fmt.Printf("[success]: compiled %s to %s\n", input, output)
}
