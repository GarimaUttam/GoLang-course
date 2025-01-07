package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
func sayHello(s string){
	s = "Hello, World"
}
func sayHelloPointer(s *string){
	*s = "Hello, world"
}
func main() {
	var greeting string ="Hello go"
	sayHello(greeting)
	//function updates the copy of greeting
	fmt.Println("After calling sayHello:", greeting)

	sayHelloPointer(&greeting)
	//function updates greeting's memory location
	fmt.Println("After calling sayhello Pointer:", greeting)

	//the zero value of all the pointer of data types - string int float bool is ----> nil
	var zeroPointer *string
	fmt.Println(zeroPointer)


	var myStringPointer *string
	var myString string
	myStringPointer = &myString
	fmt.Println(myStringPointer)


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

	// taking users input
	fmt.Println("what is your name?")
	var myname string
	fmt.Scanln(&myname)

	fmt.Printf("Hello, %v\n",myname)

	// conditional if else
	isFast := false
	isSlow := false

	if isFast {
		fmt.Println("You are going too fast slow down")
	}else if isSlow{
		fmt.Println("Thanku for not speeding")	
	}else{
		fmt.Println("Happy journey")
	}

	if !isFast && !isSlow {
		fmt.Println("You are good driver")
	}
	
	//loops
	for i := 0; i < 5; i++{
		fmt.Println("garima uttam", i)
	}

	//data structures
	var arr [3]string
	fmt.Println(arr) // in go nop need of for loop to print the array

	arr[0] ="Garima"
	arr[1] ="uttam"
	arr[2] ="patel"
	fmt.Println(arr)

	arr1 := [3]string{"Garima", "Uttam","Patel"}
	fmt.Println(arr1)

	var twoD [2][3]string
	fmt.Println(twoD)


	// for thge unknown size of the arrays called slices
	var arr2 []string
	fmt.Println(len(arr2))

	arr3 := []string{"Garima", "uttam"}
	arr3 = append(arr3, "Patel")
	fmt.Println(arr3)

	arr4 := make([]string, 3)
	fmt.Println(len(arr4))
	fmt.Println("the slice is nil:", arr4 == nil)
	arr4 = append(arr4, "garima", "gaurav", "uttam", "mammy", "patel")
	fmt.Println(arr4)

	// cars := make(map[string]int)
	// cars["sedan"] = 23
	// cars["SUV"] = 12
	// cars["convertible"] = 10

	// alternative way
	cars := map[string]int{
		"sedan" : 32,
		"SUV" : 34,
		"convertible" : 93,
	}

	fmt.Println("Total number of cars :", cars)
	// the below expression also returns the bool weather a key is present or not
	numConvertibles, isConvertibleFound := cars["convertible"]
	fmt.Println("Convertible found:", isConvertibleFound)
	if isConvertibleFound{
		fmt.Printf("i have %v convertibles\n", numConvertibles)
	}

	// another way to check 
	if numberOfConvertibles , ok := cars["convertible"]; ok{
		fmt.Printf("I have %v convertibles\n", numberOfConvertibles)
	}

	// range returns two parameters index and item
	for i, person := range arr {
		fmt.Printf("Index: %v, Name: %v\n", i, person)
	}
	for _, person := range arr {
		fmt.Println(person)
	}

	// iterating in a map here also range resturns 2 params but key and value
	for key, value := range cars {
		fmt.Printf("%v -> %v\n", key,value)
	}
	// tp print only the keys
	for typeofCars := range cars{
		fmt.Println(typeofCars)
	}
	// to print only values
	for _, values := range cars{
		fmt.Println(values)
	}

}