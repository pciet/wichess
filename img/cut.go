package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// The cut program changes an input image into
// 64 equally sized separate images.
func main() {
	if len(os.Args) != 4 {
		fmt.Println("need file input, folder output, and piece name arguments")
		os.Exit(1)
	}

	img, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(os.Args[1], err)
		os.Exit(1)
	}

	defer img.Close()

	board, format, err := image.Decode(img)
	if err != nil {
		fmt.Println(os.Args[1], err)
		os.Exit(1)
	}
	if format != "png" {
		fmt.Println(os.Args[1], "not png")
		os.Exit(1)
	}

	bounds := board.Bounds()
	if (bounds.Min.X != 0) || (bounds.Min.Y != 0) || (bounds.Max.X != bounds.Max.Y) {
		fmt.Println(os.Args[1], "bad bounds", bounds)
		os.Exit(1)
	}

	if board.ColorModel() != color.RGBAModel {
		fmt.Println(os.Args[1], "color model", board.ColorModel(), "not", color.RGBAModel)
		os.Exit(1)
	}

	if bounds.Max.X%8 != 0 {
		fmt.Println(os.Args[1], "dim not evenly divisible by 8")
		os.Exit(1)
	}

	rgbaBoard := board.(*image.RGBA)

	stride := bounds.Max.X / 8

	dir := os.Args[2]
	name := os.Args[3]

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			xaddr := x * stride
			yaddr := y * stride
			rect := image.Rectangle{
				image.Point{xaddr, yaddr},
				image.Point{xaddr + stride, yaddr + stride},
			}

			f, err := os.Create(
				fmt.Sprintf("%s/%s_%d_%d.png", dir, name, x, 7-y))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = png.Encode(f, rgbaBoard.SubImage(rect))
			if err != nil {
				f.Close()
				fmt.Println(err)
				os.Exit(1)
			}

			err = f.Close()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
