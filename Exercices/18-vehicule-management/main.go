package main

import "fmt"



type Insurable interface {
	CalculateInsurance() 
}

type Printable interface {
	Details() 
}

type Car struct {
	Vehicule      Vehicule
	NumberofDoors int
}

type Truck struct {
	Vehicule        Vehicule
	PayLoadCapacity int
}

func (c Car) CalculateInsurance() {
	fmt.Println("This is a car and the insurace is: ",500*c.NumberofDoors)
}

func (t Truck) CalculateInsurance() {
	fmt.Println("This is a truck and the insurace is: ",500*t.PayLoadCapacity)
}
type Vehicule struct {
	Make  string
	Model string
	Year  int
}
func (c Car) Details() {
	fmt.Println("This is a car and \nthe model is: ",c.Vehicule.Model)
	fmt.Println("the make is: ",c.Vehicule.Make)
	fmt.Println("the year is: ",c.Vehicule.Year)
	c.CalculateInsurance()
}

func (t Truck) Details() {
	fmt.Println("This is a car and \nthe model is: ",t.Vehicule.Model)
	fmt.Println("the make is: ",t.Vehicule.Make)
	fmt.Println("the year is: ",t.Vehicule.Year)
	t.CalculateInsurance()
}
func PrintVehicleInsurance(v []Insurable) {
	for _, val := range v {
		switch vehicle := val.(type) {
		case Insurable:
			vehicle.CalculateInsurance()
		default:
			fmt.Println("This vehicule is not insurable")
		}}
}

func PrintVehicleDetails(v []Printable) {
	for _, val := range v {
		switch vehicle := val.(type) {
		case Printable:
			vehicle.Details()
		default:
			fmt.Println("This vehicule is not insurable")
		}}
}

func main() {
	myTruck := Truck{
		Vehicule: Vehicule{Make: "Toyoto", Model:"01",Year:2002},
		PayLoadCapacity:400,
	}
	myTruck2 := Truck{
		Vehicule: Vehicule{Make: "Toyoto", Model:"01",Year:2002},
		PayLoadCapacity:450,
	}
	car :=Car{
		Vehicule: Vehicule{Make: "Dacia", Model:"02",Year:2012},
		NumberofDoors:450,
	}
	fmt.Println("Insurance")
	var listInterface []Insurable
	listInterface = append(listInterface,myTruck, myTruck2,car)
	fmt.Println("Printable")
	var listInterface2 []Printable
	listInterface2 = append(listInterface2,myTruck, myTruck2,car)
	PrintVehicleInsurance(listInterface)
	PrintVehicleDetails(listInterface2)
}