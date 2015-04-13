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

#ifndef R3DSDKDEFINITIONS_H
#define R3DSDKDEFINITIONS_H

namespace R3DSDK {

// Clip load status
enum LoadStatus
{
	LSClipLoaded = 0,
	LSPathNotFound = 1,					// could not find the clip
	LSFailedToOpenFile = 2,				// could not open the clip
	LSNotAnR3DFile = 3,					// clip does not appear to be a(n) (valid) R3D file
	LSClipIsEmpty = 4,					// clip doesn't have any video frames in it
	LSOutOfMemory = 5,					// no more memory could be allocated
	LSUnknownError = 6,					// unknown error (shouldn't happen)
	LSNoClipOpen = 7,					// initial status, no clip has been loaded yet
	LSNotInitialized = 8				// library wasn't loaded properly (if obj-c, try new() instead)
};

// Clip create status
enum CreateStatus
{
	CSStarted = 0,						// trim has started but has not finished yet
	CSFrameAdded = 1,					// a frame has been added to the output clip, but not finished yet
	CSDone = 2,							// trim is done

	CSOutOfMemory = 3,					// no more memory could be allocated
	CSRequestOutOfRange = 4,			// the requested start or end frame does not exist in the clip
	CSInvalidParameter = 5,
	CSFailedToGetSourceFrame = 6,		// unable to load a frame from the source clip
	CSFailedToCreateDestination = 7,	// unable to create output clip
	CSFailedToWriteToDestination = 8,	// unable to write to output clip
	CSUnknownError = 9,
	CSInvalidSourceClip = 10,			// source clip cannot be used for trim (RED ONE firmware build 15 and below not supported)
	CSInvalidPath = 11					// output path is invalid (see trim requirements)
};

// Video decode status
enum DecodeStatus
{
	DSDecodeOK = 0,
	DSOutputBufferInvalid = 1,			// null pointer, too small or not aligned
	DSBytesPerRowInvalid = 2,			// too small or not aligned (video decode only)
	DSRequestOutOfRange = 3,			// the requested frame does not exist in the clip
	DSInvalidParameter = 4,				// one of the supplied parameters is not valid
	DSIsDroppedFrame = 5,				// the requested frame is a dropped frame and can't be decoded (video decode only)
	DSDecodeFailed = 6,					// decode failed because of a corrupt frame
	DSOutOfMemory = 7,					// no more memory could be allocated
	DSUnknownError = 8,					// unknown error (shouldn't happen)
	DSNoClipOpen = 9,					// open a clip first before trying to decode a frame
	DSCannotReadFromFile = 10,			// I/O error reading the frame
	DSInvalidPixelType = 11,			// cannot decode to specified pixel type
	DSNotAnHDRxClip = 12,				// cannot decode HDRx as this is not an HDRx clip
	DSCancelled = 13					// user requested decode to be cancelled
};

// Possible marker types
enum MarkerType
{
	MarkerSubClip = 0,					// region, most likely has different start & end frame numbers
	MarkerStillFrame = 1				// single frame marker, start & end frame number will be the same
};

// Marker type and frame range
struct Marker
{
	MarkerType	Type;
	size_t		StartFrameNumber;
	size_t		EndFrameNumber;

	Marker();
	bool IsValid() const;
};

// Gamma curve to apply to the image. The current RED ONE firmware release
// supports REDgamma, older firmwares have created clips set to Rec.709 or
// REDspace. These three are the only values clips will come in with.
enum ImageGammaCurve
{
	ImageGammaREDgamma4 = 12,			// only available with color version 2, otherwise ignored (REDspace will be used)
	ImageGammaREDgamma3 = 11,			// only available with color version 2, otherwise ignored (REDspace will be used)
	ImageGammaREDlogFilm = 9,			// only available with color version 2, otherwise ignored (REDlog will be used)

	ImageGammaLinear = 0,
	ImageGammaRec709 = 1,
	ImageGammaSRGB = 8,

	// Below gamma curves are deprecated

	ImageGammaPDlog685 = 4,
	ImageGammaPDlog985 = 5,
	ImageGammaCustomPDlog = 6,			// if used CustomPDLogBlackPoint, WhitePoint & Gamma must be set!

	ImageGammaREDspace = 2,
	ImageGammaREDlog = 3,
	ImageGammaREDgamma = 7,				// only available with color version 2, otherwise ignored (REDspace will be used)
	ImageGammaREDgamma2 = 10			// only available with color version 2, otherwise ignored (REDspace will be used)
};

// Color space to output the image in. The current RED ONE firmware release
// supports REDcolor, older firmwares have created clips set to Rec.709 or
// REDspace. These three are the only values clips will come in with.
enum ImageColorSpace
{
	ImageColorDRAGONcolor = 9,			// only available with color version 2, otherwise ignored (REDspace will be used)
	ImageColorREDcolor3 = 8,			// only available with color version 2, otherwise ignored (REDspace will be used)

	ImageColorRec709 = 1,
	ImageColorSRGB = 4,
	ImageColorAdobe1998 = 5,

	// Below color spaces are deprecated

	ImageColorCameraRGB = 0,
	ImageColorREDspace = 2,
	ImageColorREDcolor = 3,				// only available with color version 2, otherwise ignored (REDspace will be used)
	ImageColorREDcolor2 = 6				// only available with color version 2, otherwise ignored (REDspace will be used)
};

// This structure gives access to the default, minimum and maximum setting
// for each of the image processing parameters. In a few cases it also
// provides a list of the possible choices. You are encouraged to use this
// data in your User Interface so new options in future versions are
// automatically added without you having to do anything.
struct ImageProcessingLimits
{
	static const float KelvinMin, KelvinMax, KelvinDefault;
	static const float TintMin, TintMax, TintDefault;
	static const float ExposureMin, ExposureMax, ExposureDefault;
	static const float GainsMin, GainsMax, GainsDefault;
	static const float SaturationMin, SaturationMax, SaturationDefault;
	static const float ContrastMin, ContrastMax, ContrastDefault;
	static const float BrightnessMin, BrightnessMax, BrightnessDefault;
	static const float DRXMin, DRXMax, DRXDefault;
	static const float ShadowMin, ShadowMax, ShadowDefault;
	static const float FLUTMin, FLUTMax, FLUTDefault;

	static const float LggLiftMin , LggLiftMax , LggLiftDefault;
	static const float LggGammaMin, LggGammaMax, LggGammaDefault;
	static const float LggGainMin , LggGainMax , LggGainDefault;

	static const size_t				ISOCount;			// Number of ISO values in ISOList below
	static const size_t				ISOList[];			// List of valid ISO values	
	static const size_t				ISODefault;			// default ISO value

	static const size_t				GammaCurveCount;	// Number of gamma curve values in GammaCurveList below
	static const char *				GammaCurveLabels[];	// ImageGammaCustomPDlog is not included because it's not a preset
	static const ImageGammaCurve	GammaCurveMap[];	// List of gamma options
	static const ImageGammaCurve	GammaCurveDefault;	// default gamma value

	static const size_t				ColorSpaceCount;	// Number of color space values in ColorSpaceList below
	static const char *				ColorSpaceLabels[];
	static const ImageColorSpace	ColorSpaceMap[];	// List of color space options
	static const ImageColorSpace	ColorSpaceDefault;	// default color space value

	static const size_t CustomPDLogBlackPointMin, CustomPDLogBlackPointMax, CustomPDLogBlackPointDefault;
	static const size_t CustomPDLogWhitePointMin, CustomPDLogWhitePointMax, CustomPDLogWhitePointDefault;
	static const float CustomPDLogGammaMin, CustomPDLogGammaMax, CustomPDLogGammaDefault;

private:
	ImageProcessingLimits();							// all data is static, don't instantiate this struct
};

// Amount of detail extraction (not sharpening!) for
// the full resolution premium software decode only!
enum ImageDetail
{
	ImageDetailLow = 0,
	ImageDetailMedium = 1,
	ImageDetailHigh = 2									// default
};

// Compensation for the OLPF (low-pass) filter for
// the half & full resolution premium software decodes
enum ImageOLPFCompensation
{
	ImageOLPFCompOff = 0,
	ImageOLPFCompLow = 1,
	ImageOLPFCompMedium = 2,
	ImageOLPFCompHigh = 3
};

// Amount of noise reduction to do for the
// full resolution premium software decode
enum ImageDenoise
{
	ImageDenoiseOff = 0,
	ImageDenoiseMinimum = 1,
	ImageDenoiseMilder = 2,
	ImageDenoiseMild = 3,
	ImageDenoiseMedium = 4,
	ImageDenoiseStrong = 5,
	ImageDenoiseMaximum = 6
};

// Which color version to use with the image processing settings.
// Default is ColorVersion2, which is the new FLUT color science.
enum ColorVersion
{
	ColorVersion1 = 1,
	ColorVersion2 = 2
};

// Lift, gamma and gain settings for a color channel. Please see
// our Redcine-X application, available on www.red.com/support on
// how these controls work.
struct LiftGammaGain
{
	float	Lift;										// -1 -- 1
	float	Gamma;										//  0 -- 4
	float	Gain;										//  0 -- 2
};

// Settings to process the image with. If these settings are not
// supplied the decoder will use the clip's default settings.
struct ImageProcessingSettings
{
	// *********************************************************************
	// Color version defaults to the latest and greatest (V2 for the current
	// SDK release). With the exception of FLUT and Shadow, all image
	// processing setings are available for both color versions. For older
	// clips you may want to offer a backwards compatibility decode setting.
	// To do this set Version to ColorVersion1. FLUT, Shadow, as well as
	// REDgamma and REDcolor options will not be available with the old color!
	// Settings out of bounds (see ImageProcessingLimits) will be clipped.
	// *********************************************************************
	ColorVersion			Version;					// if ColorVersion1 is set for a ColorVersion2 clip it will be forced to ColorVersion2

	float					Kelvin;						// Kelvin color temperature for white balance
	float					Tint;						// Tint for white balance
	float					ExposureCompensation;
	float					GainRed, GainGreen, GainBlue;
	float					Saturation;
	float					Contrast;
	float					Brightness;
	float					DRX;						// highlight extension control. No effect for Dragon clips.

	// See ImageProcessingLimits::ISOList for the valid list of ISO values
	size_t					ISO;						// will be reset to 320 if it's set to an invalid value
	ImageGammaCurve			GammaCurve;
	ImageColorSpace			ColorSpace;

	// Set all of these fields when GammaCurve is set to ImageGammaCustomPDlog 
	size_t					CustomPDLogBlackPoint;
	size_t					CustomPDLogWhitePoint;
	float					CustomPDLogGamma;

	// User defined 5-point luma spline curve. Each of the 5 points consists of
	// an X & Y coordinate (0 -- 1) from lower left to upper right. Array
	// goes: black X, black Y, low X, low Y, mid X, mid Y, high X, high Y, white X, white Y
	// The curve coordinate values should be constrained in both UI and
	// generating code so that the x coordinates always lie such that:
	// black X < low X < mid X < high X < white X. If this is not the case
	// the curve will default to a straight line (0.0, 0.0, 0.25, 0.25,
	// 0.50, 0.50, 0.75, 0.75, 1.0, 1.0).
	float					UserCurve[10];

	// Following setting is only used for full & half res premium decode
	// for Dragon footage. THIS SETTING IS IGNORED WHEN DECODING THROUGH
	// RED ROCKET(-X) OR FOOTAGE IS NOT DRAGON!!
	bool					DEB;						// Dragon Enhanced Blacks

	// Following setting is only used for full res premium decode
	// THIS SETTING IS IGNORED WHEN DECODING THROUGH RED ROCKET!
	ImageDetail				Detail;
	
	// Following is only used for full & half res premium decodes
	// THIS SETTING IS IGNORED WHEN DECODING THROUGH RED ROCKET!
	ImageOLPFCompensation	OLPFCompensation;

	// Following setting is only used for the full res premium decode
	// THIS SETTING IS IGNORED WHEN DECODING THROUGH RED ROCKET!
	ImageDenoise			Denoise;

	// *********************************************************************
	// Following settings are only available for ColorVersion2 which was
	// introduced with R3D SDK 3.0 (this is the default version).
	// *********************************************************************

	float					Shadow;
	float					FLUT;						// fine ISO control

	// Red, Green & Blue user defined spline curves. Works the same as the luma UserCurve
	float					UserCurveRed[10];
	float					UserCurveGreen[10];
	float					UserCurveBlue[10];

	// post lift, gamma & gain for Red, Green and Blue channels
	LiftGammaGain			LggRed;
	LiftGammaGain			LggGreen;
	LiftGammaGain			LggBlue;

	// *********************************************************************
	// Function definitions
	// *********************************************************************

	// Sets all fields to their defaults. Keep in mind that you should
	// load the actual clip settings from metadata before setting up
	// the User Interface or using these settings to decode a frame!
	ImageProcessingSettings();

	// Check to see if all parameters are within the proper range and clip any that aren't
	void CheckBounds();
};

// The different resolutions and qualities the clip can be
// decoded at. This list expand over time.
enum VideoDecodeMode
{
	DECODE_FULL_RES_PREMIUM		= 0x44465250,			// 'DFRP', full resolution, slow, but highest resolution & quality
	
	DECODE_HALF_RES_PREMIUM		= 0x44485250,			// 'DHRP', half resolution, slow, but highest quality 
	DECODE_HALF_RES_GOOD		= 0x44485247,			// 'DHRG', half resolution, fast, still very good quality
	
	DECODE_QUARTER_RES_GOOD		= 0x44515247,			// 'DQRG', quarter resolution, fast, very good quality
	DECODE_EIGHT_RES_GOOD		= 0x44455247,			// 'DERG', eight resolution, fast, good quality
	DECODE_SIXTEENTH_RES_GOOD	= 0x44535247,			// 'DSRG', sixteenth resolution, fast, good quality 
	
	// This can only be set for a RED Rocket decode, not for a software decode!
	DECODE_ROCKET_CUSTOM_RES    = 0x44524352			// 'DRCR', arbitrary scale, very good to good quality
};

// The different pixel types to decode images at. The 16-bit RGB
// planar format has always existed in the SDK but is not supported
// for RED Rocket decoding. The other two formats are interleaved
// and are supported for both software and RED Rocket decoding.
enum VideoPixelType
{
	// Supported by software decode only
	PixelType_16Bit_RGB_Planar			= 0x52423650,	// Planar RGB decoding in 16-bits per pixel

	// Supported by software and RED Rocket decode
	PixelType_16Bit_RGB_Interleaved		= 0x52423649,	// Interleaved RGB decoding in 16-bits per pixel
														// 4K decodes in this 16-bit pixel type using the
														// Rocket will not be real time. Use the 12- or
														// 10-bit pixel types below for 4K real time needs

	PixelType_8Bit_BGRA_Interleaved		= 0x42524138,	// Interleaved BGRA decoding in 8 bits per pixel, alpha channel = 0xFF

	PixelType_10Bit_DPX_MethodB			= 0x44503042,	// Interleaved RGB 10-bit DPX Method B
	
	PixelType_12Bit_BGR_Interleaved		= 0x42523132,	// Interleaved BGR 12-bit (2 pixels packed in 9 bytes)
														// BBBBBBBB BBBBGGGG GGGGGGGG RRRRRRRR RRRRBBBB BBBBBBBB GGGGGGGGG GGGGRRRR RRRRRRRR
														// This 12-bit pixel type allows for 4K real time decoding
														// on the Rocket without sacrificing bit depth

	PixelType_8Bit_BGR_Interleaved		= 0x42475238,	// Interleaved BGR 8-bit

	PixelType_HalfFloat_RGB_Interleaved = 0x52424846,	// 16-bit half-float decoding. These ImageProcessingSettings fields are ignored:
														//		- Version (will always be set to ColorVersion2)
														//		- RGB gains
														//		- Contrast
														//		- Brightness
														//		- GammaCurve
														//		- CustomPDLogBlackPoint, CustomPDLogWhitePoint & CustomPDLogGamma
														//		- UserCurve, UserCurveRed, UserCurveGreen & UserCurveBlue
														//		- Shadow
														//		- LggRed, LggGreen & LggBlue

	PixelType_HalfFloat_RGB_ACES_Int	= 0x52424841	// 16-bit half-float ACES decoding. These ImageProcessingSettings fields are ignored:
														//		- Same fields as PixelType_HalfFloat_RGB_Interleaved pixel type
														//		- ColorSpace (will always be set to ACES)
};

// HDRx blending algorithm to use when doing HDR blending
enum HdrBlendAlgorithm
{
	HDRx_SIMPLE_BLEND			= 0x48445253,			// Simple blend, exactly as implemented in 12_SimpleHDRxBlend
														// On the Rocket this will ALWAYS decode in 16-bit to do the blend
														// This can result in slow downs at full resolution. If you need a
														// preview option in 8-bit then decode 2 images in 8-bit and
														// implement the simple blend as per sample code 12_SimpleHDRxBlend
	
	HDRx_MAGIC_MOTION			= 0x4844524D			// Magic Motion, more sophisticated blend
};

// This structure gives access to the default, minimum and maximum setting
// for each of the HDR processing parameters. In a few cases it also
// provides a list of the possible choices. You are encouraged to use this
// data in your User Interface so new options in future versions are
// automatically added without you having to do anything.
struct HdrProcessingLimits
{
	static const size_t					BlendAlgorithmCount;
	static const char *					BlendAlgorithmLabels[];
	static const HdrBlendAlgorithm		BlendAlgorithmMap[];
	static const HdrBlendAlgorithm		BlendAlgorithmDefault;

	static const float					BiasMin, BiasMax, BiasDefault;

private:
	HdrProcessingLimits();								// all data is static, don't instantiate this struct
};

// HDR blending settings. This structure needs to be constructed
// to submit a HDR blending job to the decoder. Every field must
// be set.
struct HdrProcessingSettings
{
	// HDRx blending algorithm to use
	HdrBlendAlgorithm			BlendAlgorithm;

	// Amount of blending for the modes, must be in range of -1 -- +1 (will be clipped if it's not!)
	float						Bias;					// Bias is mode dependent. For Simple Blend it controls the blend
														// between the two images, +1 being full normal exposure, -1 being
														// full highlight track exposure and 0 being an equal blend of the two.
														// For Magic Motion, it adjusts the amount of highlight rolloff in the
														// blended HDR image. +1 is no highlight rolloff and -1 is full rolloff

	// Sets all fields to their defaults.
	HdrProcessingSettings();

	// Check to see if all parameters are within the proper range and clip any that aren't
	void CheckBounds();
};

// HDR read status & write setting when reading & writing RMD sidecar files
enum HdrMode
{
	HDR_USE_TRACKNO,									// Use A / main (#0) or X (#1) track
	HDR_DO_BLEND										// Blend A and X tracks using specified settings
};

// The decode job structure. This structure needs to be constructed
// to submit a decode job to the decoder. Every field must be set,
// with the exception of 'ImageProcessing' & "HdrProcessing'
struct VideoDecodeJob
{
	// Resolution/speed to decode the image at. This will also 
	// influence how much memory is needed for the output buffer
	VideoDecodeMode				Mode;

	// Which pixel type to decode the image in. See the VideoPixelType
	// enum for the available options. To get the image in the
	// original pixel type supported by earlier SDK's set
	// PixelType_16Bit_RGB_Planar
	VideoPixelType				PixelType;

	// Pointer to the buffer to store the image in. This cannot be
	// NULL otherwise the decode will fail. The buffer must be aligned
	// on a 16-byte boundary (see sample code).
	void *						OutputBuffer;

	// Number of bytes in a row. This value must also be a multiple
	// of 16 for alignment reasons. If it is too small or not a
	// multiple of 16 the decode will fail.
	size_t						BytesPerRow;

	// Size ouf the output buffer (this should at least be
	// BytesPerRow * 3 * Heigh_of_Image). If this value is too
	// small the decode will fail.
	size_t						OutputBufferSize;
	
	// Image processing settings to apply to the decode. This
	// parameter is optional (set it to NULL if not used!).
	// The decoder will simply use the clip's default settings
	// if no image processing structure is used to override.
	ImageProcessingSettings	*	ImageProcessing;

	// HDRx processing settings to apply to the decode. This
	// parameter is optional (set it to NULL if not used!).
	// If the clip is not an HDRx clip, the decode will fail.
	// If the clip is an HDRx clip and no settings are supplied
	// the main track (A track) will be decoded, or the specified
	// track when using VideoTrackDecodeFrame(). Video track
	// number is ignored when this field is not NULL and you're
	// using VideoTrackDecodeFrame() API!
	HdrProcessingSettings *		HdrProcessing;

	VideoDecodeJob();
};

}
#endif

