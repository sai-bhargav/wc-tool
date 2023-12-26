package line

import (
	"bufio"
	"log"
	"os"
)

type lineParser struct {
	filePath string
}

func Parser(filePath string) lineParser {
	return lineParser{filePath: filePath}
}

func (l lineParser) GetLines() []string {
	file, err := os.Open(l.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
