package checkoutpage

import (
	"net/url"
	"strconv"
	"strings"
)

// Structure for build payment URL
type CheckoutPage struct {
	// Payment domain with path
	baseUrl string

	// Signature handler for generate signature
	signatureHandler SignatureHandler
}

// Method for set base checkout page URL
func (p *CheckoutPage) SetBaseUrl(baseUrl string) *CheckoutPage {
	p.baseUrl = baseUrl

	return p
}

// Method build payment URL
func (p *CheckoutPage) GetUrl(payment Payment) string {
	signature := p.signatureHandler.Sign(payment.GetParams())

	queryArray := []string{}

	for key, value := range payment.GetParams() {
		preparedValue := ""

		switch value := value.(type) {
		case string:
			preparedValue = value
		case int:
			preparedValue = strconv.Itoa(value)
		case bool:
			preparedValue = strconv.FormatBool(value)
		}

		queryArray = append(queryArray, concat(concat(key, "="), url.QueryEscape(preparedValue)))
	}

	queryString := strings.Join(queryArray, "&")
	queryString = concat(queryString, concat("&signature=", url.QueryEscape(signature)))

	return concat(p.baseUrl, concat("?", queryString))
}

// Constructor for CheckoutPage structure
func NewCheckoutPage(signatureHandler SignatureHandler) *CheckoutPage {
	checkoutPage := CheckoutPage{"https://checkout.1d3.com/payment", signatureHandler}

	return &checkoutPage
}
