package main

type Course struct {
	DisplayName string
	ID          string
	Chapters    []Chapter
}

type Chapter struct {
	DisplayName string
	ID          string
	Sequentials []Sequential
}

type Sequential struct {
	DisplayName string
	ID          string
	Verticals   []Vertical
}

type Vertical struct {
	DisplayName     string
	ID              string
	Problems        []Problem
	Discussions     []Discussion
	Htmls           []Html
	OpenAssessments []OpenAssessment
	LibraryContents []LibraryContent
}

type LibraryContent struct {
	DisplayName string
	ID          string
	Problems    []Problem
}

type OpenAssessment struct {
	DisplayName string
	ID          string
}

type Html struct {
	DisplayName string
	ID          string
}

type Video struct {
	DisplayName string
	ID          string
}

type Discussion struct {
	ID          string
	DisplayName string
}

type Problem struct {
	ID          string
	DisplayName string
}
