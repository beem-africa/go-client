package beemafrica

import (
	"encoding/base64"
	"fmt"
	"os"
)

// generateHeader creates the header for the requests.
func generateHeader(a, b string) string {
	s := fmt.Sprintf("%s:%s", a, b)
	s = base64.StdEncoding.EncodeToString([]byte(s))

	return fmt.Sprintf("Basic %s", s)
}

func populate() (string, string) {
	a := os.Getenv("BEEM_API_KEY")
	b := os.Getenv("BEEM_SECRET_KEY")

	return a, b
}
