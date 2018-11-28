package standard

import (
	"math"

	"github.com/MagicalTux/goro/core"
	"github.com/MagicalTux/goro/core/phpv"
)

// WARNING: This file is auto-generated. DO NOT EDIT

func init() {
	core.RegisterExt(&core.Ext{
		Name:    "standard",
		Version: core.VERSION,
		Classes: []*core.ZClass{},
		Functions: map[string]*core.ExtFunction{
			"abs":                      &core.ExtFunction{Func: mathAbs, Args: []*core.ExtFunctionArg{}},
			"acos":                     &core.ExtFunction{Func: mathAcos, Args: []*core.ExtFunctionArg{}},
			"acosh":                    &core.ExtFunction{Func: mathACosh, Args: []*core.ExtFunctionArg{}},
			"array_merge":              &core.ExtFunction{Func: fncArrayMerge, Args: []*core.ExtFunctionArg{}},
			"asin":                     &core.ExtFunction{Func: mathAsin, Args: []*core.ExtFunctionArg{}},
			"asinh":                    &core.ExtFunction{Func: mathAsinh, Args: []*core.ExtFunctionArg{}},
			"atan":                     &core.ExtFunction{Func: mathAtan, Args: []*core.ExtFunctionArg{}},
			"atan2":                    &core.ExtFunction{Func: mathAtan2, Args: []*core.ExtFunctionArg{}},
			"atanh":                    &core.ExtFunction{Func: mathAtanh, Args: []*core.ExtFunctionArg{}},
			"base64_decode":            &core.ExtFunction{Func: fncBase64Decode, Args: []*core.ExtFunctionArg{}},
			"base64_encode":            &core.ExtFunction{Func: fncBase64Encode, Args: []*core.ExtFunctionArg{}},
			"bin2hex":                  &core.ExtFunction{Func: fncBin2hex, Args: []*core.ExtFunctionArg{}},
			"boolval":                  &core.ExtFunction{Func: fncBoolval, Args: []*core.ExtFunctionArg{}},
			"chdir":                    &core.ExtFunction{Func: fncChdir, Args: []*core.ExtFunctionArg{}},
			"constant":                 &core.ExtFunction{Func: constant, Args: []*core.ExtFunctionArg{}},
			"cos":                      &core.ExtFunction{Func: mathCos, Args: []*core.ExtFunctionArg{}},
			"cosh":                     &core.ExtFunction{Func: mathCosh, Args: []*core.ExtFunctionArg{}},
			"decbin":                   &core.ExtFunction{Func: fncDecbin, Args: []*core.ExtFunctionArg{}},
			"dechex":                   &core.ExtFunction{Func: fncDechex, Args: []*core.ExtFunctionArg{}},
			"decoct":                   &core.ExtFunction{Func: fncDecoct, Args: []*core.ExtFunctionArg{}},
			"deg2rad":                  &core.ExtFunction{Func: mathDeg2rad, Args: []*core.ExtFunctionArg{}},
			"die":                      &core.ExtFunction{Func: die, Args: []*core.ExtFunctionArg{}},
			"dirname":                  &core.ExtFunction{Func: fncDirname, Args: []*core.ExtFunctionArg{}},
			"dl":                       &core.ExtFunction{Func: stdFuncDl, Args: []*core.ExtFunctionArg{}},
			"doubleval":                &core.ExtFunction{Func: fncDoubleval, Args: []*core.ExtFunctionArg{}},
			"eval":                     &core.ExtFunction{Func: stdFuncEval, Args: []*core.ExtFunctionArg{}},
			"exit":                     &core.ExtFunction{Func: exit, Args: []*core.ExtFunctionArg{}},
			"exp":                      &core.ExtFunction{Func: mathExp, Args: []*core.ExtFunctionArg{}},
			"expm1":                    &core.ExtFunction{Func: mathExpm1, Args: []*core.ExtFunctionArg{}},
			"extension_loaded":         &core.ExtFunction{Func: stdFunc, Args: []*core.ExtFunctionArg{}},
			"floatval":                 &core.ExtFunction{Func: fncFloatval, Args: []*core.ExtFunctionArg{}},
			"flush":                    &core.ExtFunction{Func: fncFlush, Args: []*core.ExtFunctionArg{}},
			"fmod":                     &core.ExtFunction{Func: mathFmod, Args: []*core.ExtFunctionArg{}},
			"function_exists":          &core.ExtFunction{Func: stdFuncFuncExists, Args: []*core.ExtFunctionArg{}},
			"get_cfg_var":              &core.ExtFunction{Func: stdFuncGetCfgVar, Args: []*core.ExtFunctionArg{}},
			"get_magic_quotes_gpc":     &core.ExtFunction{Func: getMagicQuotesGpc, Args: []*core.ExtFunctionArg{}},
			"get_magic_quotes_runtime": &core.ExtFunction{Func: getMagicQuotesRuntime, Args: []*core.ExtFunctionArg{}},
			"getcwd":                   &core.ExtFunction{Func: fncGetcwd, Args: []*core.ExtFunctionArg{}},
			"getenv":                   &core.ExtFunction{Func: getenv, Args: []*core.ExtFunctionArg{}},
			"gettype":                  &core.ExtFunction{Func: fncGettype, Args: []*core.ExtFunctionArg{}},
			"hrtime":                   &core.ExtFunction{Func: stdFuncHrTime, Args: []*core.ExtFunctionArg{}},
			"hypot":                    &core.ExtFunction{Func: mathHypot, Args: []*core.ExtFunctionArg{}},
			"intval":                   &core.ExtFunction{Func: fncIntval, Args: []*core.ExtFunctionArg{}},
			"is_array":                 &core.ExtFunction{Func: fncIsArray, Args: []*core.ExtFunctionArg{}},
			"is_bool":                  &core.ExtFunction{Func: fncIsBool, Args: []*core.ExtFunctionArg{}},
			"is_double":                &core.ExtFunction{Func: fncIsDouble, Args: []*core.ExtFunctionArg{}},
			"is_float":                 &core.ExtFunction{Func: fncIsFloat, Args: []*core.ExtFunctionArg{}},
			"is_int":                   &core.ExtFunction{Func: fncIsInt, Args: []*core.ExtFunctionArg{}},
			"is_integer":               &core.ExtFunction{Func: fncIsInteger, Args: []*core.ExtFunctionArg{}},
			"is_long":                  &core.ExtFunction{Func: fncIsLong, Args: []*core.ExtFunctionArg{}},
			"is_null":                  &core.ExtFunction{Func: fncIsNull, Args: []*core.ExtFunctionArg{}},
			"is_numeric":               &core.ExtFunction{Func: fncIsNumeric, Args: []*core.ExtFunctionArg{}},
			"is_object":                &core.ExtFunction{Func: fncIsObject, Args: []*core.ExtFunctionArg{}},
			"is_real":                  &core.ExtFunction{Func: fncIsReal, Args: []*core.ExtFunctionArg{}},
			"is_resource":              &core.ExtFunction{Func: fncIsResource, Args: []*core.ExtFunctionArg{}},
			"is_scalar":                &core.ExtFunction{Func: fncIsScalar, Args: []*core.ExtFunctionArg{}},
			"is_string":                &core.ExtFunction{Func: fncIsString, Args: []*core.ExtFunctionArg{}},
			"microtime":                &core.ExtFunction{Func: fncMicrotime, Args: []*core.ExtFunctionArg{}},
			"ob_clean":                 &core.ExtFunction{Func: fncObClean, Args: []*core.ExtFunctionArg{}},
			"ob_end_clean":             &core.ExtFunction{Func: fncObEndClean, Args: []*core.ExtFunctionArg{}},
			"ob_end_flush":             &core.ExtFunction{Func: fncObEndFlush, Args: []*core.ExtFunctionArg{}},
			"ob_flush":                 &core.ExtFunction{Func: fncObFlush, Args: []*core.ExtFunctionArg{}},
			"ob_get_clean":             &core.ExtFunction{Func: fncObGetClean, Args: []*core.ExtFunctionArg{}},
			"ob_get_contents":          &core.ExtFunction{Func: fncObGetContents, Args: []*core.ExtFunctionArg{}},
			"ob_get_flush":             &core.ExtFunction{Func: fncObGetFlush, Args: []*core.ExtFunctionArg{}},
			"ob_get_level":             &core.ExtFunction{Func: fncObGetLevel, Args: []*core.ExtFunctionArg{}},
			"ob_implicit_flush":        &core.ExtFunction{Func: fncObImplicitFlush, Args: []*core.ExtFunctionArg{}},
			"ob_start":                 &core.ExtFunction{Func: fncObStart, Args: []*core.ExtFunctionArg{}},
			"php_sapi_name":            &core.ExtFunction{Func: stdFuncSapiName, Args: []*core.ExtFunctionArg{}},
			"php_uname":                &core.ExtFunction{Func: fncUname, Args: []*core.ExtFunctionArg{}},
			"pi":                       &core.ExtFunction{Func: mathPi, Args: []*core.ExtFunctionArg{}},
			"print_r":                  &core.ExtFunction{Func: fncPrintR, Args: []*core.ExtFunctionArg{}},
			"putenv":                   &core.ExtFunction{Func: putenv, Args: []*core.ExtFunctionArg{}},
			"rawurlencode":             &core.ExtFunction{Func: fncRawurlencode, Args: []*core.ExtFunctionArg{}},
			"set_time_limit":           &core.ExtFunction{Func: fncSetTimeLimit, Args: []*core.ExtFunctionArg{}},
			"sleep":                    &core.ExtFunction{Func: stdFuncSleep, Args: []*core.ExtFunctionArg{}},
			"sprintf":                  &core.ExtFunction{Func: fncSprintf, Args: []*core.ExtFunctionArg{}},
			"str_replace":              &core.ExtFunction{Func: stdStrReplace, Args: []*core.ExtFunctionArg{}},
			"str_rot13":                &core.ExtFunction{Func: fncStrRot13, Args: []*core.ExtFunctionArg{}},
			"strtolower":               &core.ExtFunction{Func: fncStrToLower, Args: []*core.ExtFunctionArg{}},
			"strval":                   &core.ExtFunction{Func: fncStrval, Args: []*core.ExtFunctionArg{}},
			"time":                     &core.ExtFunction{Func: fncTime, Args: []*core.ExtFunctionArg{}},
			"urlencode":                &core.ExtFunction{Func: fncUrlencode, Args: []*core.ExtFunctionArg{}},
			"usleep":                   &core.ExtFunction{Func: stdFuncUsleep, Args: []*core.ExtFunctionArg{}},
			"var_dump":                 &core.ExtFunction{Func: stdFuncVarDump, Args: []*core.ExtFunctionArg{}},
		},
		Constants: map[phpv.ZString]*phpv.ZVal{
			"INF":                 phpv.ZFloat(math.Inf(0)).ZVal(),
			"M_1_PI":              phpv.ZFloat(1 / math.Pi).ZVal(),
			"M_2_PI":              phpv.ZFloat(2 / math.Pi).ZVal(),
			"M_2_SQRTPI":          phpv.ZFloat(2 / math.Sqrt(math.Pi)).ZVal(),
			"M_E":                 phpv.ZFloat(math.E).ZVal(),
			"M_EULER":             phpv.ZFloat(0.57721566490153286061).ZVal(),
			"M_LN2":               phpv.ZFloat(math.Ln2).ZVal(),
			"M_LNPI":              phpv.ZFloat(math.Log(math.Pi)).ZVal(),
			"M_LOG10E":            phpv.ZFloat(math.Log10E).ZVal(),
			"M_LOG2E":             phpv.ZFloat(math.Log2E).ZVal(),
			"M_PHI":               phpv.ZFloat(math.Phi).ZVal(),
			"M_PI":                phpv.ZFloat(math.Pi).ZVal(),
			"M_PI_2":              phpv.ZFloat(math.Pi / 2).ZVal(),
			"M_PI_4":              phpv.ZFloat(math.Pi / 4).ZVal(),
			"M_SQRT1_2":           phpv.ZFloat(1 / math.Sqrt(2)).ZVal(),
			"M_SQRT2":             phpv.ZFloat(math.Sqrt(2)).ZVal(),
			"M_SQRT3":             phpv.ZFloat(math.Sqrt(3)).ZVal(),
			"M_SQRTPI":            phpv.ZFloat(math.Sqrt(math.Pi)).ZVal(),
			"NAN":                 phpv.ZFloat(math.NaN()).ZVal(),
			"PHP_ROUND_HALF_DOWN": phpv.ZInt(2).ZVal(),
			"PHP_ROUND_HALF_EVEN": phpv.ZInt(3).ZVal(),
			"PHP_ROUND_HALF_ODD":  phpv.ZInt(4).ZVal(),
			"PHP_ROUND_HALF_UP":   phpv.ZInt(1).ZVal(),
		},
	})
}
