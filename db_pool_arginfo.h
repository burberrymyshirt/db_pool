/* This is a generated file, edit the .stub.php file instead.
 * Stub hash: 103845d9850f3b9c83d23e0ffb704dd0449c8f70 */

ZEND_BEGIN_ARG_WITH_RETURN_TYPE_INFO_EX(arginfo_async, 0, 2, IS_STRING, 0)
	ZEND_ARG_TYPE_INFO(0, fn, IS_CALLABLE, 0)
	ZEND_ARG_TYPE_INFO(0, args, IS_ARRAY, 0)
ZEND_END_ARG_INFO()

ZEND_BEGIN_ARG_WITH_RETURN_TYPE_INFO_EX(arginfo_await, 0, 1, IS_MIXED, 0)
	ZEND_ARG_TYPE_INFO(0, key, IS_STRING, 0)
ZEND_END_ARG_INFO()

ZEND_FUNCTION(async);
ZEND_FUNCTION(await);

static const zend_function_entry ext_functions[] = {
	ZEND_FE(async, arginfo_async)
	ZEND_FE(await, arginfo_await)
	ZEND_FE_END
};
