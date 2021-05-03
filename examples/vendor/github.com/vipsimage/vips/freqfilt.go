package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"

// FwFFT transform an image to Fourier space.
// VIPS uses the fftw Fourier Transform library.
// If this library was not available when VIPS
// was configured, these functions will fail.
func (th *Image) FwFFT() (err error) {
	if C.vipsimage_fwfft(th.vipsImage, &th.vipsImage) != 0 {
		return Error()
	}

	return
}

// InvFFT transform an image from Fourier space to real space.
// The result is complex.
func (th *Image) InvFFT() (err error) {
	if C.vipsimage_invfft(th.vipsImage, &th.vipsImage) != 0 {
		return Error()
	}

	return
}

// FreqMult transformed to Fourier space, multipled with mask,
// then transformed back to real space. If in is already a complex image,
// just multiply then inverse transform.
func (th *Image) FreqMult(mask *Image) (err error) {
	if C.vipsimage_freqmult(th.vipsImage, mask.vipsImage, &th.vipsImage) != 0 {
		return Error()
	}

	return
}

// Spectrum make a displayable (ie. 8-bit unsigned int) power spectrum.
func (th *Image) Spectrum() (err error) {
	if C.vipsimage_spectrum(th.vipsImage, &th.vipsImage) != 0 {
		return Error()
	}

	return
}

// Phasecor convert the two input images to Fourier space,
// calculate phase-correlation, back to real space.
func (th *Image) Phasecor(in *Image) (err error) {
	if C.vipsimage_phasecor(th.vipsImage, in.vipsImage, &th.vipsImage) != 0 {
		return Error()
	}

	return
}
