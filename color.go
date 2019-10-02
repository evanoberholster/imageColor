package imagecolor

import (
	"fmt"
	"image"
	"math"
	"sort"
	"sync"

	"gonum.org/v1/gonum/stat"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	hueValue uint8 = iota
	saturationValue
	lightValue
)

// ColorHSL - Color defined in the HSL colorspace (Hue, Saturation, Value)
type ColorHSL [3]float64

// DistanceHue is the distance between 2 Hues on the HSL colorspace (Hue, Saturation, Lightness)
func (c ColorHSL) DistanceHue(c2 ColorHSL) float64 {
	dist1 := math.Abs(c[hueValue] - c2[hueValue])
	dist2 := c2[hueValue] + math.Abs(c[hueValue]-360)
	dist3 := math.Abs(c2[hueValue]-360) + c[hueValue]
	if dist2 < dist1 || dist3 < dist1 {
		return math.Min(dist2, dist3)
	}
	return math.Abs(dist1)
}

// Distance is the equludian distance between 2 Colors in the HSL colorspace (Hue, Saturation, Lightness)
func (c ColorHSL) Distance(c2 ColorHSL) float64 {
	hDistance := c.DistanceHue(c2)
	return math.Sqrt((hDistance * hDistance) + ((c[1] - c2[1]) * (c[1] - c2[1])) + ((c[2] - c2[2]) * (c[2] - c2[2])))
}

// ClosestMaterialColor -
func closestMaterialColor(c ColorHSL) (colorName MaterialColor) {
	// Check for black pixels
	if c[lightValue] < 0.05 {
		colorName = materialBlack
		return
	}
	// Check for white pixels
	if c[saturationValue] < 0.018 && c[lightValue] > 0.95 {
		colorName = materialWhite
		return
	}
	// Check for grey pixels
	if c[saturationValue] == 0.0 && c[hueValue] == 0.0 && c[lightValue] > 0.05 && c[lightValue] < 0.95 {
		colorName = materialGrey
		return
	}
	minDist := 200.0
	for _, series := range materialColorsSeries {
		for mc, ch := range series {
			dist := c.Distance(ch)
			if dist < minDist {
				minDist = dist
				colorName = mc
			}
		}
	}
	return colorName
}

// ProminentColors - Return Prominent Colors
// (limit) percentage limit of promiment colors to return
func (ic *ImageColors) ProminentColors(limit float64) ProminentColors {
	var wg sync.WaitGroup
	var pc ProminentColors

	go calcColors(&wg, ic, limit, &pc)
	go calcSaturation(&wg, ic, &pc)
	go calcLightness(&wg, ic, &pc)
	wg.Add(3)
	wg.Wait()

	pc.Colorfulness = math.Sqrt(pc.Saturation[0] + pc.Saturation[1])
	pc.Qlightness = ic.QuantileLightness()

	return pc
}

func calcSaturation(wg *sync.WaitGroup, ic *ImageColors, pc *ProminentColors) {
	a, b := ic.MeanSaturation()
	pc.Saturation = [2]float64{a, b}
	wg.Done()
}

func calcLightness(wg *sync.WaitGroup, ic *ImageColors, pc *ProminentColors) {
	a, b := ic.MeanLightness()
	pc.Lightness = [2]float64{a, b}
	wg.Done()
}

func calcColors(wg *sync.WaitGroup, ic *ImageColors, limit float64, pc *ProminentColors) {
	totalColors := 0
	var name MaterialColor
	result := make(map[MaterialColor]int)

	for name := range materialColorsName {
		result[name] = 0
	}

	for _, x := range *ic {
		for _, c1 := range x {
			name = closestMaterialColor(c1)
			result[name]++
			//totalColors++
		}
	}
	for _, w := range result {
		totalColors += w
	}
	for name, num := range result {
		w := float64(num) / float64(totalColors)
		if w > limit {
			pc.Colors = append(pc.Colors, ProminentColor{Color: name, W: w})
		}
	}
	sort.Sort(pc)
	wg.Done()
}

func (ic ImageColors) hues() []float64 {
	return ic.getValues(hueValue)
}

func (ic ImageColors) saturations() []float64 {
	return ic.getValues(saturationValue)
}

func (ic ImageColors) lightness() []float64 {
	return ic.getValues(lightValue)
}

func (ic ImageColors) getValues(value uint8) []float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	var items []float64
	for x := 0; x < len(ic); x++ {
		for y := 0; y < len(ic[x]); y++ {
			items = append(items, ic[x][y][int(value)])
		}
	}
	return items
}

func (ic ImageColors) lightnessSkew() float64 {
	val := ic.getValues(lightValue)
	return stat.Skew(val, nil)
}

func (ic ImageColors) focusedPixels() (float64, float64, float64, float64) {
	var focused []float64
	val := ic.getValues(lightValue)
	for _, v := range val {
		if v > 0.02 {
			focused = append(focused, v)
		}
	}
	sort.Float64s(focused)
	mean, std := stat.MeanStdDev(focused, nil)
	skew := stat.Skew(focused, nil)
	per := float64(len(focused)) / float64(len(val))
	return mean, std, skew, per
}

// MeanHue -
func (ic ImageColors) MeanHue() (float64, float64) {
	return stat.MeanStdDev(ic.hues(), nil)
}

// MeanSaturation -
func (ic ImageColors) MeanSaturation() (float64, float64) {
	return stat.MeanStdDev(ic.saturations(), nil)
}

// MeanLightness -
func (ic ImageColors) MeanLightness() (float64, float64) {
	return stat.MeanStdDev(ic.lightness(), nil)
}

// QuantileSaturation -
func (ic ImageColors) QuantileSaturation() float64 {
	sats := ic.saturations()
	sort.Float64s(sats)
	return stat.Quantile(0.50, stat.Empirical, sats, nil)
}

// QuantileLightness -
func (ic ImageColors) QuantileLightness() float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	light := ic.lightness()
	sort.Float64s(light)
	return stat.Quantile(0.70, stat.Empirical, light, nil)
}

// ImageColors - Array of ColorHSL
type ImageColors [][]ColorHSL

// defineSize - Define the size of the ImageColors array
func (ic *ImageColors) defineSize(width, height int) {
	(*ic) = make([][]ColorHSL, width)
	for x := 0; x < width; x++ {
		(*ic)[x] = make([]ColorHSL, height)
	}
}

// GetImageColors - Create ImageColors array from an image
func GetImageColors(m image.Image) *ImageColors {
	bounds := m.Bounds()
	minX, minY := bounds.Min.X, bounds.Min.Y
	width, height := bounds.Max.X-minX, bounds.Max.Y-minY
	var ic ImageColors
	ic = make([][]ColorHSL, width)
	for x := 0; x < width; x++ {
		ic[x] = make([]ColorHSL, height)
		for y := 0; y < height; y++ {
			cf := newColorful(m.At(x+minX, y+minY).RGBA())
			//cf, _ := colorful.MakeColor(m.At(x+minX, y+minY))
			//colorHSL := NewColorHSL(cf)
			//ic.AddHSL(x, y, colorHSL)
			ic.AddColor(x, y, cf)
		}
	}
	return &ic
}

// AddHSL - Add ColorHSL to Coordinates x and y of ImageColors
func (ic *ImageColors) AddHSL(x, y int, hsl ColorHSL) {
	(*ic)[x][y] = hsl
}

// AddColor - Add Color to ColorHSL in ImageColors with Coords x and y
func (ic *ImageColors) AddColor(x, y int, cf colorful.Color) {
	h, s, l := cf.Hsl()
	(*ic)[x][y] = ColorHSL{h, s, l}
}

// NewColorHSL - Create ColorHSL from colorful.Color
func NewColorHSL(cf colorful.Color) ColorHSL {
	h, s, l := cf.Hsl()
	return ColorHSL{h, s, l}
}

func newColorful(r, g, b, a uint32) colorful.Color {
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
