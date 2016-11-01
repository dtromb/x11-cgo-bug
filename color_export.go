package x11

/*
	#include <X11/Xlib.h>
	#include <X11/Xcms.h>
	#include <string.h>
*/
import "C"

//export CmsCgoCompressionFuncProxy
func CmsCgoCompressionFuncProxy(ccc C.XcmsCCC, colors_in_out *C.XcmsColor, ncolors, index C.uint, compression_flags_return *C.Bool) C.Status {
	return 0
}

//export CmsCgoWhiteAdjustFuncProxy
func CmsCgoWhiteAdjustFuncProxy(ccc C.XcmsCCC, initial, target *C.XcmsColor, format C.XcmsColorFormat, colors_in_ou *C.XcmsColor, ncolors C.int, compression_flags_resturen *C.Bool) C.Status {
	return 0
}
