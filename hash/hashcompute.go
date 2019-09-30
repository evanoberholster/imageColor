// Copyright 2017 The goimagehash Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Altered 2019
// by Evan Oberholster

package hash

import (
	"errors"
	"image"
	"math"

	"github.com/corona10/goimagehash/transforms"
)

// Errors
const (
	ErrorImageToLarge    = "Image is wrong size"
	ErrorNoImage         = "Image object can not be nil"
	ErrorImageWrongSize  = "Image needs to support 64bit or 256bit Hashing"
	ErrorImageDimensions = "Width * Height should be the power of 2"
)

// AverageHash fuction returns a hash computation of average hash.
// Implementation follows
// http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html
// Recommended Image size 8x8
func AverageHash(img image.Image) (*ImageHash, error) {
	if img == nil {
		return nil, errors.New(ErrorNoImage)
	}
	// Check size of image.
	b := img.Bounds()
	if b.Max.X != 8 && b.Max.Y != 8 {
		return nil, errors.New(ErrorImageToLarge)
	}

	ahash := NewImageHash(0, AHash)
	pixels := transforms.Rgb2Gray(img)
	flattens := transforms.FlattenPixels(pixels, b.Max.X, b.Max.Y)
	avg := MeanOfPixels(flattens)

	for idx, p := range flattens {
		if p > avg {
			ahash.leftShiftSet(len(flattens) - idx - 1)
		}
	}

	return ahash, nil
}

// DifferenceHash function returns a hash computation of difference hash.
// Implementation follows
// http://www.hackerfactor.com/blog/?/archives/529-Kind-of-Like-That.html
// Recommended Image size 9x8
func DifferenceHash(img image.Image) (*ImageHash, error) {
	if img == nil {
		return nil, errors.New(ErrorNoImage)
	}
	// Check size of image.
	b := img.Bounds()
	if b.Max.X != 9 || b.Max.Y != 8 {
		return nil, errors.New(ErrorImageToLarge)
	}

	dhash := NewImageHash(0, DHash)
	pixels := transforms.Rgb2Gray(img)
	idx := 0
	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[i])-1; j++ {
			if pixels[i][j] < pixels[i][j+1] {
				dhash.leftShiftSet(64 - idx - 1)
			}
			idx++
		}
	}

	return dhash, nil
}

// PerceptionHash function returns a hash computation of phash.
// Implementation follows
// http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html
// Recommended Image size 64x64
func PerceptionHash(img image.Image) (*ImageHash, error) {
	if img == nil {
		return nil, errors.New(ErrorNoImage)
	}
	// Check size of image.
	b := img.Bounds()
	if b.Max.X != 64 || b.Max.Y != 64 {
		return nil, errors.New(ErrorImageToLarge)
	}

	phash := NewImageHash(0, PHash)
	pixels := transforms.Rgb2Gray(img)
	dct := transforms.DCT2D(pixels, 64, 64)
	flattens := transforms.FlattenPixels(dct, 8, 8)
	median := MedianOfPixels(flattens)

	for idx, p := range flattens {
		if p > median {
			phash.leftShiftSet(len(flattens) - idx - 1)
		}
	}
	return phash, nil
}

// ExtPerceptionHash function returns phash of which the size can be set larger than uint64
// Some variable name refer to https://github.com/JohannesBuchner/imagehash/blob/master/imagehash/__init__.py
// Support 64bits phash (width=8, height=8) and 256bits phash (width=16, height=16)
// Important: width * height should be the power of 2
func ExtPerceptionHash(img image.Image) (*ExtImageHash, error) {
	if img == nil {
		return nil, errors.New(ErrorNoImage)
	}
	// Check size of image.
	b := img.Bounds()
	width := int(math.Sqrt(float64(b.Max.X)))
	imgSize := width * width
	if imgSize <= 0 || imgSize&(imgSize-1) != 0 {
		return nil, errors.New(ErrorImageDimensions)
	}
	var phash []uint64
	pixels := transforms.Rgb2Gray(img)
	dct := transforms.DCT2D(pixels, imgSize, imgSize)
	flattens := transforms.FlattenPixels(dct, width, width)
	median := MedianOfPixels(flattens)

	lenOfUnit := 64
	if imgSize%lenOfUnit == 0 {
		phash = make([]uint64, imgSize/lenOfUnit)
	} else {
		phash = make([]uint64, imgSize/lenOfUnit+1)
	}
	for idx, p := range flattens {
		indexOfArray := idx / lenOfUnit
		indexOfBit := lenOfUnit - idx%lenOfUnit - 1
		if p > median {
			phash[indexOfArray] |= 1 << uint(indexOfBit)
		}
	}
	return NewExtImageHash(phash, PHash, imgSize), nil
}

// ExtAverageHash function returns ahash of which the size can be set larger than uint64
// Support 64bits ahash (width=8, height=8) and 256bits ahash (width=16, height=16)
func ExtAverageHash(img image.Image) (*ExtImageHash, error) {
	if img == nil {
		return nil, errors.New(ErrorNoImage)
	}
	// Check size of image.
	b := img.Bounds()
	width := int(math.Sqrt(float64(b.Max.X)))
	imgSize := width * width
	if imgSize == 64 || imgSize == 256 {

	} else {
		return nil, errors.New(ErrorImageWrongSize)
	}
	var ahash []uint64

	pixels := transforms.Rgb2Gray(img)
	flattens := transforms.FlattenPixels(pixels, width, width)
	avg := MeanOfPixels(flattens)

	lenOfUnit := 64
	if imgSize%lenOfUnit == 0 {
		ahash = make([]uint64, imgSize/lenOfUnit)
	} else {
		ahash = make([]uint64, imgSize/lenOfUnit+1)
	}
	for idx, p := range flattens {
		indexOfArray := idx / lenOfUnit
		indexOfBit := lenOfUnit - idx%lenOfUnit - 1
		if p > avg {
			ahash[indexOfArray] |= 1 << uint(indexOfBit)
		}
	}
	return NewExtImageHash(ahash, AHash, imgSize), nil
}

// ExtDifferenceHash function returns dhash of which the size can be set larger than uint64
// Support 64bits dhash (width=8, height=8) and 256bits dhash (width=16, height=16)
func ExtDifferenceHash(img image.Image) (*ExtImageHash, error) {
	if img == nil {
		return nil, errors.New(ErrorNoImage)
	}
	// Check size of image.
	b := img.Bounds()
	width := int(math.Sqrt(float64(b.Max.X)))
	imgSize := width * width

	var dhash []uint64

	pixels := transforms.Rgb2Gray(img)

	lenOfUnit := 64
	if imgSize%lenOfUnit == 0 {
		dhash = make([]uint64, imgSize/lenOfUnit)
	} else {
		dhash = make([]uint64, imgSize/lenOfUnit+1)
	}
	idx := 0
	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[i])-1; j++ {
			indexOfArray := idx / lenOfUnit
			indexOfBit := lenOfUnit - idx%lenOfUnit - 1
			if pixels[i][j] < pixels[i][j+1] {
				dhash[indexOfArray] |= 1 << uint(indexOfBit)
			}
			idx++
		}
	}
	return NewExtImageHash(dhash, DHash, imgSize), nil
}
