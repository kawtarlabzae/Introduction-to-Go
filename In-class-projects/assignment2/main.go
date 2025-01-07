package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)
func countOccurrences(names []Person) map[string]int {
	map1 := make(map[string]int)

	for _, val := range names {
		_, exists := map1[val.Eductaion]
		if exists {
			map1[val.Eductaion]++
		} else {
			map1[val.Eductaion] = 1
		}
	}

	return map1
}
type Person struct {
	Name  string `json:"name"`
	Age   int `json:"age"`
	Eductaion string `json:"education"`
	Salary int `json:"salary"`
}

func main() {
	f, err := os.Create("file.json")
	file, err := os.ReadFile("file-person.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var people []Person
	err = json.Unmarshal(file, &people)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}



	var average float64
	average=0
	for _, person := range people {
		average= average + float64(person.Age)
	}
	average=average/float64(len(people))
	fmt.Println("The average age is: ",average)

	YoungestPerson:= make([]Person, len(people))
	ages:= make([]int, len(people)) 
	for i, person := range people {
		ages[i]=person.Age
	}
	var smallage int
	smallage=ages[0]
	for _, age := range ages{
		if smallage>=age{
			smallage=age
		}
	}
	fmt.Println("Names of the youngest persons")
	for _, person := range people {
		if person.Age==smallage{
		YoungestPerson= append(YoungestPerson, person)
		fmt.Println(person.Name)

		}
	}




	OlderPerson:= make([]Person, len(people))
	
	var bigage int
	bigage=ages[0]
	for _, age := range ages{
		if bigage<=age{
			bigage=age
		}
	}
	fmt.Println("Names of the oldest persons")
	for _, person := range people {
		if person.Age==bigage{
			OlderPerson= append(OlderPerson, person)
			fmt.Println(person.Name)}
	}
	

	var averageSalary float64
	averageSalary=0
	for _, person := range people {
		averageSalary= averageSalary + float64(person.Salary)
	}
	averageSalary=averageSalary/float64(len(people))
	fmt.Println("Average Salary is: ",averageSalary)

	

	RichPerson:= make([]Person, len(people))
	fmt.Println("Names of the persons with biggest salary")
	var bigsalary int
	salary:= make([]int, len(people)) 
	for i, person := range people {
		salary[i]=person.Salary
	}
	bigsalary=salary[0]
	for _, salar := range salary{
		if bigsalary<=salar{
			bigsalary=salar
		}
	}
	for _, person := range people {
		if person.Salary==bigsalary{
			RichPerson= append(RichPerson, person)
			fmt.Println(person.Name)}
	}
	
	fmt.Println("Names of the persons with smallest salary")
	poorPerson:= make([]Person, len(people))

	var smallsalary int
	
	for i, person := range people {
		salary[i]=person.Salary
	}
	smallsalary=salary[0]
	for _, salar := range salary{
		if smallsalary>=salar{
			smallsalary=salar
			
		}
	}
	for _, person := range people {
		if person.Salary==smallsalary{
			poorPerson= append(poorPerson, person)
			fmt.Println(person.Name)}
	}

	occurences:=countOccurrences(people)
	fmt.Println("occurences related to education")
	for key, value := range occurences {
		fmt.Printf("Education: %s, Value: %d\n", key, value)
	}




	b, err := json.Marshal(occurences)
    if err != nil {
        log.Fatalf("Unable to marshal due to %s\n", err)
    }

	
	
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(string(b))
	if err != nil {
		fmt.Println(err)
        f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
}
