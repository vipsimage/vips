package main

import (
	"fmt"

	"github.com/vipsimage/vips"
)

func testMerge() {
	img1, _ := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_2.jpg")
	img2, _ := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_6.jpg")

	err := img2.AutoRot()
	if err != nil {
		panic(err)
	}

	err = img1.Merge(img2, vips.DirectionHorizontal, 1800, 0)
	if err != nil {
		panic(err)
	}

	err = img1.Save2file("merge.jpg")
	if err != nil {
		panic(err)
	}
}
func testMosaic() {
	img1, _ := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_2.jpg")
	img2, _ := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_6.jpg")

	err := img2.AutoRot()
	if err != nil {
		panic(err)
	}
	fmt.Println(img1.Width(), img1.Height())

	err = img1.Mosaic(img2, vips.DirectionHorizontal, 0, 0, img1.Width()-100, 0)
	if err != nil {
		panic(err)
	}

	err = img1.Save2file("merge.jpg")
	if err != nil {
		panic(err)
	}
}
func testMosaic1() {
	img1, _ := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_2.jpg")
	img2, _ := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_6.jpg")

	err := img2.AutoRot()
	if err != nil {
		panic(err)
	}

	err = img1.Mosaic1(img2, vips.DirectionHorizontal, 0, 0, 0, 0, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}

	err = img1.Save2file("merge.jpg")
	if err != nil {
		panic(err)
	}
}
func testGlobalBalance() {
	img1, _ := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/photo.jpg")
	img2, _ := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_6.jpg")

	err := img2.AutoRot()
	if err != nil {
		panic(err)
	}

	err = img1.Merge(img2, vips.DirectionHorizontal, 900, 0)
	if err != nil {
		panic(err)
	}

	err = img1.Save2file("testGlobalBalance.jpg")
	if err != nil {
		panic(err)
	}

	err = img1.ReMosaic("old", "new")
	if err != nil {
		panic(err)
	}
}
func testWEPB() {
	vips.ProgressSet(true)

	out, err := vips.WEBPLoad("/Users/mo/Go/vipsimage/vips/examples/images/away.webp")
	if err != nil {
		panic(err)
	}

	// out.ImageSetProgress(true)

	// err = out.Save2file("test.webp")
	err = out.JPEGSave("test-webp.jpg")
	if err != nil {
		panic(err)
	}

	// buf, err := ioutil.ReadFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_6.jpg")
	// if err != nil {
	// 	panic(err)
	// }
}
func testFull() {
	var err error

	out, err := vips.JPEGLoad("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_6.jpg")
	if err != nil {
		panic(err)
	}

	out.AutoRot()

	ink := []float64{255, 255, 255}

	out.DrawRect(ink, 0, 0, 300, 300)

	out.Save2file("test-load.jpg")
	return
	// return

	// testGlobalBalance()
	// return

	// name := vips.ForeignFindLoad("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_6.jpg")
	// fmt.Println(name)
	// return

	// xyz, err := vips.Zone(100, 100)

	// xyz, err := vips.Identity()
	// xyz, err := vips.Grey(200, 200)
	// xyz, err := vips.Perlin(800, 800)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// xyz.Save2file("xyz.png")
	// return

	// black, err := vips.Black(100, 100)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// black.Save2file("black.png")

	// text, err := vips.Text("you have a message")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// text.Save2file("text.png")

	// exif := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_6.jpg")
	//
	// fmt.Println(exif.GetAngle())
	// err = exif.AutoRot()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// exif.Save2file("exif-end.png")

	// buf, err := ioutil.ReadFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_6.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// out, err := vips.JPEGLoadBuffer(buf)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// out.AutoRot()

	// out.Save2file("JPEGLoadBuffer.jpg")

	// b, z, err := out.JPEGSaveBuffer()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println(len(b), z, len(b) == z)
	// return

	// err = out.JPEGSaveMime()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// return

	// name := vips.ForeignFindLoadBuffer(buf)
	// fmt.Println(name)

	// out, err := vips.Thumbnail("/Users/mo/Go/vipsimage/vips/examples/images/photo.jpg", 300)
	// out, err := vips.ThumbnailBuffer(buf, 600)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// out.Save2file("test.png")

	// img := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/photo.jpg")
	// wm := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/watermark.png")

	// err = wm.ThumbnailImage(200)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// err = wm.Save2file("wm.png")
	// if err != nil {
	// 	panic(err)
	// }

	// err = wm.Resize(float64(img.Width()) / float64(wm.Width()))
	// if err != nil {
	// 	panic(err)
	// }
	//
	// err = wm.Rotate(30)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// err = wm.Replicate(img.Width()/wm.Width()+1, img.Height()/wm.Height()+1)
	// if err != nil {
	// 	panic(err)
	// }
	// img.Replicate(10, 10)

	// fmt.Println(img.Height(), img.Width())

	// err = img.Gravity(vips.DirectionSouthEast, 2000, 2000)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.Composite2(wm, vips.BlendModeOver)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.Embed(250, 250, 4500, 3500)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.ExtractArea(250, 250, 1500, 2000)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.Crop(1000, 0, img.Width()/10, img.Height())
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.Grid(300, 10, 1000)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.Scale()
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.SubSample(5, 5)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.Zoom(3, 3)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.Wrap()
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.ExtractBand(1)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.SmartCrop(2000, 2500)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.Flip(vips.DirectionVertical)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.InvertLut()
	// if err != nil {
	// 	panic(err)
	// }

	// ink := []float64{255, 255, 255}
	// for i := 1; i < 100; i++ {
	// 	err = img.DrawPoint(ink, 100+i, 100+i)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// mk := vips.NewFromFile("/Users/mo/Go/vipsimage/vips/examples/images/Landscape_6.jpg")
	// err = mk.AutoRot()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// ink := []float64{255, 255, 255}
	// err = img.DrawMask1(255, mk, 500, 500)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// err = img.DrawLine1(ink[0], 100, 100, 1000, 1000)
	// if err != nil {
	// 	panic(err)
	// }
	// img.DrawLine1(ink[0], 101, 100, 1001, 1000)
	// img.DrawLine1(ink[0], 102, 100, 1002, 1000)

	// img.DrawSmudge(100, 100, 100, 100)

	// img.FormatSizeof()
	//
	// err = img.Save2file("save.jpeg")
	// if err != nil {
	// 	panic(err)
	// }
}

func main() {
	img, _ := vips.NewFromFile("./images/Landscape_2.jpg")
	wm, _ := vips.NewFromFile("./images/away.webp")

	err := wm.ThumbnailImage(200)
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
	img.Replicate(2, 2)

	fmt.Println(img.Height(), img.Width())

	imgCopy, err := img.Copy()
	if err != nil {
		panic(err)
	}
	fmt.Println(imgCopy.Width())

	err = img.Save2file("out.png")
	if err != nil {
		panic(err)
	}
}
