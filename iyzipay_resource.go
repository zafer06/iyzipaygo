package iyzipaygo

import (
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const Version = "0.3.1"

type Options struct {
	ApiKey    string
	SecretKey string
	BaseUrl   string
}

func connect(method string, url string, options Options, request string, pkiString string) string {
	var randomString string = time.Now().UTC().Format("2006-01-02T15:04:05-0700")
	var authString = prepareAuthorizationString(pkiString, randomString, options)

	payload := strings.NewReader(request)

	req, _ := http.NewRequest(method, url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", authString)
	req.Header.Add("x-iyzi-rnd", randomString)
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	/*
	   fmt.Println("-------------------------------------------------------------")
	   fmt.Println("PKI STRING -> ", pkiString)
	   fmt.Println("-------------------------------------------------------------")
	   fmt.Println("AUTHORIZATION STRING -> ", authString)
	   fmt.Println("-------------------------------------------------------------")
	*/
	return string(body)
}

func prepareAuthorizationString(pkiString string, randomString string, o Options) string {
	var h = sha1.New()
	h.Write([]byte(o.ApiKey + randomString + o.SecretKey + pkiString))
	var sha1 []byte = h.Sum(nil)
	return "IYZWS " + o.ApiKey + ":" + base64.StdEncoding.EncodeToString(sha1)
}
