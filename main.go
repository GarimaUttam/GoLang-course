package main

import (
	"fmt"
	"math"
	"strconv"
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

	//unsigned int 0 to 256
	var tinynum uint8 = 32
	fmt.Println(tinynum)

	// signed int -128 to 127
	var smallnegNumber int8 = -23
	fmt.Println(smallnegNumber)

	//float32 and float 64
	//math functions
	var number float64 = 2342.0232
	fmt.Println(math.Round(number))
	fmt.Println(math.Ceil(number))
	fmt.Println(math.Floor(number))

	//typecasting
	fmt.Println(342.0)
	fmt.Printf("%.3f\n",342.000)

	str := string(80)
	fmt.Println(str)
	str11 := fmt.Sprint(80)
	fmt.Println(str11)
	// type of  the data
	fmt.Printf("%T\n", str)

	str22 := strconv.Itoa(10)
	fmt.Println(str22)

	num1 := "32432"
	var stringtonumber, _ = strconv.Atoi(num1)
	fmt.Println(stringtonumber)

	var myfloat, _ = strconv.ParseFloat(num1,64)
	fmt.Println(myfloat)
	fmt.Printf("%T\n",myfloat)

	//const and iota
	const (
		_ = iota
		num11
		num2
		num3
		num4
		num5
	)

	fmt.Println(num11)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)


	

	
}