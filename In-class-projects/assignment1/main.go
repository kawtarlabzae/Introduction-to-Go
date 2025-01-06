package main

import("fmt"
"math/rand/v2")

func main(){
	fmt.Println("Hello there, please put the limit of the trials")
	var limitTry int
	fmt.Scan(&limitTry)
	randomNumber:=rand.IntN(100)
	bestScore:=-1
	var numGuess int
	var infinity=true
	


	for infinity{
		fmt.Println("Now you can start, please put the number that you guess is right!!")
		for numGuess=1;numGuess<=limitTry;numGuess++{
		var guess int
		fmt.Scan(&guess)
		if guess<randomNumber{
			fmt.Println("too low, enter another number")
			if(numGuess==limitTry){
				fmt.Println("You lost :(")
			}
			continue
		}else if guess>randomNumber{
			fmt.Println("too high, enter another number")
			if(numGuess==limitTry){
				fmt.Println("You lost :(")
			}
			continue
		}else{
			fmt.Printf("You won!! the number is %d with number of attempt %d\n", guess,numGuess)
			if numGuess<bestScore || bestScore==-1{
				bestScore=numGuess
				fmt.Printf("Your best score now is %d \n",numGuess)
			}else{fmt.Printf("Your best score was %d \n",bestScore) }
			numGuess=1
			break
		}
		}
	fmt.Println("You can retry if you want, this is done by entering 1 if not please enter anything on the clickboard!!")
	var a int
	fmt.Scan(&a)
	if a != 1{
		fmt.Println("It was nice playing with you, goodbye!!")
		break
	}
	numGuess=1
	}
	
}
















		

