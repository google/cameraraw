package d2x

import "math/big"

type NefIFD0 struct {
	NewSubfileType            uint32     `tiff:"field,tag=254"`
	ImageWidth                uint32     `tiff:"field,tag=256"`
	ImageLength               uint32     `tiff:"field,tag=257"`
	BitsPerSample             []uint16   `tiff:"field,tag=258"`
	Compression               uint16     `tiff:"field,tag=259"`
	PhotometricInterpretation uint16     `tiff:"field,tag=262"`
	Make                      string     `tiff:"field,tag=271"`
	Model                     string     `tiff:"field,tag=272"`
	StripOffsets              uint32     `tiff:"field,tag=273"`
	Orientation               uint16     `tiff:"field,tag=274"`
	SamplesPerPixel           uint16     `tiff:"field,tag=277"`
	RowsPerStrip              uint32     `tiff:"field,tag=278"`
	StripByteCounts           uint32     `tiff:"field,tag=279"`
	XResolution               *big.Rat   `tiff:"field,tag=282"`
	YResolution               *big.Rat   `tiff:"field,tag=283"`
	PlanarConfiguration       uint16     `tiff:"field,tag=284"`
	ResolutionUnit            uint16     `tiff:"field,tag=296"`
	Software                  string     `tiff:"field,tag=305"`
	DateTime                  string     `tiff:"field,tag=306"`
	SubIFDs                   []uint32   `tiff:"field,tag=330"`
	ReferenceBlackWhite       []*big.Rat `tiff:"field,tag=532"`
	ExifIFD                   uint32     `tiff:"field,tag=34665"`
	GPSIFD                    uint32     `tiff:"field,tag=34853"`
	DateTimeOriginal          string     `tiff:"field,tag=36867"`
	TIFFEPStandardID          []byte     `tiff:"field,tag=37398"`
}

type NefIFD0SubIFD0 struct {
	NewSubfileType              uint32   `tiff:"field,tag=254"`
	Compression                 uint16   `tiff:"field,tag=259"`
	XResolution                 *big.Rat `tiff:"field,tag=282"`
	YResolution                 *big.Rat `tiff:"field,tag=283"`
	ResolutionUnit              uint16   `tiff:"field,tag=296"`
	JPEGInterchangeFormat       uint32   `tiff:"field,tag=513"`
	JPEGInterchangeFormatLength uint32   `tiff:"field,tag=514"`
	YCbCrPositioning            uint16   `tiff:"field,tag=531"`
}

type NefIFD0SubIFD1 struct {
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
	ExposureTime          *big.Rat `tiff:"field,tag=33434"`
	FNumber               *big.Rat `tiff:"field,tag=33437"`
	ExposureProgram       uint16   `tiff:"field,tag=34850"`
	ISOSpeedRatings       uint16   `tiff:"field,tag=34855"`
	DateTimeOriginal      string   `tiff:"field,tag=36867"`
	DateTimeDigitized     string   `tiff:"field,tag=36868"`
	ExposureBiasValue     *big.Rat `tiff:"field,tag=37380"`
	MaxApertureValue      *big.Rat `tiff:"field,tag=37381"`
	MeteringMode          uint16   `tiff:"field,tag=37383"`
	LightSource           uint16   `tiff:"field,tag=37384"`
	Flash                 uint16   `tiff:"field,tag=37385"`
	FocalLength           *big.Rat `tiff:"field,tag=37386"`
	MakerNote             []byte   `tiff:"field,tag=37500"`
	UserComment           []byte   `tiff:"field,tag=37510"`
	SubsecTime            string   `tiff:"field,tag=37520"`
	SubsecTimeOriginal    string   `tiff:"field,tag=37521"`
	SubsecTimeDigitized   string   `tiff:"field,tag=37522"`
	SensingMethod         uint16   `tiff:"field,tag=41495"`
	FileSource            byte     `tiff:"field,tag=41728"`
	SceneType             byte     `tiff:"field,tag=41729"`
	CFAPattern            []byte   `tiff:"field,tag=41730"`
	CustomRendered        uint16   `tiff:"field,tag=41985"`
	ExposureMode          uint16   `tiff:"field,tag=41986"`
	WhiteBalance          uint16   `tiff:"field,tag=41987"`
	DigitalZoomRatio      *big.Rat `tiff:"field,tag=41988"`
	FocalLengthIn35mmFilm uint16   `tiff:"field,tag=41989"`
	SceneCaptureType      uint16   `tiff:"field,tag=41990"`
	GainControl           uint16   `tiff:"field,tag=41991"`
	Contrast              uint16   `tiff:"field,tag=41992"`
	Saturation            uint16   `tiff:"field,tag=41993"`
	Sharpness             uint16   `tiff:"field,tag=41994"`
	SubjectDistanceRange  uint16   `tiff:"field,tag=41996"`
}

type NefIFD0GPSIFD struct {
	GPSVersionID []byte `tiff:"field,tag=0"`
}

type Makernote struct {
	Version                   []byte     `tiff:"field,tag=1"`
	ISOSpeed                  []uint16   `tiff:"field,tag=2"`
	Quality                   string     `tiff:"field,tag=4"`
	WhiteBalance              string     `tiff:"field,tag=5"`
	Sharpening                string     `tiff:"field,tag=6"`
	FocusMode                 string     `tiff:"field,tag=7"`
	FlashSetting              string     `tiff:"field,tag=8"`
	FlashDevice               string     `tiff:"field,tag=9"`
	WhiteBalanceFineTune      int16      `tiff:"field,tag=11"`
	ProgramShift              []byte     `tiff:"field,tag=13"`
	ExposureDifference        []byte     `tiff:"field,tag=14"`
	PreviewIFD                uint32     `tiff:"field,tag=17"`
	ISOSetting                []uint16   `tiff:"field,tag=19"`
	ExternalFlashExposureComp []byte     `tiff:"field,tag=23"`
	FlashExposureBracketValue []byte     `tiff:"field,tag=24"`
	ExposureBracketValue      *big.Rat   `tiff:"field,tag=25"`
	CropHiSpeed               []uint16   `tiff:"field,tag=27"`
	ExposureTuning            []byte     `tiff:"field,tag=28"`
	SerialNumber              string     `tiff:"field,tag=29"`
	ColorSpace                uint16     `tiff:"field,tag=30"`
	ToneComp                  string     `tiff:"field,tag=129"`
	LensType                  byte       `tiff:"field,tag=131"`
	Lens                      []*big.Rat `tiff:"field,tag=132"`
	FlashMode                 byte       `tiff:"field,tag=135"`
	AFInfo                    []byte     `tiff:"field,tag=136"`
	ShootingMode              uint16     `tiff:"field,tag=137"`
	AutoBracketRelease        uint16     `tiff:"field,tag=138"`
	LensFStops                []byte     `tiff:"field,tag=139"`
	ContrastCurve             []byte     `tiff:"field,tag=140"`
	ColorHue                  string     `tiff:"field,tag=141"`
	LightSource               string     `tiff:"field,tag=144"`
	ShotInfo                  []byte     `tiff:"field,tag=145"`
	HueAdjustment             int16      `tiff:"field,tag=146"`
	NoiseReduction            string     `tiff:"field,tag=149"`
	ColorBalance              []byte     `tiff:"field,tag=151"`
	LensData                  []byte     `tiff:"field,tag=152"`
	RawImageCenter            []uint16   `tiff:"field,tag=153"`
	SensorPixelSize           []*big.Rat `tiff:"field,tag=154"`
	UNKNOWN_TAG_163           byte       `tiff:"field,tag=163"`
	UNKNOWN_TAG_164           []byte     `tiff:"field,tag=164"`
	ImageCount                uint32     `tiff:"field,tag=165"`
	DeletedImageCount         uint32     `tiff:"field,tag=166"`
	ShutterCount              uint32     `tiff:"field,tag=167"`
	FlashInfo                 []byte     `tiff:"field,tag=168"`
	Saturation                string     `tiff:"field,tag=170"`
	MultiExposure             []byte     `tiff:"field,tag=176"`
	HighISONoiseReduction     uint16     `tiff:"field,tag=177"`
}
