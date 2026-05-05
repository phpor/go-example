package sdk_by_go


// extern int goCallbackHandler(int, int);
//
// static int doAdd(int a, int b) {
//     return goCallbackHandler(a, b);
// }
import "C"

//export goCallbackHandler
func goCallbackHandler(a, b C.int) C.int {
	return a + b
}

// This is the public function, callable from outside this package.
// It forwards the parameters to C.doAdd(), which in turn forwards
// them back to goCallbackHandler(). This one performs the addition
// and yields the result.
func MyAdd(a, b int) int {
	return int( C.doAdd( C.int(a), C.int(b)) )
}
