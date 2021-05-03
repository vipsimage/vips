package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"

// Free VipsImage
func Free(in *Image) {
	C.vipsimage_free(in.vipsImage)
}

// Free Image
func (th *Image) Free() {
	Free(th)
}
