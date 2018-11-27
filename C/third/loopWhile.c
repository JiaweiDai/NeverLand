#include <stdio.h>

int main(){
	/**A while statement executes a statement repeatedly as long as the controlling expression is true
	 * while (expression ) {
	 * 		statement
	 * }
	 * The while statement is a top-driven loop: first, the loop condition (i.e., the controlling expression) is evaluated. 
	 * If it yields true, the loop body is executed, and then the controlling expression is evaluated again. 
	 * If the condition is false, program execution continues with the statement that follows the loop body.
	 */
	
	int a = 0;
	int b = 10;
	while( a < 10){
		printf("a = %d\n", a);
		//a = a + 1;
		//a++; 
		a += 1;
	}

	/**The do…while statement is a bottom-driven loop:
	 *	do {
	 *		statement
	 *		} while ( expression )
	 * The loop body statement is executed once before the controlling expression is evaluated for the first time. 
	 * Unlike the while and for statements, do…while ensures that at least one iteration of the loop body is performed. 
	 * If the controlling expression yields true, then another iteration follows. If false, the loop is finished.
	 */
	
	do{
		printf("b = %d\n" , b);
		b++;
	} while( b < 20 );
	


	return 0;
}
