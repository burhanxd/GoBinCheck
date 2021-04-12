package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Bins struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    struct {
		Bin         string `json:"bin"`
		Vendor      string `json:"vendor"`
		Type        string `json:"type"`
		Level       string `json:"level"`
		Bank        string `json:"bank"`
		Country     string `json:"country"`
		Countryinfo struct {
			Code     string `json:"code"`
		} `json:"countryInfo"`
	} `json:"data"`
}

func main() {
	var bin string
	fmt.Printf("Bin: ")
	fmt.Scanf("%s",&bin)

	url := ("https://binssuapi.vercel.app/api/"+bin)

	client := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")
	req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	// log.Println(string(body))
	if err != nil {
		log.Fatal(err)
	}
	var result Bins
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Vendor:", result.Data.Vendor)
	fmt.Println("Type:", result.Data.Type)
	fmt.Println("Level:", result.Data.Level)
   	fmt.Println("Bank:", result.Data.Bank)
   	fmt.Println("Country:", result.Data.Country)
   	fmt.Println("Code:", result.Data.Countryinfo.Code)
}
