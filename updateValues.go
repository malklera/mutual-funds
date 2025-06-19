package main

import (
	"context"
	"encoding/json"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Fund struct {
	Name  string       `json:"name"`
	URL   string       `json:"url"`
	Risk  string       `json:"risk"`
	Value []ValueEntry `json:"value"`
}

type ValueEntry struct {
	Date  string  `json:"date"`
	Price float64 `json:"price"`
}

const (
	name  = `[data-testid="titleDetailDesktop"]`
	risk  = `[data-testid="fundRiskDetailName"]`
	value = `[data-testid="currentShareValueType"]`
)

func updateValues() error {
	stat, err := os.Stat(fundsFile)
	if err != nil {
		log.Printf("Error reading the stats of %s\n", fundsFile)
		return err
	}
	fileModTime := stat.ModTime()
	now := time.Now()
	fileModDay := time.Date(fileModTime.Year(), fileModTime.Month(),
		fileModTime.Day(), 0, 0, 0, 0, fileModTime.Location())
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// NOTE: for the future i will like to work with RFC3339 format, i just do
	// not like the way i would do it now

	var add bool

	if fileModDay.Before(today) {
		add = true
	} else {
		closeMarket := time.Date(now.Year(), now.Month(), now.Day(), hourCloseMarket,
			0, 0, 0, now.Location())

		if fileModTime.Before(closeMarket) {
			add = false
		} else {
			return nil
		}
	}

	data, err := os.ReadFile(fundsFile)
	if err != nil {
		log.Printf("Error reading file '%s'\n", fundsFile)
		return err
	}

	var funds []Fund
	err = json.Unmarshal(data, &funds)
	if err != nil {
		log.Printf("Error unmarshaling file '%s'\n", fundsFile)
		return err
	}

	var newFunds []Fund

	for _, fund := range funds {
		err := getInfo(&fund, add)
		if err != nil {
			log.Printf("Error getting info of fund: %s : %v\n", fund.Name, err)
		}
	}

	updatedFunds, err := json.MarshalIndent(newFunds, "", "\t")
	if err != nil {
		log.Println("Error marshaling json from updatedFunds")
		return err
	}

	err = os.WriteFile(fundsFile, updatedFunds, 0666)
	if err != nil {
		log.Printf("Error writing file '%s'\n", fundsFile)
		return err
	}
	return nil
}

func getInfo(fund *Fund, add bool) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var resName string
	var resRisk string
	var resValue string
	date := time.Now().Format(time.DateOnly)
	err := chromedp.Run(ctx,
		chromedp.Navigate(fund.URL),
		chromedp.Text(name, &resName, chromedp.NodeVisible),
		chromedp.Text(risk, &resRisk, chromedp.NodeVisible),
		chromedp.Text(value, &resValue, chromedp.NodeVisible),
	)

	if err != nil {
		log.Printf("Error with url:\n%s\n", fund.URL)
		return err
	}

	if resName != fund.Name {
		log.Printf("Error with url:\n%s\n", fund.URL)
		log.Printf("Name of fund has changed form '%s' to '%s'\n", fund.Name, resName)
	}

	if resRisk != fund.Risk {
		log.Printf("Risk of fund change from '%s' to '%s'\n", fund.Risk, resRisk)
		fund.Risk = resRisk
	}

	resValue = strings.TrimPrefix(resValue, "$ ")
	resValue = strings.ReplaceAll(resValue, ".", "")
	resValue = strings.ReplaceAll(resValue, ",", ".")
	resValueFloat, err := strconv.ParseFloat(resValue, 64)
	if err != nil {
		log.Printf("Error trying to convert %s to float\n", strings.TrimPrefix(resValue, "$ "))
		return err
	}

	// This determines if i add a new entry or update the last one
	if add {
		fund.Value = append(fund.Value, ValueEntry{Date: date, Price: resValueFloat})
	} else {
		fund.Value[len(fund.Value)-1] = ValueEntry{Date: date, Price: resValueFloat}
	}

	return nil
}

func validURL(nameFund string, url string) (bool, string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var resName string

	err := chromedp.Run(ctx, chromedp.Navigate(url),
		chromedp.Text(name, &resName, chromedp.NodeVisible))

	if err != nil {
		log.Printf("Error with url:\n%s\n", url)
		return false, "", err
	}

	if nameFund != resName {
		return false, "", nil
	} else {
		return true, resName, nil
	}
}
