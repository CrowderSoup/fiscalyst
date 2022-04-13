package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Transaction struct {
	Date        time.Time
	Number      string
	Description string
	Debit       float64
	Credit      float64
}

func getFloat(floatString string) float64 {
	debit, err := strconv.ParseFloat(floatString, 64)
	if err != nil {
		log.Print(err)
		return 0
	}
	return debit
}

func createTransactionList(data [][]string) []Transaction {
	var transactions []Transaction
	for _, row := range data {
		transactionDate, err := time.Parse("01/02/2006", row[0])
		if err != nil {
			log.Print(err)
			continue;
		}

		transactions = append(transactions, Transaction{
			Date:        transactionDate,
			Number:      row[1],
			Description: row[2],
			Debit:       getFloat(row[3]),
			Credit:      getFloat(row[4]),
		})
	}
	return transactions
}

func main() {
	fmt.Println("Hello, welcome to fiscalyst")

	f, err := os.Open("transactions.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	transactions := createTransactionList(data)
	fmt.Println(transactions)
}
