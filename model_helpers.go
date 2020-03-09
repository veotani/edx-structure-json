package parser

// There are objects that shouldn't look like their XML. This is why it's necessary
// to declare helper objects that have same structure as theese XML files and will
// serve for parsing. After parsing into theese objects their data is copied into
// destination object.
// In other words instead of having object `Video` with another object `VideoAsset`
// field we will have object Video, but parsing is going to be made by
// `VideoHelper`.

// VideoHelper is a description of XML files in course/video directory
type VideoHelper struct {
	DisplayName string       `xml:"display_name,attr"`
	URLName     string       `xml:"url_name,attr"`
	VideoAssets []VideoAsset `xml:"video_asset"`
}

// VideoAsset is a description of object within `video` element in course/video
// directory that describes video file. It is necessary to parse it to get
// duration of that video.
type VideoAsset struct {
	Duration string `xml:"duration,attr"`
}

// ToVideo converts VideoHelper to Video object.
// There is only 1 video asset and we haven't met any cases where there
// are more. If this situation happends, then this method should be rewritten
// and duration must not simple take first VideoAsset object's duration
func (videoHelper VideoHelper) ToVideo() Video {
	return Video{
		DisplayName: videoHelper.DisplayName,
		Duration:    videoHelper.VideoAssets[0].Duration,
		URLName:     videoHelper.URLName,
	}
}
