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

package nx100

import "math/big"

const Make = "SAMSUNG"
const Model = "NX100"

type IFD0 struct {
	Make        string   `tiff:"field,tag=271"`
	Model       string   `tiff:"field,tag=272"`
	Orientation uint16   `tiff:"field,tag=274"`
	DateTime    string   `tiff:"field,tag=306"`
	SubIFDs     []uint32 `tiff:"field,tag=330"`
	ExifIFD     uint32   `tiff:"field,tag=34665"`
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
	ImageWidth          uint16   `tiff:"field,tag=256"`
	ImageLength         uint16   `tiff:"field,tag=257"`
	BitsPerSample       uint16   `tiff:"field,tag=258"`
	Compression         uint16   `tiff:"field,tag=259"`
	StripOffsets        uint32   `tiff:"field,tag=273"`
	SamplesPerPixel     uint16   `tiff:"field,tag=277"`
	RowsPerStrip        uint32   `tiff:"field,tag=278"`
	StripByteCounts     uint32   `tiff:"field,tag=279"`
	CFARepeatPatternDim []uint16 `tiff:"field,tag=33421"`
	CFAPattern          []byte   `tiff:"field,tag=33422"`
}

type IFD0ExifIFD struct {
	ExposureTime      *big.Rat `tiff:"field,tag=33434"`
	FNumber           *big.Rat `tiff:"field,tag=33437"`
	ExposureProgram   uint16   `tiff:"field,tag=34850"`
	ISOSpeedRatings   uint16   `tiff:"field,tag=34855"`
	ExifVersion       []byte   `tiff:"field,tag=36864"`
	DateTimeOriginal  string   `tiff:"field,tag=36867"`
	ExposureBiasValue *big.Rat `tiff:"field,tag=37380"`
	MeteringMode      uint16   `tiff:"field,tag=37383"`
	LightSource       uint16   `tiff:"field,tag=37384"`
	Flash             uint16   `tiff:"field,tag=37385"`
	FocalLength       *big.Rat `tiff:"field,tag=37386"`
	MakerNote         []byte   `tiff:"field,tag=37500"`
	FlashpixVersion   []byte   `tiff:"field,tag=40960"`
	ColorSpace        uint16   `tiff:"field,tag=40961"`
	PixelXDimension   uint32   `tiff:"field,tag=40962"`
	PixelYDimension   uint32   `tiff:"field,tag=40963"`
	CFAPattern        []byte   `tiff:"field,tag=41730"`
	WhiteBalance      uint16   `tiff:"field,tag=41987"`
	SceneCaptureType  uint16   `tiff:"field,tag=41990"`
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
