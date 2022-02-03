package email

import (
	"io/ioutil"
	"log"
	"simulator/pkg/data"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func Result() [][]EmailData {
	var res [][]EmailData
	for _, v := range emailMap() {
		for _, val := range v {
			res = append(res, val)
		}
	}
	return res
}

func emailMap() map[string][][]EmailData {

	countryMap := make(map[string][][]EmailData)
	for _, v := range unique() {
		countryMap[v] = make([][]EmailData, 2)
	}
	emailData := emailSortByDelTime(emailContent())
	for i := 0; i < len(emailData); i++ {
		if len(countryMap[emailData[i].Country][0]) < 3 {
			countryMap[emailData[i].Country][0] = append(countryMap[emailData[i].Country][0], emailData[i])
		}
		if len(countryMap[emailData[len(emailData)-1-i].Country][1]) < 3 {
			countryMap[emailData[len(emailData)-1-i].Country][1] = append(countryMap[emailData[len(emailData)-1-i].Country][1], emailData[len(emailData)-1-i])
		}
	}
	return countryMap
}

func emailContent() []EmailData {
	content, err := ioutil.ReadFile("email.data")

	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	emailData := make([]EmailData, 0, len(lines))
	for _, v := range lines {
		if strings.Count(v, ";") == 2 {
			str := strings.Split(v, ";")
			if _, ok := data.Countries[str[0]]; ok {
				if _, ok := data.EmailProviders[str[1]]; ok {
					delTime, err := strconv.Atoi(str[2])
					if err != nil {
						log.Fatal(err)
					}
					e := EmailData{Country: str[0], Provider: str[1], DeliveryTime: delTime}
					emailData = append(emailData, e)
				}
			}
		}
	}
	return emailData
}

func emailSortByDelTime(b []EmailData) []EmailData {
	sort.Slice(b, func(i, j int) bool {
		return b[i].DeliveryTime < b[j].DeliveryTime
	})
	return b
}

func unique() []string {
	eCont := emailContent()
	countries := make([]string, 0, len(eCont))
	for _, v := range eCont {
		countries = appendIfMissing(countries, v.Country)
	}
	return countries
}

func appendIfMissing(slice []string, str string) []string {
	for _, v := range slice {
		if v == str {
			return slice
		}
	}
	return append(slice, str)
}
