package main

import "fmt"
//directory path
import "CreateNewPackage/utility"

func main() {
	fmt.Println("Hello", "World")

	fmt.Println(utility.GetName())
	fmt.Println(utility.GetMyName())
}