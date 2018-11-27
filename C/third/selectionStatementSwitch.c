#include <stdio.h>
 
#define GRADE 'A' 

/** switch( expression ) {}  expression has an integer type(char , int , e), and statement is the switch body, 
 *	which contains case labels and at most one default label
 *
 * 
 * switch( expression ){
 *	case constant-expression  :
 *     statement(s);
 *     break;
 *	case constant-expression  :
 *     statement(s);
 *     break;
 *  default :
 *  	statement(s);
 * 
 *}
 * 
 */

int main (){
   
   switch(GRADE){
   case 'A' :
      printf("excellent\n" );
      break;
   case 'B' :
   case 'C' :
      printf("well done\n" );
      break;
   case 'D' :
      printf("you passed\n" );
      break;
   case 'F' :
      printf("you have to try again\n" );
      break;
   default :
      printf("invalid grade\n" );
   }
   printf("your score is %c\n", GRADE );
 
   return 0;
}
