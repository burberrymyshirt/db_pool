package db_pool

// #include <Zend/zend_types.h>
import "C"
import (
	"crypto/rand"
	"fmt"
	"unsafe"

	"github.com/dunglas/frankenphp"
)

type Callback func(...any);
type CallbackResult chan any;
var callbacks map[string]CallbackResult = make(map[string]CallbackResult);

//export_php:function async(callable $fn, array $args): string
func async(fn *C.zval, args *C.zend_array) (unsafe.Pointer) {
	channel := make(CallbackResult);
	go_args, err := frankenphp.GoPackedArray[any](unsafe.Pointer(args));
	if (nil != err) {
		panic("couldn't convert zend array to golang array");
	}
	c := func() {
		x := frankenphp.CallPHPCallable(unsafe.Pointer(fn), go_args);
		fmt.Printf("got value: %v\n", x);
		channel <- x;
	};
	var key string = rand.Text();
	go c();
	callbacks[key] = channel;
	return frankenphp.PHPString(key, false);
}

//export_php:function await(string $key): mixed
func await(key *C.zend_string) (any) {
	go_key := frankenphp.GoString(unsafe.Pointer(key));
	channel := callbacks[go_key]
	res := <-channel;
	return frankenphp.PHPValue(res);
}
