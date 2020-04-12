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
	DisplayName string    `xml:"display_name,attr" json:"display_name"`
	Chapters    []Chapter `xml:"chapter" json:"chapters"`
	CourseCode  string    `json:"course_code"`
	CourseRun   string    `json:"course_run"`
}

// Chapter is a description of CoursePath + "/chapter/" + ChapterID + ".xml" XML file
type Chapter struct {
	URLName     string       `xml:"url_name,attr" json:"url_name"`
	DisplayName string       `xml:"display_name,attr" json:"display_name"`
	Sequentials []Sequential `xml:"sequential" json:"sequentials"`
}

// Sequential is a description of CoursePath + "/sequential/" + SequentialID + ".xml" XML file
type Sequential struct {
	DisplayName string     `xml:"display_name,attr" json:"display_name"`
	URLName     string     `xml:"url_name,attr" json:"url_name"`
	Verticals   []Vertical `xml:"vertical" json:"verticals"`
}

// Vertical is a description of CoursePath + "/vertical/" + VerticalID + ".xml" XML file
type Vertical struct {
	DisplayName     string           `xml:"display_name,attr" json:"display_name"`
	URLName         string           `xml:"url_name,attr" json:"url_name"`
	Problems        []Problem        `xml:"problem" json:"problems"`
	Discussions     []Discussion     `xml:"discussion" json:"discussions"`
	Htmls           []Html           `xml:"html" json:"htmls"`
	OpenAssessments []OpenAssessment `xml:"openassessment" json:"open_assessments"`
	LibraryContents []LibraryContent `xml:"library_content" json:"library_contents"`
	Videos          []Video          `xml:"video" json:"videos"`
}

// LibraryContent is a block with (usually) many problems within it. It serves as a randomizer
// for problems. LibraryContent XML files are in CoursePath + "/library_content/" folder.
type LibraryContent struct {
	DisplayName string    `xml:"display_name,attr" json:"display_name"`
	URLName     string    `xml:"url_name,attr" json:"url_name"`
	Problems    []Problem `xml:"problem" json:"problems"`
}

// OpenAssessment is a special type of problems, where course learner gives free answer and leaves
// his peer review for another learner's answer (e. g. essay). His answer is also being graded by
// another students. These objects are usually described within `vertical` element.
// TODO: add other OpenAssessment fields
type OpenAssessment struct {
	URLName string `xml:"url_name,attr" json:"url_name"`
}

// Html is an HTML block. Is contains HTML code within it
type Html struct {
	DisplayName string `xml:"display_name,attr" json:"display_name"`
	URLName     string `xml:"url_name,attr" json:"url_name"`
}

// Video describes video. VideoHelper is used to parse it.
type Video struct {
	DisplayName string `json:"display_name"`
	URLName     string `xml:"url_name,attr" json:"url_name"`
	Duration    string `json:"duration"`
}

// Discussion is a forum discussion.
type Discussion struct {
	URLName string `xml:"url_name,attr" json:"url_name"`
}

// Problem have difficult structure within it, so only name and
// ID are being parsed.
type Problem struct {
	URLName     string `xml:"url_name,attr" json:"url_name"`
	DisplayName string `xml:"display_name,attr" json:"display_name"`
}
