package parser

import (
	"archive/tar"
	"compress/gzip"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// CoursePath is a path to folder with course structure in XML format
const CoursePath string = "course/"

// ParseCourse parses course structure recursively from root to leaves
func ParseCourse(courseStructurePath string) (Course, error) {
	err := decompressCourseStructure(courseStructurePath)
	if err != nil {
		return Course{}, err
	}
	course, err := parseCourseRoot()
	if err != nil {
		return Course{}, err
	}
	err = cleanCourseStructure()
	return course, err
}

func decompressCourseStructure(courseStructurePath string) error {
	f, err := os.Open(courseStructurePath)
	if err != nil {
		return err
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(gzf)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		name := header.Name

		nameSplitByDot := strings.Split(name, ".")
		extension := nameSplitByDot[len(nameSplitByDot)-1]

		switch header.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(name, 0777)
		case tar.TypeReg:
			if extension == "xml" {
				newFile, err := os.Create(name)
				if err != nil {
					log.Printf("Unable to create file %v\n", name)
					return err
				}
				_, err = io.Copy(newFile, tarReader)
				if err != nil {
					log.Printf("Unable to rewrite file %v\n", name)
					return err
				}
			}
		default:
			log.Printf("%s : %c %s %s\n",
				"Yikes! Unable to figure out type",
				header.Typeflag,
				"in file",
				name,
			)
		}
	}
	return nil
}

func cleanCourseStructure() error {
	err := os.RemoveAll("course/")
	if err != nil {
		return err
	}
	return nil
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
