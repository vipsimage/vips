package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"
import (
	"os"
	"unsafe"
)

// Image wrap *C.VipsImage to go
type Image struct {
	vipsImage *C.VipsImage
}

// New empty VipsImage
func New() *Image {
	return &Image{vipsImage: C.vips_image_new()}
}

// NewMemory creates a new VipsImage.
// when written to, will create a memory image.
func NewMemory() *Image {
	return &Image{C.vips_image_new_memory()}
}

// NewFromFile opens filename for reading.
// It can load files in many image formats,
// including VIPS, TIFF, PNG, JPEG, FITS,
// Matlab, OpenEXR, CSV, WebP, Radiance,
// RAW, PPM and others.
func NewFromFile(filename string) (out *Image, err error) {
	_, err = os.Stat(filename)
	if err != nil {
		return
	}

	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = &Image{vipsImage: C.vipsimage_image_new_from_file(name)}
	return
}

// NewFromBuffer loads an image from the formatted area of memory buf.
func NewFromBuffer(buf []byte, optionString string) *Image {
	var oName *C.char = C.CString(optionString)
	defer C.free(unsafe.Pointer(oName))

	return &Image{
		vipsImage: C.vipsimage_image_new_from_buffer(
			unsafe.Pointer(&buf[0]), C.ulong(len(buf)), oName,
		),
	}
}

// Height return image height
func (th *Image) Height() int {
	return int(th.vipsImage.Ysize)
}

// Width return image width
func (th *Image) Width() int {
	return int(th.vipsImage.Xsize)
}

// ProgressSet ff set, vips will print messages about the progress of computation to stdout.
// This can also be enabled with the --vips-progress option, or by setting the environment
// variable VIPS_PROGRESS.
func ProgressSet(progress bool) {
	C.vips_progress_set(btoi(progress))
}

// ImageSetProgress vips signals evaluation progress via the “preeval”,
// “eval” and “posteval” signals. Progress is signalled on the most-downstream
// image for which vips_image_set_progress() was called.
func (th *Image) ImageSetProgress(progress bool) {
	C.vips_image_set_progress(th.vipsImage, btoi(progress))
}

// btoi transform go bool to C.int
func btoi(b bool) C.int {
	if b {
		return 1
	}
	return 0
}
