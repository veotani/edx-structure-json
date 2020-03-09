package parser

//                          Course
//                             |
//                          Chapter
//                             |
//                         Sequential
//                             |
//            --------------Vertical---------------
//           /         /       |         \         \
//       Problem   Discussion Html OpenAssessment Video

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
	OpenAssessments []OpenAssessment `xml:"openassessment"`
	LibraryContents []LibraryContent `xml:"library_content"`
	Videos          []Video          `xml:"video"`
}

// LibraryContent is a block with (usually) many problems within it. It serves as a randomizer
// for problems. LibraryContent XML files are in CoursePath + "/library_content/" folder.
type LibraryContent struct {
	DisplayName string    `xml:"display_name,attr"`
	URLName     string    `xml:"url_name,attr"`
	Problems    []Problem `xml:"problem"`
}

// OpenAssessment is a special type of problems, where course learner gives free answer and leaves
// his peer review for another learner's answer (e. g. essay). His answer is also being graded by
// another students. These objects are usually described within `vertical` element.
// TODO: add other OpenAssessment fields
type OpenAssessment struct {
	URLName string `xml:"url_name,attr"`
}

// Html is an HTML block. Is contains HTML code within it
type Html struct {
	DisplayName string `xml:"display_name,attr"`
	URLName     string `xml:"url_name,attr"`
}

// Video describes video. VideoHelper is used to parse it.
type Video struct {
	DisplayName string
	URLName     string `xml:"url_name,attr"`
	Duration    string
}

// Discussion is a forum discussion.
type Discussion struct {
	URLName     string `xml:"url_name,attr"`
	DisplayName string `xml:"display_name,attr"`
}

// Problem have difficult structure within it, so only name and
// ID are being parsed.
type Problem struct {
	URLName     string `xml:"url_name,attr"`
	DisplayName string `xml:"display_name,attr"`
}
