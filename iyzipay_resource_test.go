package iyzipaygo

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"os"
	"testing"
)

var options Options

func init() {
	options = Options{
		os.Getenv("API_KEY"),
		os.Getenv("SECRET_KEY"),
		"https://sandbox-api.iyzipay.com",
	}
}

func TestConnect(t *testing.T) {
	reqURL := options.BaseUrl + "/payment/test"
	request := ""
	pkiString := ""

	result := connect("GET", reqURL, options, request, pkiString)

	var res map[string]interface{}
	json.Unmarshal([]byte(result), &res)

	if res["status"].(string) != "success" {
		t.Errorf("Expecting success, got '%s'\n", result)
	}
}

func TestPrepareAuthorizationString(t *testing.T) {
	pkiString := "[locale=tr,conversationId=123456789,paymentId=10639181,ip=84.34.78.112]"
	randomString := "2018-10-23T13:00:42+0000"

	var h = sha1.New()
	h.Write([]byte(options.ApiKey + randomString + options.SecretKey + pkiString))
	var sha1 []byte = h.Sum(nil)
	var authString = "IYZWS " + options.ApiKey + ":" + base64.StdEncoding.EncodeToString(sha1)

	result := prepareAuthorizationString(pkiString, randomString, options)

	if authString != result {
		t.Errorf("Expecting %s, got '%s'\n", authString, result)
	}
}
