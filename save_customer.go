package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewSaveCustomerRequest() SaveCustomerRequest {
	r := SaveCustomerRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewSaveCustomerQueryParams()
	r.pathParams = r.NewSaveCustomerPathParams()
	r.requestBody = r.NewSaveCustomerRequestBody()
	return r
}

type SaveCustomerRequest struct {
	client      *Client
	queryParams *SaveCustomerQueryParams
	pathParams  *SaveCustomerPathParams
	method      string
	headers     http.Header
	requestBody SaveCustomerRequestBody
}

func (r SaveCustomerRequest) NewSaveCustomerQueryParams() *SaveCustomerQueryParams {
	return &SaveCustomerQueryParams{}
}

type SaveCustomerQueryParams struct {
	CompanyID int `schema:"CompanyId"`
}

func (p SaveCustomerQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SaveCustomerRequest) QueryParams() *SaveCustomerQueryParams {
	return r.queryParams
}

func (r SaveCustomerRequest) NewSaveCustomerPathParams() *SaveCustomerPathParams {
	return &SaveCustomerPathParams{}
}

type SaveCustomerPathParams struct {
}

func (p *SaveCustomerPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SaveCustomerRequest) PathParams() *SaveCustomerPathParams {
	return r.pathParams
}

func (r *SaveCustomerRequest) SetMethod(method string) {
	r.method = method
}

func (r *SaveCustomerRequest) Method() string {
	return r.method
}

func (r SaveCustomerRequest) NewSaveCustomerRequestBody() SaveCustomerRequestBody {
	return SaveCustomerRequestBody{}
}

type SaveCustomerRequestBody NewCustomer

func (r *SaveCustomerRequest) RequestBody() *SaveCustomerRequestBody {
	return &r.requestBody
}

func (r *SaveCustomerRequest) SetRequestBody(body SaveCustomerRequestBody) {
	r.requestBody = body
}

func (r *SaveCustomerRequest) NewResponseBody() *SaveCustomerResponseBody {
	return &SaveCustomerResponseBody{}
}

type SaveCustomerResponseBody Customer

func (r *SaveCustomerRequest) URL() url.URL {
	return r.client.GetEndpointURL("/Customer/Save", r.PathParams())
}

func (r *SaveCustomerRequest) Do() (SaveCustomerResponseBody, error) {
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
