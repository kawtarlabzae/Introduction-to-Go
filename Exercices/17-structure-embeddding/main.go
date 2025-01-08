package main

import "fmt"

// Base struct for shared vehicle properties
type Vehicle struct {
	Speed        int    // Speed in km/h
	FuelEfficiency float64 // Fuel efficiency in km/l
}

// Method of the embedded Vehicle struct
func (v Vehicle) Info() string {
	return fmt.Sprintf("Speed: %d km/h, Fuel Efficiency: %.2f km/l", v.Speed, v.FuelEfficiency)
}

// Car struct embedding Vehicle
type Car struct {
	Vehicle
	PassengerCapacity int // Specific to cars
}

// Bike struct embedding Vehicle
type Bike struct {
	Vehicle
	HasCarrier bool // Specific to bikes
}

// Truck struct embedding Vehicle
type Truck struct {
	Vehicle
	LoadCapacity int // Specific to trucks
}

// Print details using embedded methods and variables
func PrintVehicleDetails(v interface{}) {
	switch vehicle := v.(type) {
	case Car:
		fmt.Printf("Car: %s, Passenger Capacity: %d\n", vehicle.Info(), vehicle.PassengerCapacity)
	case Bike:
		fmt.Printf("Bike: %s, Has Carrier: %t\n", vehicle.Info(), vehicle.HasCarrier)
	case Truck:
		fmt.Printf("Truck: %s, Load Capacity: %d kg\n", vehicle.Info(), vehicle.LoadCapacity)
	default:
		fmt.Println("Unknown vehicle type")
	}
}

func main() {
	// Create instances of different vehicles
	myCar := Car{
		Vehicle:           Vehicle{Speed: 150, FuelEfficiency: 15},
		PassengerCapacity: 5,
	}
	myBike := Bike{
		Vehicle:   Vehicle{Speed: 80, FuelEfficiency: 60},
		HasCarrier: true,
	}
	myTruck := Truck{
		Vehicle:     Vehicle{Speed: 100, FuelEfficiency: 8},
		LoadCapacity: 1500,
	}

	// Print details of each vehicle
	PrintVehicleDetails(myCar)
	PrintVehicleDetails(myBike)
	PrintVehicleDetails(myTruck)
}
