package main

import "fmt"

const LoginToken string = "1234567890"

// public variable in Go is written in PascalCase and private variable is written in camelCase

func main() {
	var username string = "Mamoon"
	fmt.Println(username)
	fmt.Printf("The type of username is: %T\n", username)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The type of smallVal is: %T\n", smallVal)

	var smallfloat float32 = 3.14267896541
	fmt.Println(smallfloat)
	fmt.Printf("The type of smallfloat is: %T\n", smallfloat)

	var smallfloat1 float64 = 3.14267896541
	fmt.Println(smallfloat1)
	fmt.Printf("The type of smallfloat1 is: %T\n", smallfloat1)

	var str string
	fmt.Println(str)
	fmt.Printf("The type of str is: %T\n", str)

	Name := "Mamoon"
	fmt.Println(Name)
	fmt.Printf("The type of Name is: %T\n", Name)

	fmt.Println(LoginToken)
	fmt.Printf("The type of LoginToken is: %T\n", LoginToken)

}
