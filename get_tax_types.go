package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/odata"
	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewGetTaxTypesRequest() GetTaxTypesRequest {
	r := GetTaxTypesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetTaxTypesQueryParams()
	r.pathParams = r.NewGetTaxTypesPathParams()
	r.requestBody = r.NewGetTaxTypesRequestBody()
	return r
}

type GetTaxTypesRequest struct {
	client      *Client
	queryParams *GetTaxTypesQueryParams
	pathParams  *GetTaxTypesPathParams
	method      string
	headers     http.Header
	requestBody GetTaxTypesRequestBody
}

func (r GetTaxTypesRequest) NewGetTaxTypesQueryParams() *GetTaxTypesQueryParams {
	return &GetTaxTypesQueryParams{
		Pagination: odata.NewPagination(),
	}
}

type GetTaxTypesQueryParams struct {
	odata.Pagination
	CompanyID int `schema:"CompanyId"`
}

func (p GetTaxTypesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTaxTypesRequest) QueryParams() *GetTaxTypesQueryParams {
	return r.queryParams
}

func (r GetTaxTypesRequest) NewGetTaxTypesPathParams() *GetTaxTypesPathParams {
	return &GetTaxTypesPathParams{}
}

type GetTaxTypesPathParams struct {
}

func (p *GetTaxTypesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTaxTypesRequest) PathParams() *GetTaxTypesPathParams {
	return r.pathParams
}

func (r *GetTaxTypesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTaxTypesRequest) Method() string {
	return r.method
}

func (r GetTaxTypesRequest) NewGetTaxTypesRequestBody() GetTaxTypesRequestBody {
	return GetTaxTypesRequestBody{}
}

type GetTaxTypesRequestBody struct{}

func (r *GetTaxTypesRequest) RequestBody() *GetTaxTypesRequestBody {
	return &r.requestBody
}

func (r *GetTaxTypesRequest) SetRequestBody(body GetTaxTypesRequestBody) {
	r.requestBody = body
}

func (r *GetTaxTypesRequest) NewResponseBody() *GetTaxTypesResponseBody {
	return &GetTaxTypesResponseBody{}
}

type GetTaxTypesResponseBody struct {
	TotalResults    int `json:"TotalResults"`
	ReturnedResults int `json:"ReturnedResults"`
	Results         TaxTypes
}

func (r *GetTaxTypesRequest) URL() url.URL {
	return r.client.GetEndpointURL("/TaxType/Get", r.PathParams())
}

func (r *GetTaxTypesRequest) Do() (GetTaxTypesResponseBody, error) {
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

type TaxTypes []TaxType

type TaxType struct {
	ID                int     `json:"ID"`
	Name              string  `json:"Name"`
	Percentage        float64 `json:"Percentage"`
	IsDefault         bool    `json:"IsDefault"`
	HasActivity       bool    `json:"HasActivity"`
	IsManualTax       bool    `json:"IsManualTax"`
	Active            bool    `json:"Active"`
	Created           string  `json:"Created"`
	Modified          string  `json:"Modified,omitempty"`
	TaxTypeDefaultUID string  `json:"TaxTypeDefaultUID,omitempty"`
}
