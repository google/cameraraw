// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wb2000

import "math/big"

const Make = "SAMSUNG"
const Model = "WB2000"

type IFD0 struct {
	Make        string   `tiff:"field,tag=271"`
	Model       string   `tiff:"field,tag=272"`
	Orientation uint16   `tiff:"field,tag=274"`
	DateTime    string   `tiff:"field,tag=306"`
	ExifIFD     uint32   `tiff:"field,tag=34665"`
	SubIFDs     []uint32 `tiff:"field,tag=330"`
}

type IFD0SubIFD0 struct {
	NewSubfileType              uint32   `tiff:"field,tag=254"`
	XResolution                 *big.Rat `tiff:"field,tag=282"`
	YResolution                 *big.Rat `tiff:"field,tag=283"`
	ResolutionUnit              uint16   `tiff:"field,tag=296"`
	JPEGInterchangeFormat       uint32   `tiff:"field,tag=513"`
	JPEGInterchangeFormatLength uint32   `tiff:"field,tag=514"`
	YCbCrPositioning            uint16   `tiff:"field,tag=531"`
}

type IFD0SubIFD1 struct {
	NewSubfileType      uint32   `tiff:"field,tag=254"`
	ImageWidth          uint32   `tiff:"field,tag=256"`
	ImageLength         uint32   `tiff:"field,tag=257"`
	BitsPerSample       uint16   `tiff:"field,tag=258"`
	Compression         uint32   `tiff:"field,tag=259"`
	StripOffsets        uint32   `tiff:"field,tag=273"`
	SamplesPerPixel     uint16   `tiff:"field,tag=277"`
	RowsPerStrip        uint32   `tiff:"field,tag=278"`
	StripByteCounts     uint32   `tiff:"field,tag=279"`
	CFARepeatPatternDim []uint16 `tiff:"field,tag=33421"`
	CFAPattern          []byte   `tiff:"field,tag=33422"`
}

type IFD0ExifIFD struct {
	ExposureTime          *big.Rat `tiff:"field,tag=33434"`
	FNumber               *big.Rat `tiff:"field,tag=33437"`
	ExposureProgram       uint16   `tiff:"field,tag=34850"`
	ISOSpeedRatings       uint16   `tiff:"field,tag=34855"`
	ExifVersion           []byte   `tiff:"field,tag=36864"`
	DateTimeOriginal      string   `tiff:"field,tag=36867"`
	ExposureBiasValue     *big.Rat `tiff:"field,tag=37380"`
	MeteringMode          uint16   `tiff:"field,tag=37383"`
	LightSource           uint16   `tiff:"field,tag=37384"`
	Flash                 uint16   `tiff:"field,tag=37385"`
	FocalLength           *big.Rat `tiff:"field,tag=37386"`
	MakerNote             []byte   `tiff:"field,tag=37500"`
	ColorSpace            uint16   `tiff:"field,tag=40961"`
	PixelXDimension       uint32   `tiff:"field,tag=40962"`
	PixelYDimension       uint32   `tiff:"field,tag=40963"`
	ExposureMode          uint16   `tiff:"field,tag=41986"`
	WhiteBalance          uint16   `tiff:"field,tag=41987"`
	DigitalZoomRatio      *big.Rat `tiff:"field,tag=41988"`
	FocalLengthIn35mmFilm uint16   `tiff:"field,tag=41989"`
	SceneCaptureType      uint16   `tiff:"field,tag=41990"`
	Contrast              uint16   `tiff:"field,tag=41992"`
	Saturation            uint16   `tiff:"field,tag=41993"`
	Sharpness             uint16   `tiff:"field,tag=41994"`
}

type IFD1 struct {
	SubIFDs uint32 `tiff:"field,tag=330"`
}

type IFD1SubIFD0 struct {
	NewSubfileType              uint32   `tiff:"field,tag=254"`
	XResolution                 *big.Rat `tiff:"field,tag=282"`
	YResolution                 *big.Rat `tiff:"field,tag=283"`
	ResolutionUnit              uint16   `tiff:"field,tag=296"`
	JPEGInterchangeFormat       uint32   `tiff:"field,tag=513"`
	JPEGInterchangeFormatLength uint32   `tiff:"field,tag=514"`
	YCbCrPositioning            uint16   `tiff:"field,tag=531"`
}
