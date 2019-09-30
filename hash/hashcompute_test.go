package hash

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"testing"

	"github.com/nfnt/resize"
)

func fetchTestImage(fileName string) (image.Image, error) {
	// open "test.jpg"
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

func BenchmarkAverageHash50(b *testing.B) {
	for _, imgFile := range []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "a:1c3cfed8f9f9f970"},
	} {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			b.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(8, 8, img, resize.Bilinear)

		// run the Fib function b.N times
		for n := 0; n < b.N; n++ {
			_, err := AverageHash(resized)
			if err != nil {
				b.Errorf("Error calculating AverageHash: %v", err)
			}
		}
	}
}

func TestAverageHash(t *testing.T) {
	for _, imgFile := range []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "a:1c3cfed8f9f9f970"},
		{"tests/test2.jpg", "a:fcfcf8f8c08098e0"},
		{"tests/test3.jpg", "a:0107fffff80040c0"},
	} {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			t.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(8, 8, img, resize.Bilinear)
		hash, err := AverageHash(resized)
		if err != nil {
			t.Errorf("Error calculating AverageHash: %v", err)
		}
		if hash.ToString() != imgFile.hash {
			t.Errorf("AverageHash was incorrect, got: %v, want: %v.", hash.ToString(), imgFile.hash)
		}
		fmt.Println(imgFile.fileName, " AverageHash", hash.ToString())
	}
}

func TestDifferenceHash(t *testing.T) {
	images := []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "d:f8d870b233b3f3b3"},
		{"tests/test2.jpg", "d:c8000121031c3802"},
		{"tests/test3.jpg", "d:f7dfe7f040c0c000"},
	}
	for _, imgFile := range images {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			t.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(9, 8, img, resize.Bilinear)
		hash, err := DifferenceHash(resized)
		if err != nil {
			t.Errorf("Error calculating DifferenceHash: %v", err)
		}
		if hash.ToString() != imgFile.hash {
			t.Errorf("DifferenceHash was incorrect, got: %v, want: %v.", hash.ToString(), imgFile.hash)
		}
		fmt.Println(imgFile.fileName, " DifferenceHash", hash.ToString())
	}
}

func BenchmarkDifferenceHash50(b *testing.B) {
	for _, imgFile := range []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "a:1c3cfed8f9f9f970"},
	} {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			b.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(9, 8, img, resize.Bilinear)

		// run the Fib function b.N times
		for n := 0; n < b.N; n++ {
			_, err := DifferenceHash(resized)
			if err != nil {
				b.Errorf("Error calculating DifferenceHash: %v", err)
			}
		}
	}
}

func TestPerceptionHash(t *testing.T) {
	images := []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "p:c996472b68fc248f"},
	}
	for _, imgFile := range images {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			t.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(64, 64, img, resize.Bilinear)
		hash, err := PerceptionHash(resized)
		if err != nil {
			t.Errorf("Error calculating DifferenceHash: %v", err)
		}
		if hash.ToString() != imgFile.hash {
			t.Errorf("PerceptionHash was incorrect, got: %v, want: %v.", hash.ToString(), imgFile.hash)
		}
		fmt.Println(imgFile.fileName, " PerceptionHash", hash.ToString())
	}
}

func BenchmarkPerceptionHash50(b *testing.B) {
	for _, imgFile := range []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "p:c996472b68fc248f"},
	} {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			b.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(64, 64, img, resize.Bilinear)

		// run the Fib function b.N times
		for n := 0; n < b.N; n++ {
			_, err := PerceptionHash(resized)
			if err != nil {
				b.Errorf("Error calculating PerceptionHash: %v", err)
			}
		}
	}
}

func TestExtPerceptionHash(t *testing.T) {
	images := []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "p:c996472b68fc248f"},
		{"tests/test2.jpg", "p:fd82c136609e39cb"},
		{"tests/test3.jpg", "p:e5ac295fddc01a62"},
	}
	for _, imgFile := range images {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			t.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(64, 64, img, resize.Bilinear)
		hash, err := ExtPerceptionHash(resized)
		if err != nil {
			t.Errorf("Error calculating ExtPerceptionHash: %v", err)
		}
		if hash.ToString() != imgFile.hash {
			t.Errorf("ExtPerceptionHash was incorrect, got: %v, want: %v.", hash.ToString(), imgFile.hash)
		}
		fmt.Println(imgFile.fileName, " ExtPerceptionHash", hash.ToString())
	}
}

func BenchmarkExtPerceptionHash100(b *testing.B) {
	for _, imgFile := range []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "p:c996472b68fc248f"},
	} {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			b.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(256, 256, img, resize.Bilinear)

		// run the Fib function b.N times
		for n := 0; n < b.N; n++ {
			_, err := ExtPerceptionHash(resized)
			if err != nil {
				b.Errorf("Error calculating ExtPerceptionHash: %v", err)
			}
		}
	}
}

func TestExtAverageHash(t *testing.T) {
	images := []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "a:0080c3ffffff1f07"},
		{"tests/test2.jpg", "a:00000404878f9f97"},
		{"tests/test3.jpg", "a:0001070f1f1f7fff"},
	}
	for _, imgFile := range images {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			t.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(64, 64, img, resize.Bilinear)
		hash, err := ExtAverageHash(resized)
		if err != nil {
			t.Errorf("Error calculating ExtAverageHash: %v", err)
		}
		if hash.ToString() != imgFile.hash {
			t.Errorf("ExtAverageHash was incorrect, got: %v, want: %v.", hash.ToString(), imgFile.hash)
		}
		fmt.Println(imgFile.fileName, " ExtAverageHash", hash.ToString())
	}
}

func BenchmarkExtAverageHash50(b *testing.B) {
	for _, imgFile := range []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "a:0080c3ffffff1f07"},
	} {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			b.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(64, 64, img, resize.Bilinear)

		// run the Fib function b.N times
		for n := 0; n < b.N; n++ {
			_, err := ExtAverageHash(resized)
			if err != nil {
				b.Errorf("Error calculating ExtAverageHash: %v", err)
			}
		}
	}
}

func TestExtDifferenceHash(t *testing.T) {
	images := []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "d:f1a191227469f300"},
		{"tests/test2.jpg", "d:81000100678c0a00"},
		{"tests/test3.jpg", "d:b73fbd041a200000"},
	}
	for _, imgFile := range images {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			t.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(8, 8, img, resize.Bilinear)
		hash, err := ExtDifferenceHash(resized)
		if err != nil {
			t.Errorf("Error calculating ExtDifferenceHash: %v", err)
		}
		if hash.ToString() != imgFile.hash {
			t.Errorf("ExtDifferenceHash was incorrect, got: %v, want: %v.", hash.ToString(), imgFile.hash)
		}
		fmt.Println(imgFile.fileName, " ExtDifferenceHash", hash.ToString())
	}
}

func BenchmarkExtDifferenceHash50(b *testing.B) {
	for _, imgFile := range []struct {
		fileName string
		hash     string
	}{
		{"tests/test1.jpg", "d:f1a191227469f300"},
	} {
		img, err := fetchTestImage(imgFile.fileName)
		if err != nil {
			b.Errorf("Error loading file: %v", err)
		}
		resized := resize.Resize(8, 8, img, resize.Bilinear)

		// run the Fib function b.N times
		for n := 0; n < b.N; n++ {
			_, err := ExtDifferenceHash(resized)
			if err != nil {
				b.Errorf("Error calculating ExtDifferenceHash: %v", err)
			}
		}
	}
}
