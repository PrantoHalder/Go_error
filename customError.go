package main

import (
	"fmt"
	"log"
)

type customError struct {
	message string
}

func (c customError) error() string {
	return c.message
}
func error4() {
	var a int
	var b int
	fmt.Println("enter 1st number")
	fmt.Scanf("%d", &a)
	fmt.Println("enter 2nd number")
	fmt.Scanf("%d", &b)
	r, err := substraction(a, b)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r)

}
func substraction(a, b int) (int, *customError) {
	if a == 0 || b == 0 {
		return a / b, &customError{
			message: "there is a error : parameters can not be zero",
		}
	}
	return 0, nil
}
