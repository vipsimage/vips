package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"
import (
	"fmt"
)

// FormatSizeof ...
func (th *Image) FormatSizeof() {
	format := C.vips_image_get_format(th.vipsImage)
	fmt.Println(format)
	fmt.Println(C.vips_format_sizeof(format))
}
