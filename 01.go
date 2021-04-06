package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Reading from file row by row")
	csvReaderRow()
}

func csvReaderRow() {
	var arrProduct = []Product{}

	recordFile, err := os.Open("product.csv")

	if err != nil {
		fmt.Println("An error encountered :: ", err)
		return
	}

	reader := csv.NewReader(recordFile)

	_, err = reader.Read()
	if err != nil {
		fmt.Println("An error encounterd :: ", err)
		return
	}

	for i:=0;;i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("An error encountered ::", err)
			return
		}

		price, err := strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			fmt.Println(err)
		}
		p := &Product{Name: record[0], Category: record[1], Price: price}

		arrProduct = append(arrProduct, *p)

	}
	
	sort.SliceStable(arrProduct, func(i, j int) bool {
		return arrProduct[i].Price < arrProduct[j].Price
	})
	
	e, err := json.Marshal(arrProduct)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(e))
}

type Product struct {
	Name string		`json:"name"`
	Price int	`json:"price"`
	Category string `json:"category"`

}
