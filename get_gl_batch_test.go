package aktiva_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/gofrs/uuid"
)

func TestGetGLBatch(t *testing.T) {
	req := client.NewGetGLBatchRequest()
	req.RequestBody().ID = uuid.FromStringOrNil("17a6d491-3ed0-4a5e-ab28-11e18359929f")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
