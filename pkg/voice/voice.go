package voice

import (
	"io/ioutil"
	"log"
	"simulator/pkg/data"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_call_time"`
}

func Result() []VoiceCallData {
	content, err := ioutil.ReadFile("voice.data")
	var voiceCallData []VoiceCallData
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, v := range lines {
		if strings.Count(v, ";") == 7 {
			str := strings.Split(v, ";")
			if _, ok := data.Countries[str[0]]; ok {
				if _, ok := data.VoiceProviders[str[3]]; ok {
					value, err := strconv.ParseFloat(str[4], 32)

					if err != nil {
						log.Fatal(err)
					}
					connectionStability := float32(value)

					voiceData := VoiceCallData{Country: str[0], Bandwidth: str[1], ResponseTime: str[2],
						Provider: str[3], ConnectionStability: connectionStability, TTFB: converter(str[5]),
						VoicePurity: converter(str[6]), MedianOfCallsTime: converter(str[7]),
					}
					voiceCallData = append(voiceCallData, voiceData)
				}
			}
		}
	}
	return voiceCallData
}

func converter(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
