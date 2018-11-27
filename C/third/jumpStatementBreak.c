#include <stdio.h>


int main(){
	/**
	 * The break statement can occur only in the body of a loop or a switch statement,
	 * and causes a jump to the first statement after the loop or switch statement in which it is immediately contained:
	 * 
	 * break;
	 * 
	 * Thus, the break statement can be used to end the execution of a loop statement at any position in the loop body.
	 */
	int a = 10;

   while( a < 20 )
   {
      printf("a = %d\n", a);
      a++;
      if( a > 15)
      {
          break;
      }
   }
 
   return 0;
}