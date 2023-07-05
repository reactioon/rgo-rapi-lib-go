package rapi 

import (
	
	"fmt"
	"time"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"crypto/tls"
	"crypto/hmac"
	"crypto/sha256"
    "encoding/hex"

)

type RAPI struct {

	ApiUrl []byte
	
	ApiKey []byte
	ApiSecret []byte

}

func Load(aApiKey []byte, aApiSecret []byte) RAPI {

	aApiUrl := []byte("https://api.reactioon.com:1357")

	r := RAPI{aApiUrl, aApiKey, aApiSecret}
	return r

}

func (r RAPI) Request(aMethod string, aPath string, params map[string]string) (string, error) {

	// reactioon network use auto signed certificate. 
	// so, we need skip unverified keys.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// construct http client module
	spaceClient := http.Client{
		Timeout: time.Second * 30,
		Transport: tr,
	}

	// Define the parameters
    urlParams := url.Values{}
	for key, value := range params {
		urlParams.Add(key, value)
	}

	pathUrlBase := string(r.ApiUrl) + "/" + aPath
	queryParams := urlParams.Encode()
	signature := GenerateSignatureHMAC(queryParams, string(r.ApiSecret))

	fmt.Println(pathUrlBase)

	var requestBody []byte
	var requestPostData io.Reader

	requestUrl := pathUrlBase
	requestMethod := aMethod

	if aMethod == "POST" {

		requestBody = []byte(queryParams)
		requestPostData = bytes.NewBuffer(requestBody)

	}

	req, err := http.NewRequest(requestMethod, requestUrl, requestPostData)

	if err != nil {
		return "", err
	}

	// add headers
	req.Header.Add("User-Agent", "Mozilla/4.0 (compatible; Go RTN API)")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-RTN-KEY", string(r.ApiKey))
	req.Header.Add("X-RTN-SIGNATURE", signature)
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	

	// do request
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return "", getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "", readErr
	}
	
	return string(body), nil

}

func GenerateSignatureHMAC(data string, secret string) string {
	
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	sha := hex.EncodeToString(h.Sum(nil))

	return sha

}