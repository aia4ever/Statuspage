package support

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"simulator/pkg/getresponse"
	"strings"
)

type Support struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func Result() []int {
	a := getSupportData()
	b := make([]int, 2)
	var load int
	for _, v := range a {
		load += v.ActiveTickets
	}
	switch {
	case load > 16:
		b[0] = 3
	case load >= 9 && load <= 16:
		b[0] = 2
	default:
		b[0] = 1
	}
	b[1] = int(float64(load) * 60 / 18)
	return b
}

func getSupportData() []Support {
	resp := getresponse.GetResponse("http://127.0.0.1:8383/support")
	dataSlice := make([]Support, 0)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		jsonStream, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		dec := json.NewDecoder(strings.NewReader(string(jsonStream)))

		_, err = dec.Token()
		if err != nil {
			log.Fatal(err)
		}

		for dec.More() {
			var tmp Support
			err := dec.Decode(&tmp)
			if err != nil {
				log.Fatal(err)
			}
			dataSlice = append(dataSlice, tmp)
		}
		//_, err = dec.Token()
		//if err != nil {
		//	log.Fatal(err)
		//}
	}
	return dataSlice
}
