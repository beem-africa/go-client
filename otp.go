package beemafrica

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OTPClient struct {
	apiKey    string
	secretKey string
	baseUrl   string
	verifyUrl string
}

func NewOTP() *OTPClient {
	a, b := populate()
	return &OTPClient{
		apiKey:    a,
		secretKey: b,
		baseUrl:   "https://apiotp.beem.africa/v1/request",
		verifyUrl: "https://apiotp.beem.africa/v1/verify",
	}
}

// Requires Mobile number in valid international number format with country code.
// No leading + sign. Example 255713507067.
// Application ID a number representing your application. Found in OTP Dashboard
func (o *OTPClient) Request(number string, appId int) (*http.Response, error) {
	// checks for empty Apikey and secretkeys
	if o.apiKey == "" || o.secretKey == "" {
		return nil, fmt.Errorf("failed to load accounts apikey or secretkey")
	}

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

	authHeader := generateHeader(o.apiKey, o.secretKey)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OTPClient) Verify(pinId string, otp string) (*http.Response, error) {
	// checks for empty Apikey and secretkeys
	if o.apiKey == "" || o.secretKey == "" {
		return nil, fmt.Errorf("failed to load accounts apikey or secretkey")
	}

	body := map[string]interface{}{
		"pinId": pinId,
		"pin":   otp,
	}

	bb, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, o.verifyUrl, bytes.NewBuffer(bb))
	if err != nil {
		return nil, err
	}

	authHeader := generateHeader(o.apiKey, o.secretKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
