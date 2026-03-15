/* This is a generated file, edit the .stub.php file instead.
 * Stub hash: a218b258a5a6c4de12b2335f8c7abd45abc811fa */

ZEND_BEGIN_ARG_WITH_RETURN_TYPE_INFO_EX(arginfo_async, 0, 2, IS_STRING, 0)
	ZEND_ARG_TYPE_INFO(0, fn, IS_CALLABLE, 0)
	ZEND_ARG_TYPE_INFO(0, args, IS_ARRAY, 0)
ZEND_END_ARG_INFO()

ZEND_FUNCTION(async);

static const zend_function_entry ext_functions[] = {
	ZEND_FE(async, arginfo_async)
	ZEND_FE_END
};
