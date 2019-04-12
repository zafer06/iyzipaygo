package iyzipaygo

import (
	"encoding/json"
	"strconv"
)

func ApiTest(options Options) string {
	reqURL := options.BaseUrl + "/payment/test"
	pkiString := ""

	return connect("GET", reqURL, options, "", pkiString)
}

func BinNumber(request string, options Options) string {
	var req map[string]interface{}
	json.Unmarshal([]byte(request), &req)

	var pki PkiBuilder
	pki.append("locale", req["locale"].(string))
	pki.append("conversationId", req["conversationId"].(string))
	pki.append("binNumber", req["binNumber"].(string))

	reqURL := options.BaseUrl + "/payment/bin/check"
	pkiString := pki.getPkiString()

	return connect("POST", reqURL, options, request, pkiString)
}

func InstallmentInfo(request string, options Options) string {
	var req map[string]interface{}
	json.Unmarshal([]byte(request), &req)

	var pki PkiBuilder
	pki.append("locale", req["locale"].(string))
	pki.append("conversationId", req["conversationId"].(string))
	pki.append("binNumber", req["binNumber"].(string))
	pki.append("price", req["price"].(string))

	reqURL := options.BaseUrl + "/payment/iyzipos/installment"
	pkiString := pki.getPkiString()

	return connect("POST", reqURL, options, request, pkiString)
}

func CreatePayment(request string, options Options) string {
	var req map[string]interface{}
	json.Unmarshal([]byte(request), &req)

	var pki PkiBuilder
	pki.append("locale", req["locale"].(string))
	pki.append("conversationId", req["conversationId"].(string))
	pki.appendPrice("price", req["price"].(string))
	pki.appendPrice("paidPrice", req["paidPrice"].(string))
	pki.append("installment", req["installment"].(string))
	pki.append("paymentChannel", req["paymentChannel"].(string))
	pki.append("basketId", req["basketId"].(string))
	pki.append("paymentGroup", req["paymentGroup"].(string))
	pki.append("paymentCard", pkiPaymentCard(req["paymentCard"].(map[string]interface{})))
	pki.append("buyer", pkiBuyer(req["buyer"].(map[string]interface{})))
	pki.append("shippingAddress", pkiAddress(req["shippingAddress"].(map[string]interface{})))
	pki.append("billingAddress", pkiAddress(req["billingAddress"].(map[string]interface{})))
	pki.append("basketItems", pkiBasketItems(req["basketItems"].([]interface{})))
	pki.append("currency", req["currency"].(string))

	reqURL := options.BaseUrl + "/payment/auth"
	pkiString := pki.getPkiString()

	return connect("POST", reqURL, options, request, pkiString)
}

func RetrievePayment(request string, options Options) string {
	var req map[string]interface{}
	json.Unmarshal([]byte(request), &req)

	var pki PkiBuilder
	pki.append("locale", req["locale"].(string))
	pki.append("conversationId", req["conversationId"].(string))
	pki.append("paymentId", req["paymentId"].(string))
	pki.append("paymentConversationId", req["paymentConversationId"].(string))

	reqURL := options.BaseUrl + "/payment/detail"
	pkiString := pki.getPkiString()

	return connect("POST", reqURL, options, request, pkiString)
}

func ThreedsInitialize(request string, options Options) string {
	var req map[string]interface{}
	json.Unmarshal([]byte(request), &req)

	var pki PkiBuilder
	pki.append("locale", req["locale"].(string))
	pki.append("conversationId", req["conversationId"].(string))
	pki.appendPrice("price", req["price"].(string))
	pki.appendPrice("paidPrice", req["paidPrice"].(string))
	pki.append("installment", req["installment"].(string))
	pki.append("paymentChannel", req["paymentChannel"].(string))
	pki.append("basketId", req["basketId"].(string))
	pki.append("paymentGroup", req["paymentGroup"].(string))
	pki.append("paymentCard", pkiPaymentCard(req["paymentCard"].(map[string]interface{})))
	pki.append("buyer", pkiBuyer(req["buyer"].(map[string]interface{})))
	pki.append("shippingAddress", pkiAddress(req["shippingAddress"].(map[string]interface{})))
	pki.append("billingAddress", pkiAddress(req["billingAddress"].(map[string]interface{})))
	pki.append("basketItems", pkiBasketItems(req["basketItems"].([]interface{})))
	pki.append("currency", req["currency"].(string))
	pki.append("callbackUrl", req["callbackUrl"].(string))

	reqURL := options.BaseUrl + "/payment/3dsecure/initialize"
	pkiString := pki.getPkiString()

	return connect("POST", reqURL, options, request, pkiString)
}

func CreateThreedsPayment(request string, options Options) string {
	var req map[string]interface{}
	json.Unmarshal([]byte(request), &req)

	var pki PkiBuilder
	pki.append("locale", req["locale"].(string))
	pki.append("conversationId", req["conversationId"].(string))
	pki.append("paymentId", req["paymentId"].(string))
	pki.append("conversationData", req["conversationData"].(string))

	reqURL := options.BaseUrl + "/payment/3dsecure/auth"
	pkiString := pki.getPkiString()

	return connect("POST", reqURL, options, request, pkiString)
}

func CreateRefund(request string, options Options) string {
	var req map[string]interface{}
	json.Unmarshal([]byte(request), &req)

	var pki PkiBuilder
	pki.append("locale", req["locale"].(string))
	pki.append("conversationId", req["conversationId"].(string))
	pki.append("paymentTransactionId", req["paymentTransactionId"].(string))
	pki.appendPrice("price", req["price"].(string))
	pki.append("ip", req["ip"].(string))
	pki.append("currency", req["currency"].(string))

	reqURL := options.BaseUrl + "/payment/refund"
	pkiString := pki.getPkiString()

	return connect("POST", reqURL, options, request, pkiString)
}

func CreateCancel(request string, options Options) string {
	var req map[string]interface{}
	json.Unmarshal([]byte(request), &req)

	var pki PkiBuilder
	pki.append("locale", req["locale"].(string))
	pki.append("conversationId", req["conversationId"].(string))
	pki.append("paymentId", req["paymentId"].(string))
	pki.append("ip", req["ip"].(string))

	reqURL := options.BaseUrl + "/payment/cancel"
	pkiString := pki.getPkiString()

	return connect("POST", reqURL, options, request, pkiString)
}

func pkiBasketItems(jsonArray []interface{}) string {
	var basketItems string
	for _, data := range jsonArray {
		item, _ := data.(map[string]interface{})

		var pki PkiBuilder
		pki.append("id", item["id"].(string))
		pki.appendPrice("price", item["price"].(string))
		pki.append("name", item["name"].(string))
		pki.append("category1", item["category1"].(string))
		pki.append("category2", item["category2"].(string))
		pki.append("itemType", item["itemType"].(string))
		basketItems += pki.getPkiString() + ", "
	}
	return "[" + basketItems[0:len(basketItems)-2] + "]"
}

func pkiAddress(json map[string]interface{}) string {
	var pki PkiBuilder
	pki.append("address", json["address"].(string))
	pki.append("zipCode", json["zipCode"].(string))
	pki.append("contactName", json["contactName"].(string))
	pki.append("city", json["city"].(string))
	pki.append("country", json["country"].(string))
	return pki.getPkiString()
}

func pkiBuyer(json map[string]interface{}) string {
	var pki PkiBuilder
	pki.append("id", json["id"].(string))
	pki.append("name", json["name"].(string))
	pki.append("surname", json["surname"].(string))
	pki.append("identityNumber", json["identityNumber"].(string))
	pki.append("email", json["email"].(string))
	pki.append("gsmNumber", json["gsmNumber"].(string))
	pki.append("registrationDate", json["registrationDate"].(string))
	pki.append("lastLoginDate", json["lastLoginDate"].(string))
	pki.append("registrationAddress", json["registrationAddress"].(string))
	pki.append("city", json["city"].(string))
	pki.append("country", json["country"].(string))
	pki.append("zipCode", json["zipCode"].(string))
	pki.append("ip", json["ip"].(string))
	return pki.getPkiString()
}

func pkiPaymentCard(json map[string]interface{}) string {
	var pki PkiBuilder
	pki.append("cardHolderName", json["cardHolderName"].(string))
	pki.append("cardNumber", json["cardNumber"].(string))
	pki.append("expireYear", json["expireYear"].(string))
	pki.append("expireMonth", json["expireMonth"].(string))
	pki.append("cvc", json["cvc"].(string))
	pki.append("registerCard", strconv.FormatFloat(json["registerCard"].(float64), 'f', 0, 64))
	return pki.getPkiString()
}
