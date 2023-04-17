package main

import (
	"fmt"
	//"os"
)

const LoginToken string = "Training" // "Training"

func main() {
	const DaysTotal int = 90
	const learning string = "Days of DevOps" // "Day of DevOps"
	var remainingDays uint = 90
	challenge := "#90DayofDevOps&DevSecOps" // "Number of Days of DevOps"

	fmt.Printf("Welcome to the %v challenge.\nThis challenge consist of %v days of training\n", challenge, DaysTotal)

	var TwitterHandle string
	var DaysCompleted uint

	// Asking for users input
	fmt.Println("Enter your Twitter handle: ")
	fmt.Scanln(&TwitterHandle)

	fmt.Println("How many days have completed the", learning)
	fmt.Scanln(&DaysCompleted)

	// Calculate the remaining days of training
	remainingDays = remainingDays - DaysCompleted

	fmt.Printf("Thank you %v for taking part and completing the %v days of training\n", TwitterHandle, DaysCompleted)
	fmt.Printf("you have %v days of training remaining for the %v challenge\n", remainingDays, challenge)
	fmt.Println("Good work come from God")

}
