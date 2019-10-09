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
	if C.vipsimage_resize(th.vipsImage, &th.vipsImage, C.double(scale)) != 0 {
		return Error()
	}
	return
}

// Rotate image
func (th *Image) Rotate(angle float64) (err error) {
	if C.vipsimage_rotate(th.vipsImage, &th.vipsImage, C.double(angle)) != 0 {
		return Error()
	}

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
	out = New()

	if C.vipsimage_thumbnail_buffer(unsafe.Pointer(&buf[0]), C.ulong(len(buf)), &out.vipsImage, C.int(width)) != 0 {
		err = Error()
		return
	}

	return
}

// ThumbnailImage Make a thumbnail from *Image
func (th *Image) ThumbnailImage(width int) (err error) {
	if C.vipsimage_thumbnail_image(th.vipsImage, &th.vipsImage, C.int(width)) != 0 {
		return Error()
	}

	return
}
