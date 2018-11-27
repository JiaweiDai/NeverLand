#include<stdio.h>

int main()
{
	char c = 255;
	
	if(c == -1){
		printf("signed");
	}
	else if(c == 255){
		printf("unsigned");
	}
	else
		printf("error"); 
		
	return 0;
} 


