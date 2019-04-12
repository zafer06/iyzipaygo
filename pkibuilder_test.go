package iyzipaygo

import (
	"testing"
)

func TestAppend(t *testing.T) {
	var pki PkiBuilder
	pki.append("locale", "tr")
	pki.append("conversationId", "123456789")
	pki.append("binNumber", "542119")

	pkiString := pki.getPkiString()
	testPkiString := "[locale=tr,conversationId=123456789,binNumber=542119]"

	if pkiString != testPkiString {
		t.Errorf("Expecting %s, got '%s'\n", testPkiString, pkiString)
	}
}

func TestAppendPrice(t *testing.T) {
	var pki PkiBuilder
	pki.appendPrice("price3", "100")
	pki.appendPrice("price4", "100.10")
	pki.appendPrice("price5", "100.123")
	pki.appendPrice("price6", "100.10100")
	pki.appendPrice("price7", "100.120")
	pki.appendPrice("price8", "100.00")

	pkiString := pki.getPkiString()
	testPkiString := "[price3=100.0,price4=100.1,price5=100.123,price6=100.101,price7=100.12,price8=100.0]"

	if pkiString != testPkiString {
		t.Errorf("Expecting %s, got '%s'\n", testPkiString, pkiString)
	}
}
