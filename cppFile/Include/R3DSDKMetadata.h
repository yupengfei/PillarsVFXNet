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

#ifndef R3DSDKMETADATA_H
#define R3DSDKMETADATA_H

namespace R3DSDK {

// Types of data that can be stored in the metadata database
enum MetadataType
{
	MetadataTypeInvalid = 0,
	MetadataTypeInt = 1,
	MetadataTypeString = 2,
	MetadataTypeFloat = 3
};

// For all metadata items an int, float & string output can be requested.
// Below a table showing the conversions that can happen:
//
// Original type | Requested type | Description
// --------------|----------------|-----------------------------------------
// Int           | Float          | Int to float
// Int           | String         | Int to string, no padding or anything
// --------------|----------------|-----------------------------------------
// Float         | Int            | Rounded float to int (23.976 becomes 24)
// Float         | String         | Float to string, no padding or anything
// --------------|----------------|-----------------------------------------
// String        | Int            | Length of string
// String        | Float          | Length of string converted to float
//
// The MetadataItemType() methods return the original data type so you
// can avoid any conversions if that's desirable.
//
// Most methods have a key as well as index variant. The former is most useful
// if you exactly know which metadata item to query, the search is case
// insensitive. The latter is most useful when enumerating through the metadata
// list. When requesting items by key make sure to call MetadataExists() to
// be sure the item actually exists in the database. Otherwise you will get
// a return value of 0, 0.0 or "" depending on which output type you request.
//
// Below a list of items that might exist in the metadata database. These keys
// can be used to quickly get a specific metadata item. You CANNOT depend on
// the metadata items existing at a specific index!! Keep in mind that only a
// couple of items are guaranteed to exist in the database.
// DO NOT COUNT ON ALL ITEMS BEING AVAILABLE
static const char RMD_ANAMORPHIC_MODE[] =				"anamorphic_mode";				// different from of pixel aspect ratio. 0 = not anamorphic, 1 = 2x, 2 = 1.3x
static const char RMD_BRIGHTNESS[] =					"brightness";					// Also available through the image processing settings structure
static const char RMD_CAMERA_FIRMWARE_REVISION[] =		"camera_firmware_revision";
static const char RMD_CAMERA_FIRMWARE_VERSION[] =		"camera_firmware_version";
static const char RMD_CAMERA_ID[] =						"camera_id";					// Camera ID as set in camera
static const char RMD_CAMERA_MODEL[] =					"camera_model";					// "RED ONE", "EPIC-M", "EPIC-X", "SCARLET-X", etc.
static const char RMD_CAMERA_MODEL_ID[] =				"camera_model_id";				// 1 = RED ONE, 2 = SCARLET 2/3" 8x FIXED ZOOM, 3 = SCARLET-X, 4 = EPIC-M,
																						// 5 = EPIC-X, 6 = EPIC-X MONOCHROME, 7 = EPIC-M MONOCHROME
static const char RMD_CAMERA_NETWORK_NAME[] =			"camera_network_name";
static const char RMD_CAMERA_PIN[] =					"camera_pin";					// Camera's PIN (serial number)
static const char RMD_CLIP_ASPECT_RATIO[] =				"clip_aspect_ratio";
static const char RMD_CLIP_ASPECT_DENOMINATOR[] =		"clip_aspect_ratio_denominator";
static const char RMD_CLIP_ASPECT_NUMERATOR[] =			"clip_aspect_ratio_numerator";
static const char RMD_CLIP_ID[] =						"clip_id";						// Clip ID as set in camera
static const char RMD_CONTRAST[] =						"contrast";						// Also available through the image processing settings structure
static const char RMD_DIGITAL_GAIN_BLUE[] =				"digital_gain_blue";			// Also available through the image processing settings structure
static const char RMD_DIGITAL_GAIN_GREEN[] =			"digital_gain_green";			// Also available through the image processing settings structure
static const char RMD_DIGITAL_GAIN_RED[] =				"digital_gain_red";				// Also available through the image processing settings structure
static const char RMD_DROPPED_FRAME_COUNT[] =			"dropped_frame_count";
static const char RMD_EXPOSURE_COMPENSATION[] =			"exposure_compensation";		// Guaranteed to exist, also available through the image processing settings structure
static const char RMD_EXPOSURE_TIME[] =					"exposure_time";				// Length the frame was exposed for in microseconds
static const char RMD_FLIP_HORIZONTAL[] =				"flip_horizontal";				// 0 if image does not need to be flipped horizontally, 1 if it should
static const char RMD_FLIP_VERTICAL[] =					"flip_vertical";				// 0 if image does not need to be flipped vertically, 1 if it should
static const char RMD_FLUT_CONTROL[] =					"flut_control";
static const char RMD_FRAME_GUIDE_ENABLED[] =			"frame_guide_enabled";			// 0 = disabled, 1 = enabled
static const char RMD_FRAME_GUIDE_NAME[] =				"frame_guide_name";
static const char RMD_FRAME_GUIDE_TOPLEFT_X[] =			"frame_guide_topleft_x";		// float, 0.0 - 1.0, multiply by image width - 1
static const char RMD_FRAME_GUIDE_TOPLEFT_Y[] =			"frame_guide_topleft_y";		// float, 0.0 - 1.0, multiply by image height - 1
static const char RMD_FRAME_GUIDE_WIDTH[] =				"frame_guide_width";			// float, 0.0 - 1.0, multiply by image width
static const char RMD_FRAME_GUIDE_HEIGHT[] =			"frame_guide_height";			// float, 0.0 - 1.0, multiply by image height
static const char RMD_FRAMERATE[] =						"framerate";					// Calculated from RMD_FRAMERATE_NUMERATOR / RMD_FRAMERATE_DENOMINATOR
static const char RMD_FRAMERATE_DENOMINATOR[] =			"framerate_denominator";		// Guaranteed to exist (also true for RMD_FRAMERATE above)
static const char RMD_FRAMERATE_NUMERATOR[] =			"framerate_numerator";			// Guaranteed to exist
static const char RMD_GENLOCK_SETTING[] =				"genlock_setting";				// 0 = not genlocked, 1 = genlocked to external source
static const char RMD_GMT_DATE[] =						"gmt_date";
static const char RMD_GMT_TIME[] =						"gmt_time";
static const char RMD_HDR_MODE[] =						"hdr_mode";						// 0 = no HDR, 2 = HDRx
static const char RMD_HDR_STOPS_OVER[] =				"hdr_stops_over";				// number f-stops the X track is over the A (main) track
static const char RMD_IMAGE_HEIGHT[] =					"image_height";					// Guaranteed to exist, also available as clip property
static const char RMD_IMAGE_WIDTH[] =					"image_width";					// Guaranteed to exist, also available as clip property
static const char RMD_ISO[] =							"iso";							// Guaranteed to exist, also available through the image processing settings structure
static const char RMD_JAMSYNC_SETTING[] =				"jamsync_setting";				// 0 = internal clock, 1 = external clock used
static const char RMD_LENS_APERTURE[] =					"lens_aperture";				// Aperture * 10, 0xFFFF = fully closed
static const char RMD_LENS_FOCAL_LENGTH[] =				"lens_focal_length";			// in mm
static const char RMD_LENS_FOCUS_DISTANCE[] =			"lens_focus_distance";			// in mm, 0xFFFFFFFF = infinite
static const char RMD_LENS_FOCUS_DISTANCE_NEAR[] =		"lens_focus_distance_near";
static const char RMD_LENS_FOCUS_DISTANCE_FAR[] =		"lens_focus_distance_far";
static const char RMD_LENS_MOUNT[] =					"lens_mount";					// if present will be the name of the mount
static const char RMD_LINKED_CAMERA_SETUP[] =			"linked_camera_setup";			// 0 = camera wasn't linked, 1 = clip was recorded on master camera, 2 = clip was recorded on (a) slave camera
static const char RMD_LOCAL_DATE[] =					"local_date";
static const char RMD_LOCAL_TIME[] =					"local_time";
static const char RMD_MEDIA_SERIAL_NUMBER[] =			"media_serial_number";
static const char RMD_MOTION_MOUNT_ND_STOPS[] =			"motion_mount_nd_stops";		// DSMC Motion Mount ND value in stops
static const char RMD_MOTION_MOUNT_SHUTTER_TYPE[] =		"motion_mount_shutter_type";	// DSMC Motion Mount Shutter Type: 0 = ND only, 1 = Soft Shutter, 2 = Square Shutter
static const char RMD_ORIGINAL_FILENAME[] =				"original_filename";			// original camera filename for first R3D file in the clip
static const char RMD_PIXEL_ASPECT_RATIO[] =			"pixel_aspect_ratio";
static const char RMD_RECORD_FRAMERATE[] =				"record_framerate";				// Framerate the recording was made at
static const char RMD_REDCODE[] =						"redcode";						// String describing the REDCODE format the clip was recorded with
static const char RMD_REEL_ID[] =						"reel_id";						// Reel ID as set in camera
static const char RMD_REEL_ID_FULL[] =					"reel_id_full";					// REDCINE-X PRO style Reel ID Full
static const char RMD_REEL_ID_8_CHAR[] =				"reel_id_8_character";			// REDCINE-X PRO style Reel ID 8 Character for EDL's
static const char RMD_RESOLUTION_FORMAT_NAME[] =		"resolution_format_name";		// Name for the clip format (resolution)
static const char RMD_SATURATION[] =					"saturation";					// Also available through the image processing settings structure
static const char RMD_SENSOR_ID[] =						"sensor_id";					// 0 = MYSTERIUM S35, 1 = MYSTERIUM-X S35, 2 = MYSTERIUM-X 2/3", 3 = MYSTERIUM-X S35 MONOCHROME,
																						// 4 = DRAGON S35, 5 = DRAGON S35 MONOCHROME
static const char RMD_SENSOR_NAME[] =					"sensor_name";					// "MYSTERIUM S35", "MYSTERIUM-X S35", "DRAGON S35", etc.
static const char RMD_SENSOR_OLPF_INTERCHANGEABLE[] =	"sensor_olpf_interchangeable";	// 0 = no, 1 = yes
static const char RMD_SENSOR_OLPF_NAME[] =				"sensor_olpf_name";				// if camera has interchangeable sensor OLPF, this will contain the name of the OLPF installed
static const char RMD_SHADOW[] =						"shadow";
static const char RMD_SHUTTER_DEGREES[] =				"shutter_degrees";				// Shutter in degrees, calculated from RMD_EXPOSURE_TIME & RMD_FRAMERATE
static const char RMD_SHUTTER_FRACTIONS[] =				"shutter_fractions";			// Shutter in fractions (1/RMD_SHUTTER_FRACTIONS), see notes for shutter degrees
static const char RMD_SHUTTER_PHASE_OFFSET[] =			"shutter_phase_offset";
static const char RMD_START_ABSOLUTE_TIMECODE[] =		"start_absolute_timecode";		// Guaranteed to exist, in the form of "01:00:00:00"
static const char RMD_START_EDGE_TIMECODE[] =			"start_edge_timecode";			// Guaranteed to exist, in the form of "01:00:00:00"
static const char RMD_STEREO_SETUP[] =					"stereo_setup";					// 0 = not a stereo clip, 1 = left eye, 2 = right eye
static const char RMD_USER_ACTION_GUIDE_ENABLED[] =		"user_action_guide_enabled";	// 0 = disabled, 1 = enabled
static const char RMD_USER_ACTION_GUIDE_TOPLEFT_X[] =	"user_action_guide_topleft_x";	// float, 0.0 - 1.0, multiply by image width - 1
static const char RMD_USER_ACTION_GUIDE_TOPLEFT_Y[] =	"user_action_guide_topleft_y";	// float, 0.0 - 1.0, multiply by image height - 1
static const char RMD_USER_ACTION_GUIDE_WIDTH[] =		"user_action_guide_width";		// float, 0.0 - 1.0, multiply by image width
static const char RMD_USER_ACTION_GUIDE_HEIGHT[] =		"user_action_guide_height";		// float, 0.0 - 1.0, multiply by image height
static const char RMD_USER_CURVE_BLACK_X[] =			"user_curve_black_x";			// Also available through the image processing settings structure
static const char RMD_USER_CURVE_BLACK_Y[] =			"user_curve_black_y";
static const char RMD_USER_CURVE_LOW_X[] =				"user_curve_low_x";	
static const char RMD_USER_CURVE_LOW_Y[] =				"user_curve_low_y";	
static const char RMD_USER_CURVE_MID_X[] =				"user_curve_mid_x";	
static const char RMD_USER_CURVE_MID_Y[] =				"user_curve_mid_y";
static const char RMD_USER_CURVE_HIGH_X[] =				"user_curve_high_x";
static const char RMD_USER_CURVE_HIGH_Y[] =				"user_curve_high_y";
static const char RMD_USER_CURVE_WHITE_X[] =			"user_curve_white_x";
static const char RMD_USER_CURVE_WHITE_Y[] =			"user_curve_white_y";
static const char RMD_USER_TIMECODE_PREFERENCE[] =		"user_timecode_preference";		// 0 = edge / run record, 1 = absolute / external timecode, guaranteed to exist
static const char RMD_USER_TITLE_GUIDE_ENABLED[] =		"user_title_guide_enabled";		// 0 = disabled, 1 = enabled
static const char RMD_USER_TITLE_GUIDE_TOPLEFT_X[] =	"user_title_guide_topleft_x";	// float, 0.0 - 1.0, multiply by image width - 1
static const char RMD_USER_TITLE_GUIDE_TOPLEFT_Y[] =	"user_title_guide_topleft_y";	// float, 0.0 - 1.0, multiply by image height - 1
static const char RMD_USER_TITLE_GUIDE_WIDTH[] =		"user_title_guide_width";		// float, 0.0 - 1.0, multiply by image width
static const char RMD_USER_TITLE_GUIDE_HEIGHT[] =		"user_title_guide_height";		// float, 0.0 - 1.0, multiply by image height

static const char RMD_USER_CAMERA_OPERATOR[] =			"user_camera_operator";
static const char RMD_USER_COPYRIGHT[] =				"user_copyright";
static const char RMD_USER_DIRECTOR[] =					"user_director";
static const char RMD_USER_DIRECTOR_OF_PHOTOGRAPHY[] =	"user_director_of_photography";
static const char RMD_USER_LOCATION[] =					"user_location";
static const char RMD_USER_PRODUCTION_NAME[] =			"user_production_name";
static const char RMD_USER_SCENE[] =					"user_scene";
static const char RMD_USER_SHOT[] =						"user_shot";
static const char RMD_USER_TAKE[] =						"user_take";
static const char RMD_USER_UNIT[] =						"user_unit";

static const char RMD_WHITE_BALANCE_KELVIN[] =			"white_balance_kelvin";			// Also available through the image processing settings structure
static const char RMD_WHITE_BALANCE_TINT[] =			"white_balance_tint";			// Also available through the image processing settings structure

// Audio metadata items if a clip has embedded audio
static const char RMD_SAMPLERATE[] =					"samplerate";					// sample rate, should always be 48000
static const char RMD_CHANNEL_MASK[] =					"channel_mask";					// audio channel mask, bit 1 = channel 1 enabled ... bit 4 = channel 4 enabled, etc
static const char RMD_SAMPLE_SIZE[] =					"sample_size";					// sample size, should always be 24. This is the actual bits of raw audio data
																						// it might be stored differently like msb aligned in a 32-bit int

// The only metadata items you can count on that exist are framerate (all three), the two start
// timecode values, image width & height, ISO, exposure compensation and user timecode preference

}

#endif
