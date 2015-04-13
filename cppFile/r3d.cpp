#include <string>
#include <string.h>
#include <iostream>
#include <stdlib.h>

#include "R3DSDK.h"
#include "r3d.h"

unsigned char * AlignedMalloc(size_t & sizeNeeded)
{
	// alloc 15 bytes more to make sure we can align the buffer in case it isn't
	unsigned char * buffer = (unsigned char *)malloc(sizeNeeded + 15U);

	if (!buffer)
		return NULL;

	sizeNeeded = 0U;

	// cast to a 32-bit or 64-bit (depending on platform) integer so we can do the math
	uintptr_t ptr = (uintptr_t)buffer;

	// check if it's already aligned, if it is we're done
	if ((ptr % 16U) == 0U)
		return buffer;

	// calculate how many bytes we need
	sizeNeeded = 16U - (ptr % 16U);

	return buffer + sizeNeeded;
}

const char * BuildDate() {
	return R3DSDK::Info::BuildDate();
}

Clip ClipInit(const char * fileName) {
	R3DSDK::Clip * ret = new R3DSDK::Clip(fileName);
    return (void *)ret;
}

int ClipVideoTrackCount(Clip clip) {
    R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
    return (int)foo->VideoTrackCount();
}

int ClipWidth(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	return (int)foo->Width();
}

int ClipHeight(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	return (int)foo->Height();
}

float ClipVideoAudioFramerate(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	return foo->VideoAudioFramerate();
}

float ClipTimecodeFramerate(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	return foo->TimecodeFramerate();
}

int ClipVideoFrameCount(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	return (int)foo->VideoFrameCount();
}

// const char * ClipStartAbsoluteTimecode(Clip);
// const char * ClipEndAbsoluteTimecode(Clip);
// const char * 
// const char * ClipEndEdgeTimecode(Clip);
// void ClipFree();

const char * ClipStartAbsoluteTimecode(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	return foo->AbsoluteTimecode(0U);
}

const char * ClipEndAbsoluteTimecode(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	return foo->AbsoluteTimecode(foo->VideoFrameCount() - 1U);
}

const char * ClipStartEdgeTimecode(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	return foo->EdgeTimecode(0U);
}

const char * ClipEndEdgeTimecode(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	return foo->EdgeTimecode(foo->VideoFrameCount() - 1U);
}

void ClipFree(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	delete foo;
	return;
}

const char * ClipMetaData(Clip clip) {
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	std::string metadata;
	metadata = "{";
	for (std::size_t i = 0U; i < foo->MetadataCount() - 1; i++)
	{	
		metadata += "\"" + foo->MetadataItemKey(i) + "\":" + "\"" + foo->MetadataItemAsString(i) + "\",";
	}
	metadata += "\"" + foo->MetadataItemKey(foo->MetadataCount() - 1) + "\":" + "\"" + foo->MetadataItemAsString(foo->MetadataCount() - 1) + "\"}";
	unsigned int metaDataCharlength = strlen(metadata.c_str());
	char * retChar = new char[metaDataCharlength + 1];
	strcpy(retChar, metadata.c_str());
	return retChar;
}

void ClipMetaDataFree(char * metadataChar) {
	delete [] metadataChar;
}

unsigned char * ClipDecodeFrame(Clip clip, int frameIndex) {
	std::cout << "1" << std::endl;
	R3DSDK::Clip * foo = (R3DSDK::Clip *) clip;
	// let the user know if this failed
	// let the user know if this failed
	if (foo->Status() != R3DSDK::LSClipLoaded)
	{
		return NULL;
	}

	// calculate how much ouput memory we're going to need
	size_t width = foo->Width();
	size_t height = foo->Height();	
	std::cout << "2" << std::endl;
	size_t memNeeded = width * height * 3U * 2U;
	size_t adjusted = memNeeded;
	unsigned char * imgbuffer = AlignedMalloc(adjusted);
	std::cout << "3" << std::endl;
	std::cout << "4" << std::endl;
	if (imgbuffer == NULL)
	{
		return NULL;
	}
	std::cout << "5" << std::endl;
	std::cout << width << std::endl;
	std::cout << height << std::endl;
	std::cout << adjusted << std::endl;

	R3DSDK::VideoDecodeJob job;
	job.BytesPerRow = width * 2U;
	job.OutputBufferSize = memNeeded;
	job.Mode = R3DSDK::DECODE_FULL_RES_PREMIUM;
	job.OutputBuffer = imgbuffer;
	job.PixelType = R3DSDK::PixelType_16Bit_RGB_Planar;
	std::cout << frameIndex << std::endl;
	if (foo->DecodeVideoFrame(frameIndex, job) != R3DSDK::DSDecodeOK)
	{
		std::cout << "6" << std::endl;
		return NULL;
	}

	// free the original pointer, not the one adjusted for alignment
	//free(imgbuffer - adjusted);
	return imgbuffer;
}

// void ClipDecodeFrameFree(unsigned char * imgbuffer) {
// 	delete imgbuffer;
// }


