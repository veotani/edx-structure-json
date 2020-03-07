package parser

import (
	"encoding/xml"
	"io/ioutil"
)

// CoursePath is a path to folder with course structure in XML format
const CoursePath string = "course/"

// ParseCourse parses course structure recursively from root to leaves
func ParseCourse() (Course, error) {
	course, err := parseCourseRoot()
	return course, err
}

func parseCourseRoot() (Course, error) {
	course := &Course{}
	err := parseByStructure(course, CoursePath+"/course/course.xml")
	if err != nil {
		return Course{}, err
	}
	for chapterNum, chapter := range course.Chapters {
		err = parseByStructure(&course.Chapters[chapterNum], CoursePath+"/chapter/"+chapter.URLName+".xml")
		if err != nil {
			return Course{}, err
		}

		for sequentialNum, sequential := range chapter.Sequentials {
			err = parseByStructure(&course.Chapters[chapterNum].Sequentials[sequentialNum], CoursePath+"/sequential/"+sequential.URLName+".xml")
			if err != nil {
				return Course{}, err
			}
		}
	}

	return *course, err
}

func parseByStructure(structureObject interface{}, filePath string) error {
	fileToParse, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(fileToParse, structureObject)
	if err != nil {
		return err
	}

	return nil
}
