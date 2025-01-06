package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Go is working")
	var name string = "Garima Uttam"
	println(name)
	fmt.Println("Welcome to", name)
	age:= 21
	fmt.Printf("My age is %v!\n", age)

	var totalCars int = 50
	fmt.Println("I have total", totalCars,"of different brands")

	startingPrice := 30.09
	fmt.Printf("And each car has the starting price of %v\n", startingPrice)
	fmt.Println("I want more cars")

	// zero values
	var name1 string
	fmt.Println("My name is", name1)

	var age1 int
	fmt.Println("My age is", age1)

	var carsPrice float64
	fmt.Printf("This car has starting price of %.1f!\n", carsPrice)

	var patience bool
	fmt.Println("I am rich I have this much patience", patience)

	str1 := "Garima"
	str2 := "garima"
	fmt.Println(strings.EqualFold(str1, str2))

	str3 := "Uttam"
	Ind := strings.Index(str3, "t")
	fmt.Println(Ind)
	fmt.Println(str1[Ind:4])

	fmt.Println(strings.Replace(str3,"Uttam","Patel",1))
	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Contains(str1,"ma"))

}