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

	// var gelir float64
	// fmt.Scan(&gelir)

	// if gelir <= 10000 {
	// 	var calculates float64 = gelir * 10 / 100
	// 	fmt.Println(calculates)

	// } else if gelir < 20000 || gelir > 10000 {
	// 	var calculates float64 = gelir * 10 / 100
	// 	fmt.Println(calculates)

	// } else {
	// 	fmt.Println("gardaş % 20 vergiyi get maliyeye öde")

	// }

	// - Write a nested if statement that determines the grade for a given score.
	//    For example, if score is greater than or equal to 90, the grade is "A".
	//    If score is between 80 and 89, the grade is "B". If score is between 70 and 79, the grade is "C".
	//    If score is between 60 and 69, the grade is "D". If score is less than 60, the grade is "F".

	var score int
	fmt.Scan(&score)

	if score >= 90 {

		fmt.Println("The grade is A")
	} else if score >= 80 || score <= 89 {
		fmt.Println("The grade is B")
	} else if score >= 70 || score <= 79 {
		fmt.Println("The grade is C")
	} else if score >= 60 || score <= 69 {
		fmt.Println("The grade is D")
	} else {
		fmt.Println("The grade is F")
	}
}
