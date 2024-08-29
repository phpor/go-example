package sdk

//#cgo linux LDFLAGS: -L${SRCDIR} -lsay_linux
//#cgo darwin LDFLAGS: -L${SRCDIR} -lsay_darwin
/*
#include<say.h>
*/
import "C"

func Say() {
	C.say()
}
