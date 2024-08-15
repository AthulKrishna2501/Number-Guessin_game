package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.NewSource(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	target := rand.Intn(10) + 1

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100. Can you guess it?")

	var guess int
	var attempts int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++
		
		if attempts == 3{
			fmt.Println("Ohh! You have only three attempts.\n     ----Game over!!!!----")
			break
		}
		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Printf("Congratulations! You guessed the number %d in %d attempts.\n", target, attempts)
		}
	}
	}

