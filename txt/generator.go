package txt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Lang string

const (
	RU = Lang("RU")
	LT = Lang("LT")
)

func GenerateText(size int, lang Lang) string {
	switch lang {
	case RU:
		return generateRuText(size)
	case LT:
		return generateLatinText(size)
	default:
		panic("unsupported language type: " + lang)
	}
}

func generateRuText(paragraphSize int) string {
	resp, err := http.Get("https://fish-text.ru/get?&type=paragraph&number=" + strconv.Itoa(paragraphSize))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	res := response{}
	err = json.Unmarshal(body, &res)

	if err != nil {
		log.Fatalln(err)
	}
	return res.Content
}

type response struct {
	Status  string `json:"status"`
	Content string `json:"text"`
}

func generateLatinText(paragraphSize int) string {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&format=text&paras=" + strconv.Itoa(paragraphSize))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
