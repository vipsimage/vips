package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"

// DrawRect paint pixels within left , top , width , height in image with ink.
// If fill is zero, just paint a 1-pixel-wide outline.
func (th *Image) DrawRect(ink []float64, left, top, width, height int) (err error) {
	if C.vipsimage_draw_rect(th.vipsImage, f2d(ink), C.int(len(ink)),
		C.int(left), C.int(top), C.int(width), C.int(height)) != 0 {
		return Error()
	}

	return
}

// DrawRect1 as DrawRect, but just take a single float64 for ink.
func (th *Image) DrawRect1(ink float64, left, top, width, height int) (err error) {
	if C.vipsimage_draw_rect1(th.vipsImage, C.double(ink),
		C.int(left), C.int(top), C.int(width), C.int(height)) != 0 {

		return Error()
	}

	return
}

// DrawPoint draw a single pixel at x, y.
func (th *Image) DrawPoint(ink []float64, x, y int) (err error) {
	if C.vipsimage_draw_point(th.vipsImage, f2d(ink), C.int(len(ink)), C.int(x), C.int(y)) != 0 {
		return Error()
	}

	return
}

// DrawPoint1 as DrawPoint, but just take a single float64 for ink.
func (th *Image) DrawPoint1(ink float64, x, y int) (err error) {
	if C.vipsimage_draw_point1(th.vipsImage, C.double(ink), C.int(x), C.int(y)) != 0 {
		return Error()
	}

	return
}

// DrawImage draw sub on top of image at position x, y.
func (th *Image) DrawImage(sub *Image, x, y int) (err error) {
	if C.vipsimage_draw_image(th.vipsImage, sub.vipsImage, C.int(x), C.int(y)) != 0 {
		return Error()
	}

	return
}

// DrawMask draw mask on the image. mask is a monochrome 8-bit image with 0/255
// for transparent or ink coloured points. Intermediate values blend the ink
// with the pixel. Use with Text() to draw text on an image. Use in a
// DrawLine() subclass to draw an object along a line.
func (th *Image) DrawMask(ink []float64, mask *Image, x, y int) (err error) {
	if C.vipsimage_draw_mask(th.vipsImage, f2d(ink), C.int(len(ink)),
		mask.vipsImage, C.int(x), C.int(y)) != 0 {
		return Error()
	}

	return
}

// DrawMask1 as DrawMask, but just take a single float64 for ink.
func (th *Image) DrawMask1(ink float64, mask *Image, x, y int) (err error) {
	if C.vipsimage_draw_mask1(th.vipsImage, C.double(ink), mask.vipsImage, C.int(x), C.int(y)) != 0 {
		return Error()
	}

	return
}

// DrawLine draws a 1-pixel-wide line on an image.
func (th *Image) DrawLine(ink []float64, x1, y1, x2, y2 int) (err error) {
	if C.vipsimage_draw_line(th.vipsImage, f2d(ink), C.int(len(ink)),
		C.int(x1), C.int(y1), C.int(x2), C.int(y2)) != 0 {
		return Error()
	}

	return
}

// DrawLine1 as DrawLine, but just take a single double for ink.
func (th *Image) DrawLine1(ink float64, x1, y1, x2, y2 int) (err error) {
	if C.vipsimage_draw_line1(th.vipsImage, C.double(ink),
		C.int(x1), C.int(y1), C.int(x2), C.int(y2)) != 0 {
		return Error()
	}
	return
}

// DrawCircle Draws a circle on image.
// If fill is TRUE then the circle is filled,
// otherwise a 1-pixel-wide perimeter is drawn.
func (th *Image) DrawCircle(ink []float64, cx, cy, radius int) (err error) {
	if C.vipsimage_draw_circle(th.vipsImage, f2d(ink), C.int(len(ink)),
		C.int(cx), C.int(cy), C.int(radius)) != 0 {

		return Error()
	}
	return
}

// DrawCircle1 as DrawCircle, but just take a single double for ink.
func (th *Image) DrawCircle1(ink float64, cx, cy, radius int) (err error) {
	if C.vipsimage_draw_circle1(th.vipsImage, C.double(ink),
		C.int(cx), C.int(cy), C.int(radius)) != 0 {

		return Error()
	}

	return
}

// DrawFlood flood-fill image with ink, starting at position x , y.
// The filled area is bounded by pixels that are equal to the ink colour,
// in other words, it searches for pixels enclosed by an edge of ink .
func (th *Image) DrawFlood(ink []float64, x, y int) (err error) {
	if C.vipsimage_draw_flood(th.vipsImage,
		f2d(ink), C.int(len(ink)), C.int(x), C.int(y)) != 0 {

		return Error()
	}

	return
}

// DrawFlood1 as DrawFlood, but just take a single double for ink.
func (th *Image) DrawFlood1(ink float64, x, y int) (err error) {
	if C.vipsimage_draw_flood1(th.vipsImage, C.double(ink), C.int(x), C.int(y)) != 0 {
		return Error()
	}

	return
}

// DrawSmudge smudge a section of image.
// Each pixel in the area left , top , width , height is replaced by
// the average of the surrounding 3x3 pixels.
func (th *Image) DrawSmudge(left, top, width, height int) (err error) {
	if C.vipsimage_draw_smudge(th.vipsImage,
		C.int(left), C.int(top), C.int(width), C.int(height)) != 0 {

		return Error()
	}

	return
}

// f2d float to C double
// it is up to you to free it.
func f2d(in []float64) *C.double {
	arr := make([]C.double, 0)
	for _, v := range in {
		arr = append(arr, C.double(v))
	}

	return &arr[0]
}
