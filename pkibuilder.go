package iyzipaygo

import "strings"

type PkiBuilder struct {
	pkiString string
}

func (p *PkiBuilder) append(key string, value string) {
	p.pkiString = p.pkiString + key + "=" + value + ","
}

func (p *PkiBuilder) appendPrice(key string, value string) {
	var formatPrice string = ""

	if strings.Index(value, ".") == -1 {
		formatPrice = value + ".0"
	} else {
		formatPrice = strings.TrimRight(value, "0")
	}

	if strings.HasSuffix(formatPrice, ".") {
		formatPrice = formatPrice + "0"
	}
	p.pkiString = p.pkiString + key + "=" + formatPrice + ","
}

func (p *PkiBuilder) getPkiString() string {
	return "[" + p.pkiString[0:len(p.pkiString)-1] + "]"
}
