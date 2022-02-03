package result

import (
	"simulator/pkg/billing"
	"simulator/pkg/email"
	"simulator/pkg/incident"
	"simulator/pkg/mms"
	"simulator/pkg/sms"
	"simulator/pkg/support"
	"simulator/pkg/voice"
)

type ResultSetT struct {
	SMS                                       [][]sms.SMSData         `json:"sms"`
	MMS                                       [][]mms.MMSData         `json:"mms"`
	VoiceCall                                 []voice.VoiceCallData   `json:"voice_call"`
	Email/* map[string]*/ [][]email.EmailData                         `json:"email"`
	Billing                                   billing.BillingData     `json:"billing"`
	Support                                   []int                   `json:"support"`
	Incidents                                 []incident.IncidentData `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type EmailData email.EmailData

type Array []EmailData

func GetRes() ResultT {

	var res ResultT

	res.Status = true

	res.Data.SMS = sms.Result()

	res.Data.MMS = mms.Result()

	res.Data.VoiceCall = voice.Result()

	res.Data.Email = email.Result()

	res.Data.Billing = billing.Result()

	res.Data.Support = support.Result()

	res.Data.Incidents = incident.Result()

	return res
}
