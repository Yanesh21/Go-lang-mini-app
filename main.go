package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Countries struct {
	Name          CountryName             `json:"name"`
	Cca2          string                  `json:"cca2"`
	Cca3          string                  `json:"cca3"`
	Region        string                  `json:"region"`
	Subregion     string                  `json:"subregion"`
	Currencies    map[string]CurrencyInfo `json:"currencies"`
	CurrencySymbo map[string]CurrencyInfo `json:"currencysymbo"`
}

type CountryName struct {
	Official string `json:"official"`
}

type CurrencyInfo struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func main() {

	GetCountriesByRegion("/Europe")
	GetCountriesByRegion("/north America")
}

func GetCountriesByRegion(region string) {
	url := "https://restcountries.com/v3.1/region"
	method := "GET"

	client := &http.Client{}
	var countries []Countries

	req, err := http.NewRequest(method, url+region, nil)

	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to request data. Error:", err)
		return
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read data. Error:", err)
		return
	}

	json.Unmarshal(responseData, &countries)

	fmt.Println("**************************************")
	fmt.Println("Region:", strings.TrimLeft(region, "/"))
	fmt.Println("**************************************")

	for i, c := range countries {
		fmt.Println("-------------------------------------------")
		fmt.Printf("Name: %s\n", c.Name.Official)
		fmt.Println("CCA2:", c.Cca2)
		fmt.Println("CCA3:", c.Cca3)
		fmt.Println("Region:", c.Region)
		fmt.Println("Subregion:", c.Subregion)

		for j, cu := range c.Currencies {
			currencyName := j + "- " + cu.Name
			currencySymbo := cu.Symbol
			fmt.Println("Currency:", currencyName)
			fmt.Println("CurrencySymbol:", currencySymbo)
		}
		i++
	}
}
