package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch in Golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("The value of diceNumber is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and you move 2 spaces")
	case 3:
		fmt.Println("Dice value is 3 and you move 3 spot")
	case 4:
		fmt.Printlm("Dice value is 4 and you move 4 spot")
	case 5:
		fmt.Println("Dice value is 5 and you move 5 spot")
	case 6:
		fmt.Println("Dice value is 6 and you move 6 spot and roll the dice again")
	}
}