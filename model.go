package parser

// Course is a description of CoursePath + "/course/course.xml" XML file
type Course struct {
	DisplayName string    `xml:"display_name,attr"`
	Chapters    []Chapter `xml:"chapter"`
}

// Chapter is a description of CoursePath + "/chapter/" + ChapterID + ".xml" XML file
type Chapter struct {
	URLName     string       `xml:"url_name,attr"`
	DisplayName string       `xml:"display_name,attr"`
	Sequentials []Sequential `xml:"sequential"`
}

// Sequential is a description of CoursePath + "/sequential/" + SequentialID + ".xml" XML file
// type Sequential struct {
// 	URLName string `xml:"url_name,attr"`
// 	// TODO: add CoursePath + "/sequential/" + SequentialID + ".xml" file structure
// }

// Sequential is a description of CoursePath + "/sequential/" + SequentialID + ".xml" XML file
type Sequential struct {
	DisplayName string     `xml:"display_name,attr"`
	URLName     string     `xml:"url_name,attr"`
	Verticals   []Vertical `xml:"vertical"`
}

// Vertical is a description of CoursePath + "/vertical/" + VerticalID + ".xml" XML file
type Vertical struct {
	DisplayName     string           `xml:"display_name,attr"`
	URLName         string           `xml:"url_name,attr"`
	Problems        []Problem        `xml:"problem"`
	Discussions     []Discussion     `xml:"discussion"`
	Htmls           []Html           `xml:"html"`
	OpenAssessments []OpenAssessment `xml:"open_assessment"`
	LibraryContents []LibraryContent `xml:"library_content"`
}

type LibraryContent struct {
	// 	DisplayName string
	URLName string `xml:"url_name,attr"`
	// 	// Problems    []Problem
}

type OpenAssessment struct {
	// DisplayName string
	URLName string `xml:"url_name,attr"`
}

type Html struct {
	// DisplayName string
	URLName string `xml:"url_name,attr"`
}

type Video struct {
	// 	DisplayName string
	URLName string `xml:"url_name,attr"`
}

type Discussion struct {
	URLName string `xml:"url_name,attr"`
	// 	DisplayName string
}

type Problem struct {
	URLName string `xml:"url_name,attr"`
	// 	DisplayName string
}
