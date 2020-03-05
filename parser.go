package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// CoursePath is a path to folder with course structure in XML format
const CoursePath string = "course/"

// ParseCourse parses course structure recursively from root to leaves
func ParseCourse() (Course, error) {
	course, err := parseCourseRoot()
	return course, err
}

// CourseXML is a description of CoursePath + "/course/course.xml" XML file
type CourseXML struct {
	DisplayName string       `xml:"display_name,attr"`
	Chapters    []ChapterXML `xml:"chapter"`
}

// ChapterXML is a description of CoursePath + "/chapter/" + ChapterID + ".xml" XML file
type ChapterXML struct {
	URLName     string          `xml:"url_name,attr"`
	DisplayName string          `xml:"display_name,attr"`
	Sequentials []SequentialXML `xml:"sequential"`
}

// SequentialXML is a description of CoursePath + "/sequential/" + SequentialID + ".xml" XML file
type SequentialXML struct {
	URLName string `xml:"url_name,attr"`
	// TODO: add CoursePath + "/sequential/" + SequentialID + ".xml" file structure
}

func parseCourseRoot() (Course, error) {
	courseRootXMLReader, err := ioutil.ReadFile(CoursePath + "/course/course.xml")
	if err != nil {
		return Course{}, err
	}

	courseXML := CourseXML{}
	err = xml.Unmarshal(courseRootXMLReader, &courseXML)
	if err != nil {
		return Course{}, err
	}

	for _, chapter := range courseXML.Chapters {
		parseChapter(chapter.URLName)
	}

	return Course{}, nil
}

func parseChapter(chapterURL string) (Chapter, error) {
	chapterXMLReader, err := ioutil.ReadFile(CoursePath + "/chapter/" + chapterURL + ".xml")
	if err != nil {
		return Chapter{}, err
	}

	chapterXML := ChapterXML{}
	err = xml.Unmarshal(chapterXMLReader, &chapterXML)
	fmt.Println(chapterXML)
	if err != nil {
		return Chapter{}, err
	}

	for _, sequential := range chapterXML.Sequentials {
		fmt.Println(sequential)
		// Parse (CoursePath + "/sequential/" + SequentialID + ".xml")
		// parseSequential(sequential)
	}

	return Chapter{}, nil
}
