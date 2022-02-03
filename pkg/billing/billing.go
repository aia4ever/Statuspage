package billing

import (
	"io/ioutil"
	"log"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func Result() BillingData {
	content, err := ioutil.ReadFile("billing.data")
	if err != nil {
		log.Fatal(err)
	}
	return BillingData{CreateCustomer: atob(content[5]), Purchase: atob(content[4]), Payout: atob(content[3]),
		Recurring: atob(content[2]), FraudControl: atob(content[1]), CheckoutPage: atob(content[0]),
	}
}

func atob(s byte) bool { return s != 48 }
