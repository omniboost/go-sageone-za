package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewGetTaxInvoicesRequest() GetTaxInvoicesRequest {
	r := GetTaxInvoicesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetTaxInvoicesQueryParams()
	r.pathParams = r.NewGetTaxInvoicesPathParams()
	r.requestBody = r.NewGetTaxInvoicesRequestBody()
	return r
}

type GetTaxInvoicesRequest struct {
	client      *Client
	queryParams *GetTaxInvoicesQueryParams
	pathParams  *GetTaxInvoicesPathParams
	method      string
	headers     http.Header
	requestBody GetTaxInvoicesRequestBody
}

func (r GetTaxInvoicesRequest) NewGetTaxInvoicesQueryParams() *GetTaxInvoicesQueryParams {
	return &GetTaxInvoicesQueryParams{}
}

type GetTaxInvoicesQueryParams struct {
	CompanyID              int  `schema:"CompanyId"`
	IncludeDetail          bool `schema:"includeDetail"`
	IncludeCustomerDetails bool `schema:"includeCustomerDetails"`
}

func (p GetTaxInvoicesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTaxInvoicesRequest) QueryParams() *GetTaxInvoicesQueryParams {
	return r.queryParams
}

func (r GetTaxInvoicesRequest) NewGetTaxInvoicesPathParams() *GetTaxInvoicesPathParams {
	return &GetTaxInvoicesPathParams{}
}

type GetTaxInvoicesPathParams struct {
}

func (p *GetTaxInvoicesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTaxInvoicesRequest) PathParams() *GetTaxInvoicesPathParams {
	return r.pathParams
}

func (r *GetTaxInvoicesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTaxInvoicesRequest) Method() string {
	return r.method
}

func (r GetTaxInvoicesRequest) NewGetTaxInvoicesRequestBody() GetTaxInvoicesRequestBody {
	return GetTaxInvoicesRequestBody{}
}

type GetTaxInvoicesRequestBody struct {
}

func (r *GetTaxInvoicesRequest) RequestBody() *GetTaxInvoicesRequestBody {
	return &r.requestBody
}

func (r *GetTaxInvoicesRequest) SetRequestBody(body GetTaxInvoicesRequestBody) {
	r.requestBody = body
}

func (r *GetTaxInvoicesRequest) NewResponseBody() *GetTaxInvoicesResponseBody {
	return &GetTaxInvoicesResponseBody{}
}

type GetTaxInvoicesResponseBody struct {
	TotalResults    int
	ReturnedResults int
	Results         TaxInvoices
}

func (r *GetTaxInvoicesRequest) URL() url.URL {
	return r.client.GetEndpointURL("/TaxInvoice/Get", r.PathParams())
}

func (r *GetTaxInvoicesRequest) Do() (GetTaxInvoicesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

type TaxInvoices []TaxInvoice

type TaxInvoice struct {
	DueDate            string `json:"DueDate"`
	FromDocument       string `json:"FromDocument"`
	FromDocumentID     int    `json:"FromDocumentId"`
	FromDocumentTypeID int    `json:"FromDocumentTypeId"`
	AllowOnlinePayment bool   `json:"AllowOnlinePayment"`
	Paid               bool   `json:"Paid"`
	Status             string `json:"Status"`
	Locked             bool   `json:"Locked"`
	CustomerID         int    `json:"CustomerId"`
	CustomerName       string `json:"CustomerName"`
	Customer           struct {
		Name     string `json:"Name"`
		Category struct {
			Description string `json:"Description"`
			ID          int    `json:"ID"`
			Modified    string `json:"Modified"`
			Created     string `json:"Created"`
		} `json:"Category"`
		SalesRepresentativeID int `json:"SalesRepresentativeId"`
		SalesRepresentative   struct {
			ID        int    `json:"ID"`
			FirstName string `json:"FirstName"`
			LastName  string `json:"LastName"`
			Name      string `json:"Name"`
			Active    bool   `json:"Active"`
			Email     string `json:"Email"`
			Mobile    string `json:"Mobile"`
			Telephone string `json:"Telephone"`
			Created   string `json:"Created"`
			Modified  string `json:"Modified"`
		} `json:"SalesRepresentative"`
		TaxReference                string  `json:"TaxReference"`
		ContactName                 string  `json:"ContactName"`
		Telephone                   string  `json:"Telephone"`
		Fax                         string  `json:"Fax"`
		Mobile                      string  `json:"Mobile"`
		Email                       string  `json:"Email"`
		WebAddress                  string  `json:"WebAddress"`
		Active                      bool    `json:"Active"`
		IsObfuscated                bool    `json:"IsObfuscated"`
		Balance                     float64 `json:"Balance"`
		CreditLimit                 float64 `json:"CreditLimit"`
		CommunicationMethod         int     `json:"CommunicationMethod"`
		PostalAddress01             string  `json:"PostalAddress01"`
		PostalAddress02             string  `json:"PostalAddress02"`
		PostalAddress03             string  `json:"PostalAddress03"`
		PostalAddress04             string  `json:"PostalAddress04"`
		PostalAddress05             string  `json:"PostalAddress05"`
		DeliveryAddress01           string  `json:"DeliveryAddress01"`
		DeliveryAddress02           string  `json:"DeliveryAddress02"`
		DeliveryAddress03           string  `json:"DeliveryAddress03"`
		DeliveryAddress04           string  `json:"DeliveryAddress04"`
		DeliveryAddress05           string  `json:"DeliveryAddress05"`
		AutoAllocateToOldestInvoice bool    `json:"AutoAllocateToOldestInvoice"`
		EnableCustomerZone          bool    `json:"EnableCustomerZone"`
		CustomerZoneGUID            string  `json:"CustomerZoneGuid"`
		CashSale                    bool    `json:"CashSale"`
		TextField1                  string  `json:"TextField1"`
		TextField2                  string  `json:"TextField2"`
		TextField3                  string  `json:"TextField3"`
		NumericField1               float64 `json:"NumericField1"`
		NumericField2               float64 `json:"NumericField2"`
		NumericField3               float64 `json:"NumericField3"`
		YesNoField1                 bool    `json:"YesNoField1"`
		YesNoField2                 bool    `json:"YesNoField2"`
		YesNoField3                 bool    `json:"YesNoField3"`
		DateField1                  string  `json:"DateField1"`
		DateField2                  string  `json:"DateField2"`
		DateField3                  string  `json:"DateField3"`
		DefaultPriceListID          int     `json:"DefaultPriceListId"`
		DefaultPriceList            struct {
			ID          int    `json:"ID"`
			Description string `json:"Description"`
			IsDefault   bool   `json:"IsDefault"`
		} `json:"DefaultPriceList"`
		DefaultPriceListName       string  `json:"DefaultPriceListName"`
		AcceptsElectronicInvoices  bool    `json:"AcceptsElectronicInvoices"`
		Modified                   string  `json:"Modified"`
		Created                    string  `json:"Created"`
		BusinessRegistrationNumber string  `json:"BusinessRegistrationNumber"`
		TaxStatusVerified          string  `json:"TaxStatusVerified"`
		CurrencyID                 int     `json:"CurrencyId"`
		CurrencySymbol             string  `json:"CurrencySymbol"`
		HasActivity                bool    `json:"HasActivity"`
		DefaultDiscountPercentage  float64 `json:"DefaultDiscountPercentage"`
		DefaultTaxTypeID           int     `json:"DefaultTaxTypeId"`
		DefaultTaxType             struct {
			ID                int     `json:"ID"`
			Name              string  `json:"Name"`
			Percentage        float64 `json:"Percentage"`
			IsDefault         bool    `json:"IsDefault"`
			HasActivity       bool    `json:"HasActivity"`
			IsManualTax       bool    `json:"IsManualTax"`
			Active            bool    `json:"Active"`
			Created           string  `json:"Created"`
			Modified          string  `json:"Modified"`
			TaxTypeDefaultUID string  `json:"TaxTypeDefaultUID"`
		} `json:"DefaultTaxType"`
		DueDateMethodID              int  `json:"DueDateMethodId"`
		DueDateMethodValue           int  `json:"DueDateMethodValue"`
		CityID                       int  `json:"CityId"`
		HasSpecialCountryTax         bool `json:"HasSpecialCountryTax"`
		AccountingAgreement          bool `json:"AccountingAgreement"`
		HasSpecialCountryTaxActivity bool `json:"HasSpecialCountryTaxActivity"`
		ID                           int  `json:"ID"`
	} `json:"Customer"`
	SalesRepresentativeID int `json:"SalesRepresentativeId"`
	SalesRepresentative   struct {
		ID        int    `json:"ID"`
		FirstName string `json:"FirstName"`
		LastName  string `json:"LastName"`
		Name      string `json:"Name"`
		Active    bool   `json:"Active"`
		Email     string `json:"Email"`
		Mobile    string `json:"Mobile"`
		Telephone string `json:"Telephone"`
		Created   string `json:"Created"`
		Modified  string `json:"Modified"`
	} `json:"SalesRepresentative"`
	StatusID             int     `json:"StatusId"`
	Modified             string  `json:"Modified"`
	Created              string  `json:"Created"`
	CustomerCurrencyID   int     `json:"Customer_CurrencyId"`
	CustomerExchangeRate float64 `json:"Customer_ExchangeRate"`
	ID                   int     `json:"ID"`
	Date                 string  `json:"Date"`
	Inclusive            bool    `json:"Inclusive"`
	DiscountPercentage   float64 `json:"DiscountPercentage"`
	TaxReference         string  `json:"TaxReference"`
	DocumentNumber       string  `json:"DocumentNumber"`
	Reference            string  `json:"Reference"`
	Message              string  `json:"Message"`
	Discount             float64 `json:"Discount"`
	Exclusive            float64 `json:"Exclusive"`
	Tax                  float64 `json:"Tax"`
	Rounding             float64 `json:"Rounding"`
	Total                float64 `json:"Total"`
	AmountDue            float64 `json:"AmountDue"`
	PostalAddress01      string  `json:"PostalAddress01"`
	PostalAddress02      string  `json:"PostalAddress02"`
	PostalAddress03      string  `json:"PostalAddress03"`
	PostalAddress04      string  `json:"PostalAddress04"`
	PostalAddress05      string  `json:"PostalAddress05"`
	DeliveryAddress01    string  `json:"DeliveryAddress01"`
	DeliveryAddress02    string  `json:"DeliveryAddress02"`
	DeliveryAddress03    string  `json:"DeliveryAddress03"`
	DeliveryAddress04    string  `json:"DeliveryAddress04"`
	DeliveryAddress05    string  `json:"DeliveryAddress05"`
	Printed              bool    `json:"Printed"`
	TaxPeriodID          int     `json:"TaxPeriodId"`
	Editable             bool    `json:"Editable"`
	HasAttachments       bool    `json:"HasAttachments"`
	HasNotes             bool    `json:"HasNotes"`
	HasAnticipatedDate   bool    `json:"HasAnticipatedDate"`
	AnticipatedDate      string  `json:"AnticipatedDate"`
	ExternalReference    string  `json:"ExternalReference"`
	UID                  string  `json:"UID"`
	Lines                []struct {
		SelectionID         int     `json:"SelectionId"`
		TaxTypeID           int     `json:"TaxTypeId"`
		ID                  int     `json:"ID"`
		Description         string  `json:"Description"`
		LineType            int     `json:"LineType"`
		Quantity            float64 `json:"Quantity"`
		UnitPriceExclusive  float64 `json:"UnitPriceExclusive"`
		Unit                string  `json:"Unit"`
		UnitPriceInclusive  float64 `json:"UnitPriceInclusive"`
		TaxPercentage       float64 `json:"TaxPercentage"`
		DiscountPercentage  float64 `json:"DiscountPercentage"`
		Exclusive           float64 `json:"Exclusive"`
		Discount            float64 `json:"Discount"`
		Tax                 float64 `json:"Tax"`
		Total               float64 `json:"Total"`
		Comments            string  `json:"Comments"`
		AnalysisCategoryID1 int     `json:"AnalysisCategoryId1"`
		AnalysisCategoryID2 int     `json:"AnalysisCategoryId2"`
		AnalysisCategoryID3 int     `json:"AnalysisCategoryId3"`
		TrackingCode        string  `json:"$TrackingCode"`
		CurrencyID          int     `json:"CurrencyId"`
		UnitCost            float64 `json:"UnitCost"`
		UID                 string  `json:"UID"`
	} `json:"Lines"`
}
