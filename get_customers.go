package aktiva

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-merit-aktiva/utils"
)

func (c *Client) NewGetCustomersRequest() GetCustomersRequest {
	r := GetCustomersRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetCustomersQueryParams()
	r.pathParams = r.NewGetCustomersPathParams()
	r.requestBody = r.NewGetCustomersRequestBody()
	return r
}

type GetCustomersRequest struct {
	client      *Client
	queryParams *GetCustomersQueryParams
	pathParams  *GetCustomersPathParams
	method      string
	headers     http.Header
	requestBody GetCustomersRequestBody
}

func (r GetCustomersRequest) NewGetCustomersQueryParams() *GetCustomersQueryParams {
	return &GetCustomersQueryParams{}
}

type GetCustomersQueryParams struct{}

func (p GetCustomersQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomersRequest) QueryParams() *GetCustomersQueryParams {
	return r.queryParams
}

func (r GetCustomersRequest) NewGetCustomersPathParams() *GetCustomersPathParams {
	return &GetCustomersPathParams{}
}

type GetCustomersPathParams struct {
}

func (p *GetCustomersPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomersRequest) PathParams() *GetCustomersPathParams {
	return r.pathParams
}

func (r *GetCustomersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomersRequest) Method() string {
	return r.method
}

func (r GetCustomersRequest) NewGetCustomersRequestBody() GetCustomersRequestBody {
	return GetCustomersRequestBody{}
}

type GetCustomersRequestBody struct {
}

func (r *GetCustomersRequest) RequestBody() *GetCustomersRequestBody {
	return &r.requestBody
}

func (r *GetCustomersRequest) SetRequestBody(body GetCustomersRequestBody) {
	r.requestBody = body
}

func (r *GetCustomersRequest) NewResponseBody() *GetCustomersResponseBody {
	return &GetCustomersResponseBody{}
}

type GetCustomersResponseBody Customers

func (r *GetCustomersRequest) URL() url.URL {
	return r.client.GetEndpointURL("getcustomers", r.PathParams())
}

func (r *GetCustomersRequest) Do() (GetCustomersResponseBody, error) {
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

type Customers []Customer

type Customer struct {
	CustomerID        string      `json:"CustomerId"`
	Name              string      `json:"Name"`
	RegNo             string      `json:"RegNo"`
	Contact           interface{} `json:"Contact"`
	PhoneNo           string      `json:"PhoneNo"`
	PhoneNo2          string      `json:"PhoneNo2"`
	Address           string      `json:"Address"`
	City              string      `json:"City"`
	County            string      `json:"County"`
	PostalCode        string      `json:"PostalCode"`
	CountryCode       string      `json:"CountryCode"`
	CountryName       string      `json:"CountryName"`
	FaxNo             string      `json:"FaxNo"`
	Email             string      `json:"Email"`
	HomePage          string      `json:"HomePage"`
	PaymentDeadLine   int         `json:"PaymentDeadLine"`
	OverdueCharge     float64     `json:"OverdueCharge"`
	CurrencyCode      string      `json:"CurrencyCode"`
	CustomerGroupName string      `json:"CustomerGroupName"`
	VatRegNo          string      `json:"VatRegNo"`
	BankName          string      `json:"BankName"`
	NotTDCustomer     bool        `json:"NotTDCustomer"`
	SalesInvLang      string      `json:"SalesInvLang"`
	RefNoase          string      `json:"RefNoase"`
}
