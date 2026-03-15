package db_pool

// #include <Zend/zend_types.h>
import "C"
import (
	"unsafe"
	"strings"

	"github.com/dunglas/frankenphp"
)

//export_php:function get_connection(string $name): string
func get_connection(name *C.zend_string) unsafe.Pointer {
	name = frankenphp.GoString(name);

	name = strings.Reverse(name);

	return frankenphp.PHPString(name, false);
}
