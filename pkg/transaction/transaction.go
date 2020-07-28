package transactions

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"time"
)

type Transaction struct {
	XMLName string    `xml:"transactions"`
	Id      int64     `xml:"id"`
	Type    string    `xml:"type"`
	Sum     int64     `xml:"sum"`
	Status  string    `xml:"status"`
	MCC     string    `xml:"mcc"`
	Date    time.Time `xml:"date"`
}

type Transactions struct {
	XMLName      string         `xml:"transactions"`
	Transactions []*Transaction `xml:"transaction"`
}

func MakeTransactions(count int) []*Transaction {
	transactions := make([]*Transaction, count)
	for index := range transactions {
		v := &Transaction{
			`xml:"transaction"`,
			int64(index),
			"transfer",
			1000,
			"in progress",
			"4921",
			time.Date(2020, time.January, index, 11, 15, 10, 0, time.UTC),
		}
		transactions[index] = v
	}
	return transactions
}

func ExportJSON(transactions []*Transaction) ([]byte, error) {
	encodedJson, err := json.Marshal(transactions)
	if err != nil {
		log.Print(err)
		return []byte{}, err
	}

	return encodedJson, nil
}

func (t *Transactions) ExportXML() ([]byte, error) {
	encodedXML, err := xml.Marshal(t)
	if err != nil {
		log.Print(err)
		return []byte{}, err
	}
	encodedXML = append([]byte(xml.Header), encodedXML...)

	return encodedXML, nil
}
