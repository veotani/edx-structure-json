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
	course := Course{
		DisplayName: displayName,
	}
	return course, nil
}
