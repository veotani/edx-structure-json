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
	course := Course{}
	err := parseByStructure(&course, CoursePath+"/course/course.xml")
	if err != nil {
		return Course{}, err
	}
	for chapterNum := range course.Chapters {
		chapter := Chapter{}
		err = parseByStructure(&chapter, CoursePath+"/chapter/"+course.Chapters[chapterNum].URLName+".xml")
		chapter.URLName = course.Chapters[chapterNum].URLName
		course.Chapters[chapterNum] = chapter
		if err != nil {
			return Course{}, err
		}

		for sequentialNum := range chapter.Sequentials {
			sequential := Sequential{}
			err = parseByStructure(&sequential, CoursePath+"/sequential/"+chapter.Sequentials[sequentialNum].URLName+".xml")
			sequential.URLName = chapter.Sequentials[sequentialNum].URLName
			course.Chapters[chapterNum].Sequentials[sequentialNum] = sequential
			if err != nil {
				return Course{}, err
			}

			for verticalNum := range sequential.Verticals {
				vertical := Vertical{}
				err = parseByStructure(&vertical, CoursePath+"/vertical/"+sequential.Verticals[verticalNum].URLName+".xml")
				vertical.URLName = sequential.Verticals[verticalNum].URLName
				course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum] = vertical
				if err != nil {
					return Course{}, err
				}

				for htmlNum := range vertical.Htmls {
					html := Html{}
					err = parseByStructure(&html, CoursePath+"/html/"+vertical.Htmls[htmlNum].URLName+".xml")
					html.URLName = vertical.Htmls[htmlNum].URLName
					course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum].Htmls[htmlNum] = html
					if err != nil {
						return Course{}, err
					}
				}

				for problemNum := range vertical.Problems {
					problem := Problem{}
					err = parseByStructure(&problem, CoursePath+"/problem/"+vertical.Problems[problemNum].URLName+".xml")
					problem.URLName = vertical.Problems[problemNum].URLName
					course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum].Problems[problemNum] = problem
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

				for libraryContentNum := range vertical.LibraryContents {
					libraryContent := LibraryContent{}
					err = parseByStructure(&libraryContent, CoursePath+"/library_content/"+vertical.LibraryContents[libraryContentNum].URLName+".xml")
					libraryContent.URLName = vertical.LibraryContents[libraryContentNum].URLName
					course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum].LibraryContents[libraryContentNum] = libraryContent
					if err != nil {
						return Course{}, err
					}

					for problemNum := range libraryContent.Problems {
						problem := Problem{}
						err = parseByStructure(&problem, CoursePath+"/problem/"+libraryContent.Problems[problemNum].URLName+".xml")
						problem.URLName = libraryContent.Problems[problemNum].URLName
						course.Chapters[chapterNum].Sequentials[sequentialNum].Verticals[verticalNum].LibraryContents[libraryContentNum].Problems[problemNum] = problem
						if err != nil {
							return Course{}, err
						}
					}
				}
			}
		}
	}
	return course, err
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
