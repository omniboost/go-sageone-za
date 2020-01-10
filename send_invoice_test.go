package aktiva_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSendInvoice(t *testing.T) {
	b := []byte(`
		{
			"Customer": {
				"Name": "Leon Bogaert",
				"RegNo": "1122334455",
				"NotTDCustomer": false,
				"VatRegNo": "11223344",
				"CurrencyCode": "EUR",
				"PaymentDeadLine": 7,
				"OverDueCharge": 0,
				"RefNoBase": 1,
				"Address": "Stadhuisplein 3",
				"CountryCode": "NL",
				"County": "Netherlands",
				"City": "Terneuzen",
				"PostalCode": "",
				"PhoneNo": "6548765",
				"PhoneNo2": "",
				"HomePage": "https://www.omniboost.io",
				"Email": "leon@omniboost.io"
			},
			"DocDate": "20200101",
			"DueDate": "20200101",
			"InvoiceNo": "123",
			"RefNo": "222",
			"DepartmentCode": "",
			"ProjectCode": "",
			"InvoiceRow": [{
				"Item": {
					"Code": "1234567",
					"Description": "Bag of goldflakes",
					"Type": 3,
					"UOMName": "kg"
				},
				"Quantity": 2,
				"Price": 1000,
				"DiscountPct": 0,
				"DiscountAmount": 0,
				"TaxId": "973a4395-665f-47a6-a5b6-5384dd24f8d0",
				"LocationCode": "1",
				"GLAccountCode": "3060"
			}],
			"TotalAmount": 2000,
			"RoundingAmount": 0,
			"TaxAmount": [],
			"HComment": "Header",
			"FComment": "Footer"
		}
	`)

	req := client.NewSendInvoiceRequest()
	err := json.Unmarshal(b, req.RequestBody())
	if err != nil {
		t.Error(err)
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
