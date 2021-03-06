// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package canon

import "math/big"

// Cr2IFD0 represents the common structure found in IFD0 for most Canon .CR2 files.
type Cr2IFD0 struct {
	ImageWidth      uint16   `tiff:"field,tag=256"`
	ImageLength     uint16   `tiff:"field,tag=257"`
	BitsPerSample   []uint16 `tiff:"field,tag=258"`
	Compression     uint16   `tiff:"field,tag=259"`
	Make            string   `tiff:"field,tag=271"`
	Model           string   `tiff:"field,tag=272"`
	StripOffsets    uint32   `tiff:"field,tag=273"`
	Orientation     uint16   `tiff:"field,tag=274"`
	StripByteCounts uint32   `tiff:"field,tag=279"`
	XResolution     *big.Rat `tiff:"field,tag=282"`
	YResolution     *big.Rat `tiff:"field,tag=283"`
	ResolutionUnit  uint16   `tiff:"field,tag=296"`
	DateTime        string   `tiff:"field,tag=306"`
	ExifIFD         uint32   `tiff:"field,tag=34665"`
	GPSIFD          uint32   `tiff:"field,tag=34853"`
}

// Cr2IFD1 represents the common structure found in IFD1 for most Canon .CR2 files.
type Cr2IFD1 struct {
	JPEGInterchangeFormat       uint32 `tiff:"field,tag=513"`
	JPEGInterchangeFormatLength uint32 `tiff:"field,tag=514"`
}
