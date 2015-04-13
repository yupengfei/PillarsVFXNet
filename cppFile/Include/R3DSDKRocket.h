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

#ifndef R3DSDKROCKET_H
#define R3DSDKROCKET_H

namespace R3DSDK {

// Red Rocket open / decode status
enum RedRocketStatus
{
	RSOK = 0,
	RSResourceNotAvailable = 1,			// resource does not exist (Rocket board not installed or right driver not loaded)
	RSResourceInUse	= 2,				// resource is in use and can't be opened right now (another application is using it)
	RSResourceNotOpened = 3,			// not opened or not opened properly
	RSUnknownError = 4,					// unknown error
	RSOutOfMemory = 5,					// no more memory could be allocated
	RSInvalidParameter = 6,				// one of the supplied parameters is not valid
	RSClipNotOpen = 7,					// no open R3D clip was passed
	RSOutputBufferInvalid = 8,			// null pointer, too small or not aligned
	RSInvalidPixelType = 9,				// cannot decode to specified pixel type
	RSPlatformNotSupported = 10,		// RED Rocket is not yet supported on Linux
	RSCannotReadFromFile = 11,			// unable to get frame from the clip
	RSOldDriverFirmware = 12,			// Driver and/or firmware too old
	RSIsDroppedFrame = 13,				// frame is a dropped frame and could not be decoded
	RSCustomWidthNotAligned = 14,		// CustomOutputWidth is not aligned, did you check RedRocket::AlignOutputWidth()?
	RSNotAnHDRxClip = 15,				// cannot decode HDRx as this is not an HDRx clip
	RSUnsupportedClipFormat = 16,		// cannot decode this type of clip which will happen in the following scenarios:
										//  1) Monochrome clip on the original Rocket. Support for this will come in a future release
										//  2) Dragon clip on the original Rocket. The original RED Rocket does not support Dragon clips.
	RSResourceLost = 17					// RED Rocket was unplugged
};

class Clip;

// The rocket decode job structure. This structure needs to be
// constructed to submit a decode job to the RED Rocket decoder.
// Every field must be set, with the exception of 'ImageProcessing',
// 'HdrProcessing', 'CustomOutputWidth' & 'CustomOutputHeight'
// (so long as Mode is not set to DECODE_ROCKET_CUSTOM_RES).
struct RocketDecodeJob
{
	// Frame to decode for which clip
	R3DSDK::Clip *				Clip;

	size_t						VideoTrackNo;			// 0 = main (A) track, 1 = EPIC/Scarlet-X higlight protection track 2 (X track)
														// ignored when doing HDRx blending
	size_t						VideoFrameNo;

	// item is the pointer to the original item as submitted in the Decode()
	// call, decodeStatus will be either RSOK, RSCannotReadFromFile or
	// RSUnknownError right now, other more specific error messages will follow later
	typedef void (*RocketDecodeCallback)(RocketDecodeJob * item, RedRocketStatus decodeStatus);

	// Callback to call when frame is available in the specified output buffer
	// You will receive a call for each frame submitted and these calls will
	// come in on seperate threads, potentially at the same time.
	// SPEND AS LITTLE TIME AS POSSIBLE IN THE CALLBACK TO NOT HOLD UP PROCESSING OTHER FRAMES!
	RocketDecodeCallback		Callback;
	void *						PrivateData;			// available for your use as you see fit, R3D SDK will not touch this variable!

	// Resolution/speed to decode the image at. This will also 
	// influence how much memory is needed for the output buffer
	VideoDecodeMode				Mode;
	
	// Custom scale & crop output, set Mode to DECODE_ROCKET_CUSTOM_RES
	// step 1: downscale before crop, no upscale is supported. If not set (0), the image will
	// get scaled to the image dimensions as set in step 2 below. If both scale & crop/pad are
	// not set, while DECODE_ROCKET_CUSTOM_RES is set, adding the job for decode will fail.
	size_t						CustomScaleWidth;
	size_t						CustomScaleHeight;
	
	// step 2: crop & pad after scale, CustomOutputWidth cannot be wider than 4096 pixels
	// (unless equal to the full image width). Crop is always a center crop and any padding
	// (CustomOutputWidth or Height is bigger than scaled image) will always center the image
	// inside a black frame. If not set (0), the image will not get cropped or padded and the
	// same size as the scale in step 1 will be used. If both scale & crop/pad are not set,
	// while DECODE_ROCKET_CUSTOM_RES is set adding, the job for decode will fail.
	size_t						CustomOutputWidth;
	size_t						CustomOutputHeight;
	
	// Which pixel type to decode the image in. See the VideoPixelType
	// enum for the available options. The PixelType_16Bit_RGB_Planar
	// pixel type is NOT supported with RED Rocket decoding and will
	// result in RSInvalidPixelType error code being returned.
	VideoPixelType				PixelType;

	// Pointer to the buffer to store the image in. This cannot be
	// NULL otherwise the decode will fail. The buffer must be aligned
	// on a 512-byte boundary for Rocket (see sample code) and 4096
	// byte boundary for Rocket-X.
	void *						OutputBuffer;

	// Number of bytes in a row
	size_t						BytesPerRow;

	// Size ouf the output buffer, must be aligned to 512-byte
	// boundary for Rocket and 4096-byte for Rocket-X.
	size_t						OutputBufferSize;
	
	// Image processing settings to apply to the decode. This
	// parameter is optional (set it to NULL!) in which case the
	// decoder will simply use the clip's default settings.
	ImageProcessingSettings	*	ImageProcessing;

	// HDRx processing settings to apply to the decode. This
	// parameter is optional (set it to NULL if not used!).
	// If the clip is not an HDRx clip, the decode will fail.
	// VideoTrackNo is ignored when this field is not NULL!
	HdrProcessingSettings *		HdrProcessing;

	RocketDecodeJob();
};

enum RocketModel
{
	RM_UNKNOWN = 0,							// Open() has not been called or open failed
	RM_ROCKET = 1,							// Original Rocket
	RM_ROCKET_X = 2							// Rocket-X
};

// Driver, Firmware & PCI information for a specific card (see RedRocket::Open)
struct RocketInfo
{
	// Model of the hardware board. Will be unknown if Open() has not been called or Open() failed
	RocketModel	BoardModel;

	// Driver & firmware version
	// (can be all zeroes if Open failed with an error code other than RSOldDriverFirmware)
	int			DriverVersion[4];			// display as %d.%d.%d.%d, same for all cards in a system
	int			FirmwareVersion[4];			// display as %d.%d.%d.%d, can be different from card to card, but should be the same

	// true if card is running at full speed. For Rocket this is PCI Express x8 gen 1,	for Rocket-X
	// this is PCI Express x16 gen 2. If false performance can go down depending on the scenario.
	bool		PCIOK;

	// Minimum driver & firmware needed for the original RED Rocket
	int			MinimumDriverVersion[4];
	int			MinimumFirmwareVersion[4];

	// Minimum driver & firmware needed for RED Rocket-X
	int			MinimumRocketXDriverVersion[4];
	int			MinimumRocketXFirmwareVersion[4];

	// Sets up the MinimumXXX fields above, no need to call Open
	RocketInfo();

	// Get minimum version numbers as a single number
	static unsigned int GetMinimumRocketDriverVersion();
	static unsigned int GetMinimumRocketFirmwareVersion();
	static unsigned int GetMinimumRocketXDriverVersion();
	static unsigned int GetMinimumRocketXFirmwareVersion();
};

// Class for detecting RED Rocket accelerator boards, opening
// & closing them and decoding frames through the board. This
// class is thread-safe so long no call to Open or Close is
// in progress.
class RedRocket
{
public:
	RedRocket();
	~RedRocket();

	// ******************************
	// BOARD MANAGEMENT
	// ******************************

	static size_t RocketsAvailable();	// 0 if not available

	// Get RocketInfo for a specific Rocket without needing to open/close it. This function
	// will work if another application has the Rocket open. Returns RSInvalidParameter
	// if info == null, RSResourceNotAvailable if for whatever reason Rocket could not
	// be opened, RSOldDriverFirmware if driver/firmware is too old (in this case driver
	// and firmware versions will still be reported) or RSOK if everything is okay.
	static RedRocketStatus GetRocketInfo(size_t boardToQuery, RocketInfo * info);

	// Supply an existing RocketInfo function when opening a board to
	// receive driver, firmware & PCI information on a succesful open
	RedRocketStatus Open(size_t boardToOpen, RocketInfo * info = NULL);	// 0 - RocketsAvailable()-1 to open a specific board
	void Close();
	
	// ******************************
	// ASYNCHRONOUS HARDWARE DECODING
	// ******************************

	// Certain camera recording modes may not supported by either
	// Rocket or Rocket-X. Call this function to determin if the specified
	// clip can or cannot be decoded on the currently open Rocket.
	// Currently not supported:
	// - Rocket: clips shot on a Monochrome or Dragon sensor
	// Support for Monochrome clips will be added in the future.
	RedRocketStatus DecodeSupportedForClip(Clip & clip) const;

	// Decode frames back to host memory, try to keep 6 - 12 frames in
	// the queue to maximize decode performance. This function will wait
	// if the queue is full (16 items are waiting to be processed)
	// Each parallel decode will allocate a buffer for I/O.
	// When an error occurs in submitting a multi-job decode, only jobs
	// up till the error have been submitted.  There will be some changes
	// for this in a future SDK. For now it's best to submit 1 job at a time
	RedRocketStatus Decode(RocketDecodeJob * jobs, size_t noOfJobs) const;
	
	// ******************************
	// IMAGE ALIGNMENT
	// ******************************

	// Combination of certain pixel types at certain (scaled / cropped)
	// resolutions require more bytes per pixel then the desired output
	// width. This function will check, and if needed, update the supplied
	// width. This function should be called for each desired output
	// resolution (after crop & scale!) before your memory buffers get
	// allocated (since the size may be adjusted). Then you'll want to
	// use the updated width. Examples of modes that need adjusting:
	// 1) 5K ANA (3296 x 2700) decoded at 1/8th (412 x 337) with the 8-bit
	// BGR pixeltype requires a width of 416 (image will be centered).
	// 2) 5K ANA decoded at 1/8th and 1/4th (824 x 675) with the 12-bit
	// RGB pixeltype requires a width of 832 (image will be centered).
	// Similar width changes may be needed for custom output sizes.
	RedRocketStatus AlignOutputWidth(size_t * requestWidth, VideoPixelType requestPixelType, bool requestIsHdr) const;

private:
	// no copying & assignment
	RedRocket(const RedRocket &) throw();
	RedRocket & operator= (const RedRocket &) throw();

	void * reserved;
};

}

#endif
