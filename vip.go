package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

// DefaultInit init vips by default
func DefaultInit() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := Init(); err != nil {
		Shutdown()

		panic(fmt.Sprintf("unable to start vips! err: %s", err.Error()))
	}
}

// Init starts up libvips
func Init() (err error) {
	name := C.CString("vipsimage")
	defer C.free(unsafe.Pointer(name))

	if C.vips_init(name) != 0 {
		return Error()
	}

	return
}

// Shutdown drop caches and close plugins. Run with "--vips-leak" to do a leak check too.
func Shutdown() {
	C.vips_shutdown()
}

// LeakSet turn on or off vips leak checking
func LeakSet(leak bool) {
	C.vips_leak_set(btoi(leak))
}

// VersionString get the VIPS version as a static string,
// including a build date and time. Do not free.
func VersionString() string {
	return C.GoString(C.vips_version_string())
}

// Version get the major, minor or micro library version, with flag values 0, 1 and 2.
func Version(flag int) int {
	return int(C.vips_version(C.int(flag)))
}

// Error Get a pointer to the start of the error buffer as a C string.
// The string is owned by the error system and must not be freed.
func Error() error {
	return errors.New(C.GoString(C.vips_error_buffer()))
}
