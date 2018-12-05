#include <stdio.h>
 
/**
 * The continue statement can be used only within the body of a loop, and causes the program flow to skip 
 * over the rest of the current iteration of the loop:
 * 
 * continue;
 * 
 * In a while or do…while loop, the program jumps to the next evaluation of the loop’s controlling expression. 
 * In a for loop, the program jumps to the next evaluation of the third expression in the for statement(adjustment), 
 * then go to next loop.
 * *****containing the operations that are performed after every loop iteration. {this is from the book but I can't understand}
 *
 */

int main ()
{
   int a = 10;
   int b;
   do{
      if( a == 15)
      {
         
         a = a + 1;
         continue;
      }
      printf("a = %d\n", a);
      a++;
     
   }while( a < 20 );

   for( b = 10 ; b < 20 ; b++){

      if(b == 15)
         continue;
      printf("b = %d\n", b );



   }
 
   return 0;
}
