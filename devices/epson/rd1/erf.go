// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rd1

import "math/big"

const Make = "SEIKO EPSON CORP."
const Model = "R-D1"

type IFD0 struct {
	NewSubfileType            uint32   `tiff:"field,tag=254"`
	ImageWidth                uint32   `tiff:"field,tag=256"`
	ImageLength               uint32   `tiff:"field,tag=257"`
	BitsPerSample             []uint16 `tiff:"field,tag=258"`
	Compression               uint16   `tiff:"field,tag=259"`
	PhotometricInterpretation uint16   `tiff:"field,tag=262"`
	ImageDescription          string   `tiff:"field,tag=270"`
	Make                      string   `tiff:"field,tag=271"`
	Model                     string   `tiff:"field,tag=272"`
	StripOffsets              uint32   `tiff:"field,tag=273"`
	Orientation               uint16   `tiff:"field,tag=274"`
	SamplesPerPixel           uint16   `tiff:"field,tag=277"`
	RowsPerStrip              uint32   `tiff:"field,tag=278"`
	StripByteCounts           uint32   `tiff:"field,tag=279"`
	XResolution               *big.Rat `tiff:"field,tag=282"`
	YResolution               *big.Rat `tiff:"field,tag=283"`
	PlanarConfiguration       uint16   `tiff:"field,tag=284"`
	ResolutionUnit            uint16   `tiff:"field,tag=296"`
	Software                  string   `tiff:"field,tag=305"`
	DateTime                  string   `tiff:"field,tag=306"`
	SubIFDs                   uint32   `tiff:"field,tag=330"`
	ExifIFD                   uint32   `tiff:"field,tag=34665"`
	DateTimeOriginal          string   `tiff:"field,tag=36867"`
	TIFF_EPStandardID         []byte   `tiff:"field,tag=37398"`
}

type IFD0SubIFD0 struct {
	NewSubfileType            uint32   `tiff:"field,tag=254"`
	ImageWidth                uint32   `tiff:"field,tag=256"`
	ImageLength               uint32   `tiff:"field,tag=257"`
	BitsPerSample             uint16   `tiff:"field,tag=258"`
	Compression               uint16   `tiff:"field,tag=259"`
	PhotometricInterpretation uint16   `tiff:"field,tag=262"`
	StripOffsets              uint32   `tiff:"field,tag=273"`
	SamplesPerPixel           uint16   `tiff:"field,tag=277"`
	RowsPerStrip              uint32   `tiff:"field,tag=278"`
	StripByteCounts           uint32   `tiff:"field,tag=279"`
	XResolution               *big.Rat `tiff:"field,tag=282"`
	YResolution               *big.Rat `tiff:"field,tag=283"`
	PlanarConfiguration       uint16   `tiff:"field,tag=284"`
	ResolutionUnit            uint16   `tiff:"field,tag=296"`
	CFARepeatPatternDim       []uint16 `tiff:"field,tag=33421"`
	CFAPattern                []byte   `tiff:"field,tag=33422"`
}

type IFD0ExifIFD struct {
	ExposureTime      *big.Rat `tiff:"field,tag=33434"`
	ExposureProgram   uint16   `tiff:"field,tag=34850"`
	ISOSpeedRatings   uint16   `tiff:"field,tag=34855"`
	ExifVersion       []byte   `tiff:"field,tag=36864"`
	DateTimeOriginal  string   `tiff:"field,tag=36867"`
	DateTimeDigitized string   `tiff:"field,tag=36868"`
	ExposureBiasValue *big.Rat `tiff:"field,tag=37380"`
	MeteringMode      uint16   `tiff:"field,tag=37383"`
	LightSource       uint16   `tiff:"field,tag=37384"`
	Flash             uint16   `tiff:"field,tag=37385"`
	MakerNote         []byte   `tiff:"field,tag=37500"`
	ColorSpace        uint16   `tiff:"field,tag=40961"`
	FileSource        byte     `tiff:"field,tag=41728"`
	SceneType         byte     `tiff:"field,tag=41729"`
	CustomRendered    uint16   `tiff:"field,tag=41985"`
	ExposureMode      uint16   `tiff:"field,tag=41986"`
	WhiteBalance      uint16   `tiff:"field,tag=41987"`
	SceneCaptureType  uint16   `tiff:"field,tag=41990"`
	GainControl       uint16   `tiff:"field,tag=41991"`
}
