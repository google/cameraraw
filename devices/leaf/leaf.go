// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leaf

import "math/big"

type IFD0 struct {
	NewSubfileType            uint32   `tiff:"field,tag=254"`
	ImageWidth                uint32   `tiff:"field,tag=256"`
	ImageLength               uint32   `tiff:"field,tag=257"`
	BitsPerSample             uint16   `tiff:"field,tag=258"`
	Compression               uint16   `tiff:"field,tag=259"`
	PhotometricInterpretation uint16   `tiff:"field,tag=262"`
	SamplesPerPixel           uint16   `tiff:"field,tag=277"`
	XResolution               *big.Rat `tiff:"field,tag=282"`
	YResolution               *big.Rat `tiff:"field,tag=283"`
	PlanarConfiguration       uint16   `tiff:"field,tag=284"`
	ResolutionUnit            uint16   `tiff:"field,tag=296"`
	Software                  string   `tiff:"field,tag=305"`
	TileWidth                 uint16   `tiff:"field,tag=322"`
	TileLength                uint16   `tiff:"field,tag=323"`
	TileOffsets               uint32   `tiff:"field,tag=324"`
	TileByteCounts            uint32   `tiff:"field,tag=325"`
	XMP                       []byte   `tiff:"field,tag=700"`
	IPTC                      []byte   `tiff:"field,tag=33723"`
	UNKNOWN_TAG_34310         string   `tiff:"field,tag=34310"`
}
