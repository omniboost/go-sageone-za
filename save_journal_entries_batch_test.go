package sage_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	sage "github.com/omniboost/go-sageone-za"
)

func TestSaveJournalEntriesBatch(t *testing.T) {
	body := sage.SaveJournalEntriesBatchRequestBody{
		{
			Date:            sage.Date{time.Now()},
			AccountID:       512694,
			Reference:       sage.Date{time.Now()}.String(),
			Description:     "test",
			TaxTypeID:       117167,
			Effect:          sage.EffectCredit,
			Exclusive:       100.0,
			Tax:             15.0,
			Total:           115.0,
			ContraAccountID: 512695,
			Debit:           0.0,
			Credit:          115.0,
		},
	}

	req := client.NewSaveJournalEntriesBatchRequest()
	req.SetRequestBody(body)
	req.QueryParams().CompanyID = companyID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
