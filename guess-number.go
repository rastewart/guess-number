package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var Target int64 //the number the user must guess
	maxnumber := 100 //biggest number
	guesses := 5     //number of guesses
	var guess int64  //user's guess
	var err error    //error
	var input string //input from user
	fmt.Println("Let's play a guessing game!")
	for {
		rand.Seed(time.Now().Unix())             //seed the random number generator
		Target = int64(rand.Intn(maxnumber) + 1) //generate a random number for the user to guess

		fmt.Println("-----------------------------------------------------------")
		fmt.Println("Guess a number between 1 and", maxnumber /*, "<<", Target, ">>"*/)

		for i := 0; i < guesses; i++ {
			fmt.Print("Please enter your guess (", i+1, " of ", guesses, "): ")
			reader := bufio.NewReader(os.Stdin)
			input, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error occurred:", err)
				return
			}

			guess, err = strconv.ParseInt(strings.TrimSpace(input), 0, 0)
			if err != nil || guess > int64(maxnumber) || guess < 1 {
				guess = 0
			}

			if guess == 0 {
				fmt.Println("you didn't enter a valid number between 1 and", maxnumber, ".  Please try again.")
			} else if guess < Target {
				fmt.Println("Your guess is too LOW! ")
			} else if guess > Target {
				fmt.Println("Your guess is too HIGH! ")
			} else if guess == Target {
				fmt.Println("Your guessed the number! ")
				break //exit the loop when the user guesses the number
			}
			if i == guesses-1 {
				fmt.Printf("Sorry, you didn't guess the number.  It was %v.\n", Target)
			}
		}
		fmt.Print("Would you like to play again? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error occurred:", err)
			return
		}
		if strings.TrimSpace(input) == "n" {
			break
		}
	}
	fmt.Print("Thanks for playing!")
}
