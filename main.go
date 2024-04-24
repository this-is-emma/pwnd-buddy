package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Define flags for email and saveToFile
	var email string
	saveToFile := flag.Bool("s", false, "save API output to file")
	flag.StringVar(&email, "email", "", "email to check for breaches")
	flag.Parse()

	
	if email == "" {
		fmt.Println("Error: Please provide an email address using the -email flag")
		return
	}


	url := fmt.Sprintf("https://haveibeenpwned.com/api/v3/breachedaccount/%s?truncateResponse=false", email)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("hibp-api-key", "79acc9218ca14abaa678b6c4c9e99327")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	if *saveToFile {

		file, err := os.Create("search-result.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()


		_, err = io.Copy(file, res.Body)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("API output saved to search-result.txt")
	} else {

		_, err := io.Copy(os.Stdout, res.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
	}
}
