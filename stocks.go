package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

func SearchTicker(ticker string) ([]Stock, error) {
	url := TickerPath + "?ticker=" + strings.ToUpper(ticker)
	body, err := fetchApi(url)
	if err != nil {
		return  []Stock{}, err
	}
	data := SearchResult{}
	if err := json.Unmarshal(body, &data); err != nil {
		return []Stock{}, err
	}

	return data.Results, nil
}

func GetDailyValues(ticker string) (Values, error) {
	yesterday := time.Now().AddDate(0, 0, -1).UTC()
	url := DailyValuesPath + "/" + strings.ToUpper(ticker) + "/" + yesterday.Format("2006-01-02")
	body, err := fetchApi(url)
	if err != nil {
		return  Values{}, err
	}
	data := Values{}
	if err := json.Unmarshal(body, &data); err != nil {
		return Values{}, err
	}

	return data, nil
}

func fetchApi(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Bearer " + ApiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

