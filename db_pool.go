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
	go_name := frankenphp.GoString((*unsafe.Pointer)(name));

	go_name = strings.ToUpper(go_name);

	return frankenphp.PHPString(go_name, false);
}
