// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package all is an import only package meant to make it easy to import and
// support multiple camera manufacturers and models without having to import a
// bunch of other packages only for their import side effects.
// The idea is that you would add this to the imports section and nothing more.
//   _ "github.com/google/cameraraw/devices/all"
package all

import (
	_ "github.com/google/cameraraw/devices/canon/all"
	_ "github.com/google/cameraraw/devices/epson/all"
	_ "github.com/google/cameraraw/devices/fujifilm/all"
	_ "github.com/google/cameraraw/devices/hasselblad/all"
	_ "github.com/google/cameraraw/devices/kodak/all"
	_ "github.com/google/cameraraw/devices/leaf/all"
	_ "github.com/google/cameraraw/devices/leica/all"
	_ "github.com/google/cameraraw/devices/mamiya/all"
	_ "github.com/google/cameraraw/devices/minolta/all"
	_ "github.com/google/cameraraw/devices/nikon/all"
	_ "github.com/google/cameraraw/devices/olympus/all"
	_ "github.com/google/cameraraw/devices/panasonic/all"
	_ "github.com/google/cameraraw/devices/pentax/all"
	_ "github.com/google/cameraraw/devices/phaseone/all"
	_ "github.com/google/cameraraw/devices/polaroid/all"
	_ "github.com/google/cameraraw/devices/ricoh/all"
	_ "github.com/google/cameraraw/devices/samsung/all"
	_ "github.com/google/cameraraw/devices/sigma/all"
	_ "github.com/google/cameraraw/devices/sony/all"
)
