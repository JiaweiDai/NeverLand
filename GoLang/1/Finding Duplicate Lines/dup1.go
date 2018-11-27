package main 

import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	counts := make(map[string]int)
	/**
	 * A map holds a set of key/value pairs and provides constant-time operations to store, retrieve, or test for an item in the set.
	 * The key may be of any type whose values can compared with ==, strings beng the most common example; the value may be of any 
	 * type at all.
	 *
	 * The built-in function make creates a new empty map, it has other uses too.
	 *
	 * We will talk the details in 4.
	 * 
	 */
	
	/**
	 * 
	 * 
	 */
	input := bufio.NewScanner(os.Stdin)
	/**
	 * Package bufio implements buffered I/O. 
	 * It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface 
	 * but provides buffering and some help for textual I/O.
	 *
	 * package bufio helps make input and output efficient and conventient.
	 * One of its most useful features is a type called Scanner that reads input and breaks it into lines or words; it's often the
	 * easiest way to process input that comes naturally in lines.
	 * 
	 * We use a short variable declaration to create a new variable 'input' that refers to a bufio.Scanner.
	 *
	 * The scanner reads from the program's standard input. Each call to input.Scan() reads the next line and removes the newline 
	 * character from the end(normally \n); the result can be retrieved by calling input.Text(). The Scan function returns true if 
	 * there is a line and false when there is no more input.
	 */
	for input.Scan() {
		counts[input.Text()]++
		/**
		 * So we can rewrite counts[input.Text]++ to 
		 * inputLine := input.Text()
		 * counts[inputLine] = counts[inputLine] + 1
		 */
	}
	//NOTE: ignoring potential errors from input.Err()

	

	for line , n := range counts {
		if n > 1{
				/**
				 * As with for , parentheses are never used around the condition in an if statement , but braces are required for the body.
				 * There can be a optional else part that is excuted if the condition is false.
				 * 
				 */
			fmt.Printf("%d\t%s\n", n , line)
		}
	}
}