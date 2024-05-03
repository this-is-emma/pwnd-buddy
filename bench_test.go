package main

import (
	"net/http"
	"fmt"
	"testing"
)

func BenchmarkHTTPRequest(b *testing.B) {
	email := "example@example.com"
	url := fmt.Sprintf("https://haveibeenpwned.com/api/v3/breachedaccount/%s?truncateResponse=false", email)
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			b.Fatalf("Error creating request: %v", err)
		}
		req.Header.Add("hibp-api-key", "79acc9218ca14abaa678b6c4c9e99327")
		_, err = http.DefaultClient.Do(req)
		if err != nil {
			b.Fatalf("Error making request: %v", err)
		}
	}
}
