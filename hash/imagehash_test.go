package hash

import (
	"bufio"
	"bytes"
	"fmt"
	"image/jpeg"
	"os"
	"testing"

	"github.com/nfnt/resize"
)

//func TestLoadExtImageHash(t *testing.T) {
//	var b bytes.Buffer
//
//	bar := bufio.NewReader(&b)
//	var err error
//	h, err := LoadExtImageHash(bar)
//	if err != nil {
//		t.Errorf("Error loading hash: %v", err)
//	}
//	fmt.Println(h)
//}

func TestDump(t *testing.T) {
	file1, _ := os.Open("tests/test1.jpg")
	defer file1.Close()
	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	img1, _ := jpeg.Decode(file1)
	resized := resize.Resize(256, 256, img1, resize.Bilinear)
	hash1, _ := ExtPerceptionHash(resized)
	//err := hash1.Dump(foo)
	//bar := bufio.NewReader(&b)
	fmt.Println(hash1.hash)
	fmt.Println(foo.Buffered())
	fmt.Println(foo.Flush())
	fmt.Println(b.Bytes())

}
