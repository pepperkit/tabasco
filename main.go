package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
    "flag"
)

func main() {
	fmt.Println("Fish Text Generator Tool")
    inputFileName := flag.String("file", "fish.txt", "a file name")
    inputExpectedSize := flag.Int("size", 0, "a file expected size in bytes")
    flag.Parse()

    if *inputExpectedSize <= 0 {
        fmt.Println("specify --size arg")
        fmt.Println("--help or -h see more details")
        os.Exit(1)
    }
	f, err := os.Create(*inputFileName)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	f.Sync()

	w := bufio.NewWriter(f)
	expectedSize := *inputExpectedSize
	totalSize := 0

	for totalSize < expectedSize {
		res := textGenerator()
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
	fmt.Printf("The file %s is %d bytes long", fi.Name(), fi.Size())
}

func textGenerator() TextResponse {
	resp, err := http.Get("https://fish-text.ru/get?&type=paragraph&number=1")
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
