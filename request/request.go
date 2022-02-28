package request

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Req map[int]string
}

func Req() string {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	return sb
}
