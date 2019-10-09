
# Go Vips  
Go Vips is a go bind to libvips C API.  

## Install  
Check if pkg-config is able to find the right libvips include:  
```
pkg-config --cflags --libs vips  
```

Package 'libffi', required by 'gobject-2.0', not found  
```
# optional, use you libffi lib path  
export PKG_CONFIG_PATH="/usr/local/opt/libffi/lib/pkgconfig"  
```

Per the security update https://groups.google.com/forum/#!topic/golang-announce/X7N1mvntnoU you may need whitelist the -Xpreprocessor flag in your environment.  
```
export CGO_CFLAGS_ALLOW='-Xpreprocessor'  
```

## examples
```golang
package main

import (
	"fmt"

	"github.com/vipsimage/vips"
)


func main() {
	img := vips.NewFromFile("./examples/images/photo.jpg")
	wm := vips.NewFromFile("./examples/images/watermark.png")

	err = wm.ThumbnailImage(200)
	if err != nil {
		panic(err)
	}

	err = wm.Resize(float64(img.Width()) / float64(wm.Width()))
	if err != nil {
		panic(err)
	}

	err = wm.Rotate(30)
	if err != nil {
		panic(err)
	}

	err = wm.Replicate(img.Width()/wm.Width()+1, img.Height()/wm.Height()+1)
	if err != nil {
		panic(err)
	}
	img.Replicate(10, 10)

	fmt.Println(img.Height(), img.Width())

	err = img.Save2file("out.png")
	if err != nil {
		panic(err)
	}
}
```

## Features  
### conversion  
- Replicate: repeats an image many times.  
- Gravity: place in within an image of size width by height at a certain gravity.  
- Composite2: composite overlay on top of base with mode .  
- GetAngle: examine the metadata on im and return the VipsAngle  
- AutoRot: look at the image metadata and rotate the image to make it upright.  
- Embed: Embed *Image within an image of size width by height at position x , y .  
- ExtractArea: extract an area from an image. The area must fit within *Image.  
- Crop: Crop a synonym for ExtractArea  
- Flip: Flip an image left-right or up-down.  
- Grid: Grid chop a tall thin image up into a set of tiles, lay the tiles out in a grid.  
- Scale: Scale search the image for the maximum and minimum value  
- SubSample: SubSample an image by an integer fraction. This is fast, nearest-neighbour shrink.  
- Zoom: Zoom an image by repeating pixels. This is fast nearest-neighbour zoom.  
- Wrap: Wrap slice an image up and move the segments  
- ExtractBand: SmartCrop crop an image down to a specified width and height by removing boring parts.  
### create  
- Black: Black make a black unsigned char image of a specified size.  
- Xyz: Xyz Create a two-band uint32 image where the elements in the first band have the value of their x coordinate and elements in the second band have their y coordinate.  
- Grey: Grey create a one-band float image with the left-most column zero and the right-most 1.  
- GaussMat: GaussMat Creates a circularly symmetric Gaussian image of radius sigma.  
- LogMat: LogMat creates a circularly symmetric Laplacian of Gaussian mask of radius sigma.  
- Text: Text draw the string text to an image.  
- GaussNoise: GaussNoise make a one band float image of gaussian noise with the specified distribution.  
- Eye: Eye create a test pattern with increasing spatial frequence in X and amplitude in Y.  
- Sines: Sines creates a float one band image of the a sine waveform in two dimensions.  
- Zone: Zone create a one-band image of a zone plate.  
- Identity: Identity creates an identity lookup table.  
- BuildLut: BuildLut this operation builds a lookup table from a set of points.  
- InvertLut: InvertLut given a mask of target values and real values, generate a LUT which will map reals to targets.  
- ToneLut: ToneLut generates a tone curve for the adjustment of image levels.  
- MaskIdeal: MaskIdeal Make an ideal high- or low-pass filter.  
- MaskIdealRing: MaskIdealRing make an ideal ring-pass or ring-reject filter.  
- MaskIdealBand: MaskIdealBand make an ideal band-pass or band-reject filter.  
- MaskButterWorth: MaskButterWorth make an butterworth high- or low-pass filter.  
- MaskButterWorthRing: MaskButterWorthRing Make a butterworth ring-pass or ring-reject filter.  
- MaskButterWorthBand: MaskButterWorthBand make an butterworth band-pass or band-reject filter  
- MaskGaussian: MaskGaussian Make a gaussian high- or low-pass filter.  
- MaskGaussianRing: MaskGaussianRing Make a gaussian ring-pass or ring-reject filter.  
- MaskGaussianBand: MaskGaussianBand Make a gaussian band-pass or band-reject filter.  
- MaskFractal: MaskFractal operation should be used to create fractal images by filtering the power spectrum of Gaussian white noise  
- FractSurf: FractSurf Generate an image of size width by height and fractal dimension fractalDimension.  
- Worley: Worley create a one-band float image of Worley noise.  
- Perlin: Perlin create a one-band float image of Perlin noise.  
### draw  
- DrawRect: DrawRect paint pixels within left , top , width , height in image with ink.  
- DrawRect1: DrawRect1 as DrawRect, but just take a single float64 for ink.  
- DrawPoint: DrawPoint draw a single pixel at x, y.  
- DrawPoint1: DrawPoint1 as DrawPoint, but just take a single float64 for ink.  
- DrawImage: DrawImage draw sub on top of image at position x, y.  
- DrawMask: DrawMask draw mask on the image. mask is a monochrome 8-bit image with 0/255 for transparent or ink coloured points.  
- DrawMask1: DrawMask1 as DrawMask, but just take a single float64 for ink.  
- DrawLine: DrawLine draws a 1-pixel-wide line on an image.  
- DrawLine1: DrawLine1 as DrawLine, but just take a single double for ink.  
- DrawCircle: DrawCircle Draws a circle on image.  
- DrawCircle1: DrawCircle1 as DrawCircle, but just take a single double for ink.  
- DrawFlood: DrawFlood flood-fill image with ink, starting at position x , y.  
- DrawFlood1: DrawFlood1 as DrawFlood, but just take a single double for ink.  
- DrawSmudge: DrawSmudge smudge a section of image.  
### foreign_save

### freqfilt
- FwFFT: FwFFT transform an image to Fourier space.
- InvFFT: InvFFT transform an image from Fourier space to real space.
- FreqMult: FreqMult transformed to Fourier space, multipled with mask.
- Spectrum: Spectrum make a displayable (ie. 8-bit unsigned int) power spectrum.
- Phasecor: Phasecor convert the two input images to Fourier space, calculate phase-correlation, back to real space.
### header
### mosaicing
- Merge: Merge joins two images left-right (with ref on the left)  or up-down (with ref above) with a smooth seam.
- Mosaic: Mosaic joins two images left-right (with ref on the left) or top-bottom (with ref above) given an approximate overlap.
- Mosaic1: Mosaic1 This operation joins two images top-bottom (with sec on the right) or left-right (with sec at the bottom) given an approximate pair of tie-points.
- Match: Match scale, rotate and translate sec so that the tie-points line up.
- GlobalBalance: GlobalBalance can be used to remove contrast differences in an assembled mosaic.
- ReMosaic: ReMosaic takes apart the mosaiced image in and rebuilds it, substituting images.
### resample
- Resize:  Resize image
- Rotate: Rotate image
- Thumbnail: Thumbnail Make a thumbnail from a file
- ThumbnailBuffer: ThumbnailBuffer Make a thumbnail from []byte
- ThumbnailImage: ThumbnailImage Make a thumbnail from *Image
### vip
- DefaultInit: DefaultInit init vips by default
- Init: Init starts up libvips
- Shutdown: Shutdown drop caches and close plugins. Run with "--vips-leak" to do a leak check too.
- LeakSet: LeakSet turn on or off vips leak checking
- VersionString: VersionString get the VIPS version as a static string, including a build date and time.
- Version: Version get the major, minor or micro library version, with flag values 0, 1 and 2.
- Error: Error Get a pointer to the start of the error buffer as a C string.
### vipsimage
- New: New empty VipsImage
- NewMemory: NewMemory creates a new VipsImage.
- NewFromFile: NewFromFile opens filename for reading.
- NewFromBuffer: NewFromBuffer loads an image from the formatted area of memory buf.
- Height: Height return image height
- Width:  Width return image width
- ProgressSet: ProgressSet ff set, vips will print messages about the progress of computation to stdout.
- ImageSetProgress: ImageSetProgress vips signals evaluation progress via the “preeval”

## Contributing
We appreciate your help!