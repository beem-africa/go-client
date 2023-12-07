package beemafrica

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SMSClient struct {
	apiKey      string
	secretKey   string
	baseUrl     string
	ballanceUrl string
	senderUrl   string
}

func NewSMS() *SMSClient {
	a, b := populate()

	return &SMSClient{
		apiKey:      a,
		secretKey:   b,
		baseUrl:     "https://apisms.beem.africa/v1/send",
		ballanceUrl: "https://apisms.beem.africa/public/v1/vendors/balance",
		senderUrl:   "https://apisms.beem.africa/public/v1/sender-names",
	}
}

// func generateHeader(a, b string) string {
// 	s := fmt.Sprintf("%s:%s", a, b)
// 	s = base64.StdEncoding.EncodeToString([]byte(s))

// 	return fmt.Sprintf("Basic %s", s)
// }

// SendSMS sends a post request to beemAfrica with appropriate message body
// It accepts the message, a slice of recipients, and a scheduled time value
// whose format is  GMT+0 timezone,(yyyy-mm-dd hh:mm).
// For sending now scheduled_time is ""
func (s *SMSClient) SendSMS(message string, recipients []string, schedule_time string) (*http.Response, error) {
	var resp *http.Response

	for i, r := range recipients {
		// Define the request body
		body := map[string]interface{}{
			"source_addr":   "INFO",
			"schedule_time": schedule_time,
			"encoding":      "0",
			"message":       message,
			"recipients": []map[string]interface{}{
				{
					"recipient_id": i + 1,
					"dest_addr":    r,
				},
			},
		}

		// convert the body to json
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		// Create a new request
		req, err := http.NewRequest(http.MethodPost, s.baseUrl, bytes.NewBuffer(jsonBody))
		if err != nil {
			return nil, err
		}
		authHeader := generateHeader(s.apiKey, s.secretKey)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", authHeader)

		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// GetBallance request for the sms ballance for a particular account
// If the error is nil, the response of type *http.Response will be returned
func (s *SMSClient) GetBallance() (*http.Response, error) {

	var resp *http.Response

	// Create a new request
	req, err := http.NewRequest(http.MethodGet, s.ballanceUrl, nil)
	if err != nil {
		return resp, err
	}

	authHeader := generateHeader(s.apiKey, s.secretKey)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *SMSClient) RequestSenderID(id, idContent string) (*http.Response, error) {
	body := map[string]string{
		"senderid":       id,
		"sample_content": idContent,
	}

	bb, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, s.senderUrl, bytes.NewBuffer(bb))
	if err != nil {
		return nil, err
	}

	authHeader := generateHeader(s.apiKey, s.secretKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SMSClient) GetSenderNames() (*http.Response, error) {
	authHeader := generateHeader(s.apiKey, s.secretKey)

	req, err := http.NewRequest(http.MethodGet, s.senderUrl, nil)
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
