package main

import (
	"fmt"
	"main/byte"
	"main/line"
	"main/printer"
	"main/word"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 || args[1] == "" {
		fmt.Printf("Please pass file name")
		return
	}

	filePath := args[1]
	linesParser := line.Parser(filePath)
	lines := linesParser.GetLines()
	wordParser := word.Parser(lines)
	bytesParser := byte.Parser(filePath)

	printer.PrintResults(len(lines), wordParser.GetWordCount(), bytesParser.GetBytesCount(), filePath)
}
