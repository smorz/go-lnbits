package lnbits

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	base = "https://legend.lnbits.com"
)

// NewLNbitsAPI creates a new LNbits instance.
//
// It requires an admin key and an invoice/read key, provided by lnbits.com.
func NewLNbitsAPI(adminKey, invoiceKey string) *LNbitsAPI {
	return &LNbitsAPI{
		Client:     http.DefaultClient,
		AdminKey:   adminKey,
		InvoiceKey: invoiceKey,
	}

}

// GetWalletDetails fetches the currently wallet details.
func (l *LNbitsAPI) GetWalletDetails() (wal WalletDetails, err error) {
	req, err := http.NewRequest("GET", base+"/api/v1/wallet", nil)
	if err != nil {
		return
	}
	req.Header.Set("X-Api-Key", l.AdminKey)
	resp, err := l.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &wal)
	if err != nil {
		return
	}
	return
}

// CreateInvoice will make an invoice request and returns invoice details
//
// The amount must be greater than zero. Otherwise returns an error before requesting.
func (l *LNbitsAPI) CreateInvoice(out bool, amount int64, memo, webhook string) (invoice Invoice, err error) {
	if amount <= 0 {
		return invoice, errors.New("the amount is not greater than zero")
	}
	b, err := json.Marshal(InvoiceRequest{
		Out:     out,
		Amount:  amount,
		Memo:    memo,
		Webhook: webhook,
	})
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", base+"/api/v1/payments", bytes.NewReader(b))
	if err != nil {
		return
	}
	req.Header.Set("X-Api-Key", l.InvoiceKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := l.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &invoice)
	if err != nil {
		log.Println(string(body))
		return
	}
	return
}

// CheckInvoice returns true in Paid field of output if the invoice is paid.
func (l *LNbitsAPI) CheckInvoice(paymentHash string) (paymentResult PaymentResult, err error) {
	req, err := http.NewRequest("GET", base+"/api/v1/payments/"+paymentHash, nil)
	if err != nil {
		return
	}
	req.Header.Set("X-Api-Key", l.InvoiceKey)
	resp, err := l.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &paymentResult)
	if err != nil {
		return
	}
	return
}
