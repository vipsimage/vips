package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"
import (
	"unsafe"
)

// Black make a black unsigned char image of a specified size.
func Black(width, height int) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_black(&out.vipsImage, C.int(width), C.int(height)) != 0 {
		err = Error()
		return
	}
	
	return
}

// Xyz Create a two-band uint32 image where the elements in the
// first band have the value of their x coordinate and elements
// in the second band have their y coordinate.
//
// You can make any image where the value of a pixel is a function
// of its (x, y) coordinate by combining this operator with the
// arithmetic operators.
func Xyz(width, height int) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_xyz(&out.vipsImage, C.int(width), C.int(height)) != 0 {
		err = Error()
		return
	}
	
	return
}

// Grey create a one-band float image with the left-most
// column zero and the right-most 1. Intermediate pixels
// are a linear ramp.
func Grey(width, height int) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_grey(&out.vipsImage, C.int(width), C.int(height)) != 0 {
		err = Error()
		return
	}
	
	return
}

// GaussMat Creates a circularly symmetric Gaussian image of radius sigma.
// The size of the mask is determined by the variable minAmpl ;
// if for instance the value .1 is entered this means that the produced mask
// is clipped at values less than 10 percent of the maximum amplitude.
func GaussMat(sigma, minAmpl float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_gaussmat(&out.vipsImage, C.double(sigma), C.double(minAmpl)) != 0 {
		err = Error()
		return
	}
	
	return
}

// LogMat creates a circularly symmetric Laplacian of Gaussian mask of radius sigma.
// The size of the mask is determined by the variable min_ampl ; if for instance the
// value .1 is entered this means that the produced mask is clipped at values within
// 10 persent of zero, and where the change between mask elements is less than 10%.
func LogMat(sigma, minAmpl float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_logmat(&out.vipsImage, C.double(sigma), C.double(minAmpl)) != 0 {
		err = Error()
		return
	}
	
	return
}

// Text draw the string text to an image.
// out is a one-band 8-bit unsigned char image,
// with 0 for no text and 255 for text.
// Values between are used for anti-aliasing.
func Text(text string) (out *Image, err error) {
	var txt *C.char = C.CString(text)
	defer C.free(unsafe.Pointer(txt))
	
	out = New()
	
	if C.vipsimage_text(&out.vipsImage, txt) != 0 {
		err = Error()
		return
	}
	
	return
}

// GaussNoise make a one band float image of gaussian noise with the specified distribution.
// The noise distribution is created by averaging 12 random numbers with the appropriate weights.
func GaussNoise(width, height int) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_gaussnoise(&out.vipsImage, C.int(width), C.int(height)) != 0 {
		err = Error()
		return
	}
	
	return
}

// Eye create a test pattern with increasing spatial frequence in X and amplitude in Y.
func Eye(width, height int) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_eye(&out.vipsImage, C.int(width), C.int(height)) != 0 {
		err = Error()
		return
	}
	
	return
}

// Sines creates a float one band image of the a sine waveform in two dimensions.
func Sines(width, height int) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_sines(&out.vipsImage, C.int(width), C.int(height)) != 0 {
		err = Error()
		return
	}
	
	return
}

// Zone create a one-band image of a zone plate.
func Zone(width, height int) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_zone(&out.vipsImage, C.int(width), C.int(height)) != 0 {
		err = Error()
		return
	}
	
	return
}

// Identity creates an identity lookup table.
// ie. one which will leave an image unchanged when applied with Maplut().
// Each entry in the table has a value equal to its position.
func Identity() (out *Image, err error) {
	out = New()
	
	if C.vipsimage_identity(&out.vipsImage) != 0 {
		err = Error()
		return
	}
	
	return
}

// BuildLut this operation builds a lookup table from a set of points.
// Intermediate values are generated by piecewise linear interpolation
func (th *Image) BuildLut() (err error) {
	var vipsImage *C.VipsImage
	
	if C.vipsimage_buildlut(th.vipsImage, &vipsImage) != 0 {
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}

// InvertLut given a mask of target values and real values, generate a
// LUT which will map reals to targets.
// Handy for linearising images from measurements of a colour chart.
// All values in [0,1]. Piecewise linear interpolation, extrapolate
// head and tail to 0 and 1.
func (th *Image) InvertLut() (err error) {
	var vipsImage *C.VipsImage
	
	if C.vipsimage_invertlut(th.vipsImage, &vipsImage) != 0 {
		return Error()
	}
	
	C.g_object_unref(C.gpointer(th.vipsImage))
	th.vipsImage = vipsImage
	return
}

// ToneLut generates a tone curve for the adjustment of image levels.
// It is mostly designed for adjusting the L* part of a LAB image in
// a way suitable for print work, but you can use it for other things too.
func ToneLut() (out *Image, err error) {
	out = New()
	
	if C.vipsimage_tonelut(&out.vipsImage) != 0 {
		err = Error()
		return
	}
	
	return
}

// MaskIdeal Make an ideal high- or low-pass filter,
// that is, one with a sharp cutoff positioned at frequencyCutoff ,
// where frequencyCutoff is in the range 0 - 1.
func MaskIdeal(width, height int, frequencyCutoff float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_mask_ideal(&out.vipsImage, C.int(width), C.int(height), C.double(frequencyCutoff)) != 0 {
		err = Error()
		return
	}
	
	return
}

// MaskIdealRing make an ideal ring-pass or ring-reject filter, that is, one with
// a sharp ring positioned at frequency_cutoff of width width , where frequencyCutoff
// and width are expressed as the range 0 - 1.
func MaskIdealRing(width, height int, frequencyCutoff, ringWidth float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_mask_ideal_ring(&out.vipsImage,
		C.int(width), C.int(height),
		C.double(frequencyCutoff), C.double(ringWidth)) != 0 {
		
		err = Error()
		return
	}
	
	return
}

// MaskIdealBand make an ideal band-pass or band-reject filter,
// that is, one with a sharp cutoff around the point frequencyCutoffX,
// frequencyCutoffY of size radius .
func MaskIdealBand(width, height int, frequencyCutoffX, frequencyCutoffY, radius float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_mask_ideal_band(&out.vipsImage,
		C.int(width), C.int(height),
		C.double(frequencyCutoffX), C.double(frequencyCutoffY), C.double(radius)) != 0 {
		
		err = Error()
		return
		
	}
	
	return
}

// MaskButterWorth make an butterworth high- or low-pass filter, that is, one with a variable,
// smooth transition positioned at frequency_cutoff , where frequency_cutoff is in range 0 - 1.
// The shape of the curve is controlled by order --- higher values give a sharper transition.
// See Gonzalez and Wintz, Digital Image Processing, 1987.
func MaskButterWorth(width, height int, order, frequencyCutoff, amplitudeCutoff float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_mask_butterworth(&out.vipsImage,
		C.int(width), C.int(height), C.double(order),
		C.double(frequencyCutoff), C.double(amplitudeCutoff)) != 0 {
		
		err = Error()
		return
	}
	
	return
}

// MaskButterWorthRing Make a butterworth ring-pass or ring-reject filter,
// that is, one with a variable, smooth transition positioned at frequencyCutoff
// of width width , where frequencyCutoff is in the range 0 - 1. The shape of
// the curve is controlled by order --- higher values give a sharper transition.
// See Gonzalez and Wintz, Digital Image Processing, 1987.
func MaskButterWorthRing(width, height int, order, frequencyCutoff, amplitudeCutoff, ringWidth float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_mask_butterworth_ring(&out.vipsImage,
		C.int(width), C.int(height), C.double(order),
		C.double(frequencyCutoff), C.double(amplitudeCutoff), C.double(ringWidth)) != 0 {
		
		err = Error()
		return
	}
	
	return
}

// MaskButterWorthBand make an butterworth band-pass or band-reject filter,
// that is, one with a variable, smooth transition positioned at frequencyCutoffX,
// frequencyCutoffY , of radius radius . The shape of the curve is controlled by
// order --- higher values give a sharper transition. See Gonzalez and Wintz,
// Digital Image Processing, 1987.
func MaskButterWorthBand(width, height int,
	order, frequencyCutoffX, frequencyCutoffY, radius, amplitudeCutoff float64) (out *Image, err error) {
	
	out = New()
	
	if C.vipsimage_mask_butterworth_band(&out.vipsImage,
		C.int(width), C.int(height), C.double(order),
		C.double(frequencyCutoffX), C.double(frequencyCutoffY),
		C.double(radius), C.double(amplitudeCutoff)) != 0 {
		
		err = Error()
		return
		
	}
	
	return
}

// MaskGaussian Make a gaussian high- or low-pass filter, that is, one with a variable,
// smooth transition positioned at frequencyCutoff .
func MaskGaussian(width, height int, frequencyCutoff, amplitudeCutoff float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_mask_gaussian(&out.vipsImage,
		C.int(width), C.int(height),
		C.double(frequencyCutoff), C.double(amplitudeCutoff)) != 0 {
		
		err = Error()
		return
	}
	
	return
}

// MaskGaussianRing Make a gaussian ring-pass or ring-reject filter,
// that is, one with a variable, smooth transition positioned at
// frequencyCutoff of width ringWidth .
func MaskGaussianRing(width, height int, frequencyCutoff, amplitudeCutoff, ringWidth float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_mask_gaussian_ring(&out.vipsImage,
		C.int(width), C.int(height),
		C.double(frequencyCutoff), C.double(amplitudeCutoff), C.double(ringWidth)) != 0 {
		
		err = Error()
		return
	}
	
	return
}

// MaskGaussianBand Make a gaussian band-pass or band-reject filter,
// that is, one with a variable, smooth transition positioned at
// frequencyCutoffX, frequencyCutoffY, of radius radius.
func MaskGaussianBand(width, height int, frequencyCutoffX, frequencyCutoffY, radius, amplitudeCutoff float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_mask_gaussian_band(&out.vipsImage,
		C.int(width), C.int(height),
		C.double(frequencyCutoffX), C.double(frequencyCutoffY), C.double(radius),
		C.double(amplitudeCutoff)) != 0 {
		
		err = Error()
		return
	}
	
	return
}

// MaskFractal operation should be used to create fractal images by
// filtering the power spectrum of Gaussian white noise. See GaussNoise().
func MaskFractal(width, height int, fractalDimension float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_mask_fractal(&out.vipsImage, C.int(width), C.int(height), C.gdouble(fractalDimension)) != 0 {
		err = Error()
		return
	}
	
	return
}

// FractSurf Generate an image of size width by height and fractal dimension fractalDimension.
// The dimension should be between 2 and 3.
func FractSurf(width, height int, fractalDimension float64) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_fractsurf(&out.vipsImage, C.int(width), C.int(height), C.double(fractalDimension)) != 0 {
		err = Error()
		return
	}
	
	return
}

// Worley create a one-band float image of Worley noise.
// See:
//
// https://en.wikipedia.org/wiki/Worley_noise
func Worley(width, height int) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_worley(&out.vipsImage, C.int(width), C.int(height)) != 0 {
		err = Error()
		return
	}
	
	return
}

// Perlin create a one-band float image of Perlin noise.
// See:
//
// https://en.wikipedia.org/wiki/Perlin_noise
func Perlin(width, height int) (out *Image, err error) {
	out = New()
	
	if C.vipsimage_perlin(&out.vipsImage, C.int(width), C.int(height)) != 0 {
		err = Error()
		return
	}
	return
}
