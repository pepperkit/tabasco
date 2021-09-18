package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pepperkit/tabasco/cmd"
	"pepperkit/tabasco/txt"
)

const byteFactor = 1024
const kiloByteParagraphSize = 8
const megaByteParagraphSize = 25

func main() {
	args := cmd.Parse()
	cmd.Info(args)
	cmd.ValidateFileSize(args)

	f, err := os.Create(args.FileName)
	checkError(err)

	defer f.Close()

	f.Sync()

	w := bufio.NewWriter(f)
	expectedSize := args.FileSize
	paragraphSize := 1 // default size

	if args.UnitKiloByte {
		expectedSize = expectedSize * byteFactor
		paragraphSize = kiloByteParagraphSize
	}

	if args.UnitMegaByte {
		expectedSize = expectedSize * byteFactor * byteFactor
		paragraphSize = megaByteParagraphSize
	}

	totalSize := 0

	fmt.Println("Generating...")
	for totalSize < expectedSize {
		res := txt.GenerateText(paragraphSize)
		potentialSize := totalSize + len(res.Content)
		if potentialSize > expectedSize {
			needBytes := expectedSize - totalSize

			if totalSize == 0 {
				needBytes = expectedSize
			}

			potentialContent := []byte(res.Content)

			str := string(potentialContent[0:needBytes])

			size, err := w.WriteString(str)
			checkError(err)

			totalSize += size
		} else {
			size, err := w.WriteString(res.Content)
			checkError(err)
			totalSize += size
		}
	}

	w.Flush()
	fi, _ := f.Stat()
	fmt.Printf("Complete! The file %s has been generated.\n", fi.Name())
	fmt.Printf("File size is %d bytes.", fi.Size())
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
