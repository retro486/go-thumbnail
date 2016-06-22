// Port of http://members.shaw.ca/el.supremo/MagickWand/resize.htm to Go
package main

import (
    "fmt"
    "flag"
	"gopkg.in/gographics/imagick.v2/imagick"
)

var input string
var output string

func init() {
    flag.StringVar(&input, "input", "", "original file path and name (absolute)")
    flag.StringVar(&output, "output", "", "output file path and name (absolute)")
}

func main() {
    flag.Parse()
    if input == "" || output == "" {
        panic("Must specify an input and an output")
    }

    fmt.Printf("Procesing given input: %s, given output: %s ...\n", input, output)

	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	err = mw.ReadImage(input)
	if err != nil {
		panic(err)
	}

	// Get original logo size
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	// Calculate half the size
	hWidth := uint(width / 2)
	hHeight := uint(height / 2)

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		panic(err)
	}

	// Set the compression quality to 95 (high quality = low compression)
	err = mw.SetImageCompressionQuality(95)
	if err != nil {
		panic(err)
	}

	err = mw.WriteImage(output)
	if err != nil {
		panic(err)
	}
}
