#include<stdio.h>

int main()
{
	printf("please enter foot and inch:");
	
	int foot, inch;
	
	scanf("%d %d" , &foot , &inch);
	
	printf("the hight is %f",(foot + inch/12) * 0.3048);
	
	return 0;
}
