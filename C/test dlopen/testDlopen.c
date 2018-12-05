/**
 * use gcc -fPIC -shared libCalculator.c -o libCalculator.so to compile libCalculator.c to a shared library
 */

#include <stdio.h>
#include <dlfcn.h>

#define LIBPATH "/root/libCalculator.so"

typedef int (*add)(int , int);
typedef int (*subtract)(int , int);
typedef int (*multiply)(int , int);
typedef int (*divide)(int , int);

typedef struct Calculator{
	void *handle;
	add addtwo;
	subtract subtracttwo;
	multiply multiplytwo;
	divide dividetwo;

}CALCULATOR;

int get_calculator(CALCULATOR * func){
	void *handle;
	char *error;

	handle = dlopen(LIBPATH , RTLD_LAZY);

	if(!handle){
		fprintf(stderr, "the error is : %s\n", dlerror());
		return -1;
	}

	func -> handle = handle;

	func -> addtwo = dlsym(handle , "add");
	if ((error = dlerror()) != NULL)  {
	   	syslog(stderr, "dlsym get add return error %s\n", error);
		return -2;
    }

	func -> subtracttwo = dlsym(handle , "subtract");
	if ((error = dlerror()) != NULL)  {
	   	syslog(stderr, "dlsym get subtract return error %s\n", error);
		return -3;
    }

	func -> multiplytwo = dlsym(handle , "multiply");
	if ((error = dlerror()) != NULL)  {
	   	syslog(stderr, "dlsym get add return error %s\n", error);
		return -4;
    }

	func -> dividetwo = dlsym(handle , "divide");
	if ((error = dlerror()) != NULL)  {
	 	syslog(stderr, "dlsym get divide return error %s\n", error);
		return -5;
    }

    return 0;

}

void clean(CALCULATOR * func){
	dlclose(func->handle);
}

int main (){
	int a , b , c , d;

	CALCULATOR funcs;

	if (get_calculator(&funcs) ) {
		fprintf(stderr, "failed to get the funcs from the library");
		return -1;
	}

	a = funcs.addtwo(10 , 20);
	b = funcs.subtracttwo(50 , 10);
	c = funcs.multiplytwo(5 , 10);
	d = funcs.dividetwo(360 , 6);

	clean(&funcs);
	printf("%d\n%d\n%d\n%d\n", a , b , c ,d);

	return 0;
}
