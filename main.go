package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	url := "https://haveibeenpwned.com/api/v3/breachedaccount/sakatia.lise%40gmail.com?truncateResponse=false"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("hibp-api-key", "79acc9218ca14abaa678b6c4c9e99327")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	_, err := io.Copy(os.Stdout, res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
}
