package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewSaveTaxInvoiceRequest() SaveTaxInvoiceRequest {
	r := SaveTaxInvoiceRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewSaveTaxInvoiceQueryParams()
	r.pathParams = r.NewSaveTaxInvoicePathParams()
	r.requestBody = r.NewSaveTaxInvoiceRequestBody()
	return r
}

type SaveTaxInvoiceRequest struct {
	client      *Client
	queryParams *SaveTaxInvoiceQueryParams
	pathParams  *SaveTaxInvoicePathParams
	method      string
	headers     http.Header
	requestBody SaveTaxInvoiceRequestBody
}

func (r SaveTaxInvoiceRequest) NewSaveTaxInvoiceQueryParams() *SaveTaxInvoiceQueryParams {
	return &SaveTaxInvoiceQueryParams{}
}

type SaveTaxInvoiceQueryParams struct {
	CompanyID int `schema:"CompanyId"`
}

func (p SaveTaxInvoiceQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SaveTaxInvoiceRequest) QueryParams() *SaveTaxInvoiceQueryParams {
	return r.queryParams
}

func (r SaveTaxInvoiceRequest) NewSaveTaxInvoicePathParams() *SaveTaxInvoicePathParams {
	return &SaveTaxInvoicePathParams{}
}

type SaveTaxInvoicePathParams struct {
}

func (p *SaveTaxInvoicePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SaveTaxInvoiceRequest) PathParams() *SaveTaxInvoicePathParams {
	return r.pathParams
}

func (r *SaveTaxInvoiceRequest) SetMethod(method string) {
	r.method = method
}

func (r *SaveTaxInvoiceRequest) Method() string {
	return r.method
}

func (r SaveTaxInvoiceRequest) NewSaveTaxInvoiceRequestBody() SaveTaxInvoiceRequestBody {
	return SaveTaxInvoiceRequestBody{}
}

type SaveTaxInvoiceRequestBody NewTaxInvoice

func (r *SaveTaxInvoiceRequest) RequestBody() *SaveTaxInvoiceRequestBody {
	return &r.requestBody
}

func (r *SaveTaxInvoiceRequest) SetRequestBody(body SaveTaxInvoiceRequestBody) {
	r.requestBody = body
}

func (r *SaveTaxInvoiceRequest) NewResponseBody() *SaveTaxInvoiceResponseBody {
	return &SaveTaxInvoiceResponseBody{}
}

type SaveTaxInvoiceResponseBody TaxInvoice

func (r *SaveTaxInvoiceRequest) URL() url.URL {
	return r.client.GetEndpointURL("/TaxInvoice/Save", r.PathParams())
}

func (r *SaveTaxInvoiceRequest) Do() (SaveTaxInvoiceResponseBody, error) {
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
