package thumb

import "github.com/disintegration/imaging"

var (
	PreRenderSize    = 3840
	MaxRenderSize    = 3840
	JpegQuality      = 95
	JpegQualitySmall = 80
	Algorithm        = ResampleLanczos
)

const (
	ResampleLanczos ResampleAlgorithm = "lanczos"
	ResampleCubic                     = "cubic"
	ResampleLinear                    = "linear"
)

type ResampleAlgorithm string

func (a ResampleAlgorithm) Filter() imaging.ResampleFilter {
	switch a {
	case ResampleLanczos:
		return imaging.Lanczos
	case ResampleCubic:
		return imaging.CatmullRom
	case ResampleLinear:
		return imaging.Linear
	default:
		return imaging.Lanczos
	}
}

const (
	ResampleFillCenter ResampleOption = iota
	ResampleFillTopLeft
	ResampleFillBottomRight
	ResampleFit
	ResampleResize
	ResampleNearestNeighbor
	ResampleDefault
	ResamplePng
)

type ResampleOption int

var ResampleMethods = map[ResampleOption]string{
	ResampleFillCenter:      "center",
	ResampleFillTopLeft:     "left",
	ResampleFillBottomRight: "right",
	ResampleFit:             "fit",
	ResampleResize:          "resize",
}

type Type struct {
	Source  string
	Width   int
	Height  int
	Public  bool
	Options []ResampleOption
}

var Types = map[string]Type{
	"tile_50":   {"tile_500", 50, 50, false, []ResampleOption{ResampleFillCenter, ResampleDefault}},
	"tile_100":  {"tile_500", 100, 100, false, []ResampleOption{ResampleFillCenter, ResampleDefault}},
	"tile_224":  {"tile_500", 224, 224, false, []ResampleOption{ResampleFillCenter, ResampleDefault}},
	"tile_500":  {"", 500, 500, false, []ResampleOption{ResampleFillCenter, ResampleDefault}},
	"colors":    {"fit_720", 3, 3, false, []ResampleOption{ResampleResize, ResampleNearestNeighbor, ResamplePng}},
	"left_224":  {"fit_720", 224, 224, false, []ResampleOption{ResampleFillTopLeft, ResampleDefault}},
	"right_224": {"fit_720", 224, 224, false, []ResampleOption{ResampleFillBottomRight, ResampleDefault}},
	"fit_720":   {"", 720, 720, true, []ResampleOption{ResampleFit, ResampleDefault}},
	"fit_1280":  {"fit_2048", 1280, 1024, true, []ResampleOption{ResampleFit, ResampleDefault}},
	"fit_1920":  {"fit_2048", 1920, 1200, true, []ResampleOption{ResampleFit, ResampleDefault}},
	"fit_2048":  {"", 2048, 2048, true, []ResampleOption{ResampleFit, ResampleDefault}},
	"fit_2560":  {"", 2560, 1600, true, []ResampleOption{ResampleFit, ResampleDefault}},
	"fit_3840":  {"", 3840, 2400, true, []ResampleOption{ResampleFit, ResampleDefault}},
}

var DefaultTypes = []string{
	"fit_3840", "fit_2560", "fit_2048", "fit_1920", "fit_1280", "fit_720", "right_224", "left_224", "colors", "tile_500", "tile_224", "tile_100", "tile_50",
}

func (t Type) ExceedsLimit() bool {
	return t.Width > MaxRenderSize || t.Height > MaxRenderSize
}

func (t Type) SkipPreRender() bool {
	return t.Width > PreRenderSize || t.Height > PreRenderSize
}
