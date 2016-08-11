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

package nikon

import "jonathanpittman.com/tiff"

var (
	MakerNoteTags     = tiff.NewTagSet("NikonMakerNote", 0, 65535)
	MakerNoteTagSpace = tiff.NewTagSpace("NikonMakerNote")
)

func init() {
	// Tag 0
	MakerNoteTags.Register(tiff.NewTag(1, "Version", nil))
	MakerNoteTags.Register(tiff.NewTag(2, "ISOSpeed", nil))
	MakerNoteTags.Register(tiff.NewTag(3, "ColorMode", nil))
	MakerNoteTags.Register(tiff.NewTag(4, "Quality", nil))
	MakerNoteTags.Register(tiff.NewTag(5, "WhiteBalance", nil))
	MakerNoteTags.Register(tiff.NewTag(6, "Sharpening", nil))
	MakerNoteTags.Register(tiff.NewTag(7, "FocusMode", nil))
	MakerNoteTags.Register(tiff.NewTag(8, "FlashSetting", nil))
	MakerNoteTags.Register(tiff.NewTag(9, "FlashDevice", nil))
	// Tag 10
	MakerNoteTags.Register(tiff.NewTag(11, "WhiteBalanceFineTune", nil))
	MakerNoteTags.Register(tiff.NewTag(12, "WB_RBLevels", nil))
	MakerNoteTags.Register(tiff.NewTag(13, "ProgramShift", nil))
	MakerNoteTags.Register(tiff.NewTag(14, "ExposureDifference", nil))
	MakerNoteTags.Register(tiff.NewTag(15, "ISOSelection", nil))
	MakerNoteTags.Register(tiff.NewTag(16, "DataDump", nil))
	MakerNoteTags.Register(tiff.NewTag(17, "PreviewIFD", nil))
	MakerNoteTags.Register(tiff.NewTag(18, "FlashExposureComp", nil))
	MakerNoteTags.Register(tiff.NewTag(19, "ISOSetting", nil))
	MakerNoteTags.Register(tiff.NewTag(20, "ColorBalanceA", nil))
	// Tag 21
	MakerNoteTags.Register(tiff.NewTag(22, "ImageBoundary", nil))
	MakerNoteTags.Register(tiff.NewTag(23, "ExternalFlashExposureComp", nil))
	MakerNoteTags.Register(tiff.NewTag(24, "FlashExposureBracketValue", nil))
	MakerNoteTags.Register(tiff.NewTag(25, "ExposureBracketValue", nil))
	MakerNoteTags.Register(tiff.NewTag(26, "ImageProcessing", nil))
	MakerNoteTags.Register(tiff.NewTag(27, "CropHiSpeed", nil))
	MakerNoteTags.Register(tiff.NewTag(28, "ExposureTuning", nil))
	MakerNoteTags.Register(tiff.NewTag(29, "SerialNumber", nil))
	MakerNoteTags.Register(tiff.NewTag(30, "ColorSpace", nil))
	MakerNoteTags.Register(tiff.NewTag(31, "VRInfo", nil))
	MakerNoteTags.Register(tiff.NewTag(32, "ImageAuthentication", nil))
	MakerNoteTags.Register(tiff.NewTag(33, "FaceDetect", nil))
	MakerNoteTags.Register(tiff.NewTag(34, "ActiveD-Lighting", nil))
	MakerNoteTags.Register(tiff.NewTag(35, "PictureControlData", nil))
	MakerNoteTags.Register(tiff.NewTag(36, "WorldTime", nil))
	MakerNoteTags.Register(tiff.NewTag(37, "ISOInfo", nil))
	// Tags 38..41
	MakerNoteTags.Register(tiff.NewTag(42, "VignetteControl", nil))
	MakerNoteTags.Register(tiff.NewTag(43, "DistortInfo", nil))
	// Tags 44..52
	MakerNoteTags.Register(tiff.NewTag(53, "HDRInfo", nil))
	// Tags 54..56
	MakerNoteTags.Register(tiff.NewTag(57, "LocationInfo", nil))
	// Tags 58..60
	MakerNoteTags.Register(tiff.NewTag(61, "BlackLevel", nil))
	// Tags 62..127
	MakerNoteTags.Register(tiff.NewTag(128, "ImageAdjustment", nil))
	MakerNoteTags.Register(tiff.NewTag(129, "ToneComp", nil))
	MakerNoteTags.Register(tiff.NewTag(130, "AuxiliaryLens", nil))
	MakerNoteTags.Register(tiff.NewTag(131, "LensType", nil))
	MakerNoteTags.Register(tiff.NewTag(132, "Lens", nil))
	MakerNoteTags.Register(tiff.NewTag(133, "ManualFocusDistance", nil))
	MakerNoteTags.Register(tiff.NewTag(134, "DigitalZoom", nil))
	MakerNoteTags.Register(tiff.NewTag(135, "FlashMode", nil))
	MakerNoteTags.Register(tiff.NewTag(136, "AFInfo", nil))
	MakerNoteTags.Register(tiff.NewTag(137, "ShootingMode", nil))
	MakerNoteTags.Register(tiff.NewTag(138, "AutoBracketRelease", nil))
	MakerNoteTags.Register(tiff.NewTag(139, "LensFStops", nil))
	MakerNoteTags.Register(tiff.NewTag(140, "ContrastCurve", nil))
	MakerNoteTags.Register(tiff.NewTag(141, "ColorHue", nil))
	// Tag 142
	MakerNoteTags.Register(tiff.NewTag(143, "SceneMode", nil))
	MakerNoteTags.Register(tiff.NewTag(144, "LightSource", nil))
	MakerNoteTags.Register(tiff.NewTag(145, "ShotInfo", nil))
	MakerNoteTags.Register(tiff.NewTag(146, "HueAdjustment", nil))
	MakerNoteTags.Register(tiff.NewTag(147, "NEFCompression", nil))
	MakerNoteTags.Register(tiff.NewTag(148, "Saturation", nil))
	MakerNoteTags.Register(tiff.NewTag(149, "NoiseReduction", nil))
	MakerNoteTags.Register(tiff.NewTag(150, "NEFLinearizationTable", nil))
	MakerNoteTags.Register(tiff.NewTag(151, "ColorBalance", nil))
	MakerNoteTags.Register(tiff.NewTag(152, "LensData", nil))
	MakerNoteTags.Register(tiff.NewTag(153, "RawImageCenter", nil))
	MakerNoteTags.Register(tiff.NewTag(154, "SensorPixelSize", nil))
	// Tag 155
	MakerNoteTags.Register(tiff.NewTag(156, "SceneAssist", nil))
	// Tag 157
	MakerNoteTags.Register(tiff.NewTag(158, "RetouchHistory", nil))
	// Tag 159
	MakerNoteTags.Register(tiff.NewTag(160, "SerialNumber", nil))
	// Tag 161
	MakerNoteTags.Register(tiff.NewTag(162, "ImageDataSize", nil))
	// Tags 163..164
	MakerNoteTags.Register(tiff.NewTag(165, "ImageCount", nil))
	MakerNoteTags.Register(tiff.NewTag(166, "DeletedImageCount", nil))
	MakerNoteTags.Register(tiff.NewTag(167, "ShutterCount", nil))
	MakerNoteTags.Register(tiff.NewTag(168, "FlashInfo", nil))
	MakerNoteTags.Register(tiff.NewTag(169, "ImageOptimization", nil))
	MakerNoteTags.Register(tiff.NewTag(170, "Saturation", nil))
	MakerNoteTags.Register(tiff.NewTag(171, "VariProgram", nil))
	MakerNoteTags.Register(tiff.NewTag(172, "ImageStabilization", nil))
	MakerNoteTags.Register(tiff.NewTag(173, "AFResponse", nil))
	// Tags 174..175
	MakerNoteTags.Register(tiff.NewTag(176, "MultiExposure", nil))
	MakerNoteTags.Register(tiff.NewTag(177, "HighISONoiseReduction", nil))
	// Tag 178
	MakerNoteTags.Register(tiff.NewTag(179, "ToningEffect", nil))
	// Tags 180..181
	MakerNoteTags.Register(tiff.NewTag(182, "PowerUpTime", nil))
	MakerNoteTags.Register(tiff.NewTag(183, "AFInfo2", nil))
	MakerNoteTags.Register(tiff.NewTag(184, "FileInfo", nil))
	MakerNoteTags.Register(tiff.NewTag(185, "AFTune", nil))
	// Tags 186..188
	MakerNoteTags.Register(tiff.NewTag(189, "PictureControlData", nil))
	// Tags 190..194
	MakerNoteTags.Register(tiff.NewTag(195, "BarometerInfo", nil))
	// Tags 196..3583
	MakerNoteTags.Register(tiff.NewTag(3584, "PrintIM", nil))
	MakerNoteTags.Register(tiff.NewTag(3585, "CaptureData", nil))
	// Tags 3586..3592
	MakerNoteTags.Register(tiff.NewTag(3593, "CaptureVersion", nil))
	// Tags 3594..3597
	MakerNoteTags.Register(tiff.NewTag(3598, "CaptureOffsets", nil))
	// Tag 3599
	MakerNoteTags.Register(tiff.NewTag(3600, "ScanIFD", nil))
	// Tags 3601..3602
	MakerNoteTags.Register(tiff.NewTag(3603, "CaptureEditVersions", nil))
	// Tags 3604..3612
	MakerNoteTags.Register(tiff.NewTag(3613, "ICCProfile", nil))
	MakerNoteTags.Register(tiff.NewTag(3614, "CaptureOutput", nil))
	// Tags 3615..3617
	MakerNoteTags.Register(tiff.NewTag(3618, "NEFBitDepth", nil))
	// Tags 3619..65535

	MakerNoteTagSpace.RegisterTagSet(MakerNoteTags)

	/*
		TODO: Find tag names for the following that we know to exist...
		Tag 163/0x00a3
		Tag 164/0x00a4
		Tag 186/0x00ba
		Tag 187/0x00bb
		Tag 188/0x00bc
	*/
	tiff.RegisterTagSpace(MakerNoteTagSpace)
}
