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
var callbacks map[string]CallbackResult;

//export_php:function async(callable $fn, array $args): string
func async(fn *C.zval, args *C.zend_array) (unsafe.Pointer) {
	channel := make(CallbackResult);
	c := func(ch CallbackResult) {
		ch <- frankenphp.CallPHPCallable(unsafe.Pointer(fn), args);
	};
	var key string = rand.Text();
	go c(channel);
	callbacks[key] = channel;
	return frankenphp.PHPString(key, true);
}

func await(key C.zend_string) (unsafe.Pointer) {
	go_key := frankenphp.GoString(unsafe.Pointer(key));
	channel := callbacks[go_key]
	res := <-channel;
	return frankenphp.PHPValue(res);
}
