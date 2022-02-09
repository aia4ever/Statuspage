package result

import (
	"simulator/pkg/billing"
	"simulator/pkg/email"
	"simulator/pkg/incident"
	"simulator/pkg/mms"
	"simulator/pkg/sms"
	"simulator/pkg/support"
	"simulator/pkg/voice"
	"sync"
)

type ResultSetT struct {
	SMS       [][]sms.SMSData         `json:"sms"`
	MMS       [][]mms.MMSData         `json:"mms"`
	VoiceCall []voice.VoiceCallData   `json:"voice_call"`
	Email     [][]email.EmailData     `json:"email"`
	Billing   billing.BillingData     `json:"billing"`
	Support   []int                   `json:"support"`
	Incidents []incident.IncidentData `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type EmailData email.EmailData

type Array []EmailData

func GetRes() ResultT {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(7)
	var res ResultT
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		smsRes := sms.Result()
		mu.Lock()
		res.Data.SMS = smsRes
		mu.Unlock()
	}(&wg, &mu)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		mmsRes := mms.Result()
		mu.Lock()
		res.Data.MMS = mmsRes
		mu.Unlock()
	}(&wg, &mu)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		voiceRes := voice.Result()
		mu.Lock()
		res.Data.VoiceCall = voiceRes
		mu.Unlock()
	}(&wg, &mu)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		emRes := email.Result()
		mu.Lock()
		res.Data.Email = emRes
		mu.Unlock()
	}(&wg, &mu)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		bilRes := billing.Result()
		mu.Lock()
		res.Data.Billing = bilRes
		mu.Unlock()
	}(&wg, &mu)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		supRes := support.Result()
		mu.Lock()
		res.Data.Support = supRes
		mu.Unlock()
	}(&wg, &mu)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		incRes := incident.Result()
		mu.Lock()
		res.Data.Incidents = incRes
		mu.Unlock()
	}(&wg, &mu)

	wg.Wait()
	res.Status = true

	return res
}
