# Codex
This is the codex literate preprocessor written in itself :D
It's so simple we don't even need a go's build system only the compiler

```go
package main
```

First we import the modules we'll use

```go
import (
	"fmt"
	"os"
	"regexp"
	"strings"
)
```

We need a error handling function because muh clean code said so

```go
func error(a ...any) {
	fmt.Fprintf(os.Stderr, "[Error]: %s\n", a)
	os.Exit(1)
}
```

We also need a separate function which just prints the usage of the program. I know, I'm really leaning into this clean code thing

```go
func usage() {
	fmt.Println("[Usage]: codex <input> <output>")
}
```

This function is actually the entirety of the program. It performs some regex magic and extracts code from the code blocks in the markdown file.
The reason we dynamically construct the backticks is to avoid them being processed by codex (This was very hard to come up with ngl)

```go
func extractCodeBlocks(md string) []string {
	backtick := "`"
	tripleBacktick := backtick + backtick + backtick
	pattern := tripleBacktick + "(?s)(.*?)" + tripleBacktick
	re := regexp.MustCompile(pattern)
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
```

Now the main function starts. We first get the arguments

```go
func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "[Error]: not enough options provided\n")
		usage()
		os.Exit(1)
	}
```

We then read the file specified by the input argument

```go
	input := args[1]
	output := args[2]
	inputFile, err := os.ReadFile(input)
	if err != nil {
		error(err)
	}
```

Now we use our extract function to get the code blocks and create the output file

```go
	codeBlock := extractCodeBlocks(string(inputFile))
	outputFile, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		error(err)
	}
	defer outputFile.Close()
```

We then loop over the codeblocks and append them to the output file and print a success message afterwards.

```go
	for _, code := range codeBlock {
		outputFile.WriteString(code + "\n")
	}
	fmt.Printf("[success]: compiled %s to %s\n", input, output)
}
```
