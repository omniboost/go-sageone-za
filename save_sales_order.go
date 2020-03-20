package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewSaveSalesOrderRequest() SaveSalesOrderRequest {
	r := SaveSalesOrderRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewSaveSalesOrderQueryParams()
	r.pathParams = r.NewSaveSalesOrderPathParams()
	r.requestBody = r.NewSaveSalesOrderRequestBody()
	return r
}

type SaveSalesOrderRequest struct {
	client      *Client
	queryParams *SaveSalesOrderQueryParams
	pathParams  *SaveSalesOrderPathParams
	method      string
	headers     http.Header
	requestBody SaveSalesOrderRequestBody
}

func (r SaveSalesOrderRequest) NewSaveSalesOrderQueryParams() *SaveSalesOrderQueryParams {
	return &SaveSalesOrderQueryParams{}
}

type SaveSalesOrderQueryParams struct {
	CompanyID int `schema:"CompanyId"`
}

func (p SaveSalesOrderQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SaveSalesOrderRequest) QueryParams() *SaveSalesOrderQueryParams {
	return r.queryParams
}

func (r SaveSalesOrderRequest) NewSaveSalesOrderPathParams() *SaveSalesOrderPathParams {
	return &SaveSalesOrderPathParams{}
}

type SaveSalesOrderPathParams struct {
}

func (p *SaveSalesOrderPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SaveSalesOrderRequest) PathParams() *SaveSalesOrderPathParams {
	return r.pathParams
}

func (r *SaveSalesOrderRequest) SetMethod(method string) {
	r.method = method
}

func (r *SaveSalesOrderRequest) Method() string {
	return r.method
}

func (r SaveSalesOrderRequest) NewSaveSalesOrderRequestBody() SaveSalesOrderRequestBody {
	return SaveSalesOrderRequestBody{}
}

type SaveSalesOrderRequestBody NewSalesOrder

func (r *SaveSalesOrderRequest) RequestBody() *SaveSalesOrderRequestBody {
	return &r.requestBody
}

func (r *SaveSalesOrderRequest) SetRequestBody(body SaveSalesOrderRequestBody) {
	r.requestBody = body
}

func (r *SaveSalesOrderRequest) NewResponseBody() *SaveSalesOrderResponseBody {
	return &SaveSalesOrderResponseBody{}
}

type SaveSalesOrderResponseBody SalesOrder

func (r *SaveSalesOrderRequest) URL() url.URL {
	return r.client.GetEndpointURL("/SalesOrder/Save", r.PathParams())
}

func (r *SaveSalesOrderRequest) Do() (SaveSalesOrderResponseBody, error) {
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
