package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	userName := readString("What is your name?")

	userAge := readInt("How old are you?")

	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", userName, userAge))
	fmt.Printf("Your name is %s. You are %d years old.\n", userName, userAge)
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

func getInput() string {
	userInput, _ := reader.ReadString('\n')
	userInput = strings.ReplaceAll(userInput, "\r\n", "")

	return strings.ReplaceAll(userInput, "\n", "")
}
