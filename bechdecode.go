package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/btcsuite/btcutil/bech32"
)

func decode(encoded string) string {
	_, _, err := bech32.Decode(encoded)
	if err != nil {
		//fmt.Println("ERROR:", err)
		return "ERR"
	} else {
		return "OK"
	}
}

func main() {
	csvfile, err := os.Open("validators")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	csvfile_result, err := os.OpenFile("validators_corrected", os.O_CREATE|os.O_RDWR, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check := decode(record[0])
		if check == "OK" {
			w := csv.NewWriter(csvfile_result)
			if err := w.Write(record); err != nil {
				log.Fatalln("ERROR: failed to write file", err)
			}
			w.Flush()
		}

		//fmt.Printf("%v::%v\n", record[0], decode(record[0]))
	}

}
