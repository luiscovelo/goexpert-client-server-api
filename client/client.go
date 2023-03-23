package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Quotation struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	endpoint := "http://localhost:8080/cotacao"
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var quotation Quotation
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		panic(err)
	}

	err = saveQuotation(&quotation)
	if err != nil {
		panic(err)
	}

	log.Println("quotation saved successful")
}

func saveQuotation(quotation *Quotation) error {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	text := fmt.Sprintf("Dólar: %s\n", quotation.Bid)
	_, err = file.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
