package main

import "fmt"
type Employee struct{
	Person
	ID int
}
type Person struct {
	Name  string
	Age   int
	Email string
}
func (p Person) Getname() string{
	return p.Name
}
func (p *Person) Setname(name string) {
	p.Name=name
}

func main() {
	p := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	s := struct {
		StatusCode int
		ErrorCode string
	}{
		StatusCode:1,
		ErrorCode:"hi",
	}
	fmt.Println(p)
	fmt.Println(s)
}
