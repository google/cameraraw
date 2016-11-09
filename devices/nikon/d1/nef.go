// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d1

import "math/big"

type NefIFD0 struct {
	NewSubfileType            uint32     `tiff:"field,tag=254"`
	ImageWidth                uint32     `tiff:"field,tag=256"`
	ImageLength               uint32     `tiff:"field,tag=257"`
	BitsPerSample             []uint16   `tiff:"field,tag=258"`
	Compression               uint16     `tiff:"field,tag=259"`
	PhotometricInterpretation uint16     `tiff:"field,tag=262"`
	ImageDescription          string     `tiff:"field,tag=270"`
	Make                      string     `tiff:"field,tag=271"`
	Model                     string     `tiff:"field,tag=272"`
	StripOffsets              uint32     `tiff:"field,tag=273"`
	SamplesPerPixel           uint16     `tiff:"field,tag=277"`
	RowsPerStrip              uint32     `tiff:"field,tag=278"`
	StripByteCounts           uint32     `tiff:"field,tag=279"`
	XResolution               *big.Rat   `tiff:"field,tag=282"`
	YResolution               *big.Rat   `tiff:"field,tag=283"`
	PlanarConfiguration       uint16     `tiff:"field,tag=284"`
	ResolutionUnit            uint16     `tiff:"field,tag=296"`
	Software                  string     `tiff:"field,tag=305"`
	DateTime                  string     `tiff:"field,tag=306"`
	SubIFDs                   uint32     `tiff:"field,tag=330"`
	ReferenceBlackWhite       []*big.Rat `tiff:"field,tag=532"`
	Copyright                 string     `tiff:"field,tag=33432"`
	ExifIFD                   uint32     `tiff:"field,tag=34665"`
	DateTimeOriginal          string     `tiff:"field,tag=36867"`
	TIFFEPStandardID          []byte     `tiff:"field,tag=37398"`
}

type NefIFD0SubIFD0 struct {
	NewSubfileType            uint32   `tiff:"field,tag=254"`
	ImageWidth                uint32   `tiff:"field,tag=256"`
	ImageLength               uint32   `tiff:"field,tag=257"`
	BitsPerSample             uint16   `tiff:"field,tag=258"`
	Compression               uint16   `tiff:"field,tag=259"`
	PhotometricInterpretation uint16   `tiff:"field,tag=262"`
	StripOffsets              []uint32 `tiff:"field,tag=273"`
	SamplesPerPixel           uint16   `tiff:"field,tag=277"`
	RowsPerStrip              uint32   `tiff:"field,tag=278"`
	StripByteCounts           []uint32 `tiff:"field,tag=279"`
	XResolution               *big.Rat `tiff:"field,tag=282"`
	YResolution               *big.Rat `tiff:"field,tag=283"`
	PlanarConfiguration       uint16   `tiff:"field,tag=284"`
	ResolutionUnit            uint16   `tiff:"field,tag=296"`
	CFARepeatPatternDim       []uint16 `tiff:"field,tag=33421"`
	CFAPattern                []byte   `tiff:"field,tag=33422"`
	SensingMethod             uint16   `tiff:"field,tag=37399"`
}

type NefIFD0ExifIFD struct {
	ExposureTime        *big.Rat `tiff:"field,tag=33434"`
	FNumber             *big.Rat `tiff:"field,tag=33437"`
	ExposureProgram     uint16   `tiff:"field,tag=34850"`
	DateTimeOriginal    string   `tiff:"field,tag=36867"`
	DateTimeDigitized   string   `tiff:"field,tag=36868"`
	ExposureBiasValue   *big.Rat `tiff:"field,tag=37380"`
	MaxApertureValue    *big.Rat `tiff:"field,tag=37381"`
	MeteringMode        uint16   `tiff:"field,tag=37383"`
	FocalLength         *big.Rat `tiff:"field,tag=37386"`
	MakerNote           []byte   `tiff:"field,tag=37500"`
	UserComment         []byte   `tiff:"field,tag=37510"`
	SubsecTime          string   `tiff:"field,tag=37520"`
	SubsecTimeOriginal  string   `tiff:"field,tag=37521"`
	SubsecTimeDigitized string   `tiff:"field,tag=37522"`
	SensingMethod       uint16   `tiff:"field,tag=41495"`
	FileSource          byte     `tiff:"field,tag=41728"`
	SceneType           byte     `tiff:"field,tag=41729"`
	CFAPattern          []byte   `tiff:"field,tag=41730"`
}
