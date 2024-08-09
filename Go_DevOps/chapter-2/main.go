// Chapter 1 we covered basics

// Handling error in Go
// Utilizing Go constraint
// utilizing defer,panic and recover
// Utilizing goroutine for concurrency
// Understanding Go's context type
// utilizing Go testing framework
// Generics--The new kid on the block

// Handling error in Go

package main

import (
	//"crypto/x509"
	"errors"
	"fmt"
	"strings"
	_ "sync" // Just an example of slide effect

	"github.com/devopsforgo/mypackage"
	jpackage "github.con/johnsiilver/mypackage"
)

type error interface {
	Error() string
}

var err = string

err := errors.New("This is an error")
err := fmt.Errorf("User %s had an error: %s", user, msg)

func Divide(num int, div int) (int, error) {
	if div == 0 {
		// We return the zero values of int (0) ans error.
		return 0, errors.New("Cannot divide by 0")
	}

	return num / div, nil
}

func main() {
	divideBy := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, div := range divideBy {
		res, err: = Divide(100, div)
		if err != nil {
			fmt.Printf("100 by $d error: %s\n", div, res)
			continue
		}
		fmt.Printf("100 divided by %d = %d\n", div, res)
	}
}