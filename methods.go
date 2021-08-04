package lnbits

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	base = "https://lnbits.com"
)

// NewLNBitsAPI creates a new LNBits instance.
//
// It requires an admin key and an invoice/read key, provided by lnbits.com.
func NewLNBitsAPI(adminKey, invoiceKey string) *LNBitsAPI {
	return &LNBitsAPI{
		Client:     http.DefaultClient,
		AdminKey:   adminKey,
		InvoiceKey: invoiceKey,
	}

}

// GetWalletDetails fetches the currently wallet details.
func (l *LNBitsAPI) GetWalletDetails() (wal WalletDetails, err error) {
	req, err := http.NewRequest("get", base+"/api/v1/wallet", nil)
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

// CreateInvoice
func (l *LNBitsAPI) CreateInvoice(out bool, amount int64, memo, webhook string) (invoice Invoice, err error) {
	b, err := json.Marshal(InvoiceRequest{
		Out:     out,
		Amount:  amount,
		Memo:    memo,
		Webhook: webhook,
	})
	if err != nil {
		return
	}
	req, err := http.NewRequest("post", base+"/api/v1/payments", bytes.NewBuffer(b))
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
		return
	}
	return
}

// CheckInvoice
func (l *LNBitsAPI) CheckInvoice(paymentHash string) (paymentResult PaymentResult, err error) {
	req, err := http.NewRequest("get", base+"/api/v1/payments/"+paymentHash, nil)
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
