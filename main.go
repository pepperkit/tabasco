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

	file, err := os.Create(args.FileName)
	checkError(err)
	defer file.Close()

	err = file.Sync()
	checkError(err)

	writer := bufio.NewWriter(file)

	generateTextBySize(args, writer)

	fi, _ := file.Stat()
	fmt.Printf("Complete! The file %s has been generated.\n", fi.Name())
	fmt.Printf("File size is %d bytes.", fi.Size())
}

func generateTextBySize(args *cmd.TabascoArgs, writer *bufio.Writer) {
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

			size, err := writer.WriteString(str)
			checkError(err)

			totalSize += size
		} else {
			size, err := writer.WriteString(res.Content)
			checkError(err)
			totalSize += size
		}
	}
	err := writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
