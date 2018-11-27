#include <stdio.h>

int main() {

	
    int i,j;
    for(i=1;i<=9;i++) {
        for(j=1;j<=9;j++){
            printf("%d * %d = %d\t", j, i, i*j);
        }
        printf("\n");
    }


    printf("#################################\n");


/*
    *
   ***
  *****     
 *******
*********

*/
	int n;
	printf("please enter a number :");
	scanf("%d",&n);
	int i; // control line
	int j; // control row 2*i - 1
	int k; // control space  n -i
	
	for (i=1;i<=n;i++) {
		for (k=1;k<=n-i;k++) {
			printf(" ");
		}
		for (j=1;j<=2*i-1;j++) {
			printf("*");
		}
		printf("\n");
	}

    return 0;
}
