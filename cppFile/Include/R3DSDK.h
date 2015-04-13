/* R3D SDK library version 5.2 header file. Do *NOT* use this
   header file with any other version of the R3D SDK library!
   
		Name: Phenom Films Co LTD
		License number: R472
  
   This header file and everything else included with the R3D
   SDK is Copyright (c) 2008-2014 RED Digital Cinema. All
   rights reserved. Redistribution of this header is prohibited!
   
   The SDK is thread-safe for the most part, but it may
   synchronize access when needed.
*/

#ifndef R3DSDK_H
#define R3DSDK_H

#include <string>

#include "R3DSDKDefinitions.h"

#include "R3DSDKMetadata.h"

#include "R3DSDKRocket.h"

namespace R3DSDK {

#define R3DSDK_H_MAJOR 5
#define R3DSDK_H_MINOR 2
	
// Get information regarding the R3D SDK version & your license
class Info
{
public:
	// Returns R3D SDK version in majorVersion & minorVersion parameters.
	static void Version(unsigned int & majorVersion, unsigned int & minorVersion);

	// Returns R3D SDK build information
	static const char * BuildNumber();
	static const char * BuildDate();

	// Returns your R3D SDK license number
	static const char * LicenseNumber();

	// Returns your (company) name the license is registered to
	static const char * Licensee();

private:
	Info();		// no need to instantiate this class
};

// Class for loading a clip, retrieving information and extracting images.
// This class is thread-safe as long as no call to LoadFrom() or Close()
// is in progress.
class Clip
{
public:
	// ******************************
	// CLIP MANAGEMENT
	// ******************************

	// Destructor will call Close() to close any open files
	~Clip();

	// First constructor form. Use this if you do not want to load a clip
	// when constructing this class. Use LoadFrom() to load a clip before
	// calling any other functions.
	Clip();

	// Second constructor form. Use this if you want to load a clip from a UTF-8
	// path when you construct the class. Load status can be checked through Status().
	Clip(const char * pathToFile);

	// Get clip load status after having used the second constructor form
	// or after a call to LoadFrom().
	LoadStatus Status() const;

	// Load the R3D file from the given UTF-8 path. If there are any other R3D
	// files that belong to the same clip these will get loaded as well.
	// If a clip was already open it will be closed before opening the
	// one specified in the pathToFile parameter.
	LoadStatus LoadFrom(const char * pathToFile);

	// Closes any open files (class destructor will also call this function).
	void Close();

	// Closes all handles for all R3D parts belonging to the clip. Handles
	// will be re-opened on a per R3D part basis as needed automatically.
	void CloseFileHandles();

	// Retrieve the UTF-8 path for the RMD sidecar file the R3D SDK will look
	// for when trying to load or update an RMD file or create a new one.
	// You will receive a valid path even if no RMD file is present
	// Returns null if no clip has been (succesfully) loaded.
	const char * GetRmdPath() const;

	// ******************************
	// SAVING CLIPS
	// ******************************
	// The callback will be called whenever a frame has been added, the
	// operation has been completed or an error has occurred. The last callback
	// you will get is either the done status or an error. If 'status' does
	// not equal CSFrameAdded the process is done or an error was encountered.
	// 
	// Return 'true' from your callback if you want the trim to continue
	// Return 'false' if you want to abort the trim
	typedef bool (*TrimCallback)(CreateStatus status, void * privateData, size_t frameNoDone, size_t lastFrameToDo);

	// Create a new clip from another one with a specific in and out point.
	// 'destination' is an existing folder where the output files will be created
	// 'destination' is a UTF-8 path
	//
	// This is an asynchronous call, it will return immediately. The 'destination'
	// parameter can be free'd or deleted as soon as this call returns. The
	// 'source' clip must remain around for the duration of this function.
	// The function returns CSStarted if the trim is under way and you will
	// receive at least one more callback when this value is returned.
	static CreateStatus CreateTrimFrom(const Clip & source, const char * destination, size_t startFrameNo, size_t endFrameNo, void * privateData, TrimCallback callback);

	// Create single frame clip from another one. 'destination' is a full UTF-8 path
	// including filename (unlike CreateTrimFrom above). The extension '.R3D'
	// must be at the end and it must be in capitals! The last character in the
	// filename part of the path CANNOT be a digit! This call returns immediately.
	static CreateStatus CreateSnapshotFrom(const Clip & source, const char * destination, size_t frameNo);

	// If you try to include a frame (either through CreateTrimFrom or CreateSnapshotFrom)
	// that is a dropped frame you will receive CSFailedToGetSourceFrame return code.
	// In this case CreateSnapshotFrom will not have created an output file at all
	// In this case CreateTrimFrom will have created a clip up to the dropped frame.

	// ******************************
	// SYNCHRONOUS SOFTWARE DECODING (MAIN TRACK)
	// ******************************

	// Decode the given video frame with the supplied decode settings,
	// output buffer and image processing settings (through decodeJob)
	DecodeStatus DecodeVideoFrame(size_t videoFrameNo, const VideoDecodeJob & decodeJob) const;

	// Decode the given audio block into the supplied buffer. This
	// audio block has 24-bit (check the metadata!) audio MSB aligned
	// (Big Endian order) in a 32-bit word. The lower 8-bits are
	// set to zero! If there is more than one audio channel (see
	// metadata again) the samples will be interleaved. outputBuffer
	// must point to a 512-byte aligned buffer that is big enough to
	// hold the requested audio block (see AudioBlockCountAndSize()
	// function below). The size of the memory block must be stored
	// in bufferSize which will be updated by the number of bytes
	// written into the audio buffer (can be less than the available
	// space in the buffer!), which will always be a multiple of
	// 4 * number_of_channels. The function will fail if no clip is
	// open, no audio is present, audioBlockNo is out of range or
	// outputBuffer and bufferSize are NULL or not aligned properly.
	// This function is currently *NOT* thread-safe!
	DecodeStatus DecodeAudioBlock(size_t audioBlockNo, void * outputBuffer, size_t * bufferSize) const;

	// Decode raw audio samples from the clip. startSample and
	// numberOfSamples are expressed per channel, do not multiply
	// them by the channel count! Like the above function this always
	// returns all channels and does so interleaved. Output buffer
	// must be 512 byte aligned and the size must be a multiple of
	// 512. Number of samples will be updated to reflect how many
	// samples were actually written into the output buffer. Normally
	// this is the same as the number of samples requested, unless
	// more samples were asked for than exist in the clip. On failure
	// numberOfSamples will be set to zero. If startSample is out of
	// range DSRequestOutOfRange will be returned. Size of the output
	// buffer must be at least: align_to_512(numberOfSamples * numberOfChannels * 4)
	// 0 <= startSample < AudioSampleCount()
	DecodeStatus DecodeAudio(unsigned long long startSample, size_t * numberOfSamples, void * outputBuffer, size_t bufferSize) const;

	// ******************************
	// MULTI-TRACK DECODING
	// Can be used to do manual HDR blending or look at a single track
	// EPIC/Scarlet-X only
	// ******************************
	
	// Returns number of video tracks in clip. This will always
	// return 1 for RED ONE clips. For Epic and Scarlet more tracks
	// may exist (in an HDRx recording for example).
	size_t VideoTrackCount() const;

	// Decode the given video frame on the given track with the supplied decode
	// settings, output buffer and image processing settings (through decodeJob)
	DecodeStatus VideoTrackDecodeFrame(size_t videoTrackNo, size_t videoFrameNo, const VideoDecodeJob & decodeJob) const;

	// ******************************
	// IMAGE PROCESSING SETTINGS
	// ******************************

	// Get either the clip image processing settings or overriding RMD sidecar metadata
	// (display this in your UI as the clip defaults!). Use GetClipImageProcessingSettings()
	// function below to always get the clip settings. PERFORMANCE WARNING: This function
	// will check for, and parse the RMD sidecar, EVERY time this function gets called!
	void GetDefaultImageProcessingSettings(ImageProcessingSettings & settingsToFill) const;

	// Get the image processing settings as they were set on camera when the clip was
	// recorded. Use this to offer "reset to clip/recorded metadata" in your application.
	void GetClipImageProcessingSettings(ImageProcessingSettings & settingsToFill) const;

	// Create or update existing RMD sidecar file. Returns true if success, false otherwise
	// The ImageProcessingSettings struct *must* have Version set to ColorVersion2 for this
	// call to succeed. ColorVersion1 is not supported within RMD sidecar files.
	// The RMD sidecar file is an XML based format that travels with the clips. In the future
	// RED cameras will create this file and third party's can add their own data.
	bool CreateOrUpdateRmd(const ImageProcessingSettings & settingsToWrite) const;

	// Get the HDRx processing settings from an existing RMD sidecar file
	// Returns blend settings OR track number to decode. If an error occurs, or no HDR
	// processing settings exist this function will return track 0 (A / main track)
	HdrMode GetRmdHdrProcessingSettings(HdrProcessingSettings & settingsToFill, size_t & trackNoToFill) const;

	// Create or update existing RMD sidecar file either a track number OR HDRx
	// processing settings. Returns true if success, false otherwise
	bool CreateOrUpdateRmd(HdrMode mode, const HdrProcessingSettings & settingsToWrite, size_t trackNoToWrite) const;

	// Whitebalance: convert a pixel area to Kelvin & Tint
	// End-user can pick a point on the image for a given frame. X and Y are in 0.0 - 1.0
	// range and go from upper left (0,0) to lower right (1,1). Function fails if x, y or
	// videoFrameNo are out of bounds, clip has not been loaded or frame is corrupt.
	// It will only update Kelvin & Tint for the supplied ImageProcessingSettings, nothing else
	DecodeStatus GetWhitebalance(ImageProcessingSettings & settingsToFill, size_t videoFrameNo, float x, float y) const;

	// Minimum color version of the ImageProcessingSettings structure required for the clip.
	// If the version is set lower in ImageProcessingSettings than the value returned by this
	// function it will be changed to the highest color version available. This function can
	// help to identify older footage that you may want to enable older image processing for
	// to match how the image used to look in an older SDK (backwards compatibility).
	ColorVersion MinimumColorVersion() const;

	// ******************************
	// CLIP INFORMATION
	// ******************************

	// Get a unique 16-byte clip identifier. There is no guarantee that this
	// uuid follows the ITU-T Rec. X.667 / ISO/IEC 9834-8:2005 standards!
	// Returns false if no clip is loaded or uuid is NULL, otherwise returns
	// true. Output uuid buffer must be able to hold 16 bytes.
	bool Uuid(unsigned char * uuid) const;

	// Width of the clip, will return 0 if no clip is loaded.
	size_t Width() const;

	// Height of the clip, will return 0 if no clip is loaded.
	size_t Height() const;

	// Number of video frames in the clip, will return 0 if no clip is loaded.
	size_t VideoFrameCount() const;

	// Number of audio channels in the clip, will return 0 if no
	// clip is loaded or if the clip does not have any audio. To
	// get the channel map (which of the 4 channels was used),
	// query the RMD_CHANNEL_MASK metadata item.
	size_t AudioChannelCount() const;

	// Number of audio blocks and their maximum size in the clip
	// Will return 0 if no clip is loaded or if maximumSize is
	// NULL, otherwise returns number of audio blocks and sets
	// maximumSize to the maximum buffer size that will be
	// needed to process all audio blocks. This size will be a
	// multiple of 512 (guaranteed)
	size_t AudioBlockCountAndSize(size_t * maximumSize) const;

	// Number of audio samples *per channel* in the clip. All
	// channels have the exact same number of samples. Returns
	// 0 if an error is encountered or if the clip has no audio
	unsigned long long AudioSampleCount() const;

	// Framerate for the video & audio 'track(s)'. Numerator & denominator
	// can be found seperately in the metadata table. Will return 0.0 if no
	// clip is loaded.
	float VideoAudioFramerate() const;	

	// Framerate for the timecode 'track'. This will be running at half the
	// video & audio 'track' rate if these are running at over 30 fps. The
	// timecode functions below will alternate '.' and ':' in that case!
	// Will return 0.0 if no clip is loaded.
	float TimecodeFramerate() const;

	// Retrieve the clip's default timecode 'track' for a given video frame.
	// This default was set on the camera when the clip was recorded. The
	// returned value is valid until the next time this function is called
	// or either AbsoluteTimecode() or EdgeTimecode() depending on which
	// 'track' is the default. Will return a NULL pointer if no clip is loaded.
	// or if videoFrameNo is out of bounds.
	const char * Timecode(size_t videoFrameNo);

	// Retrieve the clip's absolute (time of day or external) timecode for
	// a given video frame. The returned value is valid until the next time
	// this function or Timecode() is called. Will return a NULL pointer if
	// no clip is loaded or if VideoFrameNo is out of bounds.
	const char * AbsoluteTimecode(size_t videoFrameNo);

	// Retrieve the clip's edge (aka run record) timecode for a given video
	// frame. The returned value is valid until the next time this function
	// or Timecode() is called. Will return a NULL pointer if no clip is
	// loaded or videoFrameNo is out of bounds.
	const char * EdgeTimecode(size_t videoFrameNo);
	
	// You can encounter two kinds of timecodes in an R3D file:
	//
	// 01:00:00:00 -> 01:00:00:01 etc. for clips running at 30 fps and below
	// 01.00.00.00 -> 01:00:00:00 etc. for clips running over 30 fps

	// Functions to retrieve clip metadata through the metadata database
	size_t MetadataCount() const;								// number of metadate items in the store
	bool MetadataExists(const char * key) const;				// check if a metadata item exists in the store
	std::string MetadataItemKey(size_t index) const;			// get key for an index

	// keep in mind that there is no guarantee which metadata item will be at which
	// index. The index option is only there to allow enumeration through the database.
	unsigned int MetadataItemAsInt(const char * key) const;		// retrieve metadata item by key and return as int
	unsigned int MetadataItemAsInt(size_t index) const;			// retrieve metadata item by index and return as int
	std::string MetadataItemAsString(const char * key) const;	// retrieve metadata item by key and return as STL string
	std::string MetadataItemAsString(size_t index) const;		// retrieve metadata item by index and return as STL string
	float MetadataItemAsFloat(const char * key) const;			// retrieve metadata item by key and return as float
	float MetadataItemAsFloat(size_t index) const;				// retrieve metadata item by index and return as float
	
	MetadataType MetadataItemType(const char * key) const;		// retrieve metadata item type by key
	MetadataType MetadataItemType(size_t index) const;			// retrieve metadata item type by index

	// Functions to get still frame markers creating during record
	size_t MarkerCount() const;
	bool GetMarker(size_t markerNo, Marker & markerToFill) const;

private:
	// no copying & assignment
	Clip(const Clip &) throw();
	Clip & operator= (const Clip &) throw();

	void * reserved;
};

// The asynchronous decompress job structure. This structure needs to be
// constructed to submit a decode job to AsyncDecoder::DecodeForGpuSdk().
// Every field must be set, with the exception of 'Callback' & 'PrivateData'.
struct AsyncDecompressJob
{
	// Set this flag to true if you want to abort processing this frame as soon as possible
	bool						AbortDecode;

	// Frame to decode for which clip
	R3DSDK::Clip *				Clip;

	size_t						VideoTrackNo;			// 0 = main (A) track, 1 = EPIC/Scarlet-X higlight protection track 2 (X track)

	size_t						VideoFrameNo;

	// item is the pointer to the original item as submitted in the Decode()
	typedef void (*AsyncDecompressCallback)(AsyncDecompressJob * item, DecodeStatus decodeStatus);

	// Callback to call when frame is available in the specified output buffer
	// You will receive a call for each frame submitted and these calls will
	// come in on seperate threads, potentially at the same time.
	// SPEND AS LITTLE TIME AS POSSIBLE IN THE CALLBACK TO NOT HOLD UP PROCESSING OTHER FRAMES!
	AsyncDecompressCallback		Callback;
	void *						PrivateData;			// available for your use as you see fit, R3D SDK will not touch this variable!

	// Resolution/speed to decode the image at. This will also 
	// influence how much memory is needed for the output buffer
	VideoDecodeMode				Mode;

	// Pointer to the buffer to store the image in. This cannot be
	// NULL otherwise the decode will fail. The buffer must be aligned
	// on a 16-byte boundary (see sample code in GPU extension).
	void *						OutputBuffer;

	// Size ouf the output buffer. If this value is too small the decode will fail.
	// see AsyncDecoder::GetSizeBufferNeeded()
	size_t						OutputBufferSize;
	
	AsyncDecompressJob();
};

// Class for asynchronously decompressing a frame for further processing on the GPU.
// This class works together with GPU extension for the R3D SDK which can be
// downloaded separately. Future updates to the SDK will allow full software
// decodes to be completely asynchronous as well.
class AsyncDecoder
{
public:
	AsyncDecoder();
	~AsyncDecoder();

	// Returns the number of threads the decoder will use if opened when
	// noOfThreads in Open() is set to 0 (the default).
	static size_t ThreadsAvailable();

	// Open the asynchronous decoder with the specified number of threads
	// If noOfThreads is not set, or set to 0, ThreadsAvailable() number
	// of threads will be used.
	void Open(size_t noOfThreads = 0U);
	
	// Close asynchronous decoder and clean up any threads & memory in use
	void Close();

	// Returns the size OutputBufferSize needs to be in the AsyncDecompressJob
	// so the appropriate size input buffer can be allocated before calling
	// the DecodeForGpuSdk() function. The following must be set on input:
	//  1) job.Clip must point to an open Clip and cannot be NULL
	//  2) job.Mode must be set appropriately
	// returns 0 if these input parameters are invalid or no clip is loaded
	static size_t GetSizeBufferNeeded(const AsyncDecompressJob & job);

	// Decompress a frame in to the supplied OutputBuffer for further
	// processing on the GPU through the GPU extension to the R3D SDK
	DecodeStatus DecodeForGpuSdk(AsyncDecompressJob & job);

private:
	void * reserved;
};

}
#endif

