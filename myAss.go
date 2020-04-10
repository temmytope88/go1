package main

import "fmt"

func main() {
	fmt.Println("please enter a number")
	var num int
	fmt.Scan(&num)
	fmt.Printf("binary number = %b", num)

}
