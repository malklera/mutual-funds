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
	Url   string       `json:"url"`
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

func updateValues() {
	data, err := os.ReadFile(fundsFile)
	if err != nil {
		log.Fatalf("Error reading file %s: %v", fundsFile, err)
	}

	var funds []Fund
	err = json.Unmarshal(data, &funds)
	if err != nil {
		log.Fatalf("Error unmarshaling file %s: %v", fundsFile, err)
	}

	var newFunds []Fund

	for _, fund := range funds {
		newFunds = append(newFunds, getInfo(fund))
	}

	updatedFunds, err := json.MarshalIndent(newFunds, "", "\t")
	if err != nil {
		log.Fatalf("Error marshaling json from funds: %v", err)
	}

	err = os.WriteFile(fundsFile, updatedFunds, 0666)
	if err != nil {
		log.Fatalf("Error writing file %s: %v", fundsFile, err)
	}
}

func getInfo(fund Fund) Fund {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var resName string
	var resRisk string
	var resValue string
	date := time.Now().Format(time.DateOnly)
	err := chromedp.Run(ctx,
		chromedp.Navigate(fund.Url),
		chromedp.Text(name, &resName, chromedp.NodeVisible),
		chromedp.Text(risk, &resRisk, chromedp.NodeVisible),
		chromedp.Text(value, &resValue, chromedp.NodeVisible),
	)

	if err != nil {
		log.Printf("Error with url:\n%s\n", fund.Url)
		log.Fatalln(err)
	}

	if resName != fund.Name {
		log.Printf("Error with url:\n%s\n", fund.Url)
		log.Printf("Name of fund has changed: %v", err)
		// TODO: Take me modify mutual fund, so i can update the name
	}

	if resRisk != fund.Risk {
		log.Printf("Risk of fund change from '%s' to '%s'", fund.Risk, resRisk)
		fund.Risk = resRisk
	}

	resValue = strings.TrimPrefix(resValue, "$ ")
	resValue = strings.ReplaceAll(resValue, ".", "")
	resValue = strings.ReplaceAll(resValue, ",", ".")
	resValueFloat, err := strconv.ParseFloat(resValue, 64)
	if err != nil {
		log.Fatalf("Error trying to convert %s to int: %v", strings.TrimPrefix(resValue, "$ "), err)
	}

	fund.Value = append(fund.Value, ValueEntry{Date: date, Price: resValueFloat})

	return fund
}
