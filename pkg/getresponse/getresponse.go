package getresponse

import (
	"log"
	"net/http"
)

func GetResponse(url string) *http.Response {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
