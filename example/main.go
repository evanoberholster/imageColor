package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"

	"github.com/evanoberholster/color"
	"github.com/lucasb-eyer/go-colorful"
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
	img, err := loadImage("../../test/img/1.jpg")
	if err != nil {
		panic(err)
	}
	imgR := resize.Resize(10, 10, img, resize.Lanczos3)

	start := time.Now()
	ic := color.ImageToValues(imgR)
	fmt.Println(time.Since(start))

	//start = time.Now()
	//fmt.Println(ic.AverageHue())
	m, std := ic.AverageLightness()
	fmt.Printf("Lightness\t Avg: %.2f\t Std: %.2f\n", m*100, std*100)
	m, std = ic.AverageChroma()
	fmt.Printf("Chroma\t\t Avg: %.2f\t Std: %.2f\n", m*100, std*100)
	fmt.Println(time.Since(start))

	//fmt.Println(ic.ProminentHues())

	mcolors := []string{
		"#F44336", // red
		"#E91E63", // pink
		"#9C27B0", // purple
		"#673AB7", // deep purple
		"#3F51B5", // indigo
		"#2196F3", // blue
		"#03A9F4", // light blue
		"#00BCD4", // cyan
		"#009688", // teal
		"#4CAF50", // green
		"#8BC34A", // lightgreen
		"#CDDC39", // lime
		"#FFEB3B", // yellow
		"#FFC107", // amber
		"#FF9800", // orange
		"#FF5722", // deep orange
		"#795548", // brown
		"#9E9E9E", // grey
		"#607D8B", // blue grey
		"#FFFFFF", // white
		"#000000", // black
	}

	for _, m := range mcolors {
		c, err := colorful.Hex(m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c.Hcl())
	}
}
