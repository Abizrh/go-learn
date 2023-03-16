package main

import (
	"fmt"
)

func main() {
	/*
	* DECLARE VARIABLES
	* there're few ways to delcare variables in go
	 */

	// declare using var
	var hello = "hello"

	fmt.Println(hello)

	// reassign variable
	hello = "Change"
	fmt.Println(hello)

	// declare var with type data 1

	var text_1 string
	text_1 = "wassup"
	fmt.Println(text_1)

	// declare var with type data 2
	var text_2 string = "another way"
	fmt.Println(text_2)

	// declare with :=
	text_3 := "heyy"
	text_3 = "heyy was changed"
	fmt.Println(text_3)

	/*
	* DATA TYPES
	* go has several data types, including:
	* - data types numeric (decimal & non-decimal)
	* - data types string
	* - data types boolean
	* for more details, please look at the following file: ./data-types.md
	 */

	/*
	 * non-decimal (non-floating-point) had two categories that you should know, including:
	 * uint
	 * int
	 */

	var positiveNum uint8 = 50
	var negativeNum = -1243423644 // compiler will automatic specify the data types as int32 (-1243423644 is the scope of int32 )

	fmt.Printf("positive number: %d\n", positiveNum)
	fmt.Printf("negative number: %d\n", negativeNum)

	/*
	 * decimal had two categories that you should know, including:
	 * float32
	 * float64
	 */

	var decimalNum = 2.62

	fmt.Printf("decimal number: %f\n", decimalNum)
	fmt.Printf("decimal number: %.3f\n", decimalNum)

	/*
		The %f template is used to format decimal numeric data to a string. The decimal digit to be generated is 6 digits long. In the example above, the format result for the decimalNumber variable is 2.620000. The number of digits that appears can be controlled using %.nf, just replace n with the desired number. Example: %.3f will produce 3 decimal digits, %.10f will produce 10 decimal digits.
	*/

	// * data types boolean
	var isExist bool = true
	fmt.Printf("is exist? : %t \n", isExist)

	// * nil value & zero value
	//? 'nil' is not a data types, however it is a value. Variable with value 'nil' that mean an empty value

	//? All the data types discussed above have a zero value (default data type). This means that even if a variable is declared with no initial value, it will still have a default value.

	// * Convert the data using 'Casting Technique'
	// you just need to add destination type according to the data types that you declare before, e.g:

	var floatNum float64 = float64(50)
	var intNum int32 = int32(24.00)		
	var str = "this is casting technique"
	var cast string = string(str[0])

	fmt.Printf("casting float : %f\n", floatNum)
	fmt.Println("casting int : ", intNum)
	fmt.Println("casting string : ", cast)



}
