package utility

import (
	"fmt"
	"image"
	"log"
	"testing"

	"code.google.com/p/graphics-go/graphics"
)

func Test_LoadImage(t *testing.T) {
	src, err := LoadImage("/home/pillars/Desktop/111.jpg")
	if err != nil {
		log.Fatal(err)
	}
	// 缩略图的大小
	dst := image.NewRGBA(image.Rect(0, 0, 360, 240))
	// 产生缩略图,等比例缩放
	err = graphics.Scale(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	// 需要保存的文件
	imgcounter := 734
	saveImage(fmt.Sprintf("/home/pillars/Desktop/%03d.png", imgcounter), dst)
}
