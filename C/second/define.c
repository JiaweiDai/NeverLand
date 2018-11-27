#include <stdio.h>

#define URL "http://blupages.ibm.com"
#define NAME "Jia Wei Dai"
#define YEAR 2014
#define MONTH 4
#define DAY 28

int main()
{
	printf("%s work in IBM from %d/%d/%d\n", NAME, YEAR, MONTH, DAY);
	printf("My name is %s\n",NAME);
	printf("%s url is %s\n", NAME, URL);

	return 0;
}

