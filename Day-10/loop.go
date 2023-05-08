package main

import "fmt"

func main() {
	fmt.Println("Learning loops in Golang")

	days := []string{"Sunday", "Momday", "Tueday", "Wednesday", "Thurday", "Friday", "Saturday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days{
		fmt.Printf("Index is %v and value %v\n", day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("value is: ", rogueValue)
		rogueValue++

		lco:
		    fmt.Println("Jumping at LearnCodeonline.io")

	}
}