package main 

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * A map is a reference to the data structure created by make. When a map is passed to a function, the function receives a copy
 * of the reference, so any changes the called function makes to the underlying data structure will be visible through the caller's
 * map reference too. In our example, the values inserted into the counts map by countLines are seen by main.(really like pointer)
 */

func countLines (f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)

	for input.Scan(){
		counts[input.Text()]++
	}
	//ignor the error from input.err().
}



func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			/**
			 * The function os.Open returns two values. The first is an open file (*os.File) that is used in subsequent reads by 
			 * the Scanner.
			 * The second result of os.Open is a value of the built-in error type. If err equals the special built-in value nil, 
			 * the file was opened successfully.The file is read, and when the end of the input is reached, Close closes the file
			 * and reases any resources.
			 * If the err is not nil, something went wrong. In that case, the error value describes the problem.
			 * Our simple-minded error handling prints a message on the standard error stream using Fprintf and the verb %v, which 
			 * displays a value of any type in a default format, and dup then carries on wiht the next file; the continue statement
			 * goes to the next iteration of the enclosing for loop.
			 *
			 * error handling is a complex part we will go into details in 5.
			 *
			 * Notice that the call to countLines precedes its decalration. Functions and other package-level entities may be decalred
			 * in any order.
			 */
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

