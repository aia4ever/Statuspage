package incident

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"simulator/pkg/getresponse"
	"sort"
	"strings"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func Result() []IncidentData {
	resp := getresponse.GetResponse("https://dataemul.herokuapp.com/accident")
	defer resp.Body.Close()
	var incidentData []IncidentData
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
			var tmp IncidentData
			err := dec.Decode(&tmp)
			if err != nil {
				log.Fatal(err)
			}
			if tmp.Status == "active" || tmp.Status == "closed" {
				incidentData = append(incidentData, tmp)
			}
		}
		//_, err = dec.Token()
		//if err != nil {
		//	log.Fatal(err)
		//}
	}
	return sortByStatus(incidentData)

}

func sortByStatus(incidentData []IncidentData) []IncidentData {
	sort.Slice(incidentData, func(i, j int) bool { return incidentData[i].Status < incidentData[j].Status })
	return incidentData
}
