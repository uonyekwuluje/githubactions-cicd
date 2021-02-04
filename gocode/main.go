package main

import "fmt"

func main() {
   fmt.Println("Working Code Sample")
   fmt.Println("---------------------")
   addNum()
}

func addNum() {
   fmt.Println("\nAddition")
   fmt.Println("-----------")
   var num1 int = 4
   var num2 int = 5
   fmt.Println("The Sum of ",num1," + ",num2," = ", (num1 + num2))
}
