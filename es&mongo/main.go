package main

import "fmt"

func main() {
	close := InitMongoDB()
	defer close()
	err := CreateMessage(1, 2, "123")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data, err := ListMessage(1, 2, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, i := range data {
		fmt.Println(i)
	}
}
