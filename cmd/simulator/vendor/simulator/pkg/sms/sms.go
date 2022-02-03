package sms

import (
	"io/ioutil"
	"log"
	"simulator/pkg/data"
	"sort"
	"strings"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func Result() [][]SMSData {
	result := make([][]SMSData, 2)
	result[0] = smsSortByCountry(smsContent())
	result[1] = smsSortByProvider(smsContent())
	return result
}

func smsContent() []SMSData {
	content, err := ioutil.ReadFile("sms.data")
	var datas []SMSData
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, v := range lines {
		if strings.Count(v, ";") == 3 {
			str := strings.Split(v, ";")
			if _, ok := data.Countries[str[0]]; ok {
				if _, ok := data.Providers[str[3]]; ok {
					smsData := SMSData{Country: data.Countries[str[0]], Bandwidth: str[1], ResponseTime: str[2], Provider: str[3]}
					datas = append(datas, smsData)
				}
			}
		}
	}
	return datas
}

func smsSortByCountry(b []SMSData) []SMSData {
	sort.Slice(b, func(i, j int) bool { return b[i].Country < b[j].Country })
	return b
}

func smsSortByProvider(b []SMSData) []SMSData {
	sort.Slice(b, func(i, j int) bool { return b[i].Provider < b[j].Provider })
	return b
}
