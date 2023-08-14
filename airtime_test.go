package beemafrica_test

import (
	"io"
	"strings"
	"testing"

	"github.com/Jkarage/beemafrica"
)

func TestTransfer(t *testing.T) {
	client := beemafrica.NewAirtime()
	client.ApiKey = "0aca47adb7050bf3"
	client.SecretKey = "YTRiYmU4MDdjMmFkOGYwZDhhZmNkZTE0Yzg5OTU1ODA4ODFhY2UwMTcwOWU5YjBkNmU1OGIwOTdjMmEzMmE5Ng=="

	testCases := []struct {
		address   string
		amount    int
		reference int
		expected  string
	}{
		{"255713507067", 2000, 21312, `{"errors":[{"code":103,"message":"Insufficient balance"}]}`},
		{"25571350706", 2000, 21312, `{"errors":[{"code":102,"message":"Invalid phone number","args":{"msisdn":"25571350706"}},{"code":103,"message":"Insufficient balance"}]}`},
		{"255713507067", -2000, 21312, `{"message":"Amount should be greater than 500 and less that 500000"}`},
	}

	for _, v := range testCases {
		resp, err := client.Transfer(v.address, v.amount, v.reference)
		if err != nil {
			t.Fatal(err.Error())
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

func TestGetBallance(t *testing.T) {
	client := beemafrica.NewAirtime()
	client.ApiKey = "0aca47adb7050bf3"
	client.SecretKey = "YTRiYmU4MDdjMmFkOGYwZDhhZmNkZTE0Yzg5OTU1ODA4ODFhY2UwMTcwOWU5YjBkNmU1OGIwOTdjMmEzMmE5Ng=="

	resp, err := client.GetBallance()
	if err != nil {
		t.Error(err)
	}

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(bb), `"credit_bal":"0.0000"`) {
		t.Errorf(`Failed, Expected the response to include credit_bal:"0.0000"`)
	}
}
