package main

import "fmt"

func areaRectangle(width float32, height float32) (float32 , bool) {
	if width<0 || height<0{
		return 0, false
	}else if width==0 ||height==0{
		return 0, false
	}
	return width * height, true
}



func main() {
	areaRectangle:= func(width float32, height float32) (float32 , bool) {
		if width<0 || height<0{
			return 0, false
		}else if width==0 ||height==0{
			return 0, false
		}
		return width * height, true
	}
	var width float32 = 5
	var height float32 = 5
	area, validation := areaRectangle(width, height)
	if validation ==false{
		fmt.Println("The calculation was done incorrectly ")
	}else{
	fmt.Println("The area of the rectangle is: ",area)}
}