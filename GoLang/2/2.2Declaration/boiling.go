/**
 * A declaration names a program entity and specifies some or all of its properties. There are four major kinds of declarations: 
 * var , const , type and func.
 *
 * A Go program is stored in one or more files whose names end in .go. Each file begins with a package declaration that says what 
 * package the file is part of. The package declaration is followed by any import declarations, and then a sequence of package-level
 * declarations of types, variables, constants and functions, in any order.
 */

package main 

 import "fmt"

 const boilingF = 212.0

 func main() {
 	var f = boilingF
 	var c = (f - 32) * 5 / 9
 	fmt.Printf("boiling point = %g F or %g C\n", f, c)
 }