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

			for verticalNum, vertical := range sequential.Verticals {
				err = parseByStructure(&course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum], CoursePath+"/vertical/"+vertical.URLName+".xml")
				if err != nil {
					return Course{}, err
				}

				for htmlNum, html := range vertical.Htmls {
					err = parseByStructure(&course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum].Htmls[htmlNum], CoursePath+"/html/"+html.URLName+".xml")
					if err != nil {
						return Course{}, err
					}
				}

				for problemNum, problem := range vertical.Problems {
					err = parseByStructure(&course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum].Problems[problemNum], CoursePath+"/problem/"+problem.URLName+".xml")
					if err != nil {
						return Course{}, err
					}
				}

				for videoNum, video := range vertical.Videos {
					videoHelper := VideoHelper{}
					err = parseByStructure(&videoHelper, CoursePath+"/video/"+video.URLName+".xml")
					if err != nil {
						return Course{}, err
					}
					course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum].Videos[videoNum] = videoHelper.ToVideo()
				}

				for libraryContentNum, libraryContent := range vertical.LibraryContents {
					err = parseByStructure(&course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum].LibraryContents[libraryContentNum], CoursePath+"/library_content/"+libraryContent.URLName+".xml")
					if err != nil {
						return Course{}, err
					}

					for problemNum, problem := range libraryContent.Problems {
						err = parseByStructure(&course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum].LibraryContents[libraryContentNum].Problems[problemNum], CoursePath+"/problem/"+problem.URLName+".xml")
						if err != nil {
							return Course{}, err
						}
					}
				}
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
