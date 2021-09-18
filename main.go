package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pepperkit/tabasco/cmd"
	"strconv"
)

const byteFactor = 1024
const kiloByteParagraphSize = 8
const megaByteParagraphSize = 25

func main() {
	args := cmd.Parse()
	cmd.Info(args)
	cmd.ValidateFileSize(args)

	f, err := os.Create(args.FileName)
	if err != nil {
		log.Fatalln(err)
	}

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
		res := textGenerator(paragraphSize)
		potentialSize := totalSize + len(res.Content)
		if potentialSize > expectedSize {
			needBytes := expectedSize - totalSize

			if totalSize == 0 {
				needBytes = expectedSize
			}

			potentialContent := []byte(res.Content)

			str := string(potentialContent[0:needBytes])

			size, err := w.WriteString(str)
			if err != nil {
				log.Fatalln(err)
			}

			totalSize += size
		} else {
			size, err := w.WriteString(res.Content)
			if err != nil {
				log.Fatalln(err)
			}
			totalSize += size
		}
	}

	w.Flush()
	fi, _ := f.Stat()
	fmt.Printf("Complete! The file %s has been generated.\n", fi.Name())
	fmt.Printf("File size is %d bytes.", fi.Size())
}

func textGenerator(paragrapSize int) TextResponse {
	resp, err := http.Get("https://fish-text.ru/get?&type=paragraph&number=" + strconv.Itoa(paragrapSize))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	res := TextResponse{}
	err = json.Unmarshal(body, &res)

	if err != nil {
		log.Fatalln(err)
	}
	return res
}

type TextResponse struct {
	Status  string `json:"status"`
	Content string `json:"text"`
}
