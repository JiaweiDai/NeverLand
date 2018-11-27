#include<stdio.h>

int main(){

	/**Like the while statement, the for statement is a top-driven loop, but with moreloop logic contained within the statement itself
	 *	for ( [expression1]; [expression2]; [expression3] ){
	 *		statements;
	 *	}
	 *
	 * The three actions that need to be executed in a typical loop are specified together at the top of the loop body:
	 * expression1 (initialization)
 	 * Evaluated only once, before the first evaluation of the controlling expression, to perform any necessary initialization.
	 * expression2 (controlling expression)
	 * Tested before each iteration. Loop execution ends when this expression evaluates to false.
	 * expression3 (adjustment)
	 * An adjustment, such as the incrementation of a counter, performed after each loop iteration and before expression2 is tested again.
	 * 
	 */
	int a;
	
	for(a = 10; a < 20; a++ )
    {
    	printf("a = %d\n", a);
    }

    /**Any of the three expressions in the head of the for loop can be omitted. This means that its shortest possible form is
     * for( ; ; )
     * A missing controlling expression is considered to be always true, and so defines an infinite loop.
     */
    
    /**
     * The following form, with no initializer and no adjustment expression, is equivalent to while ( expression ):
	 * for ( ;expression; )
	 * In fact, every for statement can also be rewritten as a while statement, and vice versa.
	 *
	 * for is generally preferable to while when the loop contains a counter or index variable that needs to be 
	 * initialized and then adjusted after each iteration.
     */
 
   return 0;
}
