/*
在这个示例中，我们首先定义了一个变量`num`并赋值为8。然后使用if语句检查`num`的值，根据条件输出相应的结果。

接着，我们定义了一个变量`grade`，并使用switch语句根据`grade`的值输出相应的结果。在本例中，当`grade`为"B"时，输出"Good!"。
*/
package main

import "fmt"

func main() {
	num := 8

	if num > 10 {
		fmt.Println("Number is greater than 10")
	} else if num < 5 {
		fmt.Println("Number is less than 5")
	} else {
		fmt.Println("Number is between 5 and 10")
	}

	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good!")
	case "C":
		fmt.Println("Average!")
	default:
		fmt.Println("Failed :(")
	}
}
