package byte

import (
	"log"
	"os"
)

type bytesParser struct {
	filePath string
}

func Parser(filePath string) bytesParser {
	return bytesParser{filePath: filePath}
}

func (b bytesParser) GetBytesCount() int64 {
	file, err := os.Open(b.filePath)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	return fi.Size()
}
