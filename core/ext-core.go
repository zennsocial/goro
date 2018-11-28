package core

import (
	"math"
	"runtime"

	"github.com/MagicalTux/goro/core/phpv"
)

// WARNING: This file is auto-generated. DO NOT EDIT

func init() {
	RegisterExt(&Ext{
		Name:    "Core",
		Version: VERSION,
		Classes: []*ZClass{
			ArrayAccess,
			Closure,
			Exception,
			Iterator,
			IteratorAggregate,
			Serializable,
			Throwable,
			Traversable,
			stdClass,
		},
		Functions: map[string]*ExtFunction{
			"count":             &ExtFunction{Func: fncCount, Args: []*ExtFunctionArg{}},
			"define":            &ExtFunction{Func: fncDefine, Args: []*ExtFunctionArg{}},
			"defined":           &ExtFunction{Func: fncDefined, Args: []*ExtFunctionArg{}},
			"echo":              &ExtFunction{Func: stdFuncEcho, Args: []*ExtFunctionArg{}},
			"error_reporting":   &ExtFunction{Func: fncErrorReporting, Args: []*ExtFunctionArg{}},
			"func_get_arg":      &ExtFunction{Func: fncFuncGetArg, Args: []*ExtFunctionArg{}},
			"func_get_args":     &ExtFunction{Func: fncFuncGetArgs, Args: []*ExtFunctionArg{}},
			"func_num_args":     &ExtFunction{Func: fncFuncNumArgs, Args: []*ExtFunctionArg{}},
			"gc_collect_cycles": &ExtFunction{Func: stdFuncGcCollectCycles, Args: []*ExtFunctionArg{}},
			"gc_disable":        &ExtFunction{Func: stdFuncGcDisable, Args: []*ExtFunctionArg{}},
			"gc_enable":         &ExtFunction{Func: stdFuncGcEnable, Args: []*ExtFunctionArg{}},
			"gc_enabled":        &ExtFunction{Func: stdFuncGcEnabled, Args: []*ExtFunctionArg{}},
			"gc_mem_caches":     &ExtFunction{Func: stdFuncGcMemCaches, Args: []*ExtFunctionArg{}},
			"include":           &ExtFunction{Func: fncInclude, Args: []*ExtFunctionArg{}},
			"include_once":      &ExtFunction{Func: fncIncludeOnce, Args: []*ExtFunctionArg{}},
			"phpversion":        &ExtFunction{Func: stdFuncPhpVersion, Args: []*ExtFunctionArg{}},
			"print":             &ExtFunction{Func: fncPrint, Args: []*ExtFunctionArg{}},
			"require":           &ExtFunction{Func: fncRequire, Args: []*ExtFunctionArg{}},
			"require_once":      &ExtFunction{Func: fncRequireOnce, Args: []*ExtFunctionArg{}},
			"strcmp":            &ExtFunction{Func: fncStrcmp, Args: []*ExtFunctionArg{}},
			"strlen":            &ExtFunction{Func: fncStrlen, Args: []*ExtFunctionArg{}},
			"zend_version":      &ExtFunction{Func: stdFuncZendVersion, Args: []*ExtFunctionArg{}},
		},
		Constants: map[phpv.ZString]*phpv.ZVal{
			"DEFAULT_INCLUDE_PATH":         phpv.ZString(".:").ZVal(),
			"DIRECTORY_SEPARATOR":          phpv.ZString("/").ZVal(),
			"E_ALL":                        phpv.ZInt(phpv.E_ALL).ZVal(),
			"E_COMPILphpv.E_ERROR":         phpv.ZInt(phpv.E_COMPILE_ERROR).ZVal(),
			"E_COMPILphpv.E_WARNING":       phpv.ZInt(phpv.E_COMPILE_WARNING).ZVal(),
			"E_CORphpv.E_ERROR":            phpv.ZInt(phpv.E_CORE_ERROR).ZVal(),
			"E_CORphpv.E_WARNING":          phpv.ZInt(phpv.E_CORE_WARNING).ZVal(),
			"E_DEPRECATED":                 phpv.ZInt(phpv.E_DEPRECATED).ZVal(),
			"E_ERROR":                      phpv.ZInt(phpv.E_ERROR).ZVal(),
			"E_NOTICE":                     phpv.ZInt(phpv.E_NOTICE).ZVal(),
			"E_PARSE":                      phpv.ZInt(phpv.E_PARSE).ZVal(),
			"E_RECOVERABLphpv.E_ERROR":     phpv.ZInt(phpv.E_RECOVERABLE_ERROR).ZVal(),
			"E_STRICT":                     phpv.ZInt(phpv.E_STRICT).ZVal(),
			"E_USER_DEPRECATED":            phpv.ZInt(phpv.E_USER_DEPRECATED).ZVal(),
			"E_USER_ERROR":                 phpv.ZInt(phpv.E_USER_ERROR).ZVal(),
			"E_USER_NOTICE":                phpv.ZInt(phpv.E_USER_NOTICE).ZVal(),
			"E_USER_WARNING":               phpv.ZInt(phpv.E_USER_WARNING).ZVal(),
			"E_WARNING":                    phpv.ZInt(phpv.E_WARNING).ZVal(),
			"FALSE":                        phpv.ZBool(false).ZVal(),
			"NULL":                         phpv.ZNull{}.ZVal(),
			"PHP_EOL":                      phpv.ZString("\n").ZVal(),
			"PHP_EXTRA_VERSION":            phpv.ZString("").ZVal(),
			"PHP_FD_SETSIZE":               phpv.ZInt(1024).ZVal(),
			"PHP_FLOAT_DIG":                phpv.ZInt(15).ZVal(),
			"PHP_FLOAT_EPSILON":            phpv.ZFloat(2.220446049250313e-16).ZVal(),
			"PHP_FLOAT_MAX":                phpv.ZFloat(math.MaxFloat64).ZVal(),
			"PHP_FLOAT_MIN":                phpv.ZFloat(math.SmallestNonzeroFloat64).ZVal(),
			"PHP_INT_MAX":                  phpv.ZInt(math.MaxInt64).ZVal(),
			"PHP_INT_MIN":                  phpv.ZInt(math.MinInt64).ZVal(),
			"PHP_INT_SIZE":                 phpv.ZInt(8).ZVal(),
			"PHP_MAJOR_VERSION":            phpv.ZInt(7).ZVal(),
			"PHP_MAXPATHLEN":               phpv.ZInt(4096).ZVal(),
			"PHP_MINOR_VERSION":            phpv.ZInt(3).ZVal(),
			"PHP_OS":                       phpv.ZString(runtime.GOOS).ZVal(),
			"PHP_OS_FAMILY":                phpv.ZString(runtime.GOOS).ZVal(),
			"PHP_OUTPUT_HANDLER_CLEAN":     phpv.ZInt(BufferClean).ZVal(),
			"PHP_OUTPUT_HANDLER_CLEANABLE": phpv.ZInt(BufferCleanable).ZVal(),
			"PHP_OUTPUT_HANDLER_CONT":      phpv.ZInt(BufferWrite).ZVal(),
			"PHP_OUTPUT_HANDLER_END":       phpv.ZInt(BufferFinal).ZVal(),
			"PHP_OUTPUT_HANDLER_FINAL":     phpv.ZInt(BufferFinal).ZVal(),
			"PHP_OUTPUT_HANDLER_FLUSH":     phpv.ZInt(BufferFlush).ZVal(),
			"PHP_OUTPUT_HANDLER_FLUSHABLE": phpv.ZInt(BufferFlushable).ZVal(),
			"PHP_OUTPUT_HANDLER_REMOVABLE": phpv.ZInt(BufferRemovable).ZVal(),
			"PHP_OUTPUT_HANDLER_START":     phpv.ZInt(BufferStart).ZVal(),
			"PHP_OUTPUT_HANDLER_STDFLAGS":  phpv.ZInt(BufferCleanable | BufferFlushable | BufferRemovable).ZVal(),
			"PHP_OUTPUT_HANDLER_WRITE":     phpv.ZInt(BufferWrite).ZVal(),
			"PHP_RELEASE_VERSION":          phpv.ZInt(0).ZVal(),
			"PHP_VERSION":                  phpv.ZString(VERSION).ZVal(),
			"PHP_VERSION_ID":               phpv.ZInt(70300).ZVal(),
			"PHP_ZTS":                      phpv.ZInt(1).ZVal(),
			"TRUE":                         phpv.ZBool(true).ZVal(),
			"ZEND_THREAD_SAFE":             phpv.ZBool(true).ZVal(),
		},
	})
}
