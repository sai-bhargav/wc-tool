package word

import "strings"

type wordParser struct {
	lines []string
}

func Parser(lines []string) wordParser {

	return wordParser{lines: lines}
}

func (w wordParser) GetWordCount() int {
	var count = 0

	for i := 0; i < len(w.lines); i++ {
		words := strings.Fields(w.lines[i])
		count += len(words)
	}

	return count
}
