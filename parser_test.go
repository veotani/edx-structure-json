package parser

// Course
//

import (
	"strconv"
	"testing"
)

func TestParseStructure(t *testing.T) {
	course, err := ParseCourse()
	if err != nil {
		t.Fatal(err)
	}

	if len(course.Chapters[0].Sequentials[0].Verticals[0].URLName) == 0 {
		t.Error("No vertical")
	}

	if course.DisplayName != "Академическое русское письмо" {
		t.Errorf("Incorrect course name")
	}
}

func TestParseProblem(t *testing.T) {
	course, err := ParseCourse()
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

func TestParseDiscussion(t *testing.T) {
	course, err := ParseCourse()
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
	course, err := ParseCourse()
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
	course, err := ParseCourse()
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
	course, err := ParseCourse()
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
