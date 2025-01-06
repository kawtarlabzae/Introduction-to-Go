package main
import("fmt"
"math/rand/v2")
func main(){
	
	randomNumber:=rand.IntN(100)
	
	var numGuess int
	fmt.Println("Hello there, please put the number that you guess is right!!")
	for numGuess=1;;numGuess++{
	var guess int
	fmt.Scan(&guess)
	if guess<randomNumber{
		fmt.Println("too low, enter another number")
		continue
	}else if guess>randomNumber{
		fmt.Println("too high, enter another number")
		continue
	}else{
		fmt.Printf("found it the number is %d with number of attempt %d", guess,numGuess)
		break
	}
}
















		

}