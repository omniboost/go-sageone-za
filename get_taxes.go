package aktiva

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-merit-aktiva/utils"
)

func (c *Client) NewGetTaxesRequest() GetTaxesRequest {
	r := GetTaxesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetTaxesQueryParams()
	r.pathParams = r.NewGetTaxesPathParams()
	r.requestBody = r.NewGetTaxesRequestBody()
	return r
}

type GetTaxesRequest struct {
	client      *Client
	queryParams *GetTaxesQueryParams
	pathParams  *GetTaxesPathParams
	method      string
	headers     http.Header
	requestBody GetTaxesRequestBody
}

func (r GetTaxesRequest) NewGetTaxesQueryParams() *GetTaxesQueryParams {
	return &GetTaxesQueryParams{}
}

type GetTaxesQueryParams struct{}

func (p GetTaxesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTaxesRequest) QueryParams() *GetTaxesQueryParams {
	return r.queryParams
}

func (r GetTaxesRequest) NewGetTaxesPathParams() *GetTaxesPathParams {
	return &GetTaxesPathParams{}
}

type GetTaxesPathParams struct {
}

func (p *GetTaxesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTaxesRequest) PathParams() *GetTaxesPathParams {
	return r.pathParams
}

func (r *GetTaxesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTaxesRequest) Method() string {
	return r.method
}

func (r GetTaxesRequest) NewGetTaxesRequestBody() GetTaxesRequestBody {
	return GetTaxesRequestBody{}
}

type GetTaxesRequestBody struct {
}

func (r *GetTaxesRequest) RequestBody() *GetTaxesRequestBody {
	return &r.requestBody
}

func (r *GetTaxesRequest) SetRequestBody(body GetTaxesRequestBody) {
	r.requestBody = body
}

func (r *GetTaxesRequest) NewResponseBody() *GetTaxesResponseBody {
	return &GetTaxesResponseBody{}
}

type GetTaxesResponseBody Taxes

func (r *GetTaxesRequest) URL() url.URL {
	return r.client.GetEndpointURL("gettaxes", r.PathParams())
}

func (r *GetTaxesRequest) Do() (GetTaxesResponseBody, error) {
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

type Taxes []Tax

type Tax struct {
}
