package imageColor

import "fmt"

// ProminentColor - Prominent Color
// Color and Weight
type ProminentColor struct {
	Color MaterialColor
	W     float64
}

func (p ProminentColor) String() string {
	return fmt.Sprintf("%v:%.2f\t", p.Color.String(), p.W*100)
}

// ProminentColors - Slice of ProminentColor
type ProminentColors struct {
	Colors       []ProminentColor
	Hue          [2]float64
	Saturation   [2]float64
	Lightness    [2]float64
	Colorfulness float64
	Qlightness   float64
}

func (pc ProminentColors) Len() int           { return len(pc.Colors) }
func (pc ProminentColors) Less(i, j int) bool { return pc.Colors[i].W > pc.Colors[j].W }
func (pc ProminentColors) Swap(i, j int)      { pc.Colors[i], pc.Colors[j] = pc.Colors[j], pc.Colors[i] }
