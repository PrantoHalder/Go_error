package main

import (
	"fmt"
	"log"
)

func error3 (){
  var a int 
  var b int 
  fmt.Println("enter the value a ")
  fmt.Scanf("%d",&a)
  fmt.Println("enter the value of b")
  fmt.Scanf("%d",&b)
  r,err := sub(a,b)
  if err != nil {
	log.Fatalln(err)
  }
  fmt.Println(r)
}
func sub (a,b int )(int,error){
	if a ==0 || b == 0{
		return 0,fmt.Errorf("the parametrs can not be zero")
	}
	return a/b , nil

}