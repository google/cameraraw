/*
Copyright 2016 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nikon

import (
	"image"
	"math/big"
)

type RawIFD struct {
	// IFD Entries (18 tags)
	NewSubfileType            uint32    `tiff:"field,tag=254"`
	ImageWidth                uint32    `tiff:"field,tag=256"`
	ImageLength               uint32    `tiff:"field,tag=257"`
	BitsPerSample             uint16    `tiff:"field,tag=258"`
	Compression               uint16    `tiff:"field,tag=259"`
	PhotometricInterpretation uint16    `tiff:"field,tag=262"`
	StripOffsets              []uint32  `tiff:"field,tag=273"`
	Orientation               uint16    `tiff:"field,tag=274"` // Not always present
	SamplesPerPixel           uint16    `tiff:"field,tag=277"`
	RowsPerStrip              uint32    `tiff:"field,tag=278"`
	StripByteCounts           []uint32  `tiff:"field,tag=279"`
	XResolution               *big.Rat  `tiff:"field,tag=282"`
	YResolution               *big.Rat  `tiff:"field,tag=283"`
	PlanarConfiguration       uint16    `tiff:"field,tag=284"`
	ResolutionUnit            uint16    `tiff:"field,tag=296"`
	CFARepeatPatternDim       [2]uint16 `tiff:"field,tag=33421"`
	CFAPattern                [4]byte   `tiff:"field,tag=33422"`
	SensingMethod             uint16    `tiff:"field,tag=37399"`

	// The image that this IFD describes.
	Img image.Image
}

type JpegIFD struct {
	// IFD Entries (8 tags)
	NewSubfileType              uint32   `tiff:"field,tag=254"`
	Compression                 uint16   `tiff:"field,tag=259"`
	XResolution                 *big.Rat `tiff:"field,tag=282"`
	YResolution                 *big.Rat `tiff:"field,tag=283"`
	ResolutionUnit              uint16   `tiff:"field,tag=296"`
	JPEGInterchangeFormat       uint32   `tiff:"field,tag=513"`
	JPEGInterchangeFormatLength uint32   `tiff:"field,tag=514"`
	YCbCrPositioning            uint16   `tiff:"field,tag=531"`

	// The image that this IFD describes.
	Img image.Image
}
