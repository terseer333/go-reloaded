package main

import (
	"log"
	"os"
)

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := cap(string(data))

	err = os.WriteFile("output.txt", []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
