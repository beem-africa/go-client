package beemafrica_test

import (
	"io"
	"testing"

	"github.com/Jkarage/beemafrica"
)

func TestSendSMS(t *testing.T) {
	client := beemafrica.NewSMS()
	client.ApiKey = "0aca47adb7050bf3"
	client.SecretKey = "YTRiYmU4MDdjMmFkOGYwZDhhZmNkZTE0Yzg5OTU1ODA4ODFhY2UwMTcwOWU5YjBkNmU1OGIwOTdjMmEzMmE5Ng=="

	testCases := []struct {
		message  string
		numbers  []string
		time     string
		expected string
	}{
		// Trying to send an empty message
		{"", []string{"255713507067"}, "", `{"code":102,"message":"Insufficient balance"}`},
		// Trying to send a message
		{"Testing message #3", []string{"255713507067"}, "", `{"code":102,"message":"Insufficient balance"}`},
		// Trying to send to a message to an invalid number
		{"Testing message", []string{"25571350706"}, "", `{"code":102,"message":"Insufficient balance"}`},
	}

	for _, v := range testCases {
		resp, err := client.SendSMS(v.message, v.numbers, v.time)
		if err != nil {
			t.Error(err)
		}

		bb, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		if string(bb) != v.expected {
			t.Errorf("failed, got %s, expected %s", string(bb), v.expected)
		}
	}
}
