#include <stdio.h>
/**
 * The goto statement causes an unconditional jump to another statement in the same function. 
 * The destination of the jump is specified by the name of a label:
 *		goto label_name;
 * A label is a name followed by a colon:
 *		label_name: statement
 * Labels have a name space of their own, which means they can have the same names as variables or types without causing conflicts. 
 * Labels may be placed before any statement, and a statement can have several labels.
 *
 * Labels serve only as destinations of goto statements, and have no effect at all if the labeled statement is reached in the
 * normal course of sequential execution.
 */




int main(){

	int a = 10 ;

	while( a < 20){
		a++;
		printf("a = %d\n", a);

		if(a == 15)
			goto EXIT;
	}
	
	printf("Hello World!"); 

	EXIT:
	printf("the while loop exit at %d\n" , a);

	return 0;

}
