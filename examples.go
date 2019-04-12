package main

import (
	"fmt"
	"os"

	"github.com/zafer06/iyzipay-go/iyzipay"
)

func main() {
	var options iyzipay.Options
	options.ApiKey = os.Getenv("API_KEY")
	options.SecretKey = os.Getenv("SECRET_KEY")
	options.BaseUrl = "https://sandbox-api.iyzipay.com"

	apiTest(options)
	//binNumber(options);
	//installmentInfo(options)
	//createPayment(options)
	//retrievePayment(options)
	//createThreedsInitialize(options)
	//createThreedsPayment(options)
	//createRefund(options)
	//createCancel(options)
}

func apiTest(options iyzipay.Options) {
	result := iyzipay.ApiTest(options)
	fmt.Println("ApiTest: ", result)
}

func binNumber(options iyzipay.Options) {
	request := `{
        "locale": "tr",
        "conversationId": "123456789",
        "binNumber": "542119"
    }`
	result := iyzipay.BinNumber(request, options)
	fmt.Println("BinNumber: ", result)
}

func installmentInfo(options iyzipay.Options) {
	request := `{
        "locale": "tr",
        "conversationId": "123456789",
        "binNumber": "542119",
		"price": "100.0"
    }`

	result := iyzipay.InstallmentInfo(request, options)
	fmt.Println("InstallmentInfo: ", result)
}

func createPayment(options iyzipay.Options) {
	request := `{
		"locale": "tr",
		"conversationId": "123456789",
		"price": "6.0",
		"paidPrice": "6.5",
		"installment": "1",
		"paymentChannel": "WEB",
		"basketId": "B67832",
		"paymentGroup": "PRODUCT",
        "paymentCard": {
            "cardHolderName": "John Doe",
            "cardNumber": "5528790000000008",
            "expireYear": "2030",
            "expireMonth": "12",
            "cvc": "123",
            "registerCard": 0
        },
        "buyer": {
    		"id": "BY789",
    		"name": "John",
    		"surname": "Doe",
    		"identityNumber": "74300864791",
    		"email": "email@email.com",
    		"gsmNumber": "+905350000000",
    		"registrationDate": "2013-04-21 15:12:09",
    		"lastLoginDate": "2015-10-05 12:43:35",
    		"registrationAddress": "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
    		"city": "Istanbul",
    		"country": "Turkey",
    		"zipCode": "34732",
    		"ip": "85.34.78.112"
        },
        "shippingAddress": {
    		"address": "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
    		"zipCode": "34742",
    		"contactName": "Jane Doe",
    		"city": "Istanbul",
    		"country": "Turkey"
        },
        "billingAddress": {
    		"address": "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
            "zipCode": "34742",
    		"contactName": "Jane Doe",
    		"city": "Istanbul",
    		"country": "Turkey"
    	},
        "basketItems": [
    		{
    			"id": "BI101",
    			"price": "3.0",
    			"name": "Binocular",
    			"category1": "Collectibles",
    			"category2": "Accessories",
    			"itemType": "PHYSICAL"
    		},
    		{
    			"id": "BI102",
    			"price": "2.0",
    			"name": "Game code",
    			"category1": "Game",
    			"category2": "Online Game Items",
    			"itemType": "VIRTUAL"
    		},
    		{
    			"id": "BI103",
    			"price": "1.0",
    			"name": "Usb",
    			"category1": "Electronics",
    			"category2": "Usb / Cable",
    			"itemType": "PHYSICAL"
    		}
    	],
        "currency": "TRY"
    }`

	result := iyzipay.CreatePayment(request, options)
	fmt.Println("CreatePayment: ", result)
}

func retrievePayment(options iyzipay.Options) {
	request := `{
        "locale": "tr",
        "conversationId": "123456789",
        "paymentId": "10631024",
        "paymentConversationId": "123456789"
    }`

	result := iyzipay.RetrievePayment(request, options)
	fmt.Println("RetrievePayment: ", result)
}

func createThreedsInitialize(options iyzipay.Options) {
	request := `{
		"locale": "tr",
		"conversationId": "123456789",
		"price": "6.0",
		"paidPrice": "6.5",
		"installment": "1",
		"paymentChannel": "WEB",
		"basketId": "B67832",
		"paymentGroup": "PRODUCT",
        "paymentCard": {
	        "cardHolderName": "John Doe",
	        "cardNumber": "5528790000000008",
	        "expireYear": "2030",
	        "expireMonth": "12",
	        "cvc": "123",
	        "registerCard": 0
	    },
        "buyer": {
    		"id": "BY789",
    		"name": "John",
    		"surname": "Doe",
    		"identityNumber": "74300864791",
    		"email": "email@email.com",
    		"gsmNumber": "+905350000000",
    		"registrationDate": "2013-04-21 15:12:09",
    		"lastLoginDate": "2015-10-05 12:43:35",
    		"registrationAddress": "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
    		"city": "Istanbul",
    		"country": "Turkey",
    		"zipCode": "34732",
    		"ip": "85.34.78.112"
        },
        "shippingAddress": {
    		"address": "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
    		"zipCode": "34742",
    		"contactName": "Jane Doe",
    		"city": "Istanbul",
    		"country": "Turkey"
        },
        "billingAddress": {
    		"address": "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
            "zipCode": "34742",
    		"contactName": "Jane Doe",
    		"city": "Istanbul",
    		"country": "Turkey"
    	},
        "basketItems": [
    		{
    			"id": "BI101",
    			"price": "3.0",
    			"name": "Binocular",
    			"category1": "Collectibles",
    			"category2": "Accessories",
    			"itemType": "PHYSICAL"
    		},
    		{
    			"id": "BI102",
    			"price": "2.0",
    			"name": "Game code",
    			"category1": "Game",
    			"category2": "Online Game Items",
    			"itemType": "VIRTUAL"
    		},
    		{
    			"id": "BI103",
    			"price": "1.0",
    			"name": "Usb",
    			"category1": "Electronics",
    			"category2": "Usb / Cable",
    			"itemType": "PHYSICAL"
    		}
    	],
        "currency": "TRY",
        "callbackUrl": "http://95.9.88.93:91/3dpage.php"
    }`

	result := iyzipay.ThreedsInitialize(request, options)
	fmt.Println("ThreedsInitialize: ", result)
}

func createThreedsPayment(options iyzipay.Options) {
	request := `{
        "locale": "tr",
        "conversationId": "123456789",
        "paymentId": "10639182",
        "conversationData": ""
    }`

	result := iyzipay.CreateThreedsPayment(request, options)
	fmt.Println("ThreedsPayment: ", result)
}

func createRefund(options iyzipay.Options) {
	request := `{
        "locale": "tr",
        "conversationId": "123456789",
        "paymentTransactionId": "11188574",
        "price": "2.17",
        "currency": "TRY",
        "ip": "84.34.78.112"
    }`

	result := iyzipay.CreateRefund(request, options)
	fmt.Println("Refund: ", result)
}

func createCancel(options iyzipay.Options) {
	request := `{
        "locale": "tr",
        "conversationId": "123456789",
        "paymentId": "10639181",
        "ip": "84.34.78.112"
    }`

	result := iyzipay.CreateCancel(request, options)
	fmt.Println("Cancel: ", result)
}
