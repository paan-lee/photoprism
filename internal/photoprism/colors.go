package photoprism

import (
	"errors"
	"image/color"
	"math"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/photoprism/photoprism/pkg/colors"
)

// Colors returns the ColorPerception of an image (only JPEG supported).
func (m *MediaFile) Colors(thumbPath string) (perception colors.ColorPerception, err error) {
	if !m.IsJpeg() {
		return perception, errors.New("no color information: not a JPEG file")
	}

	img, err := m.Resample(thumbPath, "colors")

	if err != nil {
		log.Printf("can't open image: %s", err.Error())

		return perception, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	pixels := float64(width * height)
	chromaSum := 0.0

	colorCount := make(map[colors.Color]uint16)
	var mainColorCount uint16

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rgb, _ := colorful.MakeColor(color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
			i := colors.Colorful(rgb)
			perception.Colors = append(perception.Colors, i)

			if _, ok := colorCount[i]; ok == true {
				colorCount[i] += colors.Weights[i]
			} else {
				colorCount[i] = colors.Weights[i]
			}

			if colorCount[i] > mainColorCount {
				mainColorCount = colorCount[i]
				perception.MainColor = i
			}

			_, c, l := rgb.Hcl()

			chromaSum += c

			perception.Luminance = append(perception.Luminance, colors.Luminance(math.Round(l*15)))
		}
	}

	perception.Chroma = colors.Chroma(math.Round((chromaSum / pixels) * 100))

	return perception, nil
}
