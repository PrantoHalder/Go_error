package main

import (
	"errors"
	"fmt"
	"log"
)

func error2 (){
	var a int
	var b int
	fmt.Println("enter 1st number ")
	fmt.Scanf("%d",&a)
	fmt.Println("enter 2nd number ")
	fmt.Scanf("%d",&b)
	r,err := addition(a,b)
    if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r)
}
func addition (a,b int) (int , error) {
	if a == 0 || b == 0{
       return 0, errors.New("the parameters cannot be zero")
	}
	return a/b,nil
}