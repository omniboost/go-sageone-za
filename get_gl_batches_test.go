package aktiva_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	aktiva "github.com/omniboost/go-merit-aktiva"
)

func TestGetGLBatches(t *testing.T) {
	req := client.NewGetGLBatchesRequest()
	req.RequestBody().PeriodStart = aktiva.Date{time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().PeriodEnd = aktiva.Date{time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
