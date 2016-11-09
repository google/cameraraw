// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eos1d

import "math/big"

type IFD0 struct {
	NewSubfileType            uint32   `tiff:"field,tag=254"`
	ImageWidth                uint16   `tiff:"field,tag=256"`
	ImageLength               uint16   `tiff:"field,tag=257"`
	BitsPerSample             []uint16 `tiff:"field,tag=258"`
	Compression               uint16   `tiff:"field,tag=259"`
	PhotometricInterpretation uint16   `tiff:"field,tag=262"`
	Make                      string   `tiff:"field,tag=271"`
	Model                     string   `tiff:"field,tag=272"`
	StripOffsets              uint32   `tiff:"field,tag=273"`
	Orientation               uint16   `tiff:"field,tag=274"`
	SamplesPerPixel           uint16   `tiff:"field,tag=277"`
	RowsPerStrip              uint16   `tiff:"field,tag=278"`
	StripByteCounts           uint32   `tiff:"field,tag=279"`
	XResolution               *big.Rat `tiff:"field,tag=282"`
	YResolution               *big.Rat `tiff:"field,tag=283"`
	PlanarConfiguration       uint16   `tiff:"field,tag=284"`
	ResolutionUnit            uint16   `tiff:"field,tag=296"`
	DateTime                  string   `tiff:"field,tag=306"`
	IPTC                      []byte   `tiff:"field,tag=33723"`
	ExifIFD                   uint32   `tiff:"field,tag=34665"`
}

type IFD0ExifIFD struct {
	ExposureTime             *big.Rat `tiff:"field,tag=33434"`
	FNumber                  *big.Rat `tiff:"field,tag=33437"`
	ExposureProgram          uint16   `tiff:"field,tag=34850"`
	ISOSpeedRatings          uint16   `tiff:"field,tag=34855"`
	DateTimeOriginal         string   `tiff:"field,tag=36867"`
	DateTimeDigitized        string   `tiff:"field,tag=36868"`
	ShutterSpeedValue        *big.Rat `tiff:"field,tag=37377"`
	ApertureValue            *big.Rat `tiff:"field,tag=37378"`
	ExposureBiasValue        *big.Rat `tiff:"field,tag=37380"`
	MeteringMode             uint16   `tiff:"field,tag=37383"`
	Flash                    uint16   `tiff:"field,tag=37385"`
	FocalLength              *big.Rat `tiff:"field,tag=37386"`
	MakerNote                []byte   `tiff:"field,tag=37500"`
	UserComment              []byte   `tiff:"field,tag=37510"`
	ColorSpace               uint16   `tiff:"field,tag=40961"`
	PixelXDimension          uint16   `tiff:"field,tag=40962"`
	PixelYDimension          uint16   `tiff:"field,tag=40963"`
	FocalPlaneXResolution    *big.Rat `tiff:"field,tag=41486"`
	FocalPlaneYResolution    *big.Rat `tiff:"field,tag=41487"`
	FocalPlaneResolutionUnit uint16   `tiff:"field,tag=41488"`
	CustomRendered           uint16   `tiff:"field,tag=41985"`
	ExposureMode             uint16   `tiff:"field,tag=41986"`
	WhiteBalance             uint16   `tiff:"field,tag=41987"`
	SceneCaptureType         uint16   `tiff:"field,tag=41990"`
}

type IFD1 struct {
	ImageWidth                uint16   `tiff:"field,tag=256"`
	ImageLength               uint16   `tiff:"field,tag=257"`
	BitsPerSample             []uint16 `tiff:"field,tag=258"`
	Compression               uint16   `tiff:"field,tag=259"`
	PhotometricInterpretation uint16   `tiff:"field,tag=262"`
	StripOffsets              uint32   `tiff:"field,tag=273"`
	SamplesPerPixel           uint16   `tiff:"field,tag=277"`
	RowsPerStrip              uint16   `tiff:"field,tag=278"`
	StripByteCounts           uint32   `tiff:"field,tag=279"`
	XResolution               *big.Rat `tiff:"field,tag=282"`
	YResolution               *big.Rat `tiff:"field,tag=283"`
	ResolutionUnit            uint16   `tiff:"field,tag=296"`
}
