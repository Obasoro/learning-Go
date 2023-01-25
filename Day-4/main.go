package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Welcome to our go learning class")
	fmt.Println("Please rate the pizza app")

	reader := bufio.NewReader(os.Stdin)

	input, _ = reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err = strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
