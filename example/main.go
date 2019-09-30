package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"

	"github.com/evanoberholster/imageColor"
	"github.com/nfnt/resize"
)

func loadImage(fileName string) (image.Image, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return img, err
}

func main() {
	img, err := loadImage("../../test/img/8.jpg")
	if err != nil {
		panic(err)
	}
	imgR := resize.Resize(0, 250, img, resize.Lanczos3)

	start := time.Now()
	ic := imageColor.GetImageColors(imgR)
	fmt.Println(time.Since(start))

	start = time.Now()
	//m, std := ic.MeanHue()
	//fmt.Printf("Hue\t\t Avg: %.2f\t Std: %.2f\n", m, std)

	pc := ic.ProminentColors(0.01)
	fmt.Printf("Lightness\t Avg: %.2f\t Std: %.2f\n", pc.Lightness[0]*100, pc.Lightness[1]*100)
	fmt.Printf("Saturation\t Avg: %.2f\t Std: %.2f\n", pc.Saturation[0]*100, pc.Saturation[1]*100)
	fmt.Printf("Colorfulness\t %.2f\n", pc.Colorfulness*100)
	fmt.Printf("Quantile light\t %.2f\n", pc.Qlightness*100)
	fmt.Println(pc.Colors)
	fmt.Println(time.Since(start))

	//start = time.Now()
	//color.CalcCompositionBoxes(img)
	//fmt.Println(time.Since(start))
}
