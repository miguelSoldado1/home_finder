package imovirtual

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// the cron job runs every hour
const offset = time.Duration(-1) * time.Hour
const baseURL = "https://www.imovirtual.com/_next/data/dpR8lHfeE74mEL00QTJdF/pt/resultados/arrendar/apartamento/lisboa.json"

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message: ", e)
	}
}

func makeApiRequest(priceMax string) Response {
	url, err := url.Parse(baseURL)
	checkNilErr(err)

	q := url.Query()
	q.Add("by", "LATEST")
	q.Add("direction", "DESC")
	q.Add("searchingCriteria", "arrendar")
	q.Add("searchingCriteria", "apartamento")
	q.Add("searchingCriteria", "muitas-localizacoes")
	q.Add("searchingCriteria", "lisboa")
	q.Add("searchingCriteria", "lisboa")
	q.Add("priceMax", priceMax)
	url.RawQuery = q.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url.String(), nil)
	checkNilErr(err)

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	checkNilErr(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	checkNilErr(err)

	var response Response
	err = json.Unmarshal(body, &response)
	checkNilErr(err)

	return response
}

func Search(priceMax string) []string {
	response := makeApiRequest(priceMax)
	messages := []string{}

	for _, listing := range response.PageProps.Data.SearchAds.Items {
		now := time.Now()
		hourAgo := now.Add(offset)
		dateCreated, err := time.Parse("2006-01-02 15:04:05", listing.DateCreated)
		checkNilErr(err)

		if dateCreated.Before(hourAgo) {
			break
		}

		message := fmt.Sprintf("%s, %d %s\n", listing.Location.Address.City.Name, listing.TotalPrice.Value, listing.TotalPrice.Currency)
		message += fmt.Sprintf("https://www.imovirtual.com/pt/anuncio/%s\n", listing.Slug)
		messages = append(messages, message)
	}

	return messages
}

type Response struct {
	PageProps struct {
		Data struct {
			SearchAds struct {
				Items []struct {
					ID       int    `json:"id"`
					Title    string `json:"title"`
					Slug     string `json:"slug"`
					Location struct {
						Address struct {
							City struct {
								Name string `json:"name"`
							} `json:"city"`
						} `json:"address"`
					} `json:"location"`
					TotalPrice struct {
						Value    int    `json:"value"`
						Currency string `json:"currency"`
					} `json:"totalPrice"`
					DateCreated string `json:"dateCreated"`
				} `json:"items"`
			}
		}
	}
}
