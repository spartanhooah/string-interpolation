package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite number?")

	fmt.Printf("Your name is %s, and you are %d years old. Your favorite number is %.2f.\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(question string) string {
	for {
		fmt.Println(question)
		prompt()

		userInput := getInput()

		if userInput == "" {
			fmt.Println("Please enter a value.")
		} else {
			return userInput
		}
	}
}

func readInt(question string) int {
	for {
		fmt.Println(question)
		prompt()

		userInput := getInput()

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(question string) float64 {
	for {
		fmt.Println(question)
		prompt()

		userInput := getInput()

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func getInput() string {
	userInput, _ := reader.ReadString('\n')

	return strings.TrimSpace(userInput)
}
