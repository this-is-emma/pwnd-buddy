package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

// ANSI color escape codes
const (
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorReset   = "\033[0m"
)

type Breach struct {
	Name        string   `json:"Name"`
	Title       string   `json:"Title"`
	Domain      string   `json:"Domain"`
	BreachDate  string   `json:"BreachDate"`
	Description string   `json:"Description"`
	DataClasses []string `json:"DataClasses"`
}

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
		var breaches []Breach
		err := json.NewDecoder(res.Body).Decode(&breaches)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		for _, breach := range breaches {
			fmt.Printf("%sName:%s %s\n%sTitle:%s %s\n%sDomain:%s %s\n%sBreachDate:%s %s\n%sDescription:%s %s\n%sDataClasses:%s %v\n\n",
				ColorRed, ColorReset, breach.Name,
				ColorGreen, ColorReset, breach.Title,
				ColorYellow, ColorReset, breach.Domain,
				ColorBlue, ColorReset, breach.BreachDate,
				ColorMagenta, ColorReset, breach.Description,
				ColorCyan, ColorReset, breach.DataClasses)
		}
	}
}
