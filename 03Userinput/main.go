package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to our resturant"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our Pizza:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ", input)

}
