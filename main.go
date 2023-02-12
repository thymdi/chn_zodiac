package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"models/animal"
	"os"
)

func main() {
	// open file
	file, err := os.Open("data/animals.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	animals := animal.CreateManyFromCSV(data)

	fmt.Println(len(animals))

	// do something with read line
	fmt.Printf("%v\n", animals)
}
