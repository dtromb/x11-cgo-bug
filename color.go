package x11

import (
	"errors"
	"reflect"
	"unsafe"
)

/*
	#include <X11/Xlib.h>
	#include <X11/Xcms.h>
	#include <string.h>

	unsigned long XColorGetPixel(XColor *xc) { return xc->pixel; }
	unsigned short XColorGetRed(XColor *xc) { return xc->red; }
	unsigned short XColorGetGreen(XColor *xc) { return xc->green; }
	unsigned short XColorGetBlue(XColor *xc) { return xc->blue; }
	int XColorGetDoRed(XColor *xc) { return (xc->flags & DoRed) > 0; }
	int XColorGetDoGreen(XColor *xc) { return (xc->flags & DoGreen) > 0; }
	int XColorGetDoBlue(XColor *xc) { return (xc->flags & DoBlue) > 0; }
	void XColorSetRed(XColor *xc, unsigned short red) { xc->red = red; }
	void XColorSetGreen(XColor *xc, unsigned short green) { xc->green = green; }
	void XColorSetBlue(XColor *xc, unsigned short blue) { xc->blue = blue; }
	void XColorSetDoRed(XColor *xc) { xc->flags |= DoRed; }
	void XColorSetDoGreen(XColor *xc) { xc->flags |= DoGreen; }
	void XColorSetDoBlue(XColor *xc) { xc->flags |= DoBlue; }
	void XColorClearDoRed(XColor *xc) { xc->flags &= ~DoRed; }
	void XColorClearDoGreen(XColor *xc) { xc->flags &= ~DoGreen; }
	void XColorClearDoBlue(XColor *xc) { xc->flags &= ~DoBlue; }
	double XcmsColorXYZGetX(XcmsColor *xc) { return xc->spec.CIEXYZ.X; }
	double XcmsColorXYZGetY(XcmsColor *xc) { return xc->spec.CIEXYZ.Y; }
	double XcmsColorXYZGetZ(XcmsColor *xc) { return xc->spec.CIEXYZ.Z; }
	void XcmsColorXYZSetX(XcmsColor *xc, double X) { xc->spec.CIEXYZ.X = X; }
	void XcmsColorXYZSetY(XcmsColor *xc, double Y) { xc->spec.CIEXYZ.Y = Y; }
	void XcmsColorXYZSetZ(XcmsColor *xc, double Z) { xc->spec.CIEXYZ.Z = Z; }
	double XcmsColoruvYGetU(XcmsColor *xc) { return xc->spec.CIEuvY.u_prime; }
	double XcmsColoruvYGetV(XcmsColor *xc) { return xc->spec.CIEuvY.v_prime; }
	double XcmsColoruvYGetY(XcmsColor *xc) { return xc->spec.CIEuvY.Y; }
	void XcmsColoruvYSetU(XcmsColor *xc, double u) { xc->spec.CIEuvY.u_prime = u; }
	void XcmsColoruvYSetV(XcmsColor *xc, double v) { xc->spec.CIEuvY.v_prime = v; }
	void XcmsColoruvYSetY(XcmsColor *xc, double Y) { xc->spec.CIEuvY.Y = Y; }
	double XcmsColorxyYGetX(XcmsColor *xc) { return xc->spec.CIExyY.x; }
	double XcmsColorxyYGetY(XcmsColor *xc) { return xc->spec.CIExyY.y; }
	double XcmsColorxyYGetYc(XcmsColor *xc) { return xc->spec.CIExyY.Y; }
	void XcmsColorxyYSetX(XcmsColor *xc, double x) { xc->spec.CIExyY.x = x; }
	void XcmsColorxyYSetY(XcmsColor *xc, double y) { xc->spec.CIExyY.y = y; }
	void XcmsColorxyYSetYc(XcmsColor *xc, double Y) { xc->spec.CIExyY.Y = Y; }
	double XcmsColorLabGetL(XcmsColor *xc) { return xc->spec.CIELab.L_star; }
	double XcmsColorLabGetA(XcmsColor *xc) { return xc->spec.CIELab.a_star; }
	double XcmsColorLabGetB(XcmsColor *xc) { return xc->spec.CIELab.b_star; }
	void XcmsColorLabSetL(XcmsColor *xc, double L) { xc->spec.CIELab.L_star = L; }
	void XcmsColorLabSetA(XcmsColor *xc, double a) { xc->spec.CIELab.a_star = a; }
	void XcmsColorLabSetB(XcmsColor *xc, double b) { xc->spec.CIELab.b_star = b; }
	double XcmsColorLuvGetL(XcmsColor *xc) { return xc->spec.CIELuv.L_star; }
	double XcmsColorLuvGetU(XcmsColor *xc) { return xc->spec.CIELuv.u_star; }
	double XcmsColorLuvGetV(XcmsColor *xc) { return xc->spec.CIELuv.v_star; }
	void XcmsColorLuvSetL(XcmsColor *xc, double L) { xc->spec.CIELuv.L_star = L; }
	void XcmsColorLuvSetU(XcmsColor *xc, double u) { xc->spec.CIELuv.u_star = u; }
	void XcmsColorLuvSetV(XcmsColor *xc, double v) { xc->spec.CIELuv.v_star = v; }
	double XcmsColorHVCGetH(XcmsColor *xc) { return xc->spec.TekHVC.H; }
	double XcmsColorHVCGetV(XcmsColor *xc) { return xc->spec.TekHVC.V; }
	double XcmsColorHVCGetC(XcmsColor *xc) { return xc->spec.TekHVC.C; }
	void XcmsColorHVCSetH(XcmsColor *xc, double H) { xc->spec.TekHVC.H = H; }
	void XcmsColorHVCSetV(XcmsColor *xc, double V) { xc->spec.TekHVC.V = V; }
	void XcmsColorHVCSetC(XcmsColor *xc, double C) { xc->spec.TekHVC.C = C; }
	double XcmsColorRGBiGetR(XcmsColor *xc) { return xc->spec.RGBi.red; }
	double XcmsColorRGBiGetG(XcmsColor *xc) { return xc->spec.RGBi.green; }
	double XcmsColorRGBiGetB(XcmsColor *xc) { return xc->spec.RGBi.blue; }
	void XcmsColorRGBiSetR(XcmsColor *xc, double R) { xc->spec.RGBi.red = R; }
	void XcmsColorRGBiSetG(XcmsColor *xc, double G) { xc->spec.RGBi.green = G; }
	void XcmsColorRGBiSetB(XcmsColor *xc, double B) { xc->spec.RGBi.blue = B; }
	unsigned short XcmsColorRGBGetR(XcmsColor *xc) { return xc->spec.RGB.red; }
	unsigned short XcmsColorRGBGetG(XcmsColor *xc) { return xc->spec.RGB.green; }
	unsigned short XcmsColorRGBGetB(XcmsColor *xc) { return xc->spec.RGB.blue; }
	void XcmsColorRGBSetR(XcmsColor *xc, unsigned short R) { xc->spec.RGB.red = R; }
	void XcmsColorRGBSetG(XcmsColor *xc, unsigned short G) { xc->spec.RGB.green = G; }
	void XcmsColorRGBSetB(XcmsColor *xc, unsigned short B) { xc->spec.RGB.blue = B; }
	unsigned long CmsColorGetPixel(XcmsColor *xc) { return xc->pixel; }
	int CmsColorGetFormat(XcmsColor *xc) { return xc->format; }
	void CopyXcmsColor(XcmsColor *dst, XcmsColor *src) { memcpy(dst,src,sizeof(XcmsColor)); }

	extern Status CmsCgoCompressionFuncProxy(XcmsCCC ccc, XcmsColor *colors_in_out, unsigned int ncolors, unsigned int index, Bool *compression_flags_return);
 	extern Status CmsCgoWhiteAdjustFuncProxy(XcmsCCC ccc, XcmsColor *initial, XcmsColor *target, XcmsColorFormat format, XcmsColor *colors_in_out, int ncolors, Bool *compression_flags_return);

	// XXX XXX XXX
	// This was an attempt to work around the problem illustrated in https://github.com/dtromb/x11-cgo-bug.
	// Written in utter desperation, these lines prevented the problem from occuring when the corresponding
	// CCC* functions were first implemented.   The problem in the current codebase still exists if they are
	// commented out - it returned after implementing some more functions.
	//
	extern XcmsWhiteAdjustProc XcmsSetWhiteAdjustProc(XcmsCCC c, XcmsWhiteAdjustProc p, char *x);
	extern XcmsCompressionProc XcmsSetCompressionProc(XcmsCCC c, XcmsCompressionProc p, char *x);


*/
import "C"

type Color C.XColor

type Colormap C.Colormap

func (xc *Color) Pixel() Pixel {
	return Pixel(C.XColorGetPixel((*C.XColor)(xc)))
}

func (xc *Color) Red() uint16 {
	return uint16(C.XColorGetRed((*C.XColor)(xc)))
}

func (xc *Color) Green() uint16 {
	return uint16(C.XColorGetGreen((*C.XColor)(xc)))
}

func (xc *Color) Blue() uint16 {
	return uint16(C.XColorGetBlue((*C.XColor)(xc)))
}

func (xc *Color) DoRed() bool {
	return C.XColorGetDoRed((*C.XColor)(xc)) > 0
}

func (xc *Color) DoGreen() bool {
	return C.XColorGetDoGreen((*C.XColor)(xc)) > 0
}

func (xc *Color) DoBlue() bool {
	return C.XColorGetDoBlue((*C.XColor)(xc)) > 0
}

func (xc *Color) SetRed(r uint16) {
	C.XColorSetRed((*C.XColor)(xc), C.ushort(r))
}

func (xc *Color) SetGreen(g uint16) {
	C.XColorSetGreen((*C.XColor)(xc), C.ushort(g))
}

func (xc *Color) SetBlue(b uint16) {
	C.XColorSetBlue((*C.XColor)(xc), C.ushort(b))
}

func (xc *Color) SetDoRed(v bool) {
	if v {
		C.XColorSetDoRed((*C.XColor)(xc))
	} else {
		C.XColorClearDoRed((*C.XColor)(xc))
	}
}

func (xc *Color) SetDoGreen(v bool) {
	if v {
		C.XColorSetDoGreen((*C.XColor)(xc))
	} else {
		C.XColorClearDoGreen((*C.XColor)(xc))
	}
}

func (xc *Color) SetDoBlue(v bool) {
	if v {
		C.XColorSetDoBlue((*C.XColor)(xc))
	} else {
		C.XColorClearDoBlue((*C.XColor)(xc))
	}
}

type CmsColorFormat C.uint

const (
	CmsUndefinedFormat CmsColorFormat = 0x00000000
	CmsCIEXYZFormat    CmsColorFormat = 0x00000001
	CmsCIEuvYFormat    CmsColorFormat = 0x00000002
	CmsCIExyYFormat    CmsColorFormat = 0x00000003
	CmsCIELabFormat    CmsColorFormat = 0x00000004
	CmsCIELuvFormat    CmsColorFormat = 0x00000005
	CmsTekHVCFormat    CmsColorFormat = 0x00000006
	CmsRGBFormat       CmsColorFormat = 0x80000000
	CmsRGBiFormat      CmsColorFormat = 0x80000001
)

type CmsColor C.XcmsColor

func (c *CmsColor) Pixel() Pixel {
	return Pixel(C.CmsColorGetPixel((*C.XcmsColor)(c)))
}

func (c *CmsColor) Format() CmsColorFormat {
	return CmsColorFormat(C.CmsColorGetFormat((*C.XcmsColor)(c)))
}

type CIEXYZColor CmsColor
type CIEuvYColor CmsColor
type CIExyYColor CmsColor
type CIELabColor CmsColor
type CIELuvColor CmsColor
type TekHVCColor CmsColor
type RGBColor CmsColor
type RGBiColor CmsColor

func (c *RGBColor) R() uint16 {
	return uint16(C.XcmsColorRGBGetR((*C.XcmsColor)(c)))
}

func (c *RGBColor) G() uint16 {
	return uint16(C.XcmsColorRGBGetG((*C.XcmsColor)(c)))
}

func (c *RGBColor) B() uint16 {
	return uint16(C.XcmsColorRGBGetB((*C.XcmsColor)(c)))
}

func (c *RGBColor) SetR(r uint16) {
	C.XcmsColorRGBSetR((*C.XcmsColor)(c), C.ushort(r))
}

func (c *RGBColor) SetG(g uint16) {
	C.XcmsColorRGBSetG((*C.XcmsColor)(c), C.ushort(g))
}

func (c *RGBColor) SetB(b uint16) {
	C.XcmsColorRGBSetB((*C.XcmsColor)(c), C.ushort(b))
}

func (c *RGBiColor) R() float64 {
	return float64(C.XcmsColorRGBiGetR((*C.XcmsColor)(c)))
}

func (c *RGBiColor) G() float64 {
	return float64(C.XcmsColorRGBiGetG((*C.XcmsColor)(c)))
}

func (c *RGBiColor) B() float64 {
	return float64(C.XcmsColorRGBiGetB((*C.XcmsColor)(c)))
}

func (c *RGBiColor) SetR(r float64) {
	C.XcmsColorRGBiSetR((*C.XcmsColor)(c), C.double(r))
}

func (c *RGBiColor) SetG(g float64) {
	C.XcmsColorRGBiSetG((*C.XcmsColor)(c), C.double(g))
}

func (c *RGBiColor) SetB(b float64) {
	C.XcmsColorRGBiSetB((*C.XcmsColor)(c), C.double(b))
}

func (c *CIEXYZColor) X() float64 {
	return float64(C.XcmsColorXYZGetX((*C.XcmsColor)(c)))
}

func (c *CIEXYZColor) Y() float64 {
	return float64(C.XcmsColorXYZGetY((*C.XcmsColor)(c)))
}

func (c *CIEXYZColor) Z() float64 {
	return float64(C.XcmsColorXYZGetZ((*C.XcmsColor)(c)))
}

func (c *CIEXYZColor) SetX(x float64) {
	C.XcmsColorXYZSetX((*C.XcmsColor)(c), C.double(x))
}

func (c *CIEXYZColor) SetY(y float64) {
	C.XcmsColorXYZSetY((*C.XcmsColor)(c), C.double(y))
}

func (c *CIEXYZColor) SetZ(z float64) {
	C.XcmsColorXYZSetZ((*C.XcmsColor)(c), C.double(z))
}

func (c *CIEuvYColor) U() float64 {
	return float64(C.XcmsColoruvYGetU((*C.XcmsColor)(c)))
}

func (c *CIEuvYColor) V() float64 {
	return float64(C.XcmsColoruvYGetV((*C.XcmsColor)(c)))
}

func (c *CIEuvYColor) Y() float64 {
	return float64(C.XcmsColoruvYGetY((*C.XcmsColor)(c)))
}

func (c *CIEuvYColor) SetU(u float64) {
	C.XcmsColoruvYSetU((*C.XcmsColor)(c), C.double(u))
}

func (c *CIEuvYColor) SetV(v float64) {
	C.XcmsColoruvYSetV((*C.XcmsColor)(c), C.double(v))
}

func (c *CIEuvYColor) SetY(y float64) {
	C.XcmsColoruvYSetY((*C.XcmsColor)(c), C.double(y))
}

func (c *CIExyYColor) X() float64 {
	return float64(C.XcmsColorxyYGetX((*C.XcmsColor)(c)))
}

func (c *CIExyYColor) Y() float64 {
	return float64(C.XcmsColorxyYGetY((*C.XcmsColor)(c)))
}

func (c *CIExyYColor) Yc() float64 {
	return float64(C.XcmsColorxyYGetYc((*C.XcmsColor)(c)))
}

func (c *CIExyYColor) SetX(x float64) {
	C.XcmsColorxyYSetX((*C.XcmsColor)(c), C.double(x))
}

func (c *CIExyYColor) SetY(y float64) {
	C.XcmsColorxyYSetY((*C.XcmsColor)(c), C.double(y))
}

func (c *CIExyYColor) SetYc(yc float64) {
	C.XcmsColorxyYSetYc((*C.XcmsColor)(c), C.double(yc))
}

func (c *CIELabColor) L() float64 {
	return float64(C.XcmsColorLabGetL((*C.XcmsColor)(c)))
}

func (c *CIELabColor) A() float64 {
	return float64(C.XcmsColorLabGetA((*C.XcmsColor)(c)))
}

func (c *CIELabColor) B() float64 {
	return float64(C.XcmsColorLabGetB((*C.XcmsColor)(c)))
}

func (c *CIELabColor) SetL(l float64) {
	C.XcmsColorLabSetL((*C.XcmsColor)(c), C.double(l))
}

func (c *CIELabColor) SetA(a float64) {
	C.XcmsColorLabSetA((*C.XcmsColor)(c), C.double(a))
}

func (c *CIELabColor) SetB(b float64) {
	C.XcmsColorLabSetB((*C.XcmsColor)(c), C.double(b))
}

func (c *CIELuvColor) L() float64 {
	return float64(C.XcmsColorLuvGetL((*C.XcmsColor)(c)))
}

func (c *CIELuvColor) U() float64 {
	return float64(C.XcmsColorLuvGetU((*C.XcmsColor)(c)))
}

func (c *CIELuvColor) V() float64 {
	return float64(C.XcmsColorLuvGetV((*C.XcmsColor)(c)))
}

func (c *CIELuvColor) SetL(l float64) {
	C.XcmsColorLuvSetL((*C.XcmsColor)(c), C.double(l))
}

func (c *CIELuvColor) SetU(u float64) {
	C.XcmsColorLuvSetU((*C.XcmsColor)(c), C.double(u))
}

func (c *CIELuvColor) SetV(v float64) {
	C.XcmsColorLuvSetV((*C.XcmsColor)(c), C.double(v))
}

func (cl *TekHVCColor) H() float64 {
	return float64(C.XcmsColorHVCGetH((*C.XcmsColor)(cl)))
}

func (cl *TekHVCColor) V() float64 {
	return float64(C.XcmsColorHVCGetV((*C.XcmsColor)(cl)))
}

func (cl *TekHVCColor) C() float64 {
	return float64(C.XcmsColorHVCGetC((*C.XcmsColor)(cl)))
}

func (cl *TekHVCColor) SetH(h float64) {
	C.XcmsColorHVCSetH((*C.XcmsColor)(cl), C.double(h))
}

func (cl *TekHVCColor) SetV(v float64) {
	C.XcmsColorHVCSetV((*C.XcmsColor)(cl), C.double(v))
}

func (cl *TekHVCColor) SetC(c float64) {
	C.XcmsColorHVCSetC((*C.XcmsColor)(cl), C.double(c))
}

func (d *Display) CreateColormap(w *Window, visual *Visual, allocate bool) Colormap {
	var alloc C.int
	if allocate {
		alloc = C.AllocAll
	} else {
		alloc = C.AllocNone
	}
	return Colormap(C.XCreateColormap(d.hnd, w.hnd, visual.hnd, alloc))
}

func (d *Display) CopyColormapAndFree(cm Colormap) Colormap {
	return Colormap(C.XCopyColormapAndFree(d.hnd, C.Colormap(cm)))
}

func (d *Display) FreeColormap(cm Colormap) {
	C.XFreeColormap(d.hnd, C.Colormap(cm))
}

func (d *Display) LookupColor(cm Colormap, colorName string) (exact *Color, screen *Color, e error) {
	var cExact, cScreen C.XColor
	st := C.XLookupColor(d.hnd, C.Colormap(cm), C.CString(colorName), &cExact, &cScreen)
	if st == 0 {
		return nil, nil, errors.New("could not resolve color name")
	}
	return (*Color)(&cExact), (*Color)(&cScreen), nil
}

func (d *Display) ParseColor(cm Colormap, spec string) (*Color, error) {
	var cColor C.XColor
	st := C.XParseColor(d.hnd, C.Colormap(cm), C.CString(spec), &cColor)
	if st == 0 {
		return nil, errors.New("could not resolve color specification")
	}
	return (*Color)(&cColor), nil
}

func (d *Display) CmsLookupColor(cm Colormap, spec string, format CmsColorFormat) (exact *CmsColor, screen *CmsColor, e error) {
	var cExact, cScreen C.XcmsColor
	st := C.XcmsLookupColor(d.hnd, C.Colormap(cm), C.CString(spec), &cExact, &cScreen, C.XcmsColorFormat(format))
	if st == 0 {
		return nil, nil, errors.New("could not resolve cms color specification")
	}
	return (*CmsColor)(&cExact), (*CmsColor)(&cScreen), nil
}

func (d *Display) AllocColor(cm Colormap, c *Color) error {
	st := C.XAllocColor(d.hnd, C.Colormap(cm), (*C.XColor)(c))
	if st == 0 {
		return errors.New("could not allocate color entry")
	}
	return nil
}

func (d *Display) CmsAllocColor(cm Colormap, c *CmsColor, format CmsColorFormat) error {
	st := C.XcmsAllocColor(d.hnd, C.Colormap(cm), (*C.XcmsColor)(c), C.XcmsColorFormat(format))
	if st == 0 {
		return errors.New("could not allocate color entry")
	}
	return nil
}

func (d *Display) AllocNamedColor(cm Colormap, name string) (exact *Color, screen *Color, e error) {
	var cExact, cScreen C.XColor
	st := C.XAllocNamedColor(d.hnd, C.Colormap(cm), C.CString(name), &cScreen, &cExact)
	if st == 0 {
		return nil, nil, errors.New("could not allocate color")
	}
	return (*Color)(&cExact), (*Color)(&cScreen), nil
}

func (d *Display) CmsAllocNamedColor(cm Colormap, name string, format CmsColorFormat) (exact *CmsColor, screen *CmsColor, e error) {
	var cExact, cScreen C.XcmsColor
	st := C.XcmsAllocNamedColor(d.hnd, C.Colormap(cm), C.CString(name), &cScreen, &cExact, C.XcmsColorFormat(format))
	if st == 0 {
		return nil, nil, errors.New("could not allocate color")
	}
	return (*CmsColor)(&cExact), (*CmsColor)(&cScreen), nil
}

func (d *Display) AllocColorCells(cm Colormap, contig bool, planes int, pixels int) ([]BitField, []Pixel, error) {
	var b C.Bool
	if contig {
		b = 1
	} else {
		b = 0
	}
	cPlanes := make([]C.ulong, planes)
	cPixels := make([]C.ulong, pixels)
	planesHdr := (*reflect.SliceHeader)(unsafe.Pointer(&cPlanes))
	pixelsHdr := (*reflect.SliceHeader)(unsafe.Pointer(&cPixels))
	st := C.XAllocColorCells(d.hnd, C.Colormap(cm), b, (*C.ulong)(unsafe.Pointer(planesHdr.Data)),
		C.uint(planes), (*C.ulong)(unsafe.Pointer(pixelsHdr.Data)), C.uint(pixels))
	if st == 0 {
		return nil, nil, errors.New("could not allocate color cells")
	}
	rPlanes := make([]BitField, planes)
	for i := 0; i < len(rPlanes); i++ {
		rPlanes[i] = BitField(uint64(cPlanes[i]))
	}
	rPixels := make([]Pixel, pixels)
	for i := 0; i < len(rPixels); i++ {
		rPixels[i] = Pixel(cPixels[i])
	}
	return rPlanes, rPixels, nil
}

func (d *Display) AllocColorPlanes(cm Colormap, contig bool, nColors int, nReds int, nGreens int, nBlues int) (p []Pixel, rm int, gm int, bm int, e error) {
	var b C.Bool
	if contig {
		b = 1
	} else {
		b = 0
	}
	var cRm, cGm, cBm C.ulong
	cPixels := make([]C.ulong, nColors)
	pixelsHdr := (*reflect.SliceHeader)(unsafe.Pointer(&cPixels))
	st := C.XAllocColorPlanes(d.hnd, C.Colormap(cm), b, (*C.ulong)(unsafe.Pointer(pixelsHdr.Data)), C.int(nColors), C.int(nReds), C.int(nGreens), C.int(nBlues), &cRm, &cGm, &cBm)
	if st == 0 {
		return nil, 0, 0, 0, errors.New("could not allocate color planes")
	}
	rPixels := make([]Pixel, nColors)
	for i := 0; i < len(rPixels); i++ {
		rPixels[i] = Pixel(cPixels[i])
	}
	return rPixels, int(cRm), int(cGm), int(cBm), nil
}

func (d *Display) FreeColors(cm Colormap, pixels []Pixel, planes BitField) {
	cPixels := make([]C.ulong, len(pixels))
	for i, p := range pixels {
		cPixels[i] = C.ulong(p)
	}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&cPixels))
	C.XFreeColors(d.hnd, C.Colormap(cm), (*C.ulong)(unsafe.Pointer(hdr.Data)), C.int(hdr.Len), C.ulong(planes))
}

func (d *Display) StoreColor(cm Colormap, c Color) {
	C.XStoreColor(d.hnd, C.Colormap(cm), (*C.XColor)(&c))
}

func (d *Display) StoreColors(cm Colormap, c []Color) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	C.XStoreColors(d.hnd, C.Colormap(cm), (*C.XColor)(unsafe.Pointer(hdr.Data)), C.int(hdr.Len))
}

func (d *Display) CmsStoreColor(cm Colormap, c CmsColor) {
	C.XcmsStoreColor(d.hnd, C.Colormap(cm), (*C.XcmsColor)(&c))
}

func (d *Display) CmsStoreColors(cm Colormap, c []CmsColor) []bool {
	cmpstat := make([]C.Bool, len(c))
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	rhdr := (*reflect.SliceHeader)(unsafe.Pointer(&cmpstat))
	C.XcmsStoreColors(d.hnd, C.Colormap(cm), (*C.XcmsColor)(unsafe.Pointer(hdr.Data)), C.uint(hdr.Len), (*C.Bool)(unsafe.Pointer(rhdr.Data)))
	stat := make([]bool, len(c))
	for i, b := range cmpstat {
		stat[i] = b > 0
	}
	return stat
}

type StoreColorFlags C.int

const (
	StoreColorDoRed   = C.DoRed
	StoreColorDoGreen = C.DoGreen
	StoreColorDoBlue  = C.DoBlue
)

func (d *Display) StoreNamedColor(cm Colormap, colorName string, p Pixel, storeColor StoreColorFlags) {
	C.XStoreNamedColor(d.hnd, C.Colormap(cm), C.CString(colorName), C.ulong(p), C.int(storeColor))
}

func (d *Display) QueryColor(cm Colormap, c *Color) {
	C.XQueryColor(d.hnd, C.Colormap(cm), (*C.XColor)(c))
}

func (d *Display) QueryColors(cm Colormap, c []Color) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	C.XQueryColors(d.hnd, C.Colormap(cm), (*C.XColor)(unsafe.Pointer(hdr.Data)), C.int(hdr.Len))
}

func (d *Display) CmsQueryColor(cm Colormap, c *CmsColor, format CmsColorFormat) error {
	st := C.XcmsQueryColor(d.hnd, C.Colormap(cm), (*C.XcmsColor)(c), C.XcmsColorFormat(format))
	if st == 0 {
		return errors.New("could not query color")
	}
	return nil
}

func (d *Display) CmsQueryColors(cm Colormap, c []CmsColor, format CmsColorFormat) error {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	st := C.XcmsQueryColors(d.hnd, C.Colormap(cm), (*C.XcmsColor)(unsafe.Pointer(hdr.Data)), C.uint(hdr.Len), C.XcmsColorFormat(format))
	if st == 0 {
		return errors.New("could not query color")
	}
	return nil
}

type CmsCCC struct {
	hnd              C.XcmsCCC
	display          *Display
	visual           *Visual
	screen           int
	hasScreen        bool
	screenWhitePoint CmsColor
	hasSWP           bool
	whitePoint       CmsColor
	hasWP            bool
}

func (cc *CmsCCC) Display() *Display {
	if cc.display == nil {
		cc.display = &Display{hnd: C.XcmsDisplayOfCCC(cc.hnd)}
	}
	return cc.display
}

func (cc *CmsCCC) Visual() *Visual {
	if cc.visual == nil {
		cc.visual = &Visual{hnd: C.XcmsVisualOfCCC(cc.hnd)}
	}
	return cc.visual
}

func (cc *CmsCCC) ScreenNumber() int {
	if !cc.hasScreen {
		cc.screen = int(C.XcmsScreenNumberOfCCC(cc.hnd))
		cc.hasScreen = true
	}
	return cc.screen
}

func (cc *CmsCCC) ScreenWhitePoint() CmsColor {
	if !cc.hasSWP {
		res := C.XcmsScreenWhitePointOfCCC(cc.hnd)
		C.CopyXcmsColor((*C.XcmsColor)(&cc.screenWhitePoint), res)
	}
	return cc.screenWhitePoint
}

func (cc *CmsCCC) ClientWhitePoint() CmsColor {
	if !cc.hasWP {
		res := C.XcmsClientWhitePointOfCCC(cc.hnd)
		C.CopyXcmsColor((*C.XcmsColor)(&cc.whitePoint), res)
	}
	return cc.whitePoint
}

func (cc *CmsCCC) SetClientWhitePoint(c CmsColor) error {
	st := C.XcmsSetWhitePoint(cc.hnd, (*C.XcmsColor)(&c))
	if st == 0 {
		return errors.New("could not set client white point")
	}
	cc.whitePoint = c
	cc.hasWP = true
	return nil
}

type CmsCompressionProc func(cc CmsCCC, colors []CmsColor, index int) ([]bool, error)

func (cc *CmsCCC) SetCompressionProc(proc CmsCompressionProc) {
	C.XcmsSetCompressionProc(cc.hnd,
		(C.XcmsCompressionProc)(unsafe.Pointer(C.CmsCgoCompressionFuncProxy)),
		(*C.char)(unsafe.Pointer(&proc)))
}

type CmsWhiteAdjustProc func(cc CmsCCC, initial *CmsColor, target *CmsColor, format CmsColorFormat, colors []CmsColor) ([]bool, error)

func (cc *CmsCCC) SetWhiteAdjustProc(proc CmsWhiteAdjustProc) {
	C.XcmsSetWhiteAdjustProc(cc.hnd,
		(C.XcmsWhiteAdjustProc)(unsafe.Pointer(C.CmsCgoWhiteAdjustFuncProxy)),
		(*C.char)(unsafe.Pointer(&proc)))
}

func (d *Display) CmsCCCOfColormap(cm Colormap) *CmsCCC {
	cc := &CmsCCC{hnd: C.XcmsCCCOfColormap(d.hnd, C.Colormap(cm))}
	cc.display = d
	return cc
}

func (d *Display) CmsDefaultCCC(screenNumber int) *CmsCCC {
	cc := &CmsCCC{hnd: C.XcmsDefaultCCC(d.hnd, C.int(screenNumber))}
	cc.display = d
	return cc
}

func (d *Display) CmsCreateCCC(screenNumber int, visual *Visual, whitePoint CmsColor,
	compressionProc CmsCompressionProc, whiteAdjustProc CmsWhiteAdjustProc) CmsCCC {
	hnd := C.XcmsCreateCCC(d.hnd, C.int(screenNumber), visual.hnd, (*C.XcmsColor)(&whitePoint),
		(C.XcmsCompressionProc)(unsafe.Pointer(C.CmsCgoCompressionFuncProxy)),
		(*C.char)(unsafe.Pointer(&compressionProc)),
		(C.XcmsWhiteAdjustProc)(unsafe.Pointer(C.CmsCgoWhiteAdjustFuncProxy)),
		(*C.char)(unsafe.Pointer(&whiteAdjustProc)))
	return &CmsCCC{
		hnd:       hnd,
		display:   d,
		visual:    visual,
		screen:    screenNumber,
		hasScreen: true,
	}
}
