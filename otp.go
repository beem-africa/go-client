package beemafrica

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type OTPClient struct {
	ApiKey    string
	SecretKey string
	baseUrl   string
	verifyUrl string
}

func NewOTP() *OTPClient {
	a, b := populate()
	return &OTPClient{
		ApiKey:    a,
		SecretKey: b,
		baseUrl:   "https://apiotp.beem.africa/v1/request",
		verifyUrl: "https://apiotp.beem.africa/v1/verify",
	}
}

func (o *OTPClient) Request(number string, appId int) (*http.Response, error) {
	body := map[string]interface{}{
		"appId":  appId,
		"msisdn": number,
	}

	bb, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, o.baseUrl, bytes.NewBuffer(bb))
	if err != nil {
		return nil, err
	}

	authHeader := generateHeader(o.ApiKey, o.SecretKey)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OTPClient) Verify(pinId string, pin string) (*http.Response, error) {
	body := map[string]interface{}{
		"pinId": pinId,
		"pin":   pin,
	}

	bb, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, o.verifyUrl, bytes.NewBuffer(bb))
	if err != nil {
		return nil, err
	}

	authHeader := generateHeader(o.ApiKey, o.SecretKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
