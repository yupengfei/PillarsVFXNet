#ifdef __cplusplus
extern "C" {
#endif

    const char * BuildDate();


    //call clipinit first and clipfree when finish usage
    typedef void * Clip;
	Clip ClipInit(const char *);
	//int ClipStatus();
	int ClipVideoTrackCount(Clip);
	int ClipWidth(Clip);
	int ClipHeight(Clip);
	float ClipVideoAudioFramerate(Clip);
	float ClipTimecodeFramerate(Clip);
	int ClipVideoFrameCount(Clip);
	const char * ClipStartAbsoluteTimecode(Clip);
	const char * ClipEndAbsoluteTimecode(Clip);
	const char * ClipStartEdgeTimecode(Clip);
	const char * ClipEndEdgeTimecode(Clip);
	const char * ClipMetaData(Clip clip);
	void ClipMetaDataFree(char * metadataChar);
	unsigned char * ClipDecodeFrame(Clip clip, int frameIndex);
	void ClipDecodeFrameFree(unsigned char * imgbuffer);
	
	void ClipFree(Clip);
	

#ifdef __cplusplus
}
#endif
