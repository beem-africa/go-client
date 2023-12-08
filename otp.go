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

// Request generates a random OTP and sends it to the provided phone number,application id.
// Requires Mobile number in valid international number format with country code.
// No leading + sign. Example 255713507067. appid is found in beem dashboard.
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

// Verify checks to see if the provided OTP matches the pinId provided.
// Returns a Valid 200 OK Response, In Both cases. Look into data for Valid or Invalid OTP
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
