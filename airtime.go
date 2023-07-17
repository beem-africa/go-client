package beemafrica

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AirtimeClient struct {
	ApiKey      string
	SecretKey   string
	baseUrl     string
	ballanceUrl string
}

func NewAirtime() *AirtimeClient {
	a, b := populate()
	return &AirtimeClient{
		ApiKey:      a,
		SecretKey:   b,
		baseUrl:     "https://apiairtime.beem.africa/v1/transfer",
		ballanceUrl: "https://apitopup.beem.africa/v1/credit-balance?app_name=AIRTIME",
	}
}

func (a *AirtimeClient) Transfer(address string, amount, reference int) (*http.Response, error) {
	body := map[string]interface{}{
		"dest_addr":    address,
		"amount":       amount,
		"reference_id": reference,
	}

	bb, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, a.baseUrl, bytes.NewBuffer(bb))
	if err != nil {
		return nil, err
	}
	authHeader := generateHeader(a.ApiKey, a.SecretKey)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AirtimeClient) GetBallance() (*http.Response, error) {
	authHeader := generateHeader(a.ApiKey, a.SecretKey)

	req, err := http.NewRequest(http.MethodGet, a.ballanceUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
