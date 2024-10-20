package main

import "fmt"

func main(){
	
	//variable
	var name string
	var age int8
	var marriage bool
	name = "kagerou"

	fmt.Println("My name is ",name)
	fmt.Println("Mys age is ",age)
	fmt.Println("And I'm not marriage yet ",marriage)

	//operator
	fmt.Println(24+13)
	fmt.Println(24-13)
	fmt.Println(24*13)
	fmt.Println(24/13)
	fmt.Println(24%13)

	//constanta
	const whoami = "New Programmer"
	const earthsGravity = 9.80665
	fmt.Println(whoami)
	fmt.Println(earthsGravity)

	var check int8
	fmt.Println(check)

	//inferred variable
	nama := "kagerou, this is my name"
	fmt.Println(nama)
	var myName, myAddress string
	myName = "valhalla"
	myAddress = "Indonesia"
	var myAge = 24
	fmt.Println(myName)
	fmt.Println(myAddress)
	fmt.Println(myAge)
}