package r3dOperation

import (
	"fmt"
	"testing"
)

func TestBuildDate(t *testing.T) {
	fmt.Println(BuildDate())
}

// Clip ClipInit(const char *);
// //int ClipStatus();
// int ClipVideoTrackCount(Clip);
// int ClipWidth(Clip);
// int ClipHeight(Clip);
// float ClipVideoAudioFramerate(Clip);
// float ClipTimecodeFramerate(Clip);
// int ClipVideoFrameCount(Clip);
// const char * ClipStartAbsoluteTimecode(Clip);
// const char * ClipEndAbsoluteTimecode(Clip);
// const char * ClipStartEdgeTimecode(Clip);
// const char * ClipEndEdgeTimecode(Clip);
// void ClipFree(Clip);
func TestClip(t *testing.T) {
	clip := ClipInit("/home/pillars/Videos/redone/B011_C001_0605L3.RDC/B011_C001_0605L3_001.R3D")
	fmt.Println("1------>", ClipVideoTrackCount(clip)) // 轨
	fmt.Println(ClipWidth(clip))                       // 宽度（尺寸1024×768）
	fmt.Println(ClipHeight(clip))                      // 高度（尺寸1024×768）
	fmt.Println(ClipVideoAudioFramerate(clip))         // 25视频音频帧速率
	fmt.Println(ClipTimecodeFramerate(clip))           // 25时间线帧速率
	fmt.Println(ClipVideoFrameCount(clip))             // 1185
	fmt.Println(ClipStartAbsoluteTimecode(clip))       // 绝对始码
	fmt.Println(ClipEndAbsoluteTimecode(clip))         // 绝对止码
	fmt.Println(ClipStartEdgeTimecode(clip))           // 始码
	fmt.Println(ClipEndEdgeTimecode(clip))             // 止码
	fmt.Println(ClipMetaData(clip))                    // 元数据
	//file, _ := os.Create("test.png")
	//defer file.Close()
	//ClipDecodeFrame(clip, 1000, file)

	//ClipDecodeFrameFree(buffer)
	ClipFree(clip)
}
