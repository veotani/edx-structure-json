# edx-structure-json
Golang library for converting edX XML structure to JSON.

# How to use
Install library: `go get github.com/veotani/edx-structure-json`

Import: `import edxparser "github.com/veotani/edx-structure-json"`

Parse: `course, err := edxparser.ParseCourse("course.tar.gz")` where `"course.tar.gz"` is a path to your exported 
course structure from edX studio.

Object `course` has type `Course`: 

![](https://github.com/veotani/edx-structure-json/blob/master/Course%20schema.png)


(see `model.go` for details). It's convertable to JSON by `encoding/json` Go package. 

Please notice, the parser only extracts object names and IDs. Few details are also parsed but there are not that much of them.
