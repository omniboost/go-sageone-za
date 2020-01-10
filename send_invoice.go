package aktiva

import (
	"net/http"
	"net/url"

	"github.com/gofrs/uuid"
	"github.com/omniboost/go-merit-aktiva/utils"
)

func (c *Client) NewSendInvoiceRequest() SendInvoiceRequest {
	r := SendInvoiceRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewSendInvoiceQueryParams()
	r.pathParams = r.NewSendInvoicePathParams()
	r.requestBody = r.NewSendInvoiceRequestBody()
	return r
}

type SendInvoiceRequest struct {
	client      *Client
	queryParams *SendInvoiceQueryParams
	pathParams  *SendInvoicePathParams
	method      string
	headers     http.Header
	requestBody SendInvoiceRequestBody
}

func (r SendInvoiceRequest) NewSendInvoiceQueryParams() *SendInvoiceQueryParams {
	return &SendInvoiceQueryParams{}
}

type SendInvoiceQueryParams struct{}

func (p SendInvoiceQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SendInvoiceRequest) QueryParams() *SendInvoiceQueryParams {
	return r.queryParams
}

func (r SendInvoiceRequest) NewSendInvoicePathParams() *SendInvoicePathParams {
	return &SendInvoicePathParams{}
}

type SendInvoicePathParams struct {
}

func (p *SendInvoicePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SendInvoiceRequest) PathParams() *SendInvoicePathParams {
	return r.pathParams
}

func (r *SendInvoiceRequest) SetMethod(method string) {
	r.method = method
}

func (r *SendInvoiceRequest) Method() string {
	return r.method
}

func (r SendInvoiceRequest) NewSendInvoiceRequestBody() SendInvoiceRequestBody {
	return SendInvoiceRequestBody{
		InvoiceRow: InvoiceRows{},
		TaxAmount:  TaxAmounts{},
	}
}

type SendInvoiceRequestBody NewInvoice

func (r *SendInvoiceRequest) RequestBody() *SendInvoiceRequestBody {
	return &r.requestBody
}

func (r *SendInvoiceRequest) SetRequestBody(body SendInvoiceRequestBody) {
	r.requestBody = body
}

func (r *SendInvoiceRequest) NewResponseBody() *SendInvoiceResponseBody {
	return &SendInvoiceResponseBody{}
}

type SendInvoiceResponseBody struct {
	CustomerID  string      `json:"CustomerId"`
	InvoiceID   string      `json:"InvoiceId"`
	InvoiceNo   string      `json:"InvoiceNo"`
	RefNo       string      `json:"RefNo"`
	NewCustomer interface{} `json:"NewCustomer"`
}

func (r *SendInvoiceRequest) URL() url.URL {
	return r.client.GetEndpointURL("sendinvoice", r.PathParams())
}

func (r *SendInvoiceRequest) Do() (SendInvoiceResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
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

type NewInvoice struct {
	Customer       NewInvoiceCustomer
	DocDate        Date
	DueDate        Date
	InvoiceNo      string
	RefNo          string
	CurrencyCode   string
	DepartmentCode string
	ProjectCode    string
	InvoiceRow     InvoiceRows
	TaxAmount      TaxAmounts
	RoundingAmount float64
	TotalAmount    float64
	Payment        *Payment
	Hcomment       string
	Fcomment       string
}

type NewInvoiceCustomer struct {
	// If filled and customer is found in the database then following fields are
	// not important. If not found, the customer is added using the following
	// fields.
	ID *uuid.UUID `json:"Id,omitempty"`
	// Required when customer is added
	Name  string `json:"Name,omitempty"`
	RegNo string `json:"RegNo,omitempty"`
	// Required when customer is added. True for physical persons and foreign
	// companies. Allowed “true” or “false” (lowercase).
	NotTDCustomer bool   `json:"NotTDCustomer,omitempty"`
	VatRegNo      string `json:"VatRegNo,omitempty"`
	CurrencyCode  string `json:"CurrencyCode,omitempty"`
	// If missing then taken from default settings.
	PaymentDeadLine int `json:"PaymentDeadLine,omitempty"`
	// If missing then taken from default settings.
	OverDueCharge float64 `json:"OverDueCharge,omitempty"`
	Address       string  `json:"Address,omitempty"`
	City          string  `json:"City,omitempty"`
	Country       string  `json:"Country,omitempty"`
	PostalCode    string  `json:"PostalCode,omitempty"`
	// Required when adding
	CountryCode string
	PhoneNo     string `json:"PhoneNo,omitempty"`
	PhoneNo2    string `json:"PhoneNo2,omitempty"`
	HomePage    string `json:"HomePage,omitempty"`
	Email       string `json:"Email,omitempty"`
	// Invoice language for this specific customer.
	SalesInvLang string `json:"SalesInvLang,omitempty"`
}

type InvoiceRows []InvoiceRow

type InvoiceRow struct {
	Item           Article
	Quantity       float64
	Price          float64
	DiscountPct    float64
	DiscountAmount float64
	TaxID          uuid.UUID `json:"TaxId"`
	LocationCode   string
	DepartmentCode string
	ItemCostAmount float64
	GLAccountCode  string
	ProjectCode    string
	CostCenterCode string
}

type Article struct {
	// Required
	Code string
	// Required
	Description string
	// 1 = stock item
	// 2 = service
	// 3 = item
	// Required.
	Type int
	// Name for the unit
	UOMName string `json:"UOMName,omitempty"`
	// If company has more than one (default) stock, stock code in this field is
	// required for all stock items.
	DefLocationCode string `json:"DefLocationCode,omitempty"`
}

type TaxAmounts []TaxAmount

type TaxAmount struct {
	// Required. Use gettaxes endpoint to detect the guid needed
	TaxID  uuid.UUID `json:"TaxId"`
	Amount float64
}

type Payment struct {
	// Name of the payment method. Must be found in the company database.
	PaymentMethod string
	PaidAmount    float64
	PaymDate      Date
}
