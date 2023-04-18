package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func main() {
	err := toJEPG(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "toJPEG: %v", err)
		os.Exit(1)
	}
}

func toJEPG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
