package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"
import (
	"math"
	"unsafe"
)

// ForeignFindLoad searches for an operation you could use to load filename.
// Any trailing options on filename are stripped and ignored.
func ForeignFindLoad(filename string) string {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	return C.GoString(C.vips_foreign_find_load(name))
}

// ForeignFindLoadBuffer searches for an operation you could use to load a memory buffer.
func ForeignFindLoadBuffer(data []byte) string {
	return C.GoString(C.vips_foreign_find_load_buffer(unsafe.Pointer(&data[0]), C.ulong(len(data))))
}

// Load read in a vips image.
func Load(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()
	if C.vipsimage_vipsload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// Save2file Write *Image to filename in VIPS format.
func (th *Image) Save2file(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_vipssave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// OpenSlideLoad Read a virtual slide supported by the OpenSlide library into a VIPS image.
// OpenSlide supports images in Aperio, Hamamatsu, MIRAX, Sakura, Trestle, and Ventana formats.
func OpenSlideLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))
	out = New()

	if C.vipsimage_openslideload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// JPEGLoad read a JPEG file into a VIPS image.
// It can read most 8-bit JPEG images, including CMYK and YCbCr.
func JPEGLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_jpegload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// JPEGLoadBuffer read a JPEG-formatted memory block into a VIPS image.
// Exactly as JPEGLoad, but read from a memory buffer.
func JPEGLoadBuffer(buf []byte) (out *Image, err error) {
	out = New()

	if C.vipsimage_jpegload_buffer(unsafe.Pointer(&buf[0]), C.ulong(len(buf)), &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// JPEGSave write a VIPS image to a file as JPEG.
func (th *Image) JPEGSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_jpegsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// JPEGSaveBuffer as JPEGSave, but save to a memory buffer.
func (th *Image) JPEGSaveBuffer() (buf []byte, size int, err error) {
	var si C.ulong
	var ptr unsafe.Pointer

	if C.vipsimage_jpegsave_buffer(th.vipsImage, &ptr, &si) != 0 {
		err = Error()
		return
	}

	size = int(si)
	// buf = C.GoBytes(ptr, C.int(si))
	// without copying the original data
	// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
	buf = (*[math.MaxInt32]byte)(ptr)[:size:size]

	return
}

// JPEGSaveMime As JPEGSave, but save as a mime jpeg on stdout.
func (th *Image) JPEGSaveMime() (err error) {
	if C.vipsimage_jpegsave_mime(th.vipsImage) != 0 {
		return Error()
	}

	return
}

// WEBPLoad read a WebP file into a VIPS image.
func WEBPLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_webpload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// WEBPLoadBuffer read a WebP-formatted memory block into a VIPS image.
// Exactly as WEBPLoad, but read from a memory buffer.
func WEBPLoadBuffer(buf []byte) (out *Image, err error) {
	out = New()

	if C.vipsimage_webpload_buffer(unsafe.Pointer(&buf[0]), C.ulong(len(buf)), &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// WEBPSave write a VIPS image to a file as WEBP.
func (th *Image) WEBPSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_webpsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// WEBPSaveBuffer as WEBPSave, but save to a memory buffer.
func (th *Image) WEBPSaveBuffer() (buf []byte, size int, err error) {
	var si C.ulong
	var ptr unsafe.Pointer

	if C.vipsimage_webpsave_buffer(th.vipsImage, &ptr, &si) != 0 {
		err = Error()
		return
	}

	size = int(si)
	buf = (*[math.MaxInt32]byte)(ptr)[:size:size]

	return
}

// WEBPSaveMime As WEBPSave, but save as a mime webp on stdout.
func (th *Image) WEBPSaveMime() (err error) {
	if C.vipsimage_webpsave_mime(th.vipsImage) != 0 {
		return Error()
	}

	return
}

// TIFFLoad read a TIFF file into a VIPS image.
// It is a full baseline TIFF 6 reader, with extensions for tiled images,
// multipage images, LAB colour space, pyramidal images and JPEG compression.
// including CMYK and YCbCr.
func TIFFLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_tiffload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// TIFFLoadBuffer read a TIFF-formatted memory block into a VIPS image.
// Exactly as TIFFLoad, but read from a memory buffer.
func TIFFLoadBuffer(buf []byte) (out *Image, err error) {
	out = New()

	if C.vipsimage_tiffload_buffer(unsafe.Pointer(&buf[0]), C.ulong(len(buf)), &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// TIFFSave write a VIPS image to a file as TIFF.
func (th *Image) TIFFSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_tiffsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// TIFFSaveBuffer as TIFFSave, but save to a memory buffer.
func (th *Image) TIFFSaveBuffer() (buf []byte, size int, err error) {
	var si C.ulong
	var ptr unsafe.Pointer

	if C.vipsimage_tiffsave_buffer(th.vipsImage, &ptr, &si) != 0 {
		err = Error()
		return
	}

	size = int(si)
	buf = (*[math.MaxInt32]byte)(ptr)[:size:size]

	return
}

// OpenEXRLoad read a OpenEXR file into a VIPS image.
func OpenEXRLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_openexrload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// FITSLoad read a FITS image file into a VIPS image.
func FITSLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_fitsload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// FITSSave write a VIPS image to a file as FITS.
func (th *Image) FITSSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_fitssave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// AnalyzeLoad load an Analyze 6.0 file.
// If filename is "fred.img", this will look for an image header called "fred.hdr"
// and pixel data in "fred.img". You can also load "fred" or "fred.hdr".
func AnalyzeLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))
	out = New()

	if C.vipsimage_analyzeload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// RawLoad mmaps the file, setting up out so that access to that image will read from the file.
func RawLoad(filename string, width, height, bands int) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_rawload(name, &out.vipsImage,
		C.int(width), C.int(height), C.int(bands)) != 0 {

		err = Error()
		return
	}

	return
}

// RawSave write a VIPS image to a file as Raw.
func (th *Image) RawSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_rawsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// CSVLoad load a CSV (comma-separated values) file.
// The output image is always 1 band (monochrome), VIPS_FORMAT_DOUBLE.
// Use BandFold() to turn RGBRGBRGB mono images into colour iamges."fred.img".
// You can also load "fred" or "fred.hdr".
func CSVLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_csvload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// CSVSave write a VIPS image to a file as CSV.
func (th *Image) CSVSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_csvsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// MatrixLoad reads a matrix from a file.
func MatrixLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_matrixload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// MatrixSave write a VIPS image to a file as Matrix.
func (th *Image) MatrixSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_matrixsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// MatrixPrint print in to stdout in matrix format.
// See MatrixLoad for a description of the format.
func (th *Image) MatrixPrint() (err error) {
	if C.vipsimage_matrixprint(th.vipsImage) != 0 {
		return Error()
	}

	return
}

// MagickLoad read in an image using libMagick, the ImageMagick library.
// This library can read more than 80 file formats, including SVG, BMP,
// EPS, DICOM and many others. The reader can handle any ImageMagick image,
// including the float and double formats. It will work with any quantum size,
// including HDR. Any metadata attached to the libMagick image is copied on
// to the VIPS image.
func MagickLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_pngload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// MagickLoadBuffer read a Magick-formatted memory block into a VIPS image.
// Exactly as MagickLoad, but read from a memory buffer.
func MagickLoadBuffer(buf []byte) (out *Image, err error) {
	out = New()

	if C.vipsimage_pngload_buffer(unsafe.Pointer(&buf[0]), C.ulong(len(buf)), &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// MagickSave write a VIPS image to a file as Magick.
func (th *Image) MagickSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_pngsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// MagickSaveBuffer as MagickSave, but save to a memory buffer.
func (th *Image) MagickSaveBuffer() (buf []byte, size int, err error) {
	var si C.ulong
	var ptr unsafe.Pointer

	if C.vipsimage_pngsave_buffer(th.vipsImage, &ptr, &si) != 0 {
		err = Error()
		return
	}

	size = int(si)
	buf = (*[math.MaxInt32]byte)(ptr)[:size:size]

	return
}

// PNGLoad read a PNG-formatted memory block into a VIPS image.
// It can read all png images, including 8- and 16-bit images, 1 and 3 channel,
// with and without an alpha channel.
func PNGLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_pngload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// PNGLoadBuffer read a PNG-formatted memory block into a VIPS image.
// Exactly as PNGLoad, but read from a memory buffer.
func PNGLoadBuffer(buf []byte) (out *Image, err error) {
	out = New()

	if C.vipsimage_pngload_buffer(unsafe.Pointer(&buf[0]), C.ulong(len(buf)), &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// PNGSave write a VIPS image to a file as PNG.
func (th *Image) PNGSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_pngsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// PNGSaveBuffer as PNGSave, but save to a memory buffer.
func (th *Image) PNGSaveBuffer() (buf []byte, size int, err error) {
	var si C.ulong
	var ptr unsafe.Pointer

	if C.vipsimage_pngsave_buffer(th.vipsImage, &ptr, &si) != 0 {
		err = Error()
		return
	}

	size = int(si)
	buf = (*[math.MaxInt32]byte)(ptr)[:size:size]

	return
}

// PPMLoad read a PPM/PBM/PGM/PFM file into a VIPS image.
func PPMLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_ppmload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// PPMSave write a VIPS image to a file as PPM.
func (th *Image) PPMSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_ppmsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// MatLoad read a Matlab save file into a VIPS image.
func MatLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_matload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// RADLoad read a Radiance (HDR) file into a VIPS image.
func RADLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_radload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// RADSave write a VIPS image to a file as RAD.
func (th *Image) RADSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_radsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// RADSaveBuffer as RADSave, but save to a memory buffer.
func (th *Image) RADSaveBuffer() (buf []byte, size int, err error) {
	var si C.ulong
	var ptr unsafe.Pointer

	if C.vipsimage_radsave_buffer(th.vipsImage, &ptr, &si) != 0 {
		err = Error()
		return
	}

	size = int(si)
	buf = (*[math.MaxInt32]byte)(ptr)[:size:size]

	return
}

// PDFLoad read a PDF-formatted memory buffer into a VIPS image. Exactly as vips_pdfload(), but read from memory.
func PDFLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_pngload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// PDFLoadBuffer read a PDF-formatted memory block into a VIPS image.
// Exactly as PDFLoad, but read from a memory buffer.
func PDFLoadBuffer(buf []byte) (out *Image, err error) {
	out = New()

	if C.vipsimage_pngload_buffer(unsafe.Pointer(&buf[0]), C.ulong(len(buf)), &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// SVGLoad render a SVG file into a VIPS image. Rendering uses the librsvg library and should be fast.
func SVGLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_pngload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// SVGLoadBuffer read a SVG-formatted memory block into a VIPS image.
// Exactly as SVGLoad, but read from a memory buffer.
func SVGLoadBuffer(buf []byte) (out *Image, err error) {
	out = New()

	if C.vipsimage_pngload_buffer(unsafe.Pointer(&buf[0]), C.ulong(len(buf)), &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// GIFLoad render a GIF file into a VIPS image. Rendering uses the librgif library and should be fast.
func GIFLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_pngload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// GIFLoadBuffer read a GIF-formatted memory block into a VIPS image.
// Exactly as GIFLoad, but read from a memory buffer.
func GIFLoadBuffer(buf []byte) (out *Image, err error) {
	out = New()

	if C.vipsimage_pngload_buffer(unsafe.Pointer(&buf[0]), C.ulong(len(buf)), &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// HEIFLoad read a HEIF image file into a VIPS image.
func HEIFLoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_heifload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// HEIFLoadBuffer read a HEIF-formatted memory block into a VIPS image.
// Exactly as HEIFLoad, but read from a memory buffer.
func HEIFLoadBuffer(buf []byte) (out *Image, err error) {
	out = New()

	if C.vipsimage_heifload_buffer(unsafe.Pointer(&buf[0]), C.ulong(len(buf)), &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// HEIFSave write a VIPS image to a file as HEIF.
func (th *Image) HEIFSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_heifsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// HEIFSaveBuffer as HEIFSave, but save to a memory buffer.
func (th *Image) HEIFSaveBuffer() (buf []byte, size int, err error) {
	var si C.ulong
	var ptr unsafe.Pointer

	if C.vipsimage_heifsave_buffer(th.vipsImage, &ptr, &si) != 0 {
		err = Error()
		return
	}

	size = int(si)
	buf = (*[math.MaxInt32]byte)(ptr)[:size:size]

	return
}

// NIFTILoad read a NIFTI image file into a VIPS image.
func NIFTILoad(filename string) (out *Image, err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	out = New()

	if C.vipsimage_niftiload(name, &out.vipsImage) != 0 {
		err = Error()
		return
	}

	return
}

// NIFTISave write a VIPS image to a file as NIFTI.
func (th *Image) NIFTISave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_niftisave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}

// DZSave save an image as a set of tiles at various resolutions.
// By default dzsave uses DeepZoom layout -- use layout to pick other conventions.
func (th *Image) DZSave(filename string) (err error) {
	var name *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	if C.vipsimage_dzsave(th.vipsImage, name) != 0 {
		return Error()
	}

	return
}
