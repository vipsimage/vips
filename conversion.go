package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"

// CompassDirection direction on a compass.
type CompassDirection int

const (
	// DirectionCentre centre
	DirectionCentre = CompassDirection(C.VIPS_COMPASS_DIRECTION_CENTRE)
	// DirectionNorth notrh
	DirectionNorth = CompassDirection(C.VIPS_COMPASS_DIRECTION_NORTH)
	// DirectionEast east
	DirectionEast = CompassDirection(C.VIPS_COMPASS_DIRECTION_EAST)
	// DirectionSouth south
	DirectionSouth = CompassDirection(C.VIPS_COMPASS_DIRECTION_SOUTH)
	// DirectionWest west
	DirectionWest = CompassDirection(C.VIPS_COMPASS_DIRECTION_WEST)
	// DirectionNorthEast north-east
	DirectionNorthEast = CompassDirection(C.VIPS_COMPASS_DIRECTION_NORTH_EAST)
	// DirectionSouthEast south-east
	DirectionSouthEast = CompassDirection(C.VIPS_COMPASS_DIRECTION_SOUTH_EAST)
	// DirectionSouthWest south-west
	DirectionSouthWest = CompassDirection(C.VIPS_COMPASS_DIRECTION_SOUTH_WEST)
	// DirectionNorthWest north-west
	DirectionNorthWest = CompassDirection(C.VIPS_COMPASS_DIRECTION_NORTH_WEST)
)

// BlendMode The various Porter-Duff and PDF blend modes. See vips_composite(), for example.
// The Cairo docs have a nice explanation of all the blend modes:
//
// https://www.cairographics.org/operators
// The non-separable modes are not implemented.
type BlendMode int

const (
	// BlendModeClear where the second object is drawn, the first is removed
	BlendModeClear = BlendMode(C.VIPS_BLEND_MODE_CLEAR)
	// BlendModeSource the second object is drawn as if nothing were below
	BlendModeSource = BlendMode(C.VIPS_BLEND_MODE_SOURCE)
	// BlendModeOver the image shows what you would expect if you held two semi-transparent slides on top of each other
	BlendModeOver = BlendMode(C.VIPS_BLEND_MODE_OVER)
	// BlendModeIn the first object is removed completely, the second is only drawn where the first was
	BlendModeIn = BlendMode(C.VIPS_BLEND_MODE_IN)
	// BlendModeOut the second is drawn only where the first isn't
	BlendModeOut = BlendMode(C.VIPS_BLEND_MODE_OUT)
	// BlendModeAtop this leaves the first object mostly intact, but mixes both objects in the overlapping area
	BlendModeAtop = BlendMode(C.VIPS_BLEND_MODE_ATOP)
	// BlendModeDest leaves the first object untouched, the second is discarded completely
	BlendModeDest = BlendMode(C.VIPS_BLEND_MODE_DEST)
	// BlendModeDestOver like OVER, but swaps the arguments
	BlendModeDestOver = BlendMode(C.VIPS_BLEND_MODE_DEST_OVER)
	// BlendModeDestIn like IN, but swaps the arguments
	BlendModeDestIn = BlendMode(C.VIPS_BLEND_MODE_DEST_IN)
	// BlendModeDestOut like OUT, but swaps the arguments
	BlendModeDestOut = BlendMode(C.VIPS_BLEND_MODE_DEST_OUT)
	// BlendModeDestAtop like ATOP, but swaps the arguments
	BlendModeDestAtop = BlendMode(C.VIPS_BLEND_MODE_DEST_ATOP)
	// BlendModeXor something like a difference operator
	BlendModeXor = BlendMode(C.VIPS_BLEND_MODE_XOR)
	// BlendModeAdd a bit like adding the two images
	BlendModeAdd = BlendMode(C.VIPS_BLEND_MODE_ADD)
	// BlendModeSaturate a bit like the darker of the two
	BlendModeSaturate = BlendMode(C.VIPS_BLEND_MODE_SATURATE)
	// BlendModeMultiply at least as dark as the darker of the two inputs
	BlendModeMultiply = BlendMode(C.VIPS_BLEND_MODE_MULTIPLY)
	// BlendModeScreen at least as light as the lighter of the inputs
	BlendModeScreen = BlendMode(C.VIPS_BLEND_MODE_SCREEN)
	// BlendModeOverlay multiplies or screens colors, depending on the lightness
	BlendModeOverlay = BlendMode(C.VIPS_BLEND_MODE_OVERLAY)
	// BlendModeDarken the darker of each component
	BlendModeDarken = BlendMode(C.VIPS_BLEND_MODE_DARKEN)
	// BlendModeLighten the lighter of each component
	BlendModeLighten = BlendMode(C.VIPS_BLEND_MODE_LIGHTEN)
	// BlendModeColourDodge brighten first by a factor second
	BlendModeColourDodge = BlendMode(C.VIPS_BLEND_MODE_COLOUR_DODGE)
	// BlendModeColourBurn darken first by a factor of second
	BlendModeColourBurn = BlendMode(C.VIPS_BLEND_MODE_COLOUR_BURN)
	// BlendModeHardLight multiply or screen, depending on lightness
	BlendModeHardLight = BlendMode(C.VIPS_BLEND_MODE_HARD_LIGHT)
	// BlendModeSoftLight darken or lighten, depending on lightness
	BlendModeSoftLight = BlendMode(C.VIPS_BLEND_MODE_SOFT_LIGHT)
	// BlendModeDifference difference of the two
	BlendModeDifference = BlendMode(C.VIPS_BLEND_MODE_DIFFERENCE)
	// BlendModeExclusion somewhat like DIFFERENCE, but lower-contrast
	BlendModeExclusion = BlendMode(C.VIPS_BLEND_MODE_EXCLUSION)
)

// Angle Fixed rotate angles.
type Angle int

const (
	// AngleD0 no rotate
	AngleD0 = Angle(C.VIPS_ANGLE_D0)
	// AngleD90 90 degrees clockwise
	AngleD90 = Angle(C.VIPS_ANGLE_D90)
	// AngleD180 180 degree rotate
	AngleD180 = Angle(C.VIPS_ANGLE_D180)
	// AngleD270 90 degrees anti-clockwise
	AngleD270 = Angle(C.VIPS_ANGLE_D270)
)

// Direction Operations like Flip() need to be told whether to flip left-right or top-bottom.
type Direction int

const (
	// DirectionHorizontal left-right
	DirectionHorizontal = Direction(C.VIPS_DIRECTION_HORIZONTAL)
	// DirectionVertical top-bottom
	DirectionVertical = Direction(C.VIPS_DIRECTION_VERTICAL)
)

// Replicate repeats an image many times.
// across: repeat input this many times across
// down: repeat input this many times down
func (th *Image) Replicate(across, down int) (err error) {
	if C.vipsimage_replicate(th.vipsImage, &th.vipsImage, C.int(across), C.int(down)) != 0 {
		return Error()
	}

	return
}

// Gravity place in within an image of size width by height at a certain gravity.
func (th *Image) Gravity(direction CompassDirection, width, height int) (err error) {
	vcd := C.VipsCompassDirection(direction)

	if C.vipsimage_gravity(th.vipsImage, &th.vipsImage, vcd, C.int(width), C.int(height)) != 0 {
		return Error()
	}

	return
}

// Composite ...
func (th *Image) Composite() {

}

// Composite2 composite overlay on top of base with mode . See Composite.
func (th *Image) Composite2(overlay *Image, mode BlendMode) (err error) {
	if C.vipsimage_composite2(th.vipsImage, overlay.vipsImage, &th.vipsImage, C.VipsBlendMode(mode)) != 0 {
		return Error()
	}

	return
}

// GetAngle examine the metadata on im and return the VipsAngle
// to rotate by to turn the image upright.
func (th *Image) GetAngle() Angle {
	return Angle(int(C.vips_autorot_get_angle(th.vipsImage)))
}

// AutoRot look at the image metadata and rotate the image to make it upright.
// Note:
//
// vips only supports the four simple rotations, it does not support the various
// mirror modes. If the image is using one of these mirror modes, the image is not
// rotated and the VIPS_META_ORIENTATION tag is not removed.
func (th *Image) AutoRot() (err error) {
	if C.vipsimage_autorot(th.vipsImage, &th.vipsImage) != 0 {
		return Error()
	}

	return
}

// Embed *Image within an image of size width by height at position x , y .
func (th *Image) Embed(x, y, width, height int) (err error) {
	if C.vipsimage_embed(th.vipsImage, &th.vipsImage, C.int(x), C.int(y), C.int(width), C.int(height)) != 0 {
		return Error()
	}

	return
}

// ExtractArea extract an area from an image. The area must fit within *Image.
func (th *Image) ExtractArea(left, top, width, height int) (err error) {
	if C.vipsimage_extract_area(th.vipsImage, &th.vipsImage, C.int(left), C.int(top), C.int(width), C.int(height)) != 0 {
		return Error()
	}
	return
}

// Crop a synonym for ExtractArea.
func (th *Image) Crop(left, top, width, height int) (err error) {
	if C.vipsimage_crop(th.vipsImage, &th.vipsImage, C.int(left), C.int(top), C.int(width), C.int(height)) != 0 {
		return Error()
	}
	return
}

// Flip an image left-right or up-down.
func (th *Image) Flip(direction Direction) (err error) {
	if C.vipsimage_flip(th.vipsImage, &th.vipsImage, C.VipsDirection(direction)) != 0 {
		return Error()
	}

	return
}

// Grid chop a tall thin image up into a set of tiles, lay the tiles out in a grid.
func (th *Image) Grid(tileHeight, across, down int) (err error) {
	if C.vipsimage_grid(th.vipsImage, &th.vipsImage, C.int(tileHeight), C.int(across), C.int(down)) != 0 {
		return
	}
	return
}

// Scale search the image for the maximum and minimum value,
// then return the image as unsigned 8-bit, scaled so that
// the maximum value is 255 and the minimum is zero.
func (th *Image) Scale() (err error) {
	if C.vipsimage_scale(th.vipsImage, &th.vipsImage) != 0 {
		return Error()
	}

	return
}

// SubSample an image by an integer fraction. This is fast, nearest-neighbour shrink.
func (th *Image) SubSample(xFac, yFac int) (err error) {
	if C.vipsimage_subsample(th.vipsImage, &th.vipsImage, C.int(xFac), C.int(yFac)) != 0 {
		return Error()
	}

	return
}

// Zoom an image by repeating pixels. This is fast nearest-neighbour zoom.
func (th *Image) Zoom(xFac, yFac int) (err error) {
	if C.vipsimage_zoom(th.vipsImage, &th.vipsImage, C.int(xFac), C.int(yFac)) != 0 {
		return Error()
	}

	return
}

// Wrap slice an image up and move the segments about so that
// the pixel that was at 0, 0 is now at x , y . If x and y
// are not set, they default to the centre of the image.
func (th *Image) Wrap() (err error) {
	if C.vipsimage_wrap(th.vipsImage, &th.vipsImage) != 0 {
		return Error()
	}

	return
}

// ExtractBand extract a band or bands from an image.
// Extracting out of range is an error.
func (th *Image) ExtractBand(band int) (err error) {
	if C.vipsimage_extract_band(th.vipsImage, &th.vipsImage, C.int(band)) != 0 {
		return Error()
	}

	return
}

// SmartCrop crop an image down to a specified width and height by removing boring parts.
func (th *Image) SmartCrop(width, height int) (err error) {
	if C.vipsimage_smartcrop(th.vipsImage, &th.vipsImage, C.int(width), C.int(height)) != 0 {
		return Error()
	}

	return
}
