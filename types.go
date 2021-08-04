package lnbits

import "net/http"

type LNBitsAPI struct {
	Client     *http.Client
	AdminKey   string
	InvoiceKey string
}

type WalletDetails struct {
	ID      string `json:"id,omitempty" bson:"id,omitempty"`
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	Balance int64  `json:"balance,omitempty" bson:"balance,omitempty"`
}

type Invoice struct {
	PaymentHash    string `json:"payment_hash,omitempty" bson:"payment_hash,omitempty"`
	PaymentRequest string `json:"payment_request,omitempty" bson:"payment_request,omitempty"`
}

type InvoiceRequest struct {
	Out     bool   `json:"out" bson:"out"`
	Amount  int64  `json:"amount" bson:"amount"`
	Memo    string `json:"memo" bson:"memo"`
	Webhook string `json:"webhook,omitempty" bson:"webhook,omitempty"`
}

type PaymentResult struct {
	Paid bool `json:"paid" bson:"paid"`
}
