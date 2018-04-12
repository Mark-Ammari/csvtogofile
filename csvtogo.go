package main 

import (
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"log"
	"io"

)

func main() {

csvFile, _ := os.Open("example.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	file, err := os.Create("csvtogoconversion.go")
	fmt.Fprintln(file, "package csvtogo", "\n")
	fmt.Fprintln(file, "var csvToGo = []string {")

	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for {
		record, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		for i := range record {
			csvToString := "\"" + record[i] + "\"" + ","
			fmt.Fprintln(file, csvToString)
		}

	}
	fmt.Fprintln(file, "}")

}