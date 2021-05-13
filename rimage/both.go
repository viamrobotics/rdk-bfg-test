package rimage

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// ReadBothFromBytes reads the given data as an image that contains depth. isAligned
// notifies the reader if the image and depth is already aligned.
func ReadBothFromBytes(allData []byte, isAligned bool) (*ImageWithDepth, error) {
	reader := bufio.NewReader(bytes.NewReader(allData))
	depth, err := ReadDepthMap(reader)
	if err != nil {
		return nil, fmt.Errorf("couldn't read depth map (both): %w", err)
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return &ImageWithDepth{ConvertImage(img), depth, isAligned, nil}, nil
}

// ReadBothFromBytes reads the given file as an image that contains depth. isAligned
// notifies the reader if the image and depth is already aligned.
func ReadBothFromFile(fn string, isAligned bool) (*ImageWithDepth, error) {
	if !strings.HasSuffix(fn, ".both.gz") {
		return nil, errors.New("bad extension")
	}

	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	in, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}

	defer in.Close()

	allData, err := ioutil.ReadAll(in)

	if err != nil {
		return nil, err
	}

	return ReadBothFromBytes(allData, isAligned)
}

// WriteBothToFile writes the image with depth to the given file.
func WriteBothToFile(i *ImageWithDepth, fn string) error {
	if !strings.HasSuffix(fn, ".both.gz") {
		return errors.New("vision.ImageWithDepth WriteTo only supports both.gz")
	}

	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	out := gzip.NewWriter(f)
	defer out.Close()

	err = EncodeBoth(i, out)
	if err != nil {
		return err
	}

	out.Flush()
	return f.Sync()
}

// EncodeBoth writes the image with depth to the given writer.
func EncodeBoth(i *ImageWithDepth, out io.Writer) error {
	err := i.Depth.WriteTo(out)
	if err != nil {
		return err
	}

	err = png.Encode(out, i.Color)
	if err != nil {
		return err
	}

	return nil
}
