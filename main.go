package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const BYTE_FACTOR = 1024

func main() {

	inputFileName := flag.String("file", "", "a file name")
	inputExpectedSize := flag.Int("size", 0, "a file expected size in bytes")
	unitKb := flag.Bool("kb", false, "flag set a size unit as KBytes")
	unitMb := flag.Bool("mb", false, "flag set a size unit as MBytes")
	flag.Parse()

	if *inputExpectedSize <= 0 && !*unitKb && !*unitMb && len(*inputFileName) <= 0 {
		info()
	}

	if *inputExpectedSize <= 0 {
		fmt.Println("specify --size arg")
		fmt.Println("Use \"tabasco --help\" for more information")
		fmt.Println("")
		fmt.Println("PepperKit(c) 2021.")
		fmt.Println("")
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
	paragrapSize:= 1

	if *unitKb {
		expectedSize = expectedSize * BYTE_FACTOR
		paragrapSize = 8
	}

	if *unitMb {
		expectedSize = expectedSize * BYTE_FACTOR * BYTE_FACTOR
		paragrapSize = 25
	}

	totalSize := 0

	fmt.Println("Generating...")
	for totalSize < expectedSize {
		res := textGenerator(paragrapSize)
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

func info() {
	fmt.Println("Tabasco is CLI tool to generate a placeholder text akka 'Lorem ipsum'.")
	fmt.Println("")
	fmt.Println("Usage: ")
	fmt.Println("\t tabasco [--arguments]")
	fmt.Println("")
	fmt.Println("The argumnets are: ")
	fmt.Println("\t file \t a file name")
	fmt.Println("\t size \t an expected file size (in bytes by default)")
	fmt.Println("\t kb \t a size will be read as KBytes")
	fmt.Println("\t mb \t a size will be read as MBytes")
	fmt.Println("")
	fmt.Println("Use \"tabasco --help\" for more information")
	fmt.Println("")
	fmt.Println("Tabasco uses https://fish-text.ru service to get a random text. It means the Internet connection is important.")
	fmt.Println("")
	fmt.Println("MIT License")
	fmt.Println("Copyright (c) 2021 PepperKit.")
	fmt.Println("")
	os.Exit(0)
}
