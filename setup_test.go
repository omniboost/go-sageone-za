package aktiva_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	aktiva "github.com/omniboost/go-merit-aktiva"
)

var (
	client *aktiva.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	apiID := os.Getenv("API_ID")
	apiKey := os.Getenv("API_KEY")
	debug := os.Getenv("DEBUG")

	client = aktiva.NewClient(nil, apiID, apiKey)
	if debug != "" {
		client.SetDebug(true)
	}
	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	client.SetDisallowUnknownFields(true)
	m.Run()
}
