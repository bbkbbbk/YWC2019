package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter number: ")
	scanner.Scan()
	numStr := scanner.Text()

	number, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	card, err := cardAt(number)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println(card)
	}
}