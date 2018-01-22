package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	jpgQuality := flag.Int("jpeg", 0, "if provided, produces a jpeg of given quality instead of a png")
	flag.Parse()
	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *jpgQuality > 0 {
		err = jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: *jpgQuality})
	} else {
		err = png.Encode(os.Stdout, img)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
