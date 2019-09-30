package imagecolor

// Experimental
// Use with caution
import (
	"fmt"
	"image"
	"math"
	"sort"

	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/transform"
)

// ImageEdges -
func (b Box) ImageEdges(img image.Image, radius float64) image.Image {
	result := transform.Crop(img, b.Rect)
	return effect.EdgeDetection(result, radius)
	//return result
}

// CalcCompositionBoxes - Experimental
func CalcCompositionBoxes(imgR image.Image) {
	//imgR := resize.Resize(500, 0, img, resize.Lanczos3)
	bounds := imgR.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	ratio := float64(width) / float64(height)
	// 4:3 500:375px image
	fmt.Println(width, height, ratio)

	boxes := ThirdsBoxes(width, height)
	for _, box := range boxes {
		edges := box.ImageEdges(imgR, 1)
		//filename := fmt.Sprintf("box%d.png", idx)
		//if err := imgio.Save(filename, edges, imgio.PNGEncoder()); err != nil {
		//	fmt.Println(err)
		//	return
		//}
		box.FocusedPixels(edges)
		box.MeanStd()
		box.Skew()
		box.FocusScore() // WIP
	}

	bigBox := NewBox(0, 0, width, height)
	var focused float64
	for _, b := range boxes {
		focused += b.Focused
		bigBox.values = append(bigBox.values, b.values...)
	}
	bigBox.Focused = focused / float64(len(boxes))
	sort.Float64s(bigBox.values)
	bigBox.MeanStd()
	bigBox.Skew()
	bigBox.FocusScore()
	boxes = append(boxes, bigBox)

	for _, b := range boxes {
		fmt.Printf("Score: %.3f\t", b.Score)
		fmt.Printf("Dist: %.3f\t", b.Distance(bigBox))
		fmt.Printf("Skew: %.3f\t", b.SkewL)
		fmt.Printf("Focused\t %.2f \t", b.Focused)
		fmt.Printf("Mean\t %.5f\t", b.MeanL)
		fmt.Printf("Std\t %.5f\t\n", b.StdL)
	}

}

// Distance -
func (b *Box) Distance(b2 *Box) float64 {
	eqludian := math.Sqrt((b.Score - b2.Score) * (b.Score - b2.Score))
	if b.Score < b2.Score {
		return -1 * eqludian
	}
	return eqludian
}
