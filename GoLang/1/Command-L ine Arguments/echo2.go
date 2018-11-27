package main

import(
	"os"
	"fmt"
)

func main(){
	s , sep := "", ""
/**
 * Another form of the for loop iterrates over a range of values from a date type like a string or a slice.
 *
 * In each iteration of the loop , range produces a pair of values : the index and the value of the element at that index.In this 
 * example , we don't need the index , but the sytax of a range loop requires that if we deal with the element,
 * we must deal with the index too.
 *
 * Go does not permit unused local variables , so we cannot use a tmporary variable like temp. If try unused local variables ,
 * there will be a compilation.
 *
 * The solution is to use the blank identifier , whose name is _(that is , an underscore).The blank identifier may be used whenever 
 * syntax requires a variable name but program logic does not.
 *
 * For echo2 , we only discard an unwanted loop index when we require only the element value.
 *
 *
 * In this program , we uses a short variable declaration to declare and initialize s and sep , but we could equally well have 
 * declared the variables separately.
 *
 * There are several ways to declare a string variable:
 * s := ""
 * var s string
 * var s = ""
 * var s string = ""
 * Above declaration are all equivalent.
 *
 * The first form , a short variable decalration , is the most compact , but it may be used only within a function, not for 
 * package-level variables.
 * The second form relies on default initialization to the zero value for strings , which is "".
 * The third form is rarely used except when declaring mutiple variables.
 * The fourth form is explicit about the variable's type , which is redundant when it is the same as that of the initial value ,
 * but necessary in other cases where they are not of the same type.
 *
 * In practice , you should generally use one of the first two forms , with explicit initialization to say that the initial value 
 * is important and implicit initialization to say that the initial value doesn't matter.
 */
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}