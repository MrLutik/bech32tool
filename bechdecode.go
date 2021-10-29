package main

import (
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/btcsuite/btcutil/bech32"
)

func decode(encoded string) string {
	_, decoded, err := bech32.Decode(encoded)
	if err != nil {
		fmt.Println("Error", err)
	}

	//fmt.Println("Decoded human-readavle part:", hrp)
	//fmt.Println("Decoded Data:", hex.EncodeToString(decoded))
	return hex.EncodeToString(decoded)
}

func main() {
	csvfile, err := os.Open("validators")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	//fmt.Printf("TEST %v:\n", decode("kira1czpk0kpsvuahurf71qqqt5krr9gu7f10etyuc0"))
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nKira address: %v\nDeocoded %v\n", record[0], decode(record[0]))
	}

}
