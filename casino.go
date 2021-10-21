package main

import (
	"fmt"
	"time"
	"math/rand"
)

var usrCoins int = 200


func diceRoll(){
	fmt.Println("")

	rand.Seed(time.Now().UnixNano())
    
    min := 1
    max := 6

	var firstNum int = rand.Intn(max - min + 1) + min
	var scdNum int = rand.Intn(max - min + 1) + min

	var cas1Num int = rand.Intn(max - min + 1) + min
	var cas2Num int = rand.Intn(max - min + 1) + min

	var usrBet int
	var usrWin int

	fmt.Print("How much would you like to bet? ")
	fmt.Scanln(&usrBet)
	
	if usrBet > usrCoins {
		fmt.Println("")
		fmt.Println("You can only bet the coins you have")
		fmt.Printf("You have %v coins \n",usrCoins)
		casino(usrCoins)
	} else {
		usrWin = usrBet * 2
	}

	fmt.Println("")
	fmt.Printf("You rolled %v and %v \n", firstNum, scdNum)
	time.Sleep(1 * time.Second)
	var usrFinal int = firstNum + scdNum

	fmt.Printf("Casino rolled %v and %v \n", cas1Num, cas2Num)
	fmt.Println("")
	time.Sleep(1 * time.Second)
	var casFinal int = cas1Num + cas2Num

	if usrFinal > casFinal {
		fmt.Printf("You won! +%v coins", usrWin)
		usrCoins += usrWin
		if firstNum == scdNum {
			fmt.Println("You rolled double!")
			fmt.Println("+10 coins")
			usrCoins += 10
		}
	} else if usrFinal < casFinal {
		fmt.Printf("You lost! -%v coins", usrBet)
		usrCoins -= usrBet
	} else {
		fmt.Printf("It's a draw! You get back your bet of %v", usrBet)
	}

	casino(usrCoins)
}

func rps(){
	fmt.Println("")

	reasons := make([]int, 0)
	reasons = append(reasons,
		1,
		2,
		3,
	)
	
	rand.Seed(time.Now().Unix()) 

	var usrPick int
	var usrPrint string
	var compPrint string
	var compPick int = reasons[rand.Intn(len(reasons))]

	var usrBet int
	var usrWin int

	fmt.Print("How much would you like to bet? ")
	fmt.Scanln(&usrBet)
	
	if usrBet > usrCoins {
		fmt.Println("")
		fmt.Println("You can only bet the coins you have")
		fmt.Printf("You have %v coins \n",usrCoins)
		casino(usrCoins)
	} else {
		usrWin = usrBet * 2
	}

	fmt.Println("")
	fmt.Println("1. Rock\n2. Paper\n3. Scissors\n")
	fmt.Print("Enter option: ")
	fmt.Scanln(&usrPick)

	switch(usrPick) {
	case 1:
		usrPrint = "rock"
	case 2:
		usrPrint = "paper"
	case 3:
		usrPrint = "scissors"
	default:
		fmt.Println("Invalid option")
		casino(usrCoins)
	}

	switch(compPick) {
	case 1:
		compPrint = "rock"
	case 2:
		compPrint = "paper"
	case 3:
		compPrint = "scissors"
	}

	fmt.Printf("You chose %v \n", usrPrint)
	time.Sleep(1 * time.Second)
	fmt.Printf("Casino chose %v \n", compPrint)
	time.Sleep(1 * time.Second)
	fmt.Println("")

	if usrPick == 1 && compPick == 2 {
		fmt.Printf("You win! +%v coins", usrWin)
		usrCoins += usrWin
		casino(usrCoins)
	} else if usrPick == 3 && compPick == 2 {
		fmt.Printf("You win! +%v coins", usrWin)
		usrCoins += usrWin
		casino(usrCoins)
	} else if usrPick == 2 && compPick == 1 {
		fmt.Printf("You win! +%v coins", usrWin)
		usrCoins += usrWin
		casino(usrCoins)
	} else if usrPick == compPick {
		fmt.Printf("It's a draw! You get back your bet of %v coins", usrBet)
		casino(usrCoins)
	} else {
		fmt.Printf("You lose! -%v coins", usrBet)
		usrCoins -= usrBet
		casino(usrCoins)
	}
}

func numGame(usrCoins int, streak int) {
    var usrGuess int
    var guess int = 0
    var opt string
	var randNum int = randInt(1, 100)

	var usrBet int
	var usrWin int

	fmt.Print("How much would you like to bet? ")
	fmt.Scanln(&usrBet)
	
	if usrBet > usrCoins {
		fmt.Println("")
		fmt.Println("You can only bet the coins you have")
		fmt.Printf("You have %v coins \n",usrCoins)
		casino(usrCoins)
	} else {
		usrWin = usrBet * 2
	}


    for guess < 5 {
    	fmt.Print("Enter guess: ")
    	fmt.Scanln(&usrGuess)

	    if usrGuess == randNum {
	        streak += 1
	        var streakBonus int = streak * 10

	        fmt.Printf("You won! +%v coins \n", usrWin)
	        fmt.Printf("Streak bonus: +%v coins \n", streakBonus)
	        
			usrCoins = usrWin + streakBonus

	        fmt.Println("Want to play again? y/n")
	        fmt.Scanln(&opt)

	        if opt == "n" {
				fmt.Println("")
	        	fmt.Println("--------------------")
	        	fmt.Println("Final streak:",streak)
	        	fmt.Println("Final coins:",usrCoins)
	        	fmt.Println("Thank you for playing")
	        	fmt.Println("--------------------")
	        	guess = 2000
	        	casino(usrCoins)
	        } else {
	        	numGame(usrCoins, streak)
	        }
	    } else {
	    	if usrGuess < randNum {
	    		fmt.Println("Try again! Number is higher than",usrGuess)
	    	} else if usrGuess > randNum {
	    		fmt.Println("Try again! Number is lower than",usrGuess)
	    	}
	    }     	
    	guess += 1
    	if guess >= 5 {
			fmt.Printf("You lost! -%v coins", usrBet)
			usrCoins -= usrBet
			fmt.Println("")
			fmt.Println("")
    		fmt.Println("--------------------")
    		fmt.Println("Number was", randNum)
    		fmt.Println("Final streak:", streak)
	        fmt.Println("Final coins:",usrCoins)
    		fmt.Println("Thank you for playing")
    		fmt.Println("--------------------")
    		casino(usrCoins)
    	}
    }

}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
    return min + rand.Intn(max-min)
}

func main() {
 	fmt.Println("Welcome to Ciaran's Casino!")
	casino(usrCoins)
}

func casino(usrCoins int) {
	fmt.Println("")

 	var opt int

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
	    fmt.Println("The rules are simple; guess the number between 1-100, you win! Don't? you lose!")
	    time.Sleep(1 * time.Second)
	    fmt.Println("---------- Let's start shall we? ---------")
 	    numGame(usrCoins, 0)
 	} else if opt == 2 {
 		rps()
 	} else if opt == 3 {
 		diceRoll()
 	}	
}
