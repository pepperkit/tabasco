package txt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const endpoint = "https://fish-text.ru/get"

func GenerateText(paragraphSize int) Response {
	resp, err := http.Get("https://fish-text.ru/get?&type=paragraph&number=" + strconv.Itoa(paragraphSize))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	res := Response{}
	err = json.Unmarshal(body, &res)

	if err != nil {
		log.Fatalln(err)
	}
	return res
}

type Response struct {
	Status  string `json:"status"`
	Content string `json:"text"`
}
