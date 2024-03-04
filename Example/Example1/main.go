package main

import "fmt"

func main() {
	/*
		- Write a nested if statement that checks the value of a variable called color and prints a message depending on the color.
		 For example, if color is "red", print "The color is red". If color is "blue", print "The color is blue".
		 If color is neither "red" nor "blue", print "The color is unknown".

	*/

	// var renk string
	// fmt.Scan(&renk)
	// if renk == "red" {
	// 	fmt.Println("The color is red")
	// } else if renk == "blue" {
	// 	fmt.Println("The color is red")

	// } else {
	// 	fmt.Println("gerisini sen yap")

	// }

	/*
		- Write a nested if statement that calculates the tax rate for a given income.
		For example, if income is less than 10,000 gelir, the tax rate is 10%. If income is between 10,000 and 20,000 gelir, the tax rate is 15%.
		If income is above 20,000 gelir, the tax rate is 20%.
	*/

	var gelir float64
	fmt.Scan(&gelir)

	if gelir <= 10000 {
		var calculates float64 = gelir * 10 / 100
		fmt.Println(calculates)

	} else if gelir < 20000 || gelir > 10000 {
		var calculates float64 = gelir * 10 / 100
		fmt.Println(calculates)

	} else {
		fmt.Println("gardaş % 20 vergiyi get maliyeye öde")

	}

}
