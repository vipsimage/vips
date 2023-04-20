package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"
import (
	"unsafe"
)

// Resize image
// upsizing (scale > 1)
func (th *Image) Resize(scale float64) (err error) {
	var vipsImage *C.VipsImage
	
	if C.vipsimage_resize(th.vipsImage, &vipsImage, C.double(scale)) != 0 {
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}

// Rotate image
func (th *Image) Rotate(angle float64) (err error) {
	var vipsImage *C.VipsImage
	
	if C.vipsimage_rotate(th.vipsImage, &vipsImage, C.double(angle)) != 0 {
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}

// Thumbnail Make a thumbnail from a file
func Thumbnail(filename string, width int) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))
	
	out = New()
	if C.vipsimage_thumbnail(name, &out.vipsImage, C.int(width)) != 0 {
		err = Error()
		return
	}
	
	return
}

// ThumbnailBuffer Make a thumbnail from []byte
func ThumbnailBuffer(buf []byte, width int) (out *Image, err error) {
	var vipsImage *C.VipsImage
	p := unsafe.Pointer(&buf[0])
	
	if C.vipsimage_thumbnail_buffer(p, C.ulong(len(buf)), &vipsImage, C.int(width)) != 0 {
		err = Error()
		return
	}
	
	return NewFromVipsImage(vipsImage), nil
}

// ThumbnailImage Make a thumbnail from *Image
func (th *Image) ThumbnailImage(width int) (err error) {
	var vipsImage *C.VipsImage
	
	if C.vipsimage_thumbnail_image(th.vipsImage, &vipsImage, C.int(width)) != 0 {
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}
