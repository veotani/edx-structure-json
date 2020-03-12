package parser

// Course
//

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
)

// func TestParseStructure(t *testing.T) {
// 	course, err := ParseCourse("course.aZ_uVd.tar.gz")
// 	if err != nil {
// 		t.Log(err)
// 	}

// 	if len(course.Chapters[0].Sequentials[0].Verticals[0].URLName) == 0 {
// 		t.Log("No vertical")
// 	}

// 	if course.DisplayName != "Академическое русское письмо" {
// 		t.Log("Incorrect course name")
// 	}
// }

func TestParseProblem(t *testing.T) {
	course, err := ParseCourse("examples/course.yEitvN.tar.gz")
	if err != nil {
		t.Errorf("%v", err)
	}

	for _, chapter := range course.Chapters {
		for _, sequential := range chapter.Sequentials {
			for _, vertical := range sequential.Verticals {
				for _, problem := range vertical.Problems {
					// All problems have name with length > 0 (edX restriction)
					if len(problem.DisplayName) == 0 {
						t.Fail()
					}

					// All problems have id with length > 0 (edX restriction)
					if len(problem.URLName) == 0 {
						t.Fail()
					}
				}
			}
		}
	}
}

func TestMarshal(t *testing.T) {
	course := Course{
		DisplayName: "course",
		Chapters: []Chapter{Chapter{
			DisplayName: "chapter",
			URLName:     "chapter_url",
			Sequentials: []Sequential{Sequential{
				DisplayName: "sequential",
				URLName:     "sequential_url",
				Verticals: []Vertical{Vertical{
					DisplayName: "vertical",
					Problems: []Problem{Problem{
						DisplayName: "problem",
					}},
				}},
			}},
		}},
	}

	jsonBytes, err := json.Marshal(course)
	if err != nil {
		t.Error(err)
	}

	unmarshaledCourse := Course{}
	json.Unmarshal(jsonBytes, &unmarshaledCourse)

	if unmarshaledCourse.DisplayName != "course" {
		t.Error("Root level data is invalid")
	}

	if unmarshaledCourse.Chapters[0].DisplayName != "chapter" {
		t.Error("Chapter level data is invalid")
	}

	if unmarshaledCourse.Chapters[0].Sequentials[0].DisplayName != "sequential" {
		t.Error("Sequential level data is invalid")
	}

	if unmarshaledCourse.Chapters[0].Sequentials[0].Verticals[0].DisplayName != "vertical" {
		t.Error("Vertical level data is invalid")
	}

	if unmarshaledCourse.Chapters[0].Sequentials[0].Verticals[0].Problems[0].DisplayName != "problem" {
		t.Error("Problem level data is invalid")
	}
}

func TestNoDirectoryAfterParserRun(t *testing.T) {
	_, err := ParseCourse("examples/course.yEitvN.tar.gz")
	if err != nil {
		t.Errorf("%v", err)
	}

	if exists("course/") {
		t.Error("Directory isn't cleaned after method run")
	}
}

func exists(filePath string) (exists bool) {
	exists = true

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exists = false
	}

	return
}

func TestParseDiscussion(t *testing.T) {
	course, err := ParseCourse("course.aZ_uVd.tar.gz")
	if err != nil {
		t.Errorf("%v", err)
	}

	for _, chapter := range course.Chapters {
		for _, sequential := range chapter.Sequentials {
			for _, vertical := range sequential.Verticals {
				for _, discussion := range vertical.Discussions {
					// All discussions have name with length > 0 (edX restriction)
					if len(discussion.DisplayName) == 0 {
						t.Fail()
					}

					// All discussions have id with length > 0 (edX restriction)
					if len(discussion.URLName) == 0 {
						t.Fail()
					}
				}
			}
		}
	}
}

func TestParseHtml(t *testing.T) {
	course, err := ParseCourse("course.aZ_uVd.tar.gz")
	if err != nil {
		t.Errorf("%v", err)
	}

	for _, chapter := range course.Chapters {
		for _, sequential := range chapter.Sequentials {
			for _, vertical := range sequential.Verticals {
				for _, html := range vertical.Htmls {
					// All htmls have name with length > 0 (edX restriction)
					if len(html.DisplayName) == 0 {
						t.Fail()
					}

					// All htmls have id with length > 0 (edX restriction)
					if len(html.URLName) == 0 {
						t.Fail()
					}
				}
			}
		}
	}
}

func TestParseOpenAssessment(t *testing.T) {
	course, err := ParseCourse("course.aZ_uVd.tar.gz")
	if err != nil {
		t.Errorf("%v", err)
	}

	for _, chapter := range course.Chapters {
		for _, sequential := range chapter.Sequentials {
			for _, vertical := range sequential.Verticals {
				for _, openAssessment := range vertical.OpenAssessments {
					// All discussions have id with length > 0 (edX restriction)
					if len(openAssessment.URLName) == 0 {
						t.Fail()
					}
				}
			}
		}
	}
}

func TestParseVideo(t *testing.T) {
	course, err := ParseCourse("course.aZ_uVd.tar.gz")
	if err != nil {
		t.Errorf("%v", err)
	}

	for _, chapter := range course.Chapters {
		for _, sequential := range chapter.Sequentials {
			for _, vertical := range sequential.Verticals {
				for _, video := range vertical.Videos {
					// All videos have id with length > 0 (edX restriction)
					if len(video.URLName) == 0 {
						t.Fail()
					}

					// All videos have name with length > 0 (edX restriction)
					if len(video.DisplayName) == 0 {
						t.Fail()
					}

					// All discussions have id with length > 0 (edX restriction)
					if len(video.URLName) == 0 {
						t.Fail()
					}

					// All videos have duration with length > 0
					if len(video.Duration) == 0 {
						t.Fail()
					}

					// Video duration must be convertable to float
					if _, err := strconv.ParseFloat(video.Duration, 32); err != nil {
						t.Fail()
					}
				}
			}
		}
	}
}
