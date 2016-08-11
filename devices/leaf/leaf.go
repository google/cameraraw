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
