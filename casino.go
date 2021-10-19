/*

usrCoins needs to be a global variable, const?
need to find a way of tracking usrCoins

dice roll, rock paper scissors

can bet coins

add rock paper scissors

need a loop so you can play again

add arguments, if coins are zero can't bet

double your bet - make sure they can bet that amount

*/


package main

import (
	"fmt"
	"time"
	"math/rand"
)

var usrCoins int = 0


func diceRoll(){
	fmt.Println("")

	rand.Seed(time.Now().UnixNano())
    
    min := 1
    max := 6

	var firstNum int = rand.Intn(max - min + 1) + min
	var scdNum int = rand.Intn(max - min + 1) + min

	var cas1Num int = rand.Intn(max - min + 1) + min
	var cas2Num int = rand.Intn(max - min + 1) + min

	fmt.Printf("You rolled %v and %v \n", firstNum, scdNum)
	time.Sleep(1 * time.Second)
	var usrFinal int = firstNum + scdNum

	fmt.Printf("Casino rolled %v and %v \n", cas1Num, cas2Num)
	time.Sleep(1 * time.Second)
	var casFinal int = cas1Num + cas2Num

	if usrFinal > casFinal {
		fmt.Println("You won! +10 coins")
		usrCoins += 10
		if firstNum == scdNum {
			fmt.Println("You rolled double!")
			fmt.Println("+10 coins")
			usrCoins += 10
		}
	} else if usrFinal < casFinal {
		fmt.Println("You lost!")
	} else {
		fmt.Println("It's a draw!")
	}

	main()
}

func rps(){
	fmt.Println("")

}

func numGame(usrCoins int, streak int) {
    var usrGuess int
    var guess int = 0
    var opt string
	var randNum int = randInt(1, 30)

    for guess < 5 {
    	fmt.Print("Enter guess: ")
    	fmt.Scanln(&usrGuess)

	    if usrGuess == randNum {
	        streak += 1
	        usrCoins = streak * 10

	        fmt.Printf("You won! +%v coins \n", usrCoins)
	        
	        fmt.Println("Want to play again? y/n")
	        fmt.Scanln(&opt)

	        if opt == "n" {
	        	fmt.Println("--------------------")
	        	fmt.Println("Final streak:",streak)
	        	fmt.Println("Final coins:",usrCoins)
	        	fmt.Println("Thank you for playing")
	        	fmt.Println("--------------------")
	        	guess = 2000
	        	break
	        } else {
	        	numGame(usrCoins, streak)
	        }
	    } else {
	    	if usrGuess < randNum {
	    		fmt.Println("You lost! Number is higher than",usrGuess)
	    	} else if usrGuess > randNum {
	    		fmt.Println("You lost! Number is lower than",usrGuess)
	    	}
	    }     	
    	guess += 1
    	if guess >= 5 {
    		fmt.Println("--------------------")
    		fmt.Println("Number was", randNum)
    		fmt.Println("Final streak:", streak)
	        fmt.Println("Final coins:",usrCoins)
    		fmt.Println("Thank you for playing")
    		fmt.Println("--------------------")
    		break
    	}
    }

}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
    return min + rand.Intn(max-min)
}

func main() {
	fmt.Println("")

 	var opt int
	
 	fmt.Println("Welcome to Ciaran's Casino!")
 	fmt.Println("")
 	fmt.Printf("Your current coins are %v",usrCoins)
 	fmt.Println("")
 	fmt.Println("\n 0. Exit \n 1. Number guesser \n 2. Rock paper scissors \n 3. Dice roll")
 	fmt.Print("\n What game would you like to play? ")
	fmt.Scanln(&opt)
	
 	if opt == 0 {
 	    fmt.Println("Thank you for visiting!")
 	} else if opt == 1 {
		fmt.Println("Welcome to number guesser!")
	    time.Sleep(1 * time.Second)
	    fmt.Println("The rules are simple; guess the number, you win! Don't? you lose!")
	    time.Sleep(1 * time.Second)
	    fmt.Println("---------- Let's start shall we? ---------")
 	    numGame(0, 0)
 	} else if opt == 2 {
 		rps()
 	} else if opt == 3 {
 		diceRoll()
 	}
  
}

