package color

import (
	"image"

	"gonum.org/v1/gonum/stat"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	hueValue uint8 = iota
	chromaValue
	lightValue
)

func (ic ImageColors) Hues() []float64 {
	return ic.getValues(hueValue)
}

func (ic ImageColors) ProminentHues() []float64 {
	var pHues []float64
	pHues = ic.getValues(hueValue)
	return pHues
}

func (ic ImageColors) Chromas() []float64 {
	return ic.getValues(chromaValue)
}

func (ic ImageColors) Lightness() []float64 {
	return ic.getValues(lightValue)
}

func (ic ImageColors) getValues(value uint8) []float64 {
	var items []float64
	for y := 0; y < len(ic); y++ {
		for x := 0; x < len(ic[y]); x++ {
			items = append(items, ic[x][y][int(value)])
		}
	}
	return items
}

func (ic ImageColors) AverageHue() (float64, float64) {
	return stat.MeanStdDev(ic.Hues(), nil)
}

func (ic ImageColors) AverageChroma() (float64, float64) {
	return stat.MeanStdDev(ic.Chromas(), nil)
}

func (ic ImageColors) AverageLightness() (float64, float64) {
	return stat.MeanStdDev(ic.Lightness(), nil)
}

type ImageColors [][][3]float64

func (ic ImageColors) Add(x, y int, h, c, l float64) {
	ic[x][y][0] = h
	ic[x][y][1] = c
	ic[x][y][2] = l
}

func (ic *ImageColors) Size(height, width int) {
	(*ic) = make([][][3]float64, height)
	for y := 0; y < height; y++ {
		(*ic)[y] = make([][3]float64, width)
	}
}

func ImageToValues(m image.Image) *ImageColors {
	bounds := m.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	ic := new(ImageColors)
	ic.Size(height, width)
	//pixels := uint(width * height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			cf := makeColorful(r, g, b, a)
			h, c, l := cf.Hcl()
			ic.Add(x, y, h, c, l)
		}
	}
	return ic
	//fmt.Println(ic)
}

func makeColorful(r32, g32, b32, a32 uint32) colorful.Color {
	// shift by 8bytes
	r32 |= r32 << 8
	g32 |= g32 << 8
	b32 |= b32 << 8
	a32 |= a32 << 8

	// Compress to uint8 (255) then convert to uint32
	r := int32(uint8(r32))
	g := int32(uint8(g32))
	b := int32(uint8(b32))
	a := int32(uint8(a32))

	// Since color.Color is alpha pre-multiplied, we need to divide the
	// RGB values by alpha again in order to get back the original RGB.
	r *= 0xffff
	r /= a
	g *= 0xffff
	g /= a
	b *= 0xffff
	b /= a

	return colorful.Color{R: float64(r) / 65535.0, G: float64(g) / 65535.0, B: float64(b) / 65535.0}
}
