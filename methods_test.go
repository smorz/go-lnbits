package lnbits

import (
	"fmt"
	"log"
	"testing"
)

const (
	adminKey   = "Admin Key String"
	invoiceKey = "Invoice Key String"
)

func TestGetWalletDetails(t *testing.T) {
	lnbits := NewLNbitsAPI(adminKey, invoiceKey)
	wal, err := lnbits.GetWalletDetails()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wal)
}

func TestCreatInvoice(t *testing.T) {
	//log.SetFlags(log.LstdFlags | log.Lshortfile)
	lnbits := NewLNbitsAPI(adminKey, invoiceKey)
	inv, err := lnbits.CreateInvoice(false, 1, "for api test", "http://smrazavian.ir/l/m/ow_q")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inv)
}

func TestCheckInvoice(t *testing.T) {
	lnbits := NewLNbitsAPI(adminKey, invoiceKey)
	r, err := lnbits.CheckInvoice("Payment Hash String")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
