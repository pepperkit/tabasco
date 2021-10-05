package main

import (
	"fmt"
	"github.com/pepperkit/tabasco/cmd"
	"github.com/pepperkit/tabasco/txt"
	"github.com/pepperkit/tabasco/writer"
	"log"
	"os"
)

const byteFactor = 1024
const kiloByteParagraphSize = 8
const megaByteParagraphSize = 25

func main() {
	args := cmd.Parse()
	cmd.Info(args)
	cmd.ValidateFileSize(args)

	documentWriter := newDocumentWriter(args)
	generateTextBySize(args, documentWriter)
	makeReport(documentWriter.FileName())
}

func newDocumentWriter(args *cmd.TabascoArgs) writer.DocumentWriter {
	if args.Docx {
		return writer.NewDocxWriter(args.FileName)
	}
	return writer.NewTxtWriter(args.FileName)
}

func makeReport(fileName string) {
	file, err := os.Open(fileName)
	checkError(err)
	fi, _ := file.Stat()
	fmt.Printf("Complete! The file %s has been generated.\n", fi.Name())
	fileSizeBytes := fi.Size()
	fmt.Printf("File size is %d bytes\n", fileSizeBytes)

	if fileSizeBytes > byteFactor {
		fileSizeKiB := float64(fileSizeBytes / byteFactor)
		fmt.Printf("File size is %.2f KiB\n", fileSizeKiB)
	}

	if fileSizeBytes > (byteFactor * byteFactor) {
		fileSizeMiB := float64((fileSizeBytes / byteFactor) / byteFactor)
		fmt.Printf("File size is %.2f MiB\n", fileSizeMiB)
	}
}

func generateTextBySize(args *cmd.TabascoArgs, wr writer.DocumentWriter) {
	expectedSize := args.FileSize
	paragraphSize := 1 // default size

	if args.UnitKiB {
		expectedSize = expectedSize * byteFactor
		paragraphSize = kiloByteParagraphSize
	}

	if args.UnitMiB {
		expectedSize = expectedSize * byteFactor * byteFactor
		paragraphSize = megaByteParagraphSize
	}

	totalSize := 0
	fmt.Println("Generating...")
	for totalSize < expectedSize {
		res := txt.GenerateText(paragraphSize, args.Language)
		potentialSize := totalSize + len(res)
		if potentialSize > expectedSize {
			needBytes := expectedSize - totalSize

			if totalSize == 0 {
				needBytes = expectedSize
			}

			potentialContent := []byte(res)

			str := string(potentialContent[0:needBytes])
			wr.WriteText(str)
			totalSize += len([]byte(res))
		} else {
			wr.WriteText(res)
			totalSize += len([]byte(res))
		}
	}

	wr.Flush()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
