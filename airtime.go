package beemafrica

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AirtimeClient struct {
	apiKey      string
	secretKey   string
	baseUrl     string
	ballanceUrl string
}

func NewAirtime() *AirtimeClient {
	a, b := populate()
	return &AirtimeClient{
		apiKey:      a,
		secretKey:   b,
		baseUrl:     "https://apiairtime.beem.africa/v1/transfer",
		ballanceUrl: "https://apitopup.beem.africa/v1/credit-balance?app_name=AIRTIME",
	}
}

// address is the phone number in format 2557135070XX,followed by the amount.
// reference is a random number for reference
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
	authHeader := generateHeader(a.apiKey, a.secretKey)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Returns ballance in your beem account for different services.
func (a *AirtimeClient) GetBallance() (*http.Response, error) {
	authHeader := generateHeader(a.apiKey, a.secretKey)

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
