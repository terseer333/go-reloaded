package main

import (
	"log"
	"os"
)

func ProcessText(s string) string {
	input := HexToDec(s)
	input = BinToDec(input)
	input = ToUp(input)
	input = ToLow(input)
	input = ToCap(input)
	input = LastTwoUpper(input)
	return input
}

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := ProcessText(string(data))

	err = os.WriteFile("output.txt", []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
