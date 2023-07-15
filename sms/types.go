package sms

type SMSClient struct {
	ApiKey      string
	SecretKey   string
	authHeader  string
	baseUrl     string
	ballanceUrl string
}
