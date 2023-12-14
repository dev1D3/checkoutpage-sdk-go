package checkoutpage

// Structure for communicate with our
type Gate struct {
	// Instance for check signature
	signatureHandler SignatureHandler

	// Instance for build payment URL
	checkoutPage CheckoutPage
}

// Method for set base checkout page URL
func (g *Gate) SetBaseUrl(url string) *Gate {
	g.checkoutPage.SetBaseUrl(url)

	return g
}

// Method build payment URL
func (g *Gate) GetCheckoutPageUrl(payment Payment) string {
	return g.checkoutPage.GetUrl(payment)
}

// Method for handling callback
func (g *Gate) HandleCallback(callbackData string) (*Callback, error) {
	callback, callbackError := NewCallback(g.signatureHandler, callbackData)

	return callback, callbackError
}

// Constructor for Gate structure
func NewGate(secret string) *Gate {
	signatureHandler := NewSignatureHandler(secret)
	checkoutPage := NewCheckoutPage(*signatureHandler)
	gate := Gate{*signatureHandler, *checkoutPage}

	return &gate
}
