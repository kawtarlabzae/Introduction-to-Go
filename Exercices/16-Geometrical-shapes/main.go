package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	return s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC)
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func PrintShapeDetails(s interface{}) error {
	shape, ok := s.(Shape)
	_, exists := s.(Circle)
	if !ok {
		return fmt.Errorf("this is not of type Shape")
	}
	fmt.Printf("The area is %f and the perimeter is %f\n", shape.Area(), shape.Perimeter())

	if exists {
		fmt.Printf("Im a circle :The area is %f and the perimeter is %f\n", shape.Area(), shape.Perimeter())
	}
	return nil
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 7}
	t := Triangle{SideA: 3, SideB: 4, SideC: 5}
	var s Shape
	s = r
	_ = PrintShapeDetails(s)
	_ = PrintShapeDetails(c)
	_ = PrintShapeDetails(t)

}
