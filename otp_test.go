package beemafrica_test

import (
	"io"
	"testing"

	"github.com/Jkarage/beemafrica"
)

func TestRequest(t *testing.T) {
	client := beemafrica.NewOTP()

	testCases := []struct {
		number   string
		appID    int
		expected string
	}{
		// The appID for this test account is not 1308.
		{"255713507067", 1308, `{"data":{"message":{"code":106,"message":"Application not found"}}}`},
		// The phone number lacks 1 digit
		{"25571350706", 1309, `{"data":{"message":{"code":102,"message":"Invalid phone number"}}}`},
	}

	for _, v := range testCases {
		resp, err := client.Request(v.number, v.appID)
		if err != nil {
			t.Error(err)
		}

		bb, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		if string(bb) != v.expected {
			t.Errorf("failed, got %s expected %s", string(bb), v.expected)
		}
	}
}
