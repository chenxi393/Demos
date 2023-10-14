package main

import (
	"fmt"
)

type Human struct{
	name string
	age int
}

func (t *Human) Eat(){
	fmt.Println("Human is eating.")
}

func (t *Human) Sleep(){
	fmt.Println("Human is sleeping.")
}

func myinterface (temp interface{}){
	fmt.Println("This is interface test")
	_,ok:=temp.(string)
	if(ok){
		fmt.Println("This is string")
	}else{
		fmt.Println("This is not string")
	}
}
func main(){
	temp:=Human{"Bob",19}
	temp.Eat()
	temp.Sleep()
	fmt.Println(temp.age)
	myinterface(1.11)
}