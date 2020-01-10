package aktiva

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-merit-aktiva/utils"
)

func (c *Client) NewGetGLBatchesRequest() GetGLBatchesRequest {
	r := GetGLBatchesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetGLBatchesQueryParams()
	r.pathParams = r.NewGetGLBatchesPathParams()
	r.requestBody = r.NewGetGLBatchesRequestBody()
	return r
}

type GetGLBatchesRequest struct {
	client      *Client
	queryParams *GetGLBatchesQueryParams
	pathParams  *GetGLBatchesPathParams
	method      string
	headers     http.Header
	requestBody GetGLBatchesRequestBody
}

func (r GetGLBatchesRequest) NewGetGLBatchesQueryParams() *GetGLBatchesQueryParams {
	return &GetGLBatchesQueryParams{}
}

type GetGLBatchesQueryParams struct {
}

func (p GetGLBatchesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetGLBatchesRequest) QueryParams() *GetGLBatchesQueryParams {
	return r.queryParams
}

func (r GetGLBatchesRequest) NewGetGLBatchesPathParams() *GetGLBatchesPathParams {
	return &GetGLBatchesPathParams{}
}

type GetGLBatchesPathParams struct {
}

func (p *GetGLBatchesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetGLBatchesRequest) PathParams() *GetGLBatchesPathParams {
	return r.pathParams
}

func (r *GetGLBatchesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetGLBatchesRequest) Method() string {
	return r.method
}

func (r GetGLBatchesRequest) NewGetGLBatchesRequestBody() GetGLBatchesRequestBody {
	return GetGLBatchesRequestBody{}
}

type GetGLBatchesRequestBody struct {
	PeriodStart Date
	PeriodEnd   Date
}

func (r *GetGLBatchesRequest) RequestBody() *GetGLBatchesRequestBody {
	return &r.requestBody
}

func (r *GetGLBatchesRequest) SetRequestBody(body GetGLBatchesRequestBody) {
	r.requestBody = body
}

func (r *GetGLBatchesRequest) NewResponseBody() *GetGLBatchesResponseBody {
	return &GetGLBatchesResponseBody{}
}

type GetGLBatchesResponseBody []struct {
	GLBID        string      `json:"GLBId"`
	BatchCode    string      `json:"BatchCode"`
	No           int         `json:"No"`
	Document     interface{} `json:"Document"`
	BatchDate    string      `json:"BatchDate"`
	CurrencyCode string      `json:"CurrencyCode"`
	CurrencyRate float64     `json:"CurrencyRate"`
	TotalAmount  float64     `json:"TotalAmount"`
	PriceInclVat int         `json:"PriceInclVat"`
}

func (r *GetGLBatchesRequest) URL() url.URL {
	return r.client.GetEndpointURL("getglbatches", r.PathParams())
}

func (r *GetGLBatchesRequest) Do() (GetGLBatchesResponseBody, error) {
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
