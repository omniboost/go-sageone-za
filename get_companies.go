package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewGetCompaniesRequest() GetCompaniesRequest {
	r := GetCompaniesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetCompaniesQueryParams()
	r.pathParams = r.NewGetCompaniesPathParams()
	r.requestBody = r.NewGetCompaniesRequestBody()
	return r
}

type GetCompaniesRequest struct {
	client      *Client
	queryParams *GetCompaniesQueryParams
	pathParams  *GetCompaniesPathParams
	method      string
	headers     http.Header
	requestBody GetCompaniesRequestBody
}

func (r GetCompaniesRequest) NewGetCompaniesQueryParams() *GetCompaniesQueryParams {
	return &GetCompaniesQueryParams{}
}

type GetCompaniesQueryParams struct{}

func (p GetCompaniesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCompaniesRequest) QueryParams() *GetCompaniesQueryParams {
	return r.queryParams
}

func (r GetCompaniesRequest) NewGetCompaniesPathParams() *GetCompaniesPathParams {
	return &GetCompaniesPathParams{}
}

type GetCompaniesPathParams struct {
}

func (p *GetCompaniesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCompaniesRequest) PathParams() *GetCompaniesPathParams {
	return r.pathParams
}

func (r *GetCompaniesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCompaniesRequest) Method() string {
	return r.method
}

func (r GetCompaniesRequest) NewGetCompaniesRequestBody() GetCompaniesRequestBody {
	return GetCompaniesRequestBody{}
}

type GetCompaniesRequestBody struct {
}

func (r *GetCompaniesRequest) RequestBody() *GetCompaniesRequestBody {
	return &r.requestBody
}

func (r *GetCompaniesRequest) SetRequestBody(body GetCompaniesRequestBody) {
	r.requestBody = body
}

func (r *GetCompaniesRequest) NewResponseBody() *GetCompaniesResponseBody {
	return &GetCompaniesResponseBody{}
}

type GetCompaniesResponseBody struct {
	TotalResults    int
	ReturnedResults int
	Results         Companies
}

func (r *GetCompaniesRequest) URL() url.URL {
	return r.client.GetEndpointURL("/Company/Get", r.PathParams())
}

func (r *GetCompaniesRequest) Do() (GetCompaniesResponseBody, error) {
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

type Companies []Company

type Company struct {
	ID                                         int    `json:"ID"`
	Name                                       string `json:"Name"`
	CurrencySymbol                             string `json:"CurrencySymbol"`
	CurrencyDecimalDigits                      int    `json:"CurrencyDecimalDigits"`
	NumberDecimalDigits                        int    `json:"NumberDecimalDigits"`
	DecimalSeparator                           string `json:"DecimalSeparator"`
	HoursDecimalDigits                         int    `json:"HoursDecimalDigits"`
	ItemCostPriceDecimalDigits                 int    `json:"ItemCostPriceDecimalDigits"`
	ItemSellingPriceDecimalDigits              int    `json:"ItemSellingPriceDecimalDigits"`
	PostalAddress1                             string `json:"PostalAddress1"`
	PostalAddress2                             string `json:"PostalAddress2"`
	PostalAddress3                             string `json:"PostalAddress3"`
	PostalAddress4                             string `json:"PostalAddress4"`
	PostalAddress5                             string `json:"PostalAddress5"`
	GroupSeparator                             string `json:"GroupSeparator"`
	RoundingValue                              int    `json:"RoundingValue"`
	TaxSystem                                  int    `json:"TaxSystem"`
	RoundingType                               int    `json:"RoundingType"`
	AgeMonthly                                 bool   `json:"AgeMonthly"`
	DisplayInactiveItems                       bool   `json:"DisplayInactiveItems"`
	WarnWhenItemCostIsZero                     bool   `json:"WarnWhenItemCostIsZero"`
	DoNotAllowProcessingIntoNegativeQuantities bool   `json:"DoNotAllowProcessingIntoNegativeQuantities"`
	WarnWhenItemQuantityIsZero                 bool   `json:"WarnWhenItemQuantityIsZero"`
	WarnWhenItemSellingBelowCost               bool   `json:"WarnWhenItemSellingBelowCost"`
	CountryID                                  int    `json:"CountryId"`
	CompanyEntityTypeID                        int    `json:"CompanyEntityTypeId"`
	TakeOnBalanceDate                          string `json:"TakeOnBalanceDate"`
	EnableCustomerZone                         bool   `json:"EnableCustomerZone"`
	EnableAutomaticBankFeedRefresh             bool   `json:"EnableAutomaticBankFeedRefresh"`
	ContactName                                string `json:"ContactName"`
	Telephone                                  string `json:"Telephone"`
	Fax                                        string `json:"Fax"`
	Mobile                                     string `json:"Mobile"`
	Email                                      string `json:"Email"`
	IsPrimarySendingEmail                      bool   `json:"IsPrimarySendingEmail"`
	PostalAddress01                            string `json:"PostalAddress01"`
	PostalAddress02                            string `json:"PostalAddress02"`
	PostalAddress03                            string `json:"PostalAddress03"`
	PostalAddress04                            string `json:"PostalAddress04"`
	PostalAddress05                            string `json:"PostalAddress05"`
	CompanyInfo01                              string `json:"CompanyInfo01"`
	CompanyInfo02                              string `json:"CompanyInfo02"`
	CompanyInfo03                              string `json:"CompanyInfo03"`
	CompanyInfo04                              string `json:"CompanyInfo04"`
	CompanyInfo05                              string `json:"CompanyInfo05"`
	IsOwner                                    bool   `json:"IsOwner"`
	UseCCEmail                                 bool   `json:"UseCCEmail"`
	CCEmail                                    string `json:"CCEmail"`
	DateFormatID                               int    `json:"DateFormatId"`
	CheckForDuplicateCustomerReferences        bool   `json:"CheckForDuplicateCustomerReferences"`
	CheckForDuplicateSupplierReferences        bool   `json:"CheckForDuplicateSupplierReferences"`
	DisplayName                                string `json:"DisplayName"`
	DisplayInactiveCustomers                   bool   `json:"DisplayInactiveCustomers"`
	DisplayInactiveSuppliers                   bool   `json:"DisplayInactiveSuppliers"`
	DisplayInactiveTimeProjects                bool   `json:"DisplayInactiveTimeProjects"`
	UseInclusiveProcessingByDefault            bool   `json:"UseInclusiveProcessingByDefault"`
	LockProcessing                             bool   `json:"LockProcessing"`
	LockProcessingDate                         string `json:"LockProcessingDate"`
	LockTimesheetProcessing                    bool   `json:"LockTimesheetProcessing"`
	LockTimesheetProcessingDate                string `json:"LockTimesheetProcessingDate"`
	LockedForConversion                        bool   `json:"LockedForConversion"`
	LockedForTransfer                          bool   `json:"LockedForTransfer"`
	TaxPeriodFrequency                         int    `json:"TaxPeriodFrequency"`
	PreviousTaxPeriodEndDate                   string `json:"PreviousTaxPeriodEndDate"`
	PreviousTaxReturnDate                      string `json:"PreviousTaxReturnDate"`
	UseNoreplyEmail                            bool   `json:"UseNoreplyEmail"`
	AgeingBasedOnDueDate                       bool   `json:"AgeingBasedOnDueDate"`
	UseLogoOnEmails                            bool   `json:"UseLogoOnEmails"`
	UseLogoOnCustomerZone                      bool   `json:"UseLogoOnCustomerZone"`
	City                                       string `json:"City"`
	State                                      string `json:"State"`
	Country                                    string `json:"Country"`
	HomeCurrencyID                             int    `json:"HomeCurrencyId"`
	CurrencyID                                 int    `json:"CurrencyId"`
	Created                                    string `json:"Created"`
	Modified                                   string `json:"Modified"`
	Active                                     bool   `json:"Active"`
	TaxNumber                                  string `json:"TaxNumber"`
	RegisteredName                             string `json:"RegisteredName"`
	RegistrationNumber                         string `json:"RegistrationNumber"`
	IsPracticeAccount                          bool   `json:"IsPracticeAccount"`
	LogoPositionID                             int    `json:"LogoPositionID"`
	Attachment                                 struct {
		ID        int    `json:"ID"`
		Image     string `json:"Image"`
		Timestamp string `json:"Timestamp"`
	} `json:"Attachment,omitempty"`
	CompanyTaxNumber                 string `json:"CompanyTaxNumber"`
	TaxOffice                        string `json:"TaxOffice"`
	CustomerZoneGUID                 string `json:"CustomerZoneGuid"`
	ClientTypeID                     int    `json:"ClientTypeId"`
	DisplayTotalTypeID               int    `json:"DisplayTotalTypeId"`
	DisplayInCompanyConsole          bool   `json:"DisplayInCompanyConsole"`
	LastLoginDate                    string `json:"LastLoginDate"`
	Status                           int    `json:"Status"`
	TaxReportingTypeID               int    `json:"TaxReportingTypeId"`
	SalesOrdersReserveItemQuantities bool   `json:"SalesOrdersReserveItemQuantities"`
	PrescribedGoodsTrader            bool   `json:"PrescribedGoodsTrader"`
	DisplayInactiveItemBundles       bool   `json:"DisplayInactiveItemBundles"`
	CompanyTransferStatus            int    `json:"CompanyTransferStatus"`
}
