#include<stdio.h>

int main()
{
	int price = 0 ;
	
	printf("please enter the price :");
	scanf("%d" , &price);
	
	int change = 100 - price ;
	
	printf("the change is : %d " , change);
	
	return 0;
	
			
}
