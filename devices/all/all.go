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

// Package all is an import only package meant to make it easy to import and
// support multiple camera manufacturers and models without having to import a
// bunch of other packages only for their import side effects.
// The idea is that you would add this to the imports section and nothing more.
//   _ "github.com/jonathanpittman/cameraraw/devices/all"
package all

import (
	_ "github.com/jonathanpittman/cameraraw/devices/canon/all"
	_ "github.com/jonathanpittman/cameraraw/devices/epson/all"
	_ "github.com/jonathanpittman/cameraraw/devices/fujifilm/all"
	_ "github.com/jonathanpittman/cameraraw/devices/hasselblad/all"
	_ "github.com/jonathanpittman/cameraraw/devices/kodak/all"
	_ "github.com/jonathanpittman/cameraraw/devices/leaf/all"
	_ "github.com/jonathanpittman/cameraraw/devices/leica/all"
	_ "github.com/jonathanpittman/cameraraw/devices/mamiya/all"
	_ "github.com/jonathanpittman/cameraraw/devices/minolta/all"
	_ "github.com/jonathanpittman/cameraraw/devices/nikon/all"
	_ "github.com/jonathanpittman/cameraraw/devices/olympus/all"
	_ "github.com/jonathanpittman/cameraraw/devices/panasonic/all"
	_ "github.com/jonathanpittman/cameraraw/devices/pentax/all"
	_ "github.com/jonathanpittman/cameraraw/devices/phaseone/all"
	_ "github.com/jonathanpittman/cameraraw/devices/polaroid/all"
	_ "github.com/jonathanpittman/cameraraw/devices/ricoh/all"
	_ "github.com/jonathanpittman/cameraraw/devices/samsung/all"
	_ "github.com/jonathanpittman/cameraraw/devices/sigma/all"
	_ "github.com/jonathanpittman/cameraraw/devices/sony/all"
)
