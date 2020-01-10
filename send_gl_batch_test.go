package aktiva_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSendGLBatch(t *testing.T) {
	b := []byte(`
		{
			"DocNo": "TEST",
			"BatchDate": "20200101",
			"EntryRow": [
				{
					"AccountCode": "1340",
					"DepartmentCode": "",
					"Debit": 0.0,
					"Credit": 100.0,
					"ProjectCode": "",
					"CostCenterCode": ""
				},
				{
					"AccountCode": "1000",
					"DepartmentCode": "",
					"Debit": 100.0,
					"Credit": 0.0,
					"ProjectCode": "",
					"CostCenterCode": ""
				}
			]
		}
	`)

	req := client.NewSendGLBatchRequest()
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
