package lnbits

import "net/http"

// LNbitsAPI allows you to interact with the LNbits API.
type LNbitsAPI struct {
	Client     *http.Client
	AdminKey   string
	InvoiceKey string
	base       string
}

// WalletDetails is a response, from GetWalletDetails.
type WalletDetails struct {
	ID      string `json:"id,omitempty" bson:"id,omitempty"`
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	Balance int64  `json:"balance,omitempty" bson:"balance,omitempty"`
}

// Invoice is a response, from CreateInvoice.
type Invoice struct {
	PaymentHash    string `json:"payment_hash,omitempty" bson:"payment_hash,omitempty"`
	PaymentRequest string `json:"payment_request,omitempty" bson:"payment_request,omitempty"`
}

// InvoiceRequest represents an invoice request parameters. ready to convert to json
type InvoiceRequest struct {
	Out     bool   `json:"out" bson:"out"`
	Amount  int64  `json:"amount" bson:"amount"`
	Memo    string `json:"memo" bson:"memo"`
	Webhook string `json:"webhook,omitempty" bson:"webhook,omitempty"`
}

// PaymentResult is a response, from CheckInvoice.
type PaymentResult struct {
	Paid bool `json:"paid" bson:"paid"`
}
