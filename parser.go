package parser

import (
	"encoding/json"
	"errors"
	"os"

	xj "github.com/basgys/goxml2json"
)

// ParseCourse into Course object
func ParseCourse() (Course, error) {
	course, err := parseCourseRoot()
	return course, err
}

func parseCourseRoot() (Course, error) {
	courseRootXMLReader, err := os.Open("course/course/course.xml")
	if err != nil {
		return Course{}, err
	}
	courseRootJSONBytes, err := xj.Convert(courseRootXMLReader)
	if err != nil {
		return Course{}, err
	}

	var courseRootJSON interface{}
	err = json.Unmarshal(courseRootJSONBytes.Bytes(), &courseRootJSON)
	if err != nil {
		return Course{}, err
	}

	courseRootJSONMap, ok := courseRootJSON.(map[string]interface{})
	if !ok {
		return Course{}, errors.New("Couldn't convert JSON bytes to map")
	}
	courseInformation, ok := courseRootJSONMap["course"]
	if !ok {
		return Course{}, errors.New("Course is not in JSON")
	}

	courseInformationMap, ok := courseInformation.(map[string]interface{})
	if !ok {
		return Course{}, errors.New("Couldn't conevrt JSON to map")
	}

	courseDisplayName, ok := courseInformationMap["-display_name"]
	if !ok {
		return Course{}, errors.New("Couldn't get display name of the course")
	}

	displayName, ok := courseDisplayName.(string)
	if !ok {
		return Course{}, errors.New("Display name is not a string")
	}

	// Chapter extraction
	chapters := make([]Chapter, 0)
	chaptersGeneric, ok := courseInformationMap["chapter"]
	if !ok {
		return Course{}, err
	}

	// fmt.Println(chaptersGeneric)

	chaptersArrGeneric, ok := chaptersGeneric.([]interface{})
	if !ok {
		// TODO: It is not necessary an array
		return Course{}, errors.New("There are not enough")
	}
	// fmt.Println(chaptersArrGeneric)

	for _, chapter := range chaptersArrGeneric {
		chapterMap, ok := chapter.(map[string]interface{})
		if !ok {
			return Course{}, errors.New("One chapter is not a map")
		}
		chapterID := chapterMap["-url_name"].(string)
		chapters = append(chapters, Chapter{
			ID: chapterID,
		})
	}

	course := Course{
		DisplayName: displayName,
		Chapters:    chapters,
	}
	return course, nil
}
