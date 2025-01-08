package main

import "fmt"

// Greeter interface defines a method for greeting
type Greeter interface {
    Greet() string
}

// Describer interface defines a method for describing
type Describer interface {
    Describe() string
}

// Animal struct acts as a base struct with common fields
type Animal struct {
    Name    string
    Species string
}

// Describe method for Animal provides a description of the animal
func (a Animal) Describe() string {
    return "This is a " + a.Species + " named " + a.Name
}

// Dog struct embeds Animal and adds a specific field for Breed
type Dog struct {
    Animal
    Breed string
}

// Greet method for Dog provides a specific greeting for dogs
func (d Dog) Greet() string {
    return "Woof! I'm " + d.Name + ", a " + d.Breed
}

// Person struct represents a human with a Name field
type Person struct {
    Name string
}

// Greet method for Person provides a specific greeting for humans
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

// Cat struct embeds Animal and adds a specific field for Color
type Cat struct {
    Animal
    Color string
}

// Greet method for Cat provides a specific greeting for cats
func (c Cat) Greet() string {
    return "Meow! I'm " + c.Name + ", a " + c.Color + " cat"
}

// ProcessGreeter function accepts any type that implements the Greeter interface
func ProcessGreeter(g Greeter) {
    fmt.Println(g.Greet())
}

// ProcessDescriber function accepts any type that implements the Describer interface
func ProcessDescriber(d Describer) {
    fmt.Println(d.Describe())
}

func main() {
    // Create a Person instance
    p := Person{Name: "Alice"}

    // Create a Dog instance with embedded Animal fields
    d := Dog{
        Animal: Animal{
            Name:    "Buddy",
            Species: "Dog",
        },
        Breed: "Golden Retriever",
    }

    // Create a Cat instance with embedded Animal fields
    c := Cat{
        Animal: Animal{
            Name:    "Whiskers",
            Species: "Cat",
        },
        Color: "white",
    }

    // Process objects through the Greeter interface
    ProcessGreeter(p)
    ProcessGreeter(d)
    ProcessGreeter(c)

    // Process objects through the Describer interface
    ProcessDescriber(d)
    ProcessDescriber(c)

    // Assigning to a variable of the Greeter interface type
    var greeter Greeter

    greeter = p
    fmt.Println("Using Greeter interface variable:", greeter.Greet())

    greeter = d
    fmt.Println("Using Greeter interface variable:", greeter.Greet())

    greeter = c
    fmt.Println("Using Greeter interface variable:", greeter.Greet())
}
