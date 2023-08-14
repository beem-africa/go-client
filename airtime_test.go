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

	resp, err := client.Transfer("255713507067", 2000, 21312)
	if err != nil {
		t.Fatal(err.Error())
	}

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}
	obtained := string(bb)
	expected := `{"errors":[{"code":103,"message":"Insufficient balance"}]}`

	if expected != obtained {
		t.Errorf("Failed, expected %s got %s", expected, obtained)
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
