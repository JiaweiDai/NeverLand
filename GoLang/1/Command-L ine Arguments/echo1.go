/**
 * This program is the same with echo
 *
 * The os package provides functions and other values for dealing wit h the operat ing system in a platform-independent fashion. 
 * Command-line arguments are avai lable to a program in a variable named Args that is par t of the os package; 
 * thus its name any where outside the os package is os.Args.
 * 
 */

package main 

import (
 		"os"
 		"fmt"
)

func main(){
	var s , sep string
	/**
	 * Variables declaration should be: 
	 * var variables-name type
	 * A variable can be initialized as part of declaration;if not explicitly initialized , 
	 * it is implicitly initialized to the zero value for its type, which is 0 for numeric types and the empty string "" for strings.
	 */
	
	for i := 1 ; i < len(os.Args) ; i++{
	/**
	 * The for loop is the only loop statement in Go.
	 * for loop has a number of forms, on of which is illustrated here:
	 * for initialization; condition; post{
	 * 		// zero or more statements
	 * }
	 * Parentheses are never used around the three componets of a for loop.
	 * The braces are mandatory, however the opening brace must be on the same line as the post statement.
	 * 
	 * 
	 * If initialization statements present, it must be a simple statement , that is :
	 * a short variable declaration , an increment or assignment statement, or a function call.
	 *
	 * condition is a boolean expression that is evaluated at the beginning of each iteration of the loop; 
	 * if it evaluates to true,the statements controlled by the loop are executed.
	 *
	 * The post statement is executed after the body of the loop, then the condition is evaluated again.
	 *
	 * Any of the three components of a for loop could be omitted.
	 * // a traditional "while" loop
	 * for condition {
	 * 		// ...
	 * }
	 *
	 * // a traditional infinite loop
	 *
	 * for {
	 * 		// ...
	 * }
	 *
	 * we will explain what is i := 1 later
	 *
	 * In Go , i++ is a statement not a expression , so it is the same as i = i + 1 ，i += 1. things lick j = i++ is illegal.
	 * only i++ and i-- is legal in Go there is no --i or ++i in go.
	 */
	 

		s += sep + os.Args[i]
	/** For numbers , Go Provides the usual arithmetic and logical oprators . 
	 * When applied to strings , the + operator concatenates the values.
	 * so the expression :
	 * 		sep + os.Args[i]
	 * represents the concatenation of the string sep and os.Args[i].
	 *
	 * I think you should know what is +=
	 * 
	 */
		sep = " "
	}
	fmt.Println(s)
}