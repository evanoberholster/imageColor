package color

import (
	"image"
	"math"
	"sort"

	"github.com/lucasb-eyer/go-colorful"
	"gonum.org/v1/gonum/stat"
)

// Box - x1, y1, x2, y2 int
type Box struct {
	Rect    image.Rectangle
	Focused float64
	Score   float64
	MeanL   float64
	StdL    float64
	SkewL   float64
	values  []float64
}

// NewBox - Create a new Box
func NewBox(x1, y1, x2, y2 int) *Box {
	return &Box{Rect: image.Rect(x1, y1, x2, y2)}
}

// ThirdsBoxes - Create boxes for thirds composition.
// 4 Boxes with centers at 1/3 and 2/3 horizontal and vertical.
func ThirdsBoxes(width, height int) (boxes []*Box) {
	x1, y1 := int(width/6), int(height/6)
	boxes = append(boxes, NewBox(x1, y1, x1*3, y1*3))
	boxes = append(boxes, NewBox(x1*3, y1, (x1*3)+(x1*2), (y1*3)))
	boxes = append(boxes, NewBox(x1, y1*3, (x1*3), (y1*3)+(y1*2)))
	boxes = append(boxes, NewBox(x1*3, y1*3, (x1*3)+(x1*2), (y1*3)+(y1*2)))
	return
}

func focusedPixels(val []float64) (float64, float64, float64, float64) {
	var focused []float64
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

// MeanStd - Get Mean and StandardVariation for values
func (b *Box) MeanStd() {
	b.MeanL, b.StdL = stat.MeanStdDev(b.values, nil)
}

// Skew - Get the skew of the normal distribution curve for values
func (b *Box) Skew() {
	b.SkewL = stat.Skew(b.values, nil)
}

// FocusedPixels - Create ImageColors array from an image
func (b *Box) FocusedPixels(m image.Image) {
	bounds := m.Bounds()
	minX, minY := bounds.Min.X, bounds.Min.Y
	width, height := bounds.Max.X-minX, bounds.Max.Y-minY
	var ic ImageColors
	ic.defineSize(width, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cf, _ := colorful.MakeColor(m.At(x+minX, y+minY))
			colorHSL := NewColorHSL(cf)
			if colorHSL[lightValue] > 0.02 {
				b.values = append(b.values, colorHSL[lightValue])
			}
		}
	}
	b.Focused = float64(len(b.values)) / float64(width*height)
}

// FocusScore -
func (b *Box) FocusScore() {
	b.Score = (math.Sqrt(b.MeanL*b.MeanL*b.StdL) - b.SkewL*b.Focused/(b.MeanL*10000)) * 10
}
