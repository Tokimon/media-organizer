package tools

import "slices"

type ImageExtensions struct {
	Camera []string
	Other  []string
	Web    []string
}

type Extensions struct {
	Audio    []string
	Book     []string
	Code     []string
	Data     []string
	Image    ImageExtensions
	Settings []string
	Vector   []string
	Video    []string
}

func (ext *Extensions) Images() []string {
	return slices.Concat(
		ext.Image.Camera,
		ext.Image.Web,
		ext.Image.Other,
	)
}

func (ext *Extensions) All() []string {
	return slices.Concat(
		ext.Audio,
		ext.Book,
		ext.Code,
		ext.Data,
		ext.Image.Camera,
		ext.Image.Web,
		ext.Image.Other,
		ext.Settings,
		ext.Vector,
		ext.Video,
	)
}

var DefaultExtensions = &Extensions{
	Audio: []string{
		"f4a", // Adobe Flash Protected Audio File
		"m4a", // MPEG-4 Audio File
		"m4p", // iTunes Music Store Audio File
	},
	Book: []string{
		"aax", // Audible Enhanced Audiobook File
		"f4b", // Adobe Flash MP4 Audio EBook
		"m4b", // MPEG-4 Audiobook File
		"pdf", // Adobe Portable Document Format
	},
	Code: []string{
		"arq", // Remedy ARS Macro
		"dvb", // AutoCAD VBA Project File
		"ps",  // PostScript file
		"vrd", // Visio Report Definition File
	},
	Data: []string{
		"crm", // CHARTrunner Multi Chart Document
		"erf", // BioWare Entity Resource File
		"gpr", // GenePix Results File
		"hif", // Quicken Online File
		"ind", // Memory Stick Formatting File
		"mie", // MapImagery Encrypted Image
		"qif", // Quicken Interchange Format File
		"xmb", // X-Wing Mission Briefing File
		"xmp", // Extensible Metadata Platform File
	},
	Image: ImageExtensions{
		Camera: []string{
			"arw",  // Sony Digital Camera RAW Image Format
			"ciff", // Camera Image File Format (Canon raw)
			"crw",  // Canon Raw CIFF Image File
			"cr2",  // Canon Raw Image File
			"cr3",  // Canon Raw 3 Image File
			"cs1",  // CaptureShop 1-shot Raw Image
			"dng",  // Digital Negative Image File
			"fff",  // Hasselblad RAW Image
			"iiq",  // Phase One RAW Image
			"mef",  // Mamiya RAW Image
			"mos",  // Leaf Camera RAW File
			"mrw",  // Minolta Raw Image File
			"nef",  // Nikon Electronic Format RAW Image
			"nrw",  // Nikon Raw Image File
			"orf",  // Olympus RAW File
			"pef",  // Pentax Digital Camera Raw Image Format
			"raf",  // Fuji RAW Image File
			"raw",  // Raw Image Data File
			"rw2",  // Panasonic RAW Image
			"rwl",  // Leica RAW Image
			"sr2",  // Sony RAW Image
			"srw",  // Samsung RAW Image
			"tif",  // Tagged Image File
			"tiff", // Tagged Image File
			"x3f",  // SIGMA X3F Camera RAW File
		},
		Other: []string{
			"360",  // 360desktop Panorama File
			"flif", // Free Lossless Image Format
			"heic", // High Efficiency Image Format (Collection of HEIF files)
			"heif", // High Efficiency Image Format
			"hdp",  // HD Photo File
			"insp", // Insta360 Panoramic Image
			"jng",  // JPEG Network Graphic (Image file format related to the .PNG format, but uses lossy compression like standard .JPG files)
			"jp2",  // JPEG 2000 Image
			"jpf",  // JPEG 2000 Image
			"jpm",  // JPEG 2000 Multi-layer Image Format (ISO 15444-6)
			"jpx",  // JPEG 2000 Extended Image Format (ISO 15444-2)
			"jxr",  // JPEG XR Image
			"mng",  // Multiple Network Graphic (An extension of the .PNG image format)
			"mpo",  // Multi Picture Object File (consists of two 2D .JPG images that are combined into one 3D image)
			"psb",  // Photoshop Large Document Format
			"psd",  // Adobe Photoshop Document
			"qti",  // QuickTime Image File
			"qtif", // QuickTime Image File
			"wdp",  // Windows Media Photo File
		},
		Web: []string{
			"apng", // Animated Portable Network Graphic (animation like .gif, but for .png)
			"avif", // AV1 Image
			"bmp",  // Bitmap Image File
			"gif",  // Graphical Interchange Format File
			"ico",  // Icon File
			"jpeg", // JPEG Image (Joint Photographic Experts Group)
			"jpg",  // JPEG Image
			"jpe",  // JPEG Image
			"png",  // Portable Network Graphic
			"svg",  // Scalable Vector Graphics File
			"webp", // WebP Image
		},
	},
	Settings: []string{
		"dcp",  // Adobe DNG Camera Profile
		"exv",  // Adobe Extension Script
		"indt", // Adobe InDesign Template
		"psdt", // Adobe Photoshop Document Template
		"thm",  // Sony Ericsson Theme File
	},
	Vector: []string{
		"eps",  // Encapsulated PostScript File
		"epsf", // Encapsulated PostScript Format File
		"indd", // Adobe InDesign Document
	},
	Video: []string{
		"3g2",  // 3GPP2 Multimedia File
		"3gp2", // 3GPP Multimedia File
		"3gp",  // 3GPP Multimedia File
		"3gpp", // 3GPP Multimedia File
		"f4p",  // Adobe Flash Protected Media File
		"f4v",  // Flash MP4 Video File
		"lrv",  // Low Resolution Video File
		"m4v",  // iTunes Video File
		"mov",  // Apple QuickTime Movie
		"mp4",  // MPEG-4 Video File
		"mqv",  // Sony Movie Format File
		"qt",   // Apple QuickTime Movie
	},
}
