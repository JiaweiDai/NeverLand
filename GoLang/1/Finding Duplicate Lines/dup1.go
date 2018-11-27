package main 

import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	//NOTE: ignoring potential errors from input.Err()
	/**
	 * As with for , parentheses are never used around the condition in an if statement , but braces are required for the body.
	 * There can be a optional else part that is excuted if the condition is false.
	 * 
	 */
	for line , n := range counts {
		if n > 1{
			fmt.Printf("%d\t%s\n", n , line)
		}
	}
}