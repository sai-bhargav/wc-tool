package printer

import "fmt"

func PrintResults(linesCount int, wordsCount int, bytesCount int64, fileName string) {

	fmt.Printf("       %d       %d      %d %s", linesCount, wordsCount, bytesCount, fileName)

}
