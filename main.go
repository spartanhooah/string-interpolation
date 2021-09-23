package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/eiannone/keyboard"
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
	user.OwnsADog = readChar("Do you own a dog (y/n)?")

	fmt.Printf("\nYour name is %s, and you are %d years old. Your favorite number is %.2f. Owns a dog: %t\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
		user.OwnsADog)
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

func readChar(question string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(question)
		prompt()
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if unicode.IsLetter(char) {
			char = unicode.ToLower(char)

			if char == 'y' {
				return true
			} else if char == 'n' {
				return false
			} else {
				fmt.Println("Please enter 'y' or 'n'.")
			}
		} else {
			fmt.Println("Please enter an alphabetic character.")
		}
	}
}

func getInput() string {
	userInput, _ := reader.ReadString('\n')

	return strings.TrimSpace(userInput)
}
