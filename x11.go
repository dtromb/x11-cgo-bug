package x11

import (
	"encoding/binary"
	"errors"
	"math"
	"reflect"
	"runtime"
	"time"
	"unsafe"
)

// X11 library bindings for Go.   This is a work in progress.

/*
	#cgo LDFLAGS: -lX11
	#include <X11/Xlib.h>
	#include <string.h>
	#include <malloc.h>

	int DisplayIsNull(Display *d) {
		return d == NULL;
	}

	int IntPtrIsNull(int *ptr) {
		return ptr == NULL;
	}

	int CharPtrIsNull(char *ptr) {
		return ptr == NULL;
	}

	int UCharPtrIsNull(unsigned char *ptr) {
		return ptr == NULL;
	}

	int IntPtrIndex(int *ptr, int idx) {
		return ptr[idx];
	}

	Atom AtomPtrIndex(Atom *atoms, int idx) {
		return atoms[idx];
	}

	int XPixmapFormatValuesIsNull(XPixmapFormatValues *x) {
		return x == NULL;
	}

	int XPixmapFormatValuesIndex_BitsPerPixel(XPixmapFormatValues *x, int idx) {
		return x[idx].bits_per_pixel;
	}

	int XPixmapFormatValuesIndex_Depth(XPixmapFormatValues *x, int idx) {
		return x[idx].depth;
	}

	int XPixmapFormayValuesIndex_ScanlinePad(XPixmapFormatValues *x, int idx) {
		return x[idx].scanline_pad;
	}

	Pixmap GetBackPixmap(XSetWindowAttributes *attrs) {
		return attrs->background_pixmap;
	}

	void SetBackPixmap(XSetWindowAttributes *attrs, Pixmap p) {
		attrs->background_pixmap = p;
	}

	unsigned long GetBackPixel(XSetWindowAttributes *attrs) {
		return attrs->background_pixel;
	}

	void SetBackPixel(XSetWindowAttributes *attrs, unsigned long p) {
		attrs->background_pixel = p;
	}

	unsigned long GetBackingPixel(XSetWindowAttributes *attrs) {
		return attrs->backing_pixel;
	}

	void SetBackingPixel(XSetWindowAttributes *attrs, unsigned long pixel) {
		attrs->backing_pixel = pixel;
	}

	unsigned long GetBackingPlanes(XSetWindowAttributes *attrs) {
		return attrs->backing_planes;
	}

	void SetBackingPlanes(XSetWindowAttributes *attrs, unsigned long planes) {
		attrs->backing_planes = planes;
	}

	int GetBackingStore(XSetWindowAttributes *attrs) {
		return attrs->backing_store;
	}

	void SetBackingStore(XSetWindowAttributes *attrs, int store) {
		attrs->backing_store = store;
	}

	int GetBitGravity(XSetWindowAttributes *attrs) {
		return attrs->bit_gravity;
	}

	void SetBitGravity(XSetWindowAttributes *attrs, int gravity) {
		attrs->bit_gravity = gravity;
	}

	unsigned long GetBorderPixel(XSetWindowAttributes *attrs) {
		return attrs->border_pixel;
	}

	void SetBorderPixel(XSetWindowAttributes *attrs, unsigned long p) {
		attrs->border_pixel = p;
	}

	Pixmap GetBorderPixmap(XSetWindowAttributes *attrs) {
		return attrs->border_pixmap;
	}

	void SetBorderPixmap(XSetWindowAttributes *attrs, Pixmap p) {
		attrs->border_pixmap = p;
	}

	Colormap GetColormap(XSetWindowAttributes *attrs) {
		return attrs->colormap;
	}

	void SetColormap(XSetWindowAttributes *attrs, Colormap c) {
		attrs->colormap = c;
	}

	Cursor GetCursor(XSetWindowAttributes *attrs) {
		return attrs->cursor;
	}

	void SetCursor(XSetWindowAttributes *attrs, Cursor c) {
		attrs->cursor = c;
	}

	int GetDoNotPropagateEventMask(XSetWindowAttributes *attrs) {
		return attrs->do_not_propagate_mask;
	}

	void SetDoNotPropagateEventMask(XSetWindowAttributes *attrs, int mask) {
		attrs->do_not_propagate_mask = mask;
	}

	Bool GetOverrideRedirect(XSetWindowAttributes *attrs) {
		return attrs->override_redirect;
	}

	void SetOverrideRedirect(XSetWindowAttributes *attrs, Bool val) {
		attrs->override_redirect = val;
	}

	int GetSaveEventMask(XSetWindowAttributes *attrs) {
		return attrs->event_mask;
	}

	void SetSaveEventMask(XSetWindowAttributes *attrs, int mask) {
		attrs->event_mask = mask;
	}

	Bool GetSaveUnder(XSetWindowAttributes *attrs) {
		return attrs->save_under;
	}

	void SetSaveUnder(XSetWindowAttributes *attrs, Bool val) {
		attrs->save_under = val;
	}

	int GetWinGravity(XSetWindowAttributes *attrs) {
		return attrs->win_gravity;
	}

	void SetWinGravity(XSetWindowAttributes *attrs, int gravity) {
		attrs->win_gravity = gravity;
	}

	int WindowChangesGetX(XWindowChanges *changes) {
		return changes->x;
	}

	void WindowChangesSetX(XWindowChanges *changes, int x) {
		changes->x = x;
	}

	int WindowChangesGetY(XWindowChanges *changes) {
		return changes->y;
	}

	void WindowChangesSetY(XWindowChanges *changes, int y) {
		changes->y = y;
	}

	int WindowChangesGetWidth(XWindowChanges *changes) {
		return changes->width;
	}

	void WindowChangesSetWidth(XWindowChanges *changes, int w) {
		changes->width = w;
	}

	int WindowChangesGetHeight(XWindowChanges *changes) {
		return changes->height;
	}

	void WindowChangesSetHeight(XWindowChanges *changes, int h) {
		changes->height = h;
	}

	int WindowChangesGetBorderWidth(XWindowChanges *changes) {
		return changes->border_width;
	}

	void WindowChangesSetBorderWidth(XWindowChanges *changes, int width) {
		changes->border_width = width;
	}

	int WindowChangesGetStackMode(XWindowChanges *changes) {
		return changes->stack_mode;
	}

	void WindowChangesSetStackMode(XWindowChanges *changes, int mode) {
		changes->stack_mode = mode;
	}

	Window WindowChangesGetSibling(XWindowChanges *changes) {
		return changes->sibling;
	}

	void WindowChangesSetSibling(XWindowChanges *changes, Window w) {
		changes->sibling = w;
	}

	Window WindowPtrIndex(Window *w, int idx) {
		return w[idx];
	}

	void SetWindowAttributesFillFromWindowAttributes(XSetWindowAttributes *swa,
													   XWindowAttributes *wa) {
		swa->bit_gravity = wa->bit_gravity;
		swa->win_gravity = wa->win_gravity;
		swa->backing_store = wa->backing_store;
		swa->backing_planes = wa->backing_planes;
		swa->backing_pixel = wa->backing_pixel;
		swa->save_under = wa->save_under;
		swa->colormap = wa->colormap;
		swa->do_not_propagate_mask = wa->do_not_propagate_mask;
		swa->override_redirect = wa->override_redirect;
	}


	void SetWindowAttributesRead(XWindowAttributes *xwa, int *x, int *y,
								  unsigned int *width, unsigned int *height,
								  unsigned int *borderWidth, unsigned int *depth,
								  Visual **visual, Window *rootWindow, Bool *mapInstalled,
								  int *mapState, unsigned long *allEvents, unsigned long *myEvents,
								  Screen **screen) {
		*x = xwa->x;
		*y = xwa->y;
		*width = xwa->width;
		*height = xwa->height;
		*borderWidth = xwa->border_width;
		*depth = xwa->depth;
		*visual = xwa->visual;
		*rootWindow = xwa->root;
		*mapInstalled = xwa->map_installed;
		*mapState = xwa->map_state;
		*allEvents = xwa->all_event_masks;
		*myEvents = xwa->your_event_mask;
		*screen = xwa->screen;
	}

	unsigned char *CreateCharBlob(size_t len) {
		return malloc(len);
	}

	void FreeCharBlob(unsigned char *x) {
		free(x);
	}

*/
import "C"

type Display struct {
	hnd *C.Display
}

func OpenDisplay(displayName string) (*Display, error) {
	nd := &Display{}
	nd.hnd = C.XOpenDisplay(C.CString(displayName))
	if C.DisplayIsNull(nd.hnd) > 0 {
		return nil, errors.New("could not open display " + displayName)
	}
	runtime.SetFinalizer(nd, displayFinalizer)
	return nd, nil
}

func displayFinalizer(d *Display) {
	if C.DisplayIsNull(d.hnd) <= 0 {
		C.XCloseDisplay(d.hnd)
	}
}

func (d *Display) Close() {
	C.XCloseDisplay(d.hnd)
	d.hnd = nil
}

func AllPlanes() C.ulong {
	return C.XAllPlanes()
}

func (d *Display) XBlackPixel(screenNumber int) C.ulong {
	return C.XBlackPixel(d.hnd, C.int(screenNumber))
}

func (d *Display) XWhitePixel(screenNumber int) C.ulong {
	return C.XBlackPixel(d.hnd, C.int(screenNumber))
}

func (d *Display) ConnectionNumber() int {
	return int(C.XConnectionNumber(d.hnd))
}

func (d *Display) DefaultColormap(screenNumber int) Colormap {
	return Colormap(C.XDefaultColormap(d.hnd, C.int(screenNumber)))
}

func (d *Display) DefaultDepth(screenNumber int) int {
	return int(C.XDefaultDepth(d.hnd, C.int(screenNumber)))
}

type Depths struct {
	hnd  *C.int
	size C.int
}

func (d *Display) ListDepths(screenNumber int) (*Depths, error) {
	nd := &Depths{}
	nd.hnd = C.XListDepths(d.hnd, C.int(screenNumber), &nd.size)
	if C.IntPtrIsNull(nd.hnd) > 0 {
		return nil, errors.New("could not get depth list")
	}
	runtime.SetFinalizer(nd, depthListFinalizer)
	return nd, nil
}

func depthListFinalizer(d *Depths) {
	if C.IntPtrIsNull(d.hnd) <= 0 {
		C.XFree(unsafe.Pointer(d.hnd))
	}
}

func (d *Depths) Len() int {
	return int(d.size)
}

func (d *Depths) Get(idx int) int {
	if idx < 0 || idx >= int(d.size) {
		panic("depth index out of bounds")
	}
	return int(C.IntPtrIndex(d.hnd, C.int(idx)))
}

func (d *Depths) Depths() []int {
	sl := make([]int, d.Len())
	for i := 0; i < d.Len(); i++ {
		sl[i] = d.Get(i)
	}
	return sl
}

/*
GC XDefaultGC (display, screen_number)
Display *display;
int screen_number;
*/

func (d *Display) DefaultRootWindow() *Window {
	return &Window{hnd: C.XDefaultRootWindow(d.hnd)}
}

func (d *Display) DefaultScreenNumber() int {
	return int(C.XDefaultScreen(d.hnd))
}

func (d *Display) Cells(screenNumber int) int {
	return int(C.XDisplayCells(d.hnd, C.int(screenNumber)))
}

func (d *Display) Planes(screenNumber int) int {
	return int(C.XDisplayPlanes(d.hnd, C.int(screenNumber)))
}

func (d *Display) String() string {
	return C.GoString(C.XDisplayString(d.hnd))
}

func (d *Display) ExtendedMaxRequestSize() uint64 {
	return uint64(C.XExtendedMaxRequestSize(d.hnd))
}

func (d *Display) MaxRequestSize() uint64 {
	return uint64(C.XMaxRequestSize(d.hnd))
}

func (d *Display) LastKnownRequestProcessed() uint64 {
	return uint64(C.XLastKnownRequestProcessed(d.hnd))
}

func (d *Display) NextRequest() uint64 {
	return uint64(C.XLastKnownRequestProcessed(d.hnd))
}

func (d *Display) ProtocolVersion() int {
	return int(C.XProtocolVersion(d.hnd))
}

func (d *Display) ProtocolRevision() int {
	return int(C.XProtocolRevision(d.hnd))
}

func (d *Display) QLength() int {
	return int(C.XQLength(d.hnd))
}

func (d *Display) RootWindow(screenNumber int) *Window {
	return &Window{hnd: C.XRootWindow(d.hnd, C.int(screenNumber))}
}

func (d *Display) ScreenCount() int {
	return int(C.XScreenCount(d.hnd))
}

func (d *Display) ServerVendor() string {
	return C.GoString(C.XServerVendor(d.hnd))
}

type Drawable interface {
	IsDrawable() bool
	Handle() C.Drawable
}
type XPixmapFormatValue struct {
	Depth        int
	BitsPerPixel int
	ScanlinePad  int
}

type XPixmapFormatValues struct {
	hnd  *C.XPixmapFormatValues
	size C.int
}

func (d *Display) ListPixmapFormats() (*XPixmapFormatValues, error) {
	xpfv := &XPixmapFormatValues{}
	xpfv.hnd = C.XListPixmapFormats(d.hnd, &xpfv.size)
	if C.XPixmapFormatValuesIsNull(xpfv.hnd) > 0 {
		return nil, errors.New("could not list pixmap formats")
	}
	runtime.SetFinalizer(xpfv, xPixmapFormatValuesFinalizer)
	return xpfv, nil
}

func xPixmapFormatValuesFinalizer(x *XPixmapFormatValues) {
	if C.XPixmapFormatValuesIsNull(x.hnd) <= 0 {
		C.XFree(unsafe.Pointer(x.hnd))
	}
}

func (xpfv *XPixmapFormatValues) Len() int {
	return int(xpfv.size)
}

func (xpfv *XPixmapFormatValues) Get(idx int) XPixmapFormatValue {
	if idx < 0 || idx >= xpfv.Len() {
		panic("pixmap format index is out of bounds")
	}
	fv := XPixmapFormatValue{
		Depth:        int(C.XPixmapFormatValuesIndex_Depth(xpfv.hnd, C.int(idx))),
		BitsPerPixel: int(C.XPixmapFormatValuesIndex_BitsPerPixel(xpfv.hnd, C.int(idx))),
		ScanlinePad:  int(C.XPixmapFormayValuesIndex_ScanlinePad(xpfv.hnd, C.int(idx))),
	}
	return fv
}

func (xpfv *XPixmapFormatValues) Values() []XPixmapFormatValue {
	vals := make([]XPixmapFormatValue, xpfv.Len())
	for i := 0; i < xpfv.Len(); i++ {
		vals[i] = xpfv.Get(i)
	}
	return vals
}

func (d *Display) ImageByteOrder() binary.ByteOrder {
	switch C.XImageByteOrder(d.hnd) {
	case C.LSBFirst:
		return binary.LittleEndian
	case C.MSBFirst:
		return binary.BigEndian
	}
	panic("unknown byte ordering")
}

func (d *Display) BitmapUnit() int {
	return int(C.XBitmapUnit(d.hnd))
}

type BitOrder C.int

const (
	BitOrderLSB BitOrder = C.LSBFirst
	BitOrderMSB          = C.MSBFirst
)

func (d *Display) BitmapBitOrder() BitOrder {
	return BitOrder(C.XBitmapBitOrder(d.hnd))
}

func (d *Display) BitmapPad() int {
	return int(C.XBitmapPad(d.hnd))
}

func (d *Display) Height(screenNumber int) int {
	return int(C.XDisplayHeight(d.hnd, C.int(screenNumber)))
}

func (d *Display) HeightMM(screenNumber int) int {
	return int(C.XDisplayHeightMM(d.hnd, C.int(screenNumber)))
}

func (d *Display) Width(screenNumber int) int {
	return int(C.XDisplayWidth(d.hnd, C.int(screenNumber)))
}

func (d *Display) WidthMM(screenNumber int) int {
	return int(C.XDisplayWidthMM(d.hnd, C.int(screenNumber)))
}

type Screen struct {
	display *Display
	hnd     *C.Screen
}

func (d *Display) DefaultScreen() *Screen {
	ns := &Screen{display: d}
	ns.hnd = C.XDefaultScreenOfDisplay(d.hnd)
	return ns
}

func (d *Display) Screen(screenNumber int) *Screen {
	ns := &Screen{display: d}
	ns.hnd = C.XScreenOfDisplay(d.hnd, C.int(screenNumber))
	return ns
}

func (s *Screen) BlackPixel() C.ulong {
	return C.XBlackPixelOfScreen(s.hnd)
}

func (s *Screen) WhitePixel() C.ulong {
	return C.XBlackPixelOfScreen(s.hnd)
}

func (s *Screen) Cells() int {
	return int(C.XCellsOfScreen(s.hnd))
}

func (s *Screen) DefaultColormap() Colormap {
	return Colormap(C.XDefaultColormapOfScreen(s.hnd))
}

func (s *Screen) DefaultDepth() int {
	return int(C.XDefaultDepthOfScreen(s.hnd))
}

/*
GC XDefaultGCOfScreen (screen)
Screen *screen;
screen Specifies the appropriate Screen structure.
*/

type BackingStoreMode C.int

const (
	BackingStoreWhenMapped = C.WhenMapped
	BackingStoreNotUseful  = C.NotUseful
	BackingStoreAlways     = C.Always
)

func (s *Screen) DoesBackingStore() BackingStoreMode {
	return BackingStoreMode(C.XDoesBackingStore(s.hnd))
}

func (s *Screen) DoesSaveUnders() bool {
	return C.XDoesSaveUnders(s.hnd) > 0
}

func (s *Screen) Display() *Display {
	return s.display
}

type EventMask C.long

func (s *Screen) ScreenNumber() int {
	return int(C.XScreenNumberOfScreen(s.hnd))
}

func (s *Screen) EventMask() EventMask {
	return EventMask(C.XEventMaskOfScreen(s.hnd))
}

func (s *Screen) Height(screenNumber int) int {
	return int(C.XHeightOfScreen(s.hnd))
}

func (s *Screen) HeightMM(screenNumber int) int {
	return int(C.XHeightMMOfScreen(s.hnd))
}

func (s *Screen) Width(screenNumber int) int {
	return int(C.XWidthOfScreen(s.hnd))
}

func (s *Screen) WidthMM(screenNumber int) int {
	return int(C.XWidthMMOfScreen(s.hnd))
}

func (s *Screen) MaxCMaps() int {
	return int(C.XMaxCmapsOfScreen(s.hnd))
}

func (s *Screen) Planes() int {
	return int(C.XPlanesOfScreen(s.hnd))
}

/*
RootWindowOfScreen (screen)
Window XRootWindowOfScreen (screen)
Screen *screen;
screen Specifies the appropriate Screen structure.
*/

func (d *Display) NoOp() {
	C.XNoOp(d.hnd)
}

type CloseDownMode C.int

const (
	CloseDownDestroyAll      CloseDownMode = C.DestroyAll
	CloseDownRetainPermanent               = C.RetainPermanent
	CloseDownRetainTemporary               = C.RetainTemporary
)

func (d *Display) SetCloseDownMode(mode CloseDownMode) {
	C.XSetCloseDownMode(d.hnd, C.int(mode))
}

func InitThreads() error {
	status := C.XInitThreads()
	if int(status) != 0 {
		return errors.New("could not initialize threads")
	}
	return nil
}

func (d *Display) LockDisplay() {
	C.XLockDisplay(d.hnd)
}

func (d *Display) UnlockDisplay() {
	C.XUnlockDisplay(d.hnd)
}

/*
typedef void (*XConnectionWatchProc) (display, client_data, fd, opening, watch_data)
Display *display;
XPointer client_data;
int fd;
Bool opening;
XPointer *watch_data;
Status XAddConnectionWatch (display, procedure, client_data)
Display *display;
XWatchProc procedure;
XPointer client_data;
display Specifies the connection to the X server.
procedure Specifies the procedure to be called.
client_data Specifies the additional client data.
*/

/*
Status XRemoveConnectionWatch (display, procedure, client_data)
Display *display;
XWatchProc procedure;
XPointer client_data;
display Specifies the connection to the X server.
procedure Specifies the procedure to be called.
client_data Specifies the additional client data.
*/

/*
void XProcessInternalConnection(display, fd)
Display *display;
int fd;
display Specifies the connection to the X server.
fd Specifies the file descriptor.
*/

/*
Status XInternalConnectionNumbers(display, fd_return, count_return)
Display *display;
int **fd_return;
int *count_return;
*/

type VisualId uint32

type Visual struct {
	hnd *C.Visual
}

func (d *Display) DefaultVisual(screenNumber int) *Visual {
	return &Visual{
		hnd: C.XDefaultVisual(d.hnd, C.int(screenNumber)),
	}
}

func (s *Screen) DefaultVisual() *Visual {
	return &Visual{
		hnd: C.XDefaultVisualOfScreen(s.hnd),
	}
}

func (v *Visual) Id() VisualId {
	return VisualId(C.XVisualIDFromVisual(v.hnd))
}

type Window struct {
	hnd C.Window
}

func (w *Window) IsDrawable() bool {
	return true
}

func (w *Window) Handle() C.Drawable {
	return C.Drawable(w.hnd)
}

type Pixel C.ulong

type WindowClass C.int

const (
	WindowClassInputOutput    WindowClass = C.InputOutput
	WindowClassInputOnly                  = C.InputOnly
	WindowClassCopyFromParent             = C.CopyFromParent
)

type WindowAttributes struct {
	hnd      C.XSetWindowAttributes
	mask     uint32
	state    WindowStateAttributes
	hasState bool
}

// Extra attribs not in the XSetWindowAttributes but in
// XWindowAttributes.   (It is /totally ridiculous/ to have these
// as separate structures, guys/gals! =/ )
//
// These are read-only.
type WindowStateAttributes struct {
	x, y, width, height int
	borderWidth, depth  int
	visual              *Visual
	rootWindow          *Window
	mapInstalled        bool
	mapState            MapState
	allEvents, myEvents EventMask
	screen              *Screen
}

type WindowAttributeType uint32

const (
	WindowAttributeTypeCWBackPixmap       WindowAttributeType = C.CWBackPixmap
	WindowAttributeTypeCWBackPixel                            = C.CWBackPixel
	WindowAttributeTypeCWBorderPixmap                         = C.CWBorderPixmap
	WindowAttributeTypeCWBorderPixel                          = C.CWBorderPixel
	WindowAttributeTypeCWBitGravity                           = C.CWBitGravity
	WindowAttributeTypeCWWinGravity                           = C.CWWinGravity
	WindowAttributeTypeCWBackingStore                         = C.CWBackingStore
	WindowAttributeTypeCWBackingPlanes                        = C.CWBackingPlanes
	WindowAttributeTypeCWBackingPixel                         = C.CWBackingPixel
	WindowAttributeTypeCWOverrideRedirect                     = C.CWOverrideRedirect
	WindowAttributeTypeCWSaveUnder                            = C.CWSaveUnder
	WindowAttributeTypeCWEventMask                            = C.CWEventMask
	WindowAttributeTypeCWDontPropagate                        = C.CWDontPropagate
	WindowAttributeTypeCWColormap                             = C.CWColormap
	WindowAttributeTypeCWCursor                               = C.CWCursor
)

type Pixmap C.Pixmap

const (
	PixmapNone           Pixmap = C.None
	PixmapParentRelative        = C.ParentRelative
	PixmapCopyFromParent        = C.CopyFromParent
)

type MapState uint32

const (
	MapStateIsUnmapped   MapState = C.IsUnmapped
	MapStateIsUnviewable          = C.IsUnviewable
	MapStateIsViewable            = C.IsViewable
)

type GravityAttribute uint32

const (
	GravityForget    GravityAttribute = C.ForgetGravity
	GravityNorthWest                  = C.NorthWestGravity
	GravityNorth                      = C.NorthGravity
	GravityNorthEast                  = C.NorthEastGravity
	GravityEast                       = C.EastGravity
	GravitySouthEast                  = C.SouthEastGravity
	GravitySouth                      = C.SouthGravity
	GravitySouthWest                  = C.SouthWestGravity
	GravityWest                       = C.WestGravity
	GravityStatic                     = C.StaticGravity
	GravityUnmap                      = C.UnmapGravity
)

type BitField uint64

func ZerosBitfield() BitField {
	return BitField(0)
}

func OnesBitField() BitField {
	return BitField(math.MaxUint64)
}

func (bf BitField) Test(idx uint) bool {
	return uint64(bf)&(uint64(1)<<idx) != 0
}

func (bf BitField) Set(idx uint) BitField {
	return BitField(uint64(bf) | (uint64(1) << idx))
}

func (bf BitField) Clear(idx uint) BitField {
	return BitField(uint64(bf) & ^(uint64(1) << idx))
}

func (bf BitField) Toggle(idx uint) BitField {
	return BitField(uint64(bf) ^ (uint64(1) << idx))
}

type EventTypeAttribute uint32

const (
	EventAttributeNone                 EventTypeAttribute = C.NoEventMask
	EventAttributeKeyPress                                = C.KeyPressMask
	EventAttributeKeyRelease                              = C.KeyReleaseMask
	EventAttributeButtonPress                             = C.ButtonPressMask
	EventAttributeButtonRelease                           = C.ButtonReleaseMask
	EventAttributeEnterWindow                             = C.EnterWindowMask
	EventAttributeLeaveWindow                             = C.LeaveWindowMask
	EventAttributePointerMotion                           = C.PointerMotionMask
	EventAttributePointerMotionHint                       = C.PointerMotionHintMask
	EventAttributeButton1Motion                           = C.Button1MotionMask
	EventAttributeButton2Motion                           = C.Button2MotionMask
	EventAttributeButton3Motion                           = C.Button3MotionMask
	EventAttributeButton4Motion                           = C.Button4MotionMask
	EventAttributeButton5Motion                           = C.Button5MotionMask
	EventAttributeButtonMotion                            = C.ButtonMotionMask
	EventAttributeKeymapState                             = C.KeymapStateMask
	EventAttributeExposure                                = C.ExposureMask
	EventAttributeVisibilityChange                        = C.VisibilityChangeMask
	EventAttributeStructureNotify                         = C.StructureNotifyMask
	EventAttributeResizeRedirect                          = C.ResizeRedirectMask
	EventAttributeSubstructureNotify                      = C.SubstructureNotifyMask
	EventAttributeSubstructureRedirect                    = C.SubstructureRedirectMask
	EventAttributeFocusChange                             = C.FocusChangeMask
	EventAttributePropertyChange                          = C.PropertyChangeMask
	EventAttributeColormapChange                          = C.ColormapChangeMask
	EventAttributeOwnerGrabButton                         = C.OwnerGrabButtonMask
)

type EventTypeAttributeSet uint64

func (es EventTypeAttributeSet) Add(et EventTypeAttribute) EventTypeAttributeSet {
	return EventTypeAttributeSet(uint64(es) | uint64(et))
}

func (es EventTypeAttributeSet) Delete(et EventTypeAttribute) EventTypeAttributeSet {
	return EventTypeAttributeSet(uint64(es) & ^uint64(et))
}

func (es EventTypeAttributeSet) Test(et EventTypeAttribute) bool {
	return uint64(es)&uint64(et) != 0
}

type Cursor C.Cursor

func NewWindowAttributes() *WindowAttributes {
	return &WindowAttributes{}
}

func (wa *WindowAttributes) IsSet(waType WindowAttributeType) bool {
	return wa.mask&uint32(waType) != 0
}

func (wa *WindowAttributes) SetBackgroundPixmap(p Pixmap) {
	C.SetBackPixmap(&wa.hnd, C.Pixmap(p))
	wa.mask |= uint32(WindowAttributeTypeCWBackPixmap)
}

func (wa *WindowAttributes) GetBackgroundPixmap() (Pixmap, bool) {
	if !wa.IsSet(WindowAttributeTypeCWBackPixmap) {
		return PixmapNone, false
	}
	return Pixmap(C.GetBackPixmap(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearBackgroundPixmap() {
	wa.mask &= ^uint32(WindowAttributeTypeCWBackPixmap)
}

func (wa *WindowAttributes) SetBackgroundPixel(p Pixel) {
	C.SetBackPixel(&wa.hnd, C.ulong(p))
	wa.mask |= uint32(WindowAttributeTypeCWBackingPixel)
}

func (wa *WindowAttributes) GetBackgroundPixel() (Pixel, bool) {
	if !wa.IsSet(WindowAttributeTypeCWBackingPixel) {
		return Pixel(0), false
	}
	return Pixel(C.GetBackPixel(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearBackgroundPixel() {
	wa.mask &= ^uint32(WindowAttributeTypeCWBackingPixel)
}

func (wa *WindowAttributes) SetBorderPixmap(p Pixmap) {
	C.SetBorderPixmap(&wa.hnd, C.Pixmap(p))
	wa.mask |= WindowAttributeTypeCWBorderPixmap
}

func (wa *WindowAttributes) GetBorderPixmap() (Pixmap, bool) {
	if !wa.IsSet(WindowAttributeTypeCWBorderPixmap) {
		return PixmapNone, false
	}
	return Pixmap(C.GetBorderPixmap(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearBorderPixmap() {
	wa.mask &= ^uint32(WindowAttributeTypeCWBorderPixmap)
}

func (wa *WindowAttributes) SetBorderPixel(p Pixel) {
	C.SetBorderPixel(&wa.hnd, C.ulong(p))
	wa.mask |= WindowAttributeTypeCWBorderPixel
}

func (wa *WindowAttributes) GetBorderPixel() (Pixel, bool) {
	if !wa.IsSet(WindowAttributeTypeCWBorderPixel) {
		return Pixel(0), false
	}
	return Pixel(C.GetBorderPixel(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearBorderPixel() {
	wa.mask &= ^uint32(WindowAttributeTypeCWBorderPixel)
}

func (wa *WindowAttributes) SetBitGravity(g GravityAttribute) {
	C.SetBitGravity(&wa.hnd, C.int(g))
	wa.mask |= WindowAttributeTypeCWBitGravity
}

func (wa *WindowAttributes) GetBitGravity() (GravityAttribute, bool) {
	if !wa.IsSet(WindowAttributeTypeCWBitGravity) {
		return GravityStatic, false
	}
	return GravityAttribute(C.GetBitGravity(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearBitGravity() {
	wa.mask &= ^uint32(WindowAttributeTypeCWBitGravity)
}

func (wa *WindowAttributes) SetWindowGravity(g GravityAttribute) {
	C.SetWinGravity(&wa.hnd, C.int(g))
	wa.mask |= WindowAttributeTypeCWWinGravity
}

func (wa *WindowAttributes) GetWindowGravity() (GravityAttribute, bool) {
	if !wa.IsSet(WindowAttributeTypeCWWinGravity) {
		return GravityStatic, false
	}
	return GravityAttribute(C.GetWinGravity(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearWindowGravity() {
	wa.mask &= ^uint32(WindowAttributeTypeCWWinGravity)
}

func (wa *WindowAttributes) SetBackingStoreMode(bsm BackingStoreMode) {
	C.SetBackingStore(&wa.hnd, C.int(bsm))
	wa.mask |= WindowAttributeTypeCWBackingStore
}

func (wa *WindowAttributes) GetBackingStoreMode() (BackingStoreMode, bool) {
	if !wa.IsSet(WindowAttributeTypeCWBackingStore) {
		return BackingStoreNotUseful, false
	}
	return BackingStoreMode(C.GetBackingStore(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearBackingStoreMode() {
	wa.mask &= ^uint32(WindowAttributeTypeCWBackingStore)
}

func (wa *WindowAttributes) SetBackingPlanes(planes BitField) {
	C.SetBackingPlanes(&wa.hnd, C.ulong(planes))
	wa.mask |= WindowAttributeTypeCWBackingPlanes
}

func (wa *WindowAttributes) GetBackingPlanes() (BitField, bool) {
	if !wa.IsSet(WindowAttributeTypeCWBackingPlanes) {
		return OnesBitField(), false
	}
	return BitField(C.GetBackingPlanes(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearBackingPlanes() {
	wa.mask &= ^uint32(WindowAttributeTypeCWBackingPlanes)
}

func (wa *WindowAttributes) SetBackingPixel(p Pixel) {
	C.SetBackingPixel(&wa.hnd, C.ulong(p))
	wa.mask |= WindowAttributeTypeCWBackPixel
}

func (wa *WindowAttributes) GetBackingPixel() (Pixel, bool) {
	if !wa.IsSet(WindowAttributeTypeCWBackingPixel) {
		return Pixel(0), false
	}
	return Pixel(C.GetBackingPixel(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearBackingPixel() {
	wa.mask &= ^uint32(WindowAttributeTypeCWBackingPixel)
}

func (wa *WindowAttributes) SetSaveUnder(val bool) {
	b := 0
	if val {
		b = 1
	}
	C.SetSaveUnder(&wa.hnd, C.Bool(b))
	wa.mask |= WindowAttributeTypeCWSaveUnder
}

func (wa *WindowAttributes) GetSaveUnder() (bool, bool) {
	if !wa.IsSet(WindowAttributeTypeCWSaveUnder) {
		return false, false
	}
	return C.GetSaveUnder(&wa.hnd) > 0, true
}

func (wa *WindowAttributes) ClearSaveUnder() {
	wa.mask &= ^uint32(WindowAttributeTypeCWSaveUnder)
}

func (wa *WindowAttributes) SetSaveEventMask(mask EventTypeAttributeSet) {
	C.SetSaveEventMask(&wa.hnd, C.int(mask))
	wa.mask |= WindowAttributeTypeCWEventMask
}

func (wa *WindowAttributes) GetSaveEventMask() (EventTypeAttributeSet, bool) {
	if !wa.IsSet(WindowAttributeTypeCWEventMask) {
		return EventTypeAttributeSet(0), false
	}
	return EventTypeAttributeSet(C.GetSaveEventMask(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearSaveEventMask() {
	wa.mask &= ^uint32(WindowAttributeTypeCWEventMask)
}

func (wa *WindowAttributes) SetDoNotPropagateEventMask(mask EventTypeAttributeSet) {
	C.SetDoNotPropagateEventMask(&wa.hnd, C.int(mask))
	wa.mask |= WindowAttributeTypeCWDontPropagate
}

func (wa *WindowAttributes) GetDoNotPropagateEventMask() (EventTypeAttributeSet, bool) {
	if !wa.IsSet(WindowAttributeTypeCWDontPropagate) {
		return EventTypeAttributeSet(0), false
	}
	return EventTypeAttributeSet(C.GetDoNotPropagateEventMask(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearDoNotPropagateEventMask() {
	wa.mask &= ^uint32(WindowAttributeTypeCWDontPropagate)
}

func (wa *WindowAttributes) SetOverrideRedirect(val bool) {
	b := 0
	if val {
		b = 1
	}
	C.SetOverrideRedirect(&wa.hnd, C.Bool(b))
	wa.mask |= WindowAttributeTypeCWOverrideRedirect
}

func (wa *WindowAttributes) GetOverrideRedirect() (bool, bool) {
	if !wa.IsSet(WindowAttributeTypeCWOverrideRedirect) {
		return false, false
	}
	return C.GetOverrideRedirect(&wa.hnd) > 0, true
}

func (wa *WindowAttributes) ClearOverrideRedirect() {
	wa.mask &= ^uint32(WindowAttributeTypeCWOverrideRedirect)
}

func (wa *WindowAttributes) SetColormap(cmap Colormap) {
	C.SetColormap(&wa.hnd, C.Colormap(cmap))
	wa.mask |= WindowAttributeTypeCWColormap
}

func (wa *WindowAttributes) GetColormap() (Colormap, bool) {
	if !wa.IsSet(WindowAttributeTypeCWColormap) {
		return Colormap(0), false
	}
	return Colormap(C.GetColormap(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearColormap() {
	wa.mask &= ^uint32(WindowAttributeTypeCWColormap)
}

func (wa *WindowAttributes) SetCursor(c Cursor) {
	C.SetCursor(&wa.hnd, C.Cursor(c))
	wa.mask |= WindowAttributeTypeCWCursor
}

func (wa *WindowAttributes) GetCursor() (Cursor, bool) {
	if !wa.IsSet(WindowAttributeTypeCWCursor) {
		return Cursor(0), false
	}
	return Cursor(C.GetCursor(&wa.hnd)), true
}

func (wa *WindowAttributes) ClearCursor() {
	wa.mask &= ^uint32(WindowAttributeTypeCWCursor)
}

func (wa *WindowAttributes) GetX() (int, bool) {
	if !wa.hasState {
		return 0, false
	}
	return wa.state.x, true
}

func (wa *WindowAttributes) GetY() (int, bool) {
	if !wa.hasState {
		return 0, false
	}
	return wa.state.y, true
}

func (wa *WindowAttributes) GetWidth() (int, bool) {
	if !wa.hasState {
		return 0, false
	}
	return wa.state.width, true
}

func (wa *WindowAttributes) GetHeight() (int, bool) {
	if !wa.hasState {
		return 0, false
	}
	return wa.state.height, true
}

func (wa *WindowAttributes) GetBorderWidth() (int, bool) {
	if !wa.hasState {
		return 0, false
	}
	return wa.state.borderWidth, true
}

func (wa *WindowAttributes) GetDepth() (int, bool) {
	if !wa.hasState {
		return 0, false
	}
	return wa.state.depth, true
}

func (wa *WindowAttributes) GetVisual() (*Visual, bool) {
	if !wa.hasState {
		return nil, false
	}
	return wa.state.visual, true
}

func (wa *WindowAttributes) GetRootWindow() (*Window, bool) {
	if !wa.hasState {
		return nil, false
	}
	return wa.state.rootWindow, true
}

func (wa *WindowAttributes) GetMapInstalled() (bool, bool) {
	if !wa.hasState {
		return false, false
	}
	return wa.state.mapInstalled, true
}

func (wa *WindowAttributes) GetMapState() (MapState, bool) {
	if !wa.hasState {
		return 0, false
	}
	return wa.state.mapState, true
}

func (wa *WindowAttributes) GetAllEvents() (EventMask, bool) {
	if !wa.hasState {
		return 0, false
	}
	return wa.state.allEvents, true
}

func (wa *WindowAttributes) GetMyEvents() (EventMask, bool) {
	if !wa.hasState {
		return 0, false
	}
	return wa.state.myEvents, true
}

func (wa *WindowAttributes) GetScreen() (*Screen, bool) {
	if !wa.hasState {
		return nil, false
	}
	return wa.state.screen, true
}

func (d *Display) CreateWindow(parent *Window, x, y, width, height int, border_width int,
	depth int, class WindowClass, visual *Visual, attribs WindowAttributes) *Window {
	if parent == nil {
		parent = &Window{hnd: 0}
	}
	if visual == nil {
		visual = &Visual{hnd: nil}
	}
	hnd := C.XCreateWindow(d.hnd, parent.hnd, C.int(x), C.int(y), C.uint(width), C.uint(height), C.uint(border_width),
		C.int(depth), C.uint(class), visual.hnd, C.ulong(attribs.mask), &attribs.hnd)
	win := &Window{hnd: hnd}
	return win
}

func (d *Display) CreateSimpleWindow(parent *Window, x, y, width, height int, border_width int,
	border Pixel, background Pixel) *Window {
	if parent == nil {
		parent = &Window{hnd: 0}
	}
	hnd := C.XCreateSimpleWindow(d.hnd, parent.hnd, C.int(x), C.int(y), C.uint(width), C.uint(height),
		C.uint(border_width), C.ulong(border), C.ulong(background))
	win := &Window{hnd: hnd}
	return win
}

func (d *Display) DestroyWindow(win *Window) {
	C.XDestroyWindow(d.hnd, win.hnd)
}

func (d *Display) DestroySubwindows(win *Window) {
	C.XDestroySubwindows(d.hnd, win.hnd)
}

func (d *Display) MapWindow(win *Window) {
	C.XMapWindow(d.hnd, win.hnd)
}

func (d *Display) MapRaised(win *Window) {
	C.XMapRaised(d.hnd, win.hnd)
}

func (d *Display) MapSubwindows(win *Window) {
	C.XMapSubwindows(d.hnd, win.hnd)
}

func (d *Display) UnmapWindow(win *Window) {
	C.XUnmapWindow(d.hnd, win.hnd)
}

func (d *Display) UnmapSubwindows(win *Window) {
	C.XUnmapSubwindows(d.hnd, win.hnd)
}

type WindowChanges struct {
	hnd     C.XWindowChanges
	sibling *Window
	mask    int
}

type WindowChangesType C.int

const (
	WindowChangesTypeCWX           WindowChangesType = C.CWX
	WindowChangesTypeCWY                             = C.CWY
	WindowChangesTypeCWWidth                         = C.CWWidth
	WindowChangesTypeCWHeight                        = C.CWHeight
	WindowChangesTypeCWBorderWidth                   = C.CWBorderWidth
	WindowChangesTypeCWSibling                       = C.CWSibling
	WindowChangesTypeCWStackMode                     = C.CWStackMode
)

type StackMode C.int

const (
	StackModeAbove    StackMode = C.Above
	StackModeBelow              = C.Below
	StackModeTopIf              = C.TopIf
	StackModeBottomIf           = C.BottomIf
	StackModeOpposite           = C.Opposite
)

func (wc *WindowChanges) IsSet(wcType WindowChangesType) bool {
	return wc.mask&int(wcType) != 0
}

func (wc *WindowChanges) SetX(x int) {
	C.WindowChangesSetX(&wc.hnd, C.int(x))
	wc.mask |= int(WindowChangesTypeCWX)
}

func (wc *WindowChanges) GetX() (int, bool) {
	if !wc.IsSet(WindowChangesTypeCWX) {
		return 0, false
	}
	return int(C.WindowChangesGetX(&wc.hnd)), true
}

func (wc *WindowChanges) ClearX() {
	wc.mask &= ^int(WindowChangesTypeCWX)
}

func (wc *WindowChanges) SetY(y int) {
	C.WindowChangesSetY(&wc.hnd, C.int(y))
	wc.mask |= WindowChangesTypeCWY
}

func (wc *WindowChanges) GetY() (int, bool) {
	if !wc.IsSet(WindowChangesTypeCWY) {
		return 0, false
	}
	return int(C.WindowChangesGetY(&wc.hnd)), true
}

func (wc *WindowChanges) ClearY() {
	wc.mask &= ^WindowChangesTypeCWY
}

func (wc *WindowChanges) SetWidth(w int) {
	C.WindowChangesSetWidth(&wc.hnd, C.int(w))
	wc.mask |= WindowChangesTypeCWWidth
}

func (wc *WindowChanges) GetWidth() (int, bool) {
	if !wc.IsSet(WindowChangesTypeCWWidth) {
		return 0, false
	}
	return int(C.WindowChangesGetWidth(&wc.hnd)), true
}

func (wc *WindowChanges) ClearWidth() {
	wc.mask &= ^WindowChangesTypeCWWidth
}

func (wc *WindowChanges) SetHeight(h int) {
	C.WindowChangesSetHeight(&wc.hnd, C.int(h))
	wc.mask |= WindowChangesTypeCWHeight
}

func (wc *WindowChanges) GetHeight() (int, bool) {
	if !wc.IsSet(WindowChangesTypeCWHeight) {
		return 0, false
	}
	return int(C.WindowChangesGetHeight(&wc.hnd)), true
}

func (wc *WindowChanges) ClearHeight() {
	wc.mask &= ^WindowChangesTypeCWHeight
}

func (wc *WindowChanges) SetBorderWidth(width int) {
	C.WindowChangesSetBorderWidth(&wc.hnd, C.int(width))
	wc.mask |= WindowChangesTypeCWBorderWidth
}

func (wc *WindowChanges) GetBorderWidth() (int, bool) {
	if !wc.IsSet(WindowChangesTypeCWBorderWidth) {
		return 0, false
	}
	return int(C.WindowChangesGetBorderWidth(&wc.hnd)), true
}

func (wc *WindowChanges) ClearBorderWidth() {
	wc.mask &= ^WindowChangesTypeCWBorderWidth
}

func (wc *WindowChanges) SetSibling(w *Window) {
	C.WindowChangesSetSibling(&wc.hnd, w.hnd)
	wc.sibling = w
	wc.mask |= WindowChangesTypeCWSibling
}

func (wc *WindowChanges) GetSibling() (*Window, bool) {
	if !wc.IsSet(WindowChangesTypeCWSibling) {
		return nil, false
	}
	return wc.sibling, true
}

func (wc *WindowChanges) ClearSibling() {
	wc.mask &= ^WindowChangesTypeCWSibling
}

func (wc *WindowChanges) SetStackMode(sm StackMode) {
	C.WindowChangesSetStackMode(&wc.hnd, C.int(sm))
	wc.mask |= WindowChangesTypeCWStackMode
}

func (wc *WindowChanges) GetStackMode() (StackMode, bool) {
	if !wc.IsSet(WindowChangesTypeCWStackMode) {
		return StackMode(0), false
	}
	return StackMode(C.WindowChangesGetStackMode(&wc.hnd)), true
}

func (wc *WindowChanges) ClearStackMode() {
	wc.mask &= ^WindowChangesTypeCWStackMode
}

func (d *Display) ConfigureWindow(w *Window, wc WindowChanges) {
	C.XConfigureWindow(d.hnd, w.hnd, C.uint(wc.mask), &wc.hnd)
}

func (d *Display) MoveWindow(w *Window, x, y int) {
	C.XMoveWindow(d.hnd, w.hnd, C.int(x), C.int(y))
}

func (d *Display) ResizeWindow(w *Window, width, height int) {
	C.XResizeWindow(d.hnd, w.hnd, C.uint(width), C.uint(height))
}

func (d *Display) MoveResizeWindow(w *Window, x, y, width, height int) {
	C.XMoveResizeWindow(d.hnd, w.hnd, C.int(x), C.int(y), C.uint(width), C.uint(height))
}

func (d *Display) SetWindowBorderWidth(w *Window, width int) {
	C.XSetWindowBorderWidth(d.hnd, w.hnd, C.uint(width))
}

func (d *Display) RaiseWindow(win *Window) {
	C.XRaiseWindow(d.hnd, win.hnd)
}

func (d *Display) LowerWindow(win *Window) {
	C.XLowerWindow(d.hnd, win.hnd)
}

type CirculateDirection C.int

const (
	CirculateDirectionRaiseLowest  CirculateDirection = C.RaiseLowest
	CirculateDirectionLowerHighest                    = C.LowerHighest
)

func (d *Display) CirculateSubwindows(w *Window, dir CirculateDirection) {
	C.XCirculateSubwindows(d.hnd, w.hnd, C.int(dir))
}

func (d *Display) CirculateSubwindowsUp(w *Window) {
	C.XCirculateSubwindowsUp(d.hnd, w.hnd)
}

func (d *Display) CirculateSubwindowsDown(w *Window) {
	C.XCirculateSubwindowsDown(d.hnd, w.hnd)
}

func (d *Display) RestackWindows(windows []*Window) {
	stack := make([]C.Window, len(windows))
	for i, w := range windows {
		stack[i] = w.hnd
	}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&stack))
	C.XRestackWindows(d.hnd, (*C.Window)(unsafe.Pointer(hdr.Data)), C.int(hdr.Len))
}

func (d *Display) ChangeWindowAttributes(w *Window, attrs *WindowAttributes) {
	C.XChangeWindowAttributes(d.hnd, w.hnd, C.ulong(attrs.mask), &attrs.hnd)
}

func (d *Display) SetWindowBackground(w *Window, p Pixel) {
	C.XSetWindowBackground(d.hnd, w.hnd, C.ulong(p))
}

func (d *Display) SetWindowBackgroundPixmap(w *Window, p Pixmap) {
	C.XSetWindowBackgroundPixmap(d.hnd, w.hnd, C.Pixmap(p))
}

func (d *Display) SetWindowBorder(w *Window, p Pixel) {
	C.XSetWindowBorder(d.hnd, w.hnd, C.ulong(p))
}

func (d *Display) SetWindowBorderPixmap(w *Window, p Pixmap) {
	C.XSetWindowBorderPixmap(d.hnd, w.hnd, C.Pixmap(p))
}

func (d *Display) SetWindowColormap(w *Window, cm Colormap) {
	C.XSetWindowColormap(d.hnd, w.hnd, C.Colormap(cm))
}

func (d *Display) DefineCursor(w *Window, c Cursor) {
	C.XDefineCursor(d.hnd, w.hnd, C.Cursor(c))
}

func (d *Display) UndefineCursor(w *Window) {
	C.XUndefineCursor(d.hnd, w.hnd)
}

func (d *Display) QueryTree(w *Window) (root *Window, parent *Window, children []*Window, e error) {
	var croot, cparent C.Window
	var cchildren *C.Window
	var count C.uint
	status := C.XQueryTree(d.hnd, w.hnd, &croot, &cparent, &cchildren, &count)
	if status == 0 {
		return nil, nil, nil, errors.New("could not get query tree info for window")
	}
	rw := &Window{hnd: croot}
	pw := &Window{hnd: cparent}
	if int(count) == 0 {
		return rw, pw, []*Window{}, nil
	}
	cw := make([]*Window, int(count))
	for i := 0; i < int(count); i++ {
		cw[i] = &Window{hnd: C.WindowPtrIndex(cchildren, C.int(i))}
	}
	return rw, pw, cw, nil
}

func (d *Display) GetWindowAttributes(w *Window) *WindowAttributes {
	nwa := &WindowAttributes{}
	var xwa C.XWindowAttributes
	C.XGetWindowAttributes(d.hnd, w.hnd, &xwa)
	C.SetWindowAttributesFillFromWindowAttributes(&nwa.hnd, &xwa)
	var cVisual *C.Visual
	var cRootWin C.Window
	var cMapInstalled C.Bool
	var cMapState C.int
	var cAllEvents, cMyEvents C.ulong
	var cScreen *C.Screen
	var x, y C.int
	var width, height, borderWidth, depth C.uint
	C.SetWindowAttributesRead(&xwa, &x, &y, &width, &height,
		&borderWidth, &depth,
		&cVisual, &cRootWin, &cMapInstalled, &cMapState,
		&cAllEvents, &cMyEvents, &cScreen)
	nwa.state.x = int(x)
	nwa.state.y = int(y)
	nwa.state.width = int(width)
	nwa.state.height = int(height)
	nwa.state.borderWidth = int(borderWidth)
	nwa.state.depth = int(depth)
	nwa.state.visual = &Visual{hnd: cVisual}
	nwa.state.rootWindow = &Window{hnd: cRootWin}
	nwa.state.mapInstalled = cMapInstalled > 0
	nwa.state.mapState = MapState(cMapState)
	nwa.state.allEvents = EventMask(cAllEvents)
	nwa.state.myEvents = EventMask(cMyEvents)
	nwa.state.screen = &Screen{hnd: cScreen}
	return nwa
}

func (d *Display) GetGeometry(dd Drawable) (rootWindow *Window, x int, y int, width int,
	height int, borderWidth int, depth int, e error) {
	var cRoot C.Window
	var cX, cY C.int
	var cWidth, cHeight C.uint
	var cBorderWidth, cDepth C.uint
	C.XGetGeometry(d.hnd, dd.Handle(), &cRoot, &cX, &cY, &cWidth, &cHeight, &cBorderWidth, &cDepth)
	return &Window{hnd: cRoot}, int(cX), int(cY), int(cWidth), int(cHeight), int(cBorderWidth), int(cDepth), nil
}

func (d *Display) TranslateCoordinates(src *Window, dst *Window, sx, sy int) (dx int, dy int, child *Window, r bool) {
	var cBool C.Bool
	var cX, cY C.int
	var cChild C.Window
	cBool = C.XTranslateCoordinates(d.hnd, src.hnd, dst.hnd, C.int(sx), C.int(sy), &cX, &cY, &cChild)
	return int(cX), int(cY), &Window{hnd: cChild}, cBool > 0
}

type ModifierMask C.uint

const (
// XXX - Fill these in!
)

func (d *Display) QueryPointer(w *Window) (root *Window, child *Window, rx int, ry int, wx int, wy int, r bool) {
	var cWx, cWy, cX, cY C.int
	var cRoot, cChild C.Window
	var cMask C.uint
	cBool := C.XQueryPointer(d.hnd, w.hnd, &cRoot, &cChild, &cX, &cY, &cWx, &cWy, &cMask)
	return &Window{hnd: cRoot}, &Window{hnd: cChild}, int(cX), int(cY), int(cWx), int(cWy), cBool > 0
}

type Atom C.Atom

type SelectionPropertyName string

func (spn SelectionPropertyName) IsName() bool   { return true }
func (spn SelectionPropertyName) String() string { return string(spn) }

const (
	SelectionPrimary   SelectionPropertyName = "PRIMARY"
	SelectionSecondary SelectionPropertyName = "SECONDARY"
)

func lookupName(name string) Name {
	switch name {
	case SelectionPrimary.String():
		return SelectionPrimary
	case SelectionSecondary.String():
		return SelectionSecondary
	case PropertyCutBuffer0.String():
		return PropertyCutBuffer0
	case PropertyCutBuffer1.String():
		return PropertyCutBuffer1
	case PropertyCutBuffer2.String():
		return PropertyCutBuffer2
	case PropertyCutBuffer3.String():
		return PropertyCutBuffer3
	case PropertyCutBuffer4.String():
		return PropertyCutBuffer4
	case PropertyCutBuffer5.String():
		return PropertyCutBuffer5
	case PropertyCutBuffer6.String():
		return PropertyCutBuffer6
	case PropertyCutBuffer7.String():
		return PropertyCutBuffer7
	case PropertyRgbBestMap.String():
		return PropertyRgbBestMap
	case PropertyRgbBlueMap.String():
		return PropertyRgbBlueMap
	case PropertyRgbDefaultMap.String():
		return PropertyRgbDefaultMap
	case PropertyRgbGrayMap.String():
		return PropertyRgbGrayMap
	case PropertyRgbGreenMap.String():
		return PropertyRgbGreenMap
	case PropertyRgbRedMap.String():
		return PropertyRgbRedMap
	case PropertyResourceManager.String():
		return PropertyResourceManager
	case PropertyWmClass.String():
		return PropertyWmClass
	case PropertyWmClientMachine.String():
		return PropertyWmClientMachine
	case PropertyWmColormapWindows.String():
		return PropertyWmColormapWindows
	case PropertyWmCommand.String():
		return PropertyWmCommand
	case PropertyWmHints.String():
		return PropertyWmHints
	case PropertyWmIconName.String():
		return PropertyWmIconName
	case PropertyWmIconSize.String():
		return PropertyWmIconSize
	case PropertyWmName.String():
		return PropertyWmName
	case PropertyWmNormalHints.String():
		return PropertyWmNormalHints
	case PropertyWmProtocols.String():
		return PropertyWmProtocols
	case PropertyWmState.String():
		return PropertyWmState
	case PropertyWmTransientFor.String():
		return PropertyWmTransientFor
	case PropertyWmZoomHints.String():
		return PropertyWmZoomHints
	case FontPropertyMinSpace.String():
		return FontPropertyMinSpace
	case FontPropertyNormSpace.String():
		return FontPropertyNormSpace
	case FontPropertyMaxSpace.String():
		return FontPropertyMaxSpace
	case FontPropertyEndSpace.String():
		return FontPropertyEndSpace
	case FontPropertySuperscriptX.String():
		return FontPropertySuperscriptX
	case FontPropertySuperscriptY.String():
		return FontPropertySuperscriptY
	case FontPropertySubscriptX.String():
		return FontPropertySubscriptX
	case FontPropertySubscriptY.String():
		return FontPropertySubscriptY
	case FontPropertyUnderlinePosition.String():
		return FontPropertyUnderlinePosition
	case FontPropertyUnderlineThickness.String():
		return FontPropertyUnderlineThickness
	case FontPropertyFontName.String():
		return FontPropertyFontName
	case FontPropertyFullName.String():
		return FontPropertyFullName
	case FontPropertyStrikeoutDescent.String():
		return FontPropertyStrikeoutDescent
	case FontPropertyStrikeoutAscent.String():
		return FontPropertyStrikeoutAscent
	case FontPropertyItalicAngle.String():
		return FontPropertyItalicAngle
	case FontPropertyXHeight.String():
		return FontPropertyXHeight
	case FontPropertyQuadWidth.String():
		return FontPropertyQuadWidth
	case FontPropertyWeight.String():
		return FontPropertyWeight
	case FontPropertyPointSize.String():
		return FontPropertyPointSize
	case FontPropertyResolution.String():
		return FontPropertyResolution
	case FontPropertyCopyright.String():
		return FontPropertyCopyright
	case FontPropertyNotice.String():
		return FontPropertyNotice
	case FontPropertyFamilyName.String():
		return FontPropertyFamilyName
	case FontPropertyCapHeight.String():
		return FontPropertyCapHeight
	default:
		{
			return &UserDefinedName{name: name}
		}
	}
}

type PropertyName string

const (
	PropertyCutBuffer0        PropertyName = "CUT_BUFFER0"
	PropertyCutBuffer1        PropertyName = "CUT_BUFFER1"
	PropertyCutBuffer2        PropertyName = "CUT_BUFFER2"
	PropertyCutBuffer3        PropertyName = "CUT_BUFFER3"
	PropertyCutBuffer4        PropertyName = "CUT_BUFFER4"
	PropertyCutBuffer5        PropertyName = "CUT_BUFFER5"
	PropertyCutBuffer6        PropertyName = "CUT_BUFFER6"
	PropertyCutBuffer7        PropertyName = "CUT_BUFFER7"
	PropertyRgbBestMap        PropertyName = "RGB_BEST_MAP"
	PropertyRgbBlueMap        PropertyName = "RGB_BLUE_MAP"
	PropertyRgbDefaultMap     PropertyName = "RGB_DEFAULT_MAP"
	PropertyRgbGrayMap        PropertyName = "RGB_GRAY_MAP"
	PropertyRgbGreenMap       PropertyName = "RGB_GREEN_MAP"
	PropertyRgbRedMap         PropertyName = "RGB_RED_MAP"
	PropertyResourceManager   PropertyName = "RESOURCE_MANAGER"
	PropertyWmClass           PropertyName = "WM_CLASS"
	PropertyWmClientMachine   PropertyName = "WM_CLIENT_MACHINE"
	PropertyWmColormapWindows PropertyName = "WM_COLORMAP_WINDOWS"
	PropertyWmCommand         PropertyName = "WM_COMMAND"
	PropertyWmHints           PropertyName = "WM_HINTS"
	PropertyWmIconName        PropertyName = "WM_ICON_NAME"
	PropertyWmIconSize        PropertyName = "WM_ICON_SIZE"
	PropertyWmName            PropertyName = "WM_NAME"
	PropertyWmNormalHints     PropertyName = "WM_NORMAL_HINTS"
	PropertyWmProtocols       PropertyName = "WM_PROTOCOLS"
	PropertyWmState           PropertyName = "WM_STATE"
	PropertyWmTransientFor    PropertyName = "WM_TRANSIENT_FOR"
	PropertyWmZoomHints       PropertyName = "WM_ZOOM_HINTS"
)

func (pn PropertyName) IsName() bool   { return true }
func (pn PropertyName) String() string { return string(pn) }

type PropertyTypeName string

const (
	PropertyTypeArc         PropertyTypeName = "ARC"
	PropertyTypeAtom        PropertyTypeName = "ATOM"
	PropertyTypeBitmap      PropertyTypeName = "BITMAP"
	PropertyTypeCardinal    PropertyTypeName = "CARDINAL"
	PropertyTypeColormap    PropertyTypeName = "COLORMAP"
	PropertyTypeCursor      PropertyTypeName = "CURSOR"
	PropertyTypeDrawable    PropertyTypeName = "DRAWABLE"
	PropertyTypeFont        PropertyTypeName = "FONT"
	PropertyTypeInteger     PropertyTypeName = "INTEGER"
	PropertyTypePIxmap      PropertyTypeName = "PIXMAP"
	PropertyTypePoint       PropertyTypeName = "POINT"
	PropertyTypeRgbColorMap PropertyTypeName = "RGB_COLOR_MAP"
	PropertyTypeRectangle   PropertyTypeName = "RECTANGLE"
	PropertyTypeString      PropertyTypeName = "STRING"
	PropertyTypeVisualId    PropertyTypeName = "VISUALID"
	PropertyTypeWindow      PropertyTypeName = "WINDOW"
	PropertyTypeWmHints     PropertyTypeName = "WM_HINTS"
	PropertyTypeWmSizeHints PropertyTypeName = "WM_SIZE_HINTS"
)

func (ptn PropertyTypeName) IsName() bool   { return true }
func (ptn PropertyTypeName) String() string { return string(ptn) }

type FontPropertyName string

const (
	FontPropertyMinSpace           FontPropertyName = "MIN_SPACE"
	FontPropertyNormSpace          FontPropertyName = "NORM_SPACE"
	FontPropertyMaxSpace           FontPropertyName = "MAX_SPACE"
	FontPropertyEndSpace           FontPropertyName = "END_SPACE"
	FontPropertySuperscriptX       FontPropertyName = "SUPERSCRIPT_X"
	FontPropertySuperscriptY       FontPropertyName = "SUPERSCRIPT_Y"
	FontPropertySubscriptX         FontPropertyName = "SUBSCRIPT_X"
	FontPropertySubscriptY         FontPropertyName = "SUBSCRIPT_Y"
	FontPropertyUnderlinePosition  FontPropertyName = "UNDERLINE_POSITION"
	FontPropertyUnderlineThickness FontPropertyName = "UNDERLINE_THICKNESS"
	FontPropertyFontName           FontPropertyName = "FONT_NAME"
	FontPropertyFullName           FontPropertyName = "FULL_NAME"
	FontPropertyStrikeoutDescent   FontPropertyName = "STRIKEOUT_DESCENT"
	FontPropertyStrikeoutAscent    FontPropertyName = "STRIKEOUT_ASCENT"
	FontPropertyItalicAngle        FontPropertyName = "ITALIC_ANGLE"
	FontPropertyXHeight            FontPropertyName = "X_HEIGHT"
	FontPropertyQuadWidth          FontPropertyName = "QUAD_WIDTH"
	FontPropertyWeight             FontPropertyName = "WEIGHT"
	FontPropertyPointSize          FontPropertyName = "POINT_SIZE"
	FontPropertyResolution         FontPropertyName = "RESOLUTION"
	FontPropertyCopyright          FontPropertyName = "COPYRIGHT"
	FontPropertyNotice             FontPropertyName = "NOTICE"
	FontPropertyFamilyName         FontPropertyName = "FAMILY_NAME"
	FontPropertyCapHeight          FontPropertyName = "CAP_HEIGHT"
)

func (fpn FontPropertyName) IsName() bool   { return true }
func (fpn FontPropertyName) String() string { return string(fpn) }

type UserDefinedName struct {
	name string
}

func (n *UserDefinedName) IsName() bool   { return true }
func (n *UserDefinedName) String() string { return n.name }

func NewUserDefinedName(name string) *UserDefinedName {
	return &UserDefinedName{name: name}
}

type Name interface {
	IsName() bool
	String() string
}

func (d *Display) InternAtom(name Name, onlyIfExists bool) Atom {
	var b C.Bool
	if onlyIfExists {
		b = 1
	} else {
		b = 0
	}
	return Atom(C.XInternAtom(d.hnd, C.CString(name.String()), b))
}

func (d *Display) InternAtoms(name []Name, onlyIfExists bool) ([]Atom, error) {
	var b C.Bool
	if onlyIfExists {
		b = 1
	} else {
		b = 0
	}
	src := make([]*C.char, len(name))
	for i, n := range name {
		src[i] = C.CString(n.String())
	}
	ret := make([]Atom, len(name))
	inHdr := *(*reflect.SliceHeader)(unsafe.Pointer(&src))
	outHdr := *(*reflect.SliceHeader)(unsafe.Pointer(&ret))
	status := C.XInternAtoms(d.hnd, unsafe.Pointer(inHdr.Data), C.int(inHdr.Len), b, (*C.Atom)(unsafe.Pointer(outHdr.Data)))
	if status == 0 {
		return ret, errors.New("one or more atoms could not be interned")
	}
	return ret, nil
}

func (d *Display) GetAtomName(a Atom) Name {
	str := C.XGetAtomName(d.hnd, C.Atom(a))
	if C.CharPtrIsNull(str) > 0 {
		return nil
	}
	gstr := C.GoString(str)
	C.XFree(unsafe.Pointer(str))
	return lookupName(gstr)
}

func (d Display) GetAtomNames(atoms []Atom) ([]Name, error) {
	var dst = make([]*C.char, len(atoms))
	inHdr := *(*reflect.SliceHeader)(unsafe.Pointer(&atoms))
	outHdr := *(*reflect.SliceHeader)(unsafe.Pointer(&dst))
	status := C.XGetAtomNames(d.hnd, (*C.Atom)(unsafe.Pointer(inHdr.Data)), C.int(inHdr.Len), (**C.char)(unsafe.Pointer(outHdr.Data)))
	ret := make([]Name, len(atoms))
	for i, cn := range dst {
		if C.CharPtrIsNull(cn) > 0 {
			ret[i] = nil
		} else {
			ret[i] = lookupName(C.GoString(cn))
			C.XFree(unsafe.Pointer(cn))
		}
	}
	if status == 0 {
		return ret, errors.New("one or more names could not be looked up")
	}
	return ret, nil
}

type PropertyDataFormat int

const (
	PropertyDataFormat8  PropertyDataFormat = 8
	PropertyDataFormat16 PropertyDataFormat = 16
	PropertyDataFormat32 PropertyDataFormat = 32
)

type PropertyData struct {
	format PropertyDataFormat
	recLen int
	cData  *C.uchar
	goData []byte
}

func destroyPropertyData(pd *PropertyData) {
	C.XFree(unsafe.Pointer(pd.cData))
}

func (pd *PropertyData) Format() PropertyDataFormat {
	return pd.format
}

func (pd *PropertyData) Data() []byte {
	if pd.goData == nil {
		pd.goData = make([]byte, int(pd.format)*pd.recLen)
		hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&pd.goData))
		C.memcpy(unsafe.Pointer(hdr.Data), unsafe.Pointer(pd.cData), C.size_t(hdr.Len))
	}
	return pd.goData
}

func NewPropertyData(format PropertyDataFormat, recLen int) *PropertyData {
	return &PropertyData{
		format: format,
		recLen: recLen,
		goData: make([]byte, int(format)*recLen),
	}
}

func (d *Display) GetWindowProperty(w *Window, prop Atom, wordOffset int, wordLength int,
	deleteProp bool, propType Atom) (rType Atom, tailBytes uint32, data *PropertyData, e error) {
	var deleteBool C.Bool
	if deleteProp {
		deleteBool = 1
	} else {
		deleteBool = 0
	}
	var cActualType C.Atom
	var cActualFormat C.int
	var cRecordLen C.ulong
	var cExtraBytes C.ulong
	var cData *C.uchar
	s := C.XGetWindowProperty(d.hnd, w.hnd, C.Atom(prop), C.long(wordOffset), C.long(wordLength),
		deleteBool, C.Atom(propType), &cActualType, &cActualFormat,
		&cRecordLen, &cExtraBytes, &cData)
	if s != C.Success {
		return Atom(0), 0, nil, errors.New("property lookup failed")
	}
	pd := &PropertyData{
		format: PropertyDataFormat(int(cActualFormat)),
		recLen: int(cRecordLen),
		cData:  cData,
	}
	runtime.SetFinalizer(pd, destroyPropertyData)
	return Atom(cActualType), uint32(cExtraBytes), pd, nil
}

func (d *Display) ListProperties(w *Window) []Atom {
	var cNumProps C.int
	cAtoms := C.XListProperties(d.hnd, w.hnd, &cNumProps)
	ret := make([]Atom, int(cNumProps))
	for i := 0; i < len(ret); i++ {
		ret[i] = Atom(C.AtomPtrIndex(cAtoms, C.int(i)))
	}
	C.XFree(unsafe.Pointer(cAtoms))
	return ret
}

type PropertyChangeMode C.int

const (
	PropertyChangeModeReplace PropertyChangeMode = C.PropModeReplace
	PropertyChangeModePrepend PropertyChangeMode = C.PropModeReplace
	PropertyChangeModeAppend  PropertyChangeMode = C.PropModeReplace
)

func (d *Display) ChangeProperty(w *Window, prop Atom, propType Atom,
	changeMode PropertyChangeMode, data PropertyData) {
	// If there is not a C data object backing this PropertyData, create one now.
	if C.UCharPtrIsNull(data.cData) > 0 {
		data.cData = C.CreateCharBlob(C.size_t(data.recLen * int(data.format)))
		runtime.SetFinalizer(data, func(x PropertyData) {
			C.FreeCharBlob(data.cData)
		})
	}
	// Copy the (modified) go byte slice into the C-managed data storage.
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&data.goData))
	C.memcpy(unsafe.Pointer(data.cData), unsafe.Pointer(hdr.Data), C.size_t(data.recLen*int(data.format)))
	// Call the property accessor.
	C.XChangeProperty(d.hnd, w.hnd, C.Atom(prop), C.Atom(propType), C.int(data.format),
		C.int(changeMode), data.cData, C.int(data.recLen))
}

func (d *Display) RotateWindowProperties(w *Window, props []Atom, offset int) {
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&props))
	C.XRotateWindowProperties(d.hnd, w.hnd, (*C.Atom)(unsafe.Pointer(hdr.Data)), C.int(hdr.Len), C.int(offset))
}

func (d *Display) DeleteProperty(w *Window, prop Atom) {
	C.XDeleteProperty(d.hnd, w.hnd, C.Atom(prop))
}

func (d *Display) SetSelectionOwner(selection Atom, newOwner *Window, tm time.Time) {
	C.XSetSelectionOwner(d.hnd, C.Atom(selection), newOwner.hnd, C.Time(tm.Unix()))
}

func (d *Display) GetSelectionOwner(selection Atom) (*Window, bool) {
	wnd := C.XGetSelectionOwner(d.hnd, C.Atom(selection))
	if wnd == C.None {
		return nil, false
	}
	return &Window{hnd: wnd}, true
}

func (d *Display) ConvertSelection(selection Atom, target Atom, prop Atom, req *Window, tm time.Time) {
	C.XConvertSelection(d.hnd, C.Atom(selection), C.Atom(target), C.Atom(prop), req.hnd, C.Time(tm.Unix()))
}

type DrawableConstant C.int

const DrawableConstantInputOnly = C.InputOnly

func (dc DrawableConstant) IsDrawable() bool { return true }

func (dc DrawableConstant) Handle() C.Drawable { return C.Drawable(dc) }

func (d *Display) CreatePixmap(dd Drawable, width, height int, depth int) Pixmap {
	return Pixmap(C.XCreatePixmap(d.hnd, dd.Handle(), C.uint(width), C.uint(height), C.uint(depth)))
}

func (d *Display) FreePixmap(pm Pixmap) {
	C.XFreePixmap(d.hnd, C.Pixmap(pm))
}

func (d *Display) CreateFontCursor(wcShape uint) Cursor {
	return Cursor(C.XCreateFontCursor(d.hnd, C.uint(wcShape)))
}

type Font C.Font

/*
func (d *Display) CreateGlyphCursor(source Font, mask Font, srcChar uint, maskChar uint) {

}
Cursor XCreateGlyphCursor(display, source_font, mask_font, source_char, mask_char,
foreground_color, background_color)
Display *display;
Font source_font, mask_font;
unsigned int source_char, mask_char;
XColor *foreground_color;
XColor *background_color;
*/
