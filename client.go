package main

import (
	"errors"
	"strconv"
)

func (hac *Client) CreateAccount(params *AccountParams) (Account, error) {
	return Account{
		Id: "",
	}, nil
}
func (hac *Client) CreateAccountLink(params *AccountLinkParams) (AccountLink, error) {
	return AccountLink{
		RedirectUrl: "",
	}, nil
}
func (hac *Client) GetCurrentOrganization() (Organization, error) {
	o := Organization{}
	page, err := hac.GetMyOrganizations()
	if err != nil {
		return o, err
	}
	if page.Count <= 0 {
		return o, errors.New("no organization for this account")
	}
	if page.Count > 1 {
		return o, errors.New("more than one organization for this account")
	}
	return page.Items[0], nil
}
func (hac *Client) GetMyOrganizations() (OrganizationPage, error) {
	o := OrganizationPage{}
	var items []Organization
	err := hac.getDocument("/v5/users/me/organizations", &items)
	if err == nil {
		o.Items = items
		o.Count = len(o.Items)
	}
	return o, err
}
func (hac *Client) GetJOCategories() (JOCategoryPage, error) {
	o := JOCategoryPage{}
	err := hac.getDocument("/v5/values/organization/categories", &o.Items)
	if err != nil {
		o.Count = len(o.Items)
	}
	return o, err
}
func (hac *Client) GetPublicTags() (PublicTagPage, error) {
	o := PublicTagPage{}
	err := hac.getDocument("/v5/values/tags", &o.Items)
	if err != nil {
		o.Count = len(o.Items)
	}
	return o, err
}
func (hac *Client) GetCompanyLegalStatuses() (CompanyLegalStatusPage, error) {
	o := CompanyLegalStatusPage{}
	err := hac.getDocument("/v5/values/company-legal-status", &o.Items)
	if err != nil {
		o.Count = len(o.Items)
	}
	return o, err
}
func (hac *Client) GetOrganizationCheckoutIntent(orgSlug string, id string) (CheckoutIntent, error) {
	o := CheckoutIntent{}
	err := hac.getDocument("/v5/organizations/"+orgSlug+"/checkout-intents/"+id, &o)
	return o, err
}
func (hac *Client) GetOrderItem(id string, opts GetOrderItemOptions) (OrderItem, error) {
	o := OrderItem{}
	query := ""
	if opts.WithDetails {
		query = "?withDetails=true"
	}
	err := hac.getDocument("/v5/items/"+id+query, &o)
	return o, err
}
func (hac *Client) GetOrder(id string) (Order, error) {
	o := Order{}
	err := hac.getDocument("/v5/orders/"+id, &o)
	return o, err
}
func (hac *Client) GetPayment(id string) (Payment, error) {
	o := Payment{}
	err := hac.getDocument("/v5/payments/"+id, &o)
	return o, err
}
func (hac *Client) GetOrganizationPayments(orgSlug string, opts GetPaymentsOptions) (PaymentPage, error) {
	o := PaymentPage{}
	query := ""
	infos := map[string]string{}
	if len(opts.From) > 0 {
		infos["from"] = opts.From
	}
	if len(opts.To) > 0 {
		infos["to"] = opts.To
	}
	if len(opts.UserSearchKey) > 0 {
		infos["userSearchKey"] = opts.UserSearchKey
	}
	if opts.PageIndex > 1 {
		infos["pageIndex"] = strconv.Itoa(opts.PageIndex)
	}
	if len(infos) > 0 {
		query = ""
		for k, v := range infos {
			if len(query) > 0 {
				query = query + "&" + k + "=" + v
			} else {
				query = "?" + k + "=" + v
			}
		}
	}
	err := hac.getDocument("/v5/organizations/"+orgSlug+"/payments"+query, &o.Items)
	if err != nil {
		o.Count = len(o.Items)
	}
	return o, err
}
func (hac *Client) CreateOrganizationCheckoutIntent(orgSlug string, params *CheckoutIntentParams) (CheckoutIntent, error) {
	ci := CheckoutIntent{}
	err := hac.createDocument("/v5/organizations/"+orgSlug+"/checkout-intents", params, &ci)
	return ci, err
}
