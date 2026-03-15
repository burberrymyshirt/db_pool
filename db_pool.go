package db_pool

// #include <Zend/zend_types.h>
import "C"
import (
	"crypto/rand"
	"unsafe"

	"github.com/dunglas/frankenphp"
)

type Callback func(...interface{});
type CallbackResult chan interface{};
var callbacks map[string]CallbackResult = make(map[string]CallbackResult);

//export_php:function async(callable $fn, array $args): string;
func async(fn *C.zval, args *C.zend_array) (unsafe.Pointer) {
	channel := make(CallbackResult);
	go_args, err := frankenphp.GoPackedArray[interface{}](unsafe.Pointer(args));
	if (nil != err) {
		panic("couldn't convert zend array to golang array");
	}
	c := func(ch CallbackResult) {
		ch <- frankenphp.CallPHPCallable(unsafe.Pointer(fn), go_args);
	};
	var key string = rand.Text();
	go c(channel);
	callbacks[key] = channel;
	return frankenphp.PHPString(key, true);
}

//export_php:function await(string $key): mixed;
func await(key *C.zend_string) (unsafe.Pointer) {
	go_key := frankenphp.GoString(unsafe.Pointer(key));
	channel := callbacks[go_key]
	res := <-channel;
	return frankenphp.PHPValue(res);
}
