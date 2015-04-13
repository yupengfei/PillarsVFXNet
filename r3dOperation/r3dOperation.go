package r3dOperation

// #cgo CFLAGS: -I ./cppFile/
// #cgo LDFLAGS:  -L /home/pillars/Go/src/PillarsPhenomVFXWeb/cppFile -L /home/pillars/Go/src/PillarsPhenomVFXWeb/cppFile/Lib/linux64 -lr3d  -lR3DSDK -lstdc++ -lpthread -luuid -lm
// #include "/home/pillars/Go/src/PillarsPhenomVFXWeb/cppFile/r3d.h"
// #include <stdlib.h>
import "C"

import "fmt"
import "unsafe"
import "image"
import "image/color"
import "image/png"
import "io"

// import "encoding/binary"

func BuildDate() string {
	var cmsg *C.char = C.BuildDate()

	var result string = C.GoString(cmsg)
	return result
}

func ClipInit(fileName string) C.Clip {
	var fileNameC *C.char = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var clip C.Clip = C.ClipInit(fileNameC)
	return clip
}

func ClipFree(clip C.Clip) {
	C.ClipFree(clip)
}

func ClipVideoTrackCount(clip C.Clip) int {
	return int(C.ClipVideoTrackCount(clip))
}

func ClipWidth(clip C.Clip) int {
	return int(C.ClipWidth(clip))
}

func ClipHeight(clip C.Clip) int {
	return int(C.ClipHeight(clip))
}

func ClipVideoAudioFramerate(clip C.Clip) float32 {
	return float32(C.ClipVideoAudioFramerate(clip))
}

func ClipTimecodeFramerate(clip C.Clip) float32 {
	return float32(C.ClipTimecodeFramerate(clip))
}

func ClipVideoFrameCount(clip C.Clip) float32 {
	return float32(C.ClipVideoFrameCount(clip))
}

func ClipStartAbsoluteTimecode(clip C.Clip) string {
	return C.GoString(C.ClipStartAbsoluteTimecode(clip))
}

func ClipEndAbsoluteTimecode(clip C.Clip) string {
	return C.GoString(C.ClipEndAbsoluteTimecode(clip))
}

func ClipStartEdgeTimecode(clip C.Clip) string {
	return C.GoString(C.ClipStartEdgeTimecode(clip))
}

func ClipEndEdgeTimecode(clip C.Clip) string {
	return C.GoString(C.ClipEndEdgeTimecode(clip))
}

func ClipMetaData(clip C.Clip) string {
	var metadata *C.char = C.ClipMetaData(clip)
	var metadataGoString = C.GoString(metadata)
	C.ClipMetaDataFree(metadata)
	return metadataGoString
}

func ClipDecodeFrame(clip C.Clip, frameIndex int, out io.Writer) *C.uchar {
	imgbuffer := C.ClipDecodeFrame(clip, C.int(frameIndex))
	fmt.Println(C.int(frameIndex))
	fmt.Println(imgbuffer)
	var width int = ClipWidth(clip)
	var height int = ClipHeight(clip)
	resultByte := C.GoBytes(unsafe.Pointer(imgbuffer), C.int(width*height*3*2))

	img := image.NewRGBA64(image.Rect(0, 0, width, height))
	var i = 0
	var j = 0
	var currentIndex = 0
	for i = 0; i < height; i++ {
		for j = 0; j < width; j++ {

			color := color.RGBA64{
				R: uint16(resultByte[currentIndex]) | uint16(resultByte[currentIndex+1])<<8,
				G: uint16(resultByte[currentIndex+width*height]) | uint16(resultByte[currentIndex+1+width*height])<<8,
				B: uint16(resultByte[currentIndex+2*width*height]) | uint16(resultByte[currentIndex+1+2*width*height])<<8,
				A: uint16(65535),
				// R: uint16(resultByte[currentIndex])<<8 | uint16(resultByte[currentIndex + 1]),
				// G: uint16(resultByte[currentIndex + width * height])<<8 | uint16(resultByte[currentIndex + 1 + width * height]),
				// B: uint16(resultByte[currentIndex + 2 * width * height])<<8 | uint16(resultByte[currentIndex + 1 + 2 * width * height]),
				// A: uint16(65535),
			}
			img.SetRGBA64(j, i, color)
			currentIndex++
		}
	}
	png.Encode(out, img)
	return imgbuffer
}

// func ClipDecodeFrameFree(imgbuffer * C.uchar) {
// 	C.ClipDecodeFrameFree(imgbuffer)
// }

//func C.CString(goString string) *C.char
//func C.GoString(cString *C.char) string
//func C.GoString(cString *C.char, length C.int) string

//C.CString() need free mannul
//var cmsg * C.char = C.CString("hi")
//defer C.free(unsafe.Pointer(cmsg))
