package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"
import (
	"unsafe"
)

// Merge joins two images left-right (with ref on the left)
// or up-down (with ref above) with a smooth seam.
func (th *Image) Merge(sec *Image, direction Direction, dx, dy int) (err error) {
	var vipsImage *C.VipsImage
	
	if C.vipsimage_merge(th.vipsImage, sec.vipsImage, &vipsImage,
		C.VipsDirection(direction), C.int(dx), C.int(dy)) != 0 {
		
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}

// Mosaic joins two images left-right (with ref on the left)
// or top-bottom (with ref above) given an approximate overlap.
func (th *Image) Mosaic(sec *Image, direction Direction, xref, yref, xsec, ysec int) (err error) {
	var vipsImage *C.VipsImage
	
	if C.vipsimage_mosaic(th.vipsImage, sec.vipsImage, &vipsImage,
		C.VipsDirection(direction), C.int(xref), C.int(yref), C.int(xsec), C.int(ysec)) != 0 {
		
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}

// Mosaic1 This operation joins two images top-bottom (with sec on the right)
// or left-right (with sec at the bottom) given an approximate pair of tie-points.
// sec is scaled and rotated as necessary before the join.
func (th *Image) Mosaic1(sec *Image, direction Direction, xr1, yr1, xs1, ys1, xr2, yr2, xs2, ys2 int) (err error) {
	var vipsImage *C.VipsImage
	
	if C.vipsimage_mosaic1(th.vipsImage, sec.vipsImage, &vipsImage, C.VipsDirection(direction),
		C.int(xr1), C.int(yr1), C.int(xs1), C.int(ys1), C.int(xr2), C.int(yr2), C.int(xs2), C.int(ys2)) != 0 {
		
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}

// Match scale, rotate and translate sec so that the tie-points line up.
func (th *Image) Match(sec *Image, xr1, yr1, xs1, ys1, xr2, yr2, xs2, ys2 int) (err error) {
	var vipsImage *C.VipsImage
	
	if C.vipsimage_match(th.vipsImage, sec.vipsImage, &vipsImage,
		C.int(xr1), C.int(yr1), C.int(xs1), C.int(ys1), C.int(xr2), C.int(yr2), C.int(xs2), C.int(ys2)) != 0 {
		
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}

// GlobalBalance can be used to remove contrast differences in an assembled mosaic.
func (th *Image) GlobalBalance() (err error) {
	var vipsImage *C.VipsImage
	
	if C.vipsimage_globalbalance(th.vipsImage, &vipsImage) != 0 {
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}

// ReMosaic takes apart the mosaiced image in and rebuilds it, substituting images.
// todo failed: class "remosaic" not found
func (th *Image) ReMosaic(oldStr, newStr string) (err error) {
	var old *C.char = C.CString(oldStr)
	defer C.free(unsafe.Pointer(old))
	var nStr *C.char = C.CString(newStr)
	defer C.free(unsafe.Pointer(nStr))
	
	var vipsImage *C.VipsImage
	if C.vipsimage_remosaic(th.vipsImage, &vipsImage,
		old, nStr) != 0 {
		
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}
