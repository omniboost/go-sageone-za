package sage_test

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"testing"

	sage "github.com/omniboost/go-sageone-za"
)

var (
	client    *sage.Client
	companyID int
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	apiKey := os.Getenv("API_KEY")
	debug := os.Getenv("DEBUG")
	id := os.Getenv("TEST_COMPANY_ID")
	companyID, err = strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	client = sage.NewClient(nil, user, password, apiKey)
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
