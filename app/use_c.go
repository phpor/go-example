package main

/*
#include <stdio.h>
#include <stdlib.h>
void print(char *s) {
	printf("%s", s);
}
*/
import "C"

func main() {
	s := "Hello world"
	cs := C.CString(s)
	C.print(cs)
}
