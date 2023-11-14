package goha

type AccountParams struct {
}
type AccountLinkParams struct {
	Account    string `json:"account,omitempty"`
	RefreshURL string `json:"refreshURL,omitempty"`
	ReturnURL  string `json:"returnURL,omitempty"`
}
type GetOrderItemOptions struct {
	WithDetails bool `json:"withDetails,omitempty"`
}
type GetPaymentsOptions struct {
	From              string
	To                string
	UserSearchKey     string
	PageIndex         int
	PageSize          int
	ContinuationToken string
	States            []string
	SortOrder         string
	SortField         string
}
type CheckoutIntentParams struct {
	TotalAmount       int64                      `json:"totalAmount,omitempty"`
	InitialAmount     int64                      `json:"initialAmount,omitempty"`
	ItemName          string                     `json:"itemName,omitempty"`
	ContainsDonation  bool                       `json:"containsDonation,omitempty"`
	Metadata          Metadata                   `json:"metadata,omitempty"`
	ReturnUrl         string                     `json:"returnUrl,omitempty"`
	BackUrl           string                     `json:"backUrl,omitempty"`
	ErrorUrl          string                     `json:"errorUrl,omitempty"`
	TrackingParameter string                     `json:"trackingParameter,omitempty"`
	Terms             *CheckoutIntentParamsTerms `json:"terms,omitempty"`
	Payer             *CheckoutIntentParamsPayer `json:"payer,omitempty"`
}
type CheckoutIntentParamsTerms struct {
	Amount int64  `json:"amount,omitempty"`
	Date   string `json:"date,omitempty"`
}
type CheckoutIntentParamsPayer struct {
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Email       string `json:"email,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`
	Country     string `json:"country,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
}
type Account struct {
	Id string `json:"id,omitempty"`
}
type AccountLink struct {
	RedirectUrl string `json:"redirectUrl,omitempty"`
}
type CheckoutIntentOrderPayer struct {
	Email       string `json:"email"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	Country     string `json:"country"`
	Company     string `json:"company"`
	DateOfBirth string `json:"dateOfBirth"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
}
type CheckoutIntentOrderItemPayment struct {
	Id          int `json:"id"`
	ShareAmount int `json:"shareAmount"`
}
type CheckoutIntentOrderItemUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type CheckoutIntentOrderItemDiscount struct {
	Code   string `json:"code"`
	Amount int    `json:"amount"`
}
type CheckoutIntentOrderItemOptionCustomField struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Answer string `json:"answer"`
}
type CheckoutIntentOrderItemOption struct {
	Name          string                                     `json:"name"`
	Amount        int                                        `json:"amount"`
	PriceCategory string                                     `json:"priceCategory"`
	IsRequired    bool                                       `json:"isRequired"`
	CustomFields  []CheckoutIntentOrderItemOptionCustomField `json:"customFields"`
}
type CheckoutIntentOrderItem struct {
	Payments        []CheckoutIntentOrderItemPayment `json:"payments"`
	Name            string                           `json:"name,omitempty"`
	User            CheckoutIntentOrderItemUser      `json:"user"`
	PriceCategory   string                           `json:"priceCategory"`
	Discount        CheckoutIntentOrderItemDiscount  `json:"discount,omitempty"`
	Options         []CheckoutIntentOrderItemOption  `json:"options,omitempty"`
	TicketUrl       string                           `json:"ticketUrl,omitempty"`
	QrCode          string                           `json:"qrCode"`
	TierDescription string                           `json:"tierDescription"`
	Id              int                              `json:"id"`
	Amount          int                              `json:"amount"`
	Type            string                           `json:"type"`
	InitialAmount   int                              `json:"initialAmount"`
	State           string                           `json:"state"`
}
type CheckoutIntentOrderPaymentItem struct {
	Id                 int `json:"id"`
	ShareAmount        int `json:"shareAmount"`
	ShareItemAmount    int `json:"shareItemAmount"`
	ShareOptionsAmount int `json:"shareOptionsAmount,omitempty"`
}
type CheckoutIntentOrderPaymentRefundOperation struct {
	Id        int    `json:"id"`
	Amount    int    `json:"amount"`
	AmountTip int    `json:"amountTip"`
	Status    string `json:"status"`
}
type CheckoutIntentOrderPayment struct {
	Items             []CheckoutIntentOrderPaymentItem            `json:"items"`
	CashOutDate       string                                      `json:"cashOutDate,omitempty"`
	CashOutState      string                                      `json:"cashOutState,omitempty"`
	PaymentReceiptUrl string                                      `json:"paymentReceiptUrl,omitempty"`
	FiscalReceiptUrl  string                                      `json:"fiscalReceiptUrl,omitempty"`
	Id                int                                         `json:"id"`
	Amount            int                                         `json:"amount"`
	Date              string                                      `json:"date"`
	PaymentMeans      string                                      `json:"paymentMeans"`
	InstallmentNumber int                                         `json:"installmentNumber"`
	State             string                                      `json:"state"`
	Meta              Metadata                                    `json:"meta"`
	RefundOperations  []CheckoutIntentOrderPaymentRefundOperation `json:"refundOperations,omitempty"`
}
type CheckoutIntentOrderAmount struct {
	Total    int `json:"total"`
	Vat      int `json:"vat"`
	Discount int `json:"discount"`
}
type Metadata map[string]string
type CheckoutIntentOrder struct {
	Payer            CheckoutIntentOrderPayer     `json:"payer"`
	Items            []CheckoutIntentOrderItem    `json:"items"`
	Payments         []CheckoutIntentOrderPayment `json:"payments"`
	Amount           CheckoutIntentOrderAmount    `json:"amount"`
	Id               int                          `json:"id"`
	Date             string                       `json:"date"`
	FormSlug         string                       `json:"formSlug"`
	FormType         string                       `json:"formType"`
	OrganizationName string                       `json:"organizationName"`
	OrganizationSlug string                       `json:"organizationSlug"`
	Meta             Metadata                     `json:"meta"`
	IsAnonymous      bool                         `json:"isAnonymous"`
	IsAmountHidden   bool                         `json:"isAmountHidden"`
}
type CheckoutIntent struct {
	Metadata    Metadata            `json:"metadata"`
	Order       CheckoutIntentOrder `json:"order"`
	Id          int                 `json:"id"`
	RedirectUrl string              `json:"redirectUrl"`
}
type OrderItemOrder struct {
	Id               int      `json:"id"`
	Date             string   `json:"date"`
	FormSlug         string   `json:"formSlug"`
	FormType         string   `json:"formType"`
	OrganizationName string   `json:"organizationName"`
	OrganizationSlug string   `json:"organizationSlug"`
	FormName         string   `json:"formName"`
	Meta             Metadata `json:"meta"`
	IsAnonymous      bool     `json:"isAnonymous"`
	IsAmountHidden   bool     `json:"isAmountHidden"`
}
type OrderItemPayer struct {
	Email       string `json:"email"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	Country     string `json:"country"`
	Company     string `json:"company"`
	DateOfBirth string `json:"dateOfBirth"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
}
type OrderItemRefundOperation struct {
	Id        int    `json:"id"`
	Amount    int    `json:"amount"`
	AmountTip int    `json:"amountTip"`
	Status    string `json:"status"`
}
type OrderItemPayment struct {
	CashOutState      string                     `json:"cashOutState,omitempty"`
	ShareAmount       int                        `json:"shareAmount"`
	Id                int                        `json:"id"`
	Amount            int                        `json:"amount"`
	Date              string                     `json:"date"`
	PaymentMeans      string                     `json:"paymentMeans"`
	InstallmentNumber int                        `json:"installmentNumber"`
	State             string                     `json:"state"`
	Meta              Metadata                   `json:"meta"`
	RefundOperations  []OrderItemRefundOperation `json:"refundOperations,omitempty"`
}
type OrderItemUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type OrderItemDiscount struct {
	Code   string `json:"code"`
	Amount int    `json:"amount"`
}
type OrderItemCustomField struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Answer string `json:"answer"`
}
type OrderItemOptionCustomField struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Answer string `json:"answer"`
}
type OrderItemOption struct {
	Name          string                       `json:"name"`
	Amount        int                          `json:"amount"`
	PriceCategory string                       `json:"priceCategory"`
	IsRequired    bool                         `json:"isRequired"`
	CustomFields  []OrderItemOptionCustomField `json:"customFields,omitempty"`
}
type Order struct {
	IsAuthenticated              bool   `json:"isAuthenticated"`
	Banner                       string `json:"banner"`
	FiscalReceiptEligibility     bool   `json:"fiscalReceiptEligibility"`
	FiscalReceiptIssuanceEnabled bool   `json:"fiscalReceiptIssuanceEnabled"`
	Type                         string `json:"type"`
	RnaNumber                    string `json:"rnaNumber"`
	Logo                         string `json:"logo"`
	Name                         string `json:"name"`
	Role                         string `json:"role"`
	City                         string `json:"city"`
	ZipCode                      string `json:"zipCode"`
	Description                  string `json:"description"`
	UpdateDate                   string `json:"updateDate"`
	Url                          string `json:"url"`
	OrganizationSlug             string `json:"organizationSlug"`
}
type PaymentOrder struct {
	Id               int      `json:"id"`
	Date             string   `json:"date"`
	FormSlug         string   `json:"formSlug"`
	FormType         string   `json:"formType"`
	OrganizationName string   `json:"organizationName"`
	OrganizationSlug string   `json:"organizationSlug"`
	FormName         string   `json:"formName"`
	Meta             Metadata `json:"meta"`
	IsAnonymous      bool     `json:"isAnonymous"`
	IsAmountHidden   bool     `json:"isAmountHidden"`
}
type PaymentPayer struct {
	Email       string `json:"email"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	Country     string `json:"country"`
	Company     string `json:"company"`
	DateOfBirth string `json:"dateOfBirth"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
}
type PaymentItem struct {
	ShareAmount        int    `json:"shareAmount"`
	ShareItemAmount    int    `json:"shareItemAmount"`
	ShareOptionsAmount int    `json:"shareOptionsAmount,omitempty"`
	Id                 int    `json:"id"`
	Amount             int    `json:"amount"`
	Type               string `json:"type"`
	State              string `json:"state"`
}
type Payment struct {
	Order             PaymentOrder  `json:"order"`
	Payer             PaymentPayer  `json:"payer"`
	Items             []PaymentItem `json:"items"`
	PaymentReceiptUrl string        `json:"paymentReceiptUrl"`
	Id                int           `json:"id"`
	Amount            int           `json:"amount"`
	Date              string        `json:"date"`
	PaymentMeans      string        `json:"paymentMeans"`
	InstallmentNumber int           `json:"installmentNumber"`
	State             string        `json:"state"`
	Meta              Metadata      `json:"meta"`
}
type OrderItem struct {
	Order           OrderItemOrder         `json:"order"`
	Payer           OrderItemPayer         `json:"payer"`
	Payments        []OrderItemPayment     `json:"payments"`
	Name            string                 `json:"name"`
	User            OrderItemUser          `json:"user"`
	PriceCategory   string                 `json:"priceCategory"`
	Discount        OrderItemDiscount      `json:"discount"`
	CustomFields    []OrderItemCustomField `json:"customFields"`
	Options         []OrderItemOption      `json:"options"`
	QrCode          string                 `json:"qrCode"`
	TierDescription string                 `json:"tierDescription"`
	Id              int                    `json:"id"`
	Amount          int                    `json:"amount"`
	Type            string                 `json:"type"`
	InitialAmount   int                    `json:"initialAmount"`
	State           string                 `json:"state"`
}
type Organization struct {
	Logo             string `json:"logo"`
	Name             string `json:"name"`
	Role             string `json:"role"`
	City             string `json:"city"`
	ZipCode          string `json:"zipCode"`
	Description      string `json:"description"`
	UpdateDate       string `json:"updateDate"`
	Url              string `json:"url"`
	OrganizationSlug string `json:"organizationSlug"`
}
type JOCategory struct {
	Id         int    `json:"id"`
	Label      string `json:"label"`
	ShortLabel string `json:"shortLabel"`
}
type PublicTag struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
type CompanyLegalStatus struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}
type OrganizationPage struct {
	Count  int            `json:"count,omitempty"`
	Cursor string         `json:"cursor,omitempty"`
	Items  []Organization `json:"items,omitempty"`
}
type PaymentPage struct {
	Count  int       `json:"count,omitempty"`
	Cursor string    `json:"cursor,omitempty"`
	Items  []Payment `json:"items,omitempty"`
}
type JOCategoryPage struct {
	Count  int          `json:"count,omitempty"`
	Cursor string       `json:"cursor,omitempty"`
	Items  []JOCategory `json:"items,omitempty"`
}
type PublicTagPage struct {
	Count  int         `json:"count,omitempty"`
	Cursor string      `json:"cursor,omitempty"`
	Items  []PublicTag `json:"items,omitempty"`
}
type CompanyLegalStatusPage struct {
	Count  int                  `json:"count,omitempty"`
	Cursor string               `json:"cursor,omitempty"`
	Items  []CompanyLegalStatus `json:"items,omitempty"`
}
