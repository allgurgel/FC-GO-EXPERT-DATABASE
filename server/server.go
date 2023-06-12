package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type Exchange struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type CurrencyRate struct {
	ID    string `json:"id"`
	Date  string `json:"date"`
	Price string `json:"price"`
	Code  string `json:"code"`
}

var db *sql.DB

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/cotacao", handleExchange)
	http.ListenAndServe(":8080", nil)

}

func initDB() error {
	var err error
	const create string = `
  CREATE TABLE IF NOT EXISTS currency_rate (
  id TEXT NOT NULL PRIMARY KEY,
  date DATETIME NOT NULL,
  price TEXT,
  code TEXT
  );`

	db, err = sql.Open("sqlite", "./currency_rate.db")
	if err != nil {
		return err
	}

	_, err = db.Exec(create)
	if err != nil {
		return err
	}

	return nil
}

func newCurrencyRate(price string, date string, code string) *CurrencyRate {
	return &CurrencyRate{
		ID:    uuid.New().String(),
		Price: price,
		Date:  date,
		Code:  code,
	}
}

func insertCurrencyRate(db *sql.DB, currencyRate *CurrencyRate) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()

	stmt, err := db.Prepare("insert into currency_rate (id,date,price,code) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, currencyRate.ID, currencyRate.Date, currencyRate.Price, currencyRate.Code)
	if err != nil {
		return err
	}
	return nil
}

func handleExchange(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Falha ao realizar requisição, tempo excedido")
		w.WriteHeader(http.StatusRequestTimeout)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var exchange Exchange
	err = json.Unmarshal(body, &exchange)
	if err != nil {
		panic(err)
	}

	currencyRate := newCurrencyRate(exchange.USDBRL.Bid, exchange.USDBRL.CreateDate, exchange.USDBRL.Code+exchange.USDBRL.Codein)
	err = insertCurrencyRate(db, currencyRate)
	if err != nil {
		fmt.Println("Falha ao salvar cotação no banco, tempo excedido")
		w.WriteHeader(http.StatusRequestTimeout)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchange)

}
