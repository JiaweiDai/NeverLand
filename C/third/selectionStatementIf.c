#include <stdio.h>

#define FLAG 20

/**
 * There are two selection statement in C : if , switch.
 *
 * In C , we consider NULL and 0 as false .other than them are true.
 */

int main()
{
	int a = 10 ;
	printf("#################################\n");

/**   if ....  (if boolean_expression is true will excute the statements)
 * if (boolean_expression) { 
 * 		statements
 * }
 */
	printf("\t if ...\n");

	if ( a > FLAG ){
		printf("a is greater than FLAG \n");
	}

	printf("#################################\n");
/**   if ... else ...(if boolean_expression is true will excute the statement1 ,otherwise it will excute statement2)
 * if (boolean_expression){
 * 		statement1
 * }
 * else{
 * 		statement2
 * }
 */
	printf("\t if ...else...\n");

	if( a > FLAG ){
		printf("a is greater\n");
	}
	else {
		printf("FLAG is greater\n");
	}

	printf("#################################\n");

/** if... else if ... else..
 *if(boolean_expression 1)
 *{
 *   excute when boolean_expression 1 is true
 *}
 *else if( boolean_expression 2)
 *{
 *   excute when boolean_expression 2 is true
 *}
 *else if( boolean_expression 3)
 *{
 *   excute when boolean_expression 3 is true
 *}
 *else 
 *{
 *   excute when above boolean_expression are false
 *}
 */

	printf("\t if ...else if...else\n");

	if ( FLAG == 10 ){
		printf("FLAG is 10\n");
	}
	else if( FLAG == 15 ){
		printf("FLAG is 15 \n");
	}
	else if( FLAG == 20 ){
		printf("FLAG is 20\n");
	}
	else {
		printf("there is no match case.\n");
	}



	printf("#################################\n");

/** nested if 
 *
 * If several if statements are nested, then an else clause always belongs to the last if
 * (on the same block nesting level) that does not yet have an else clause
 * 
 */

	printf("\t nested if\n");

	if ( FLAG > 0){
		if( ( FLAG % 2 ) == 0){
			printf("FLAG is positive and even\n");
		}
		else{
			printf("FLAG is positive and odd\n");
		}
	}
	else
		printf("n is negative or zero\n");

	printf("#################################\n");

/** Ternary operator  :  Exp1 ? Exp2 : Exp3;
 *  if Exp1 is true then excute Exp2 , otherwise excute Exp3
 */
	printf("\t Ternary operator\n");

	printf("%d\n", 0 ? 10 : 20) ;
	printf("%d\n", 1 ? 10 : 20) ;
	printf("%d\n", (20 < 10) ? 10 :20);



	return 0;
}
