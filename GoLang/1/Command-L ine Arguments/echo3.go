package main

import(
	"os"
	"fmt"
	"strings"
)
/**
 * In echo2 , each time around the loop , the string s gets completely new contents . The += statement maks a new string by 
 * concatenating the old string, a pace character and the next argument, then assigns the new string to s .The old contents of s are
 * no longer in use , so they will be garbage-collected in due course.
 *
 *
 * If the amount of data invoved is large, this could be costly. A simpler and more efficient solution would be to use the Join 
 * function from the strings package.
 *
 * Finally , if we don't care about format but just want to see the values , perhaps for debugging, we can let Println format the 
 * results for us.
 */

func main() {
	fmt.Println(strings.Join(os.Args[1:] , " "))

	fmt.Println(os.Args[1:])
}