package main
import("fmt")
func main(){
	
		
		var sum float64
	
		fmt.Println("Enter your grades (0-100), exit by writing -1")

		i := true
		counter:=0.
		gradeCount:=0
		for i ==true { 
			var a float64
			fmt.Printf("Grade %d: ", gradeCount+1)
			fmt.Scan(&a)
			
			if a == -1. || a<0 || a>100{
				break
			}
			sum += a
			gradeCount++
			counter++
		}
	
		average := sum / counter
		fmt.Printf("Average Grade: %.2f\n", average)

		switch {
		case average >= 90:
			fmt.Println("=>A")
		case average >= 80:
			fmt.Println("=>B")
		case average >= 70:
			fmt.Println("=>C")
		case average >= 60:
			fmt.Println("=>D")
		default:
			fmt.Println("=>F")
		}

		

}