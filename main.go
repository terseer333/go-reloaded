package main

import (
	"fmt"
	"go-reloaded2/helper"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("ERROR: invalid command")
		fmt.Println("USAGE: go run . sample.txt result.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	data, err := os.ReadFile(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	result := helper.Complier(string(data))

	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
