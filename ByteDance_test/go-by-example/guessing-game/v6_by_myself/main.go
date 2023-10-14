package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	maxNum := 100
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	//rand.Seed(time.Now().UnixNano()) 已弃用

	secretNumber := random.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	//reader := bufio.NewReader(os.Stdin)

	var input string
	for {
		_, err := fmt.Scanf("%d", &input)

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
