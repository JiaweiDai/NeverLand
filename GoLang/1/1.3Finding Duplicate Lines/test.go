package main

import "fmt"
/**
 * This program is just for test purpose.
 * We defined i and initial i with the value 1 , but when we use i in if , it won't work.
 * 
 */
func main() {
	for i := 1 ; i <10 ; i++{
		fmt.Println(i)
	}

	if i > 0 {
		fmt.Println("Hello world!")
	}
}