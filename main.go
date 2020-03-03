package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	xj "github.com/basgys/goxml2json"
)

func main() {
	// get the arguments from the command line
	numPtr := flag.Int("n", 4, "an integer")
	flag.Parse()

	sourceFile := flag.Arg(0)

	if sourceFile == "" {
		fmt.Println("Dude, you didn't pass in a tar file!")
		os.Exit(1)
	}

	fmt.Println("arg 1: ", flag.Arg(0))

	processFile(sourceFile, *numPtr)
}

func processFile(srcFile string, num int) {
	f, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tarReader := tar.NewReader(gzf)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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
					log.Fatal(err)
				}
				_, err = io.Copy(newFile, tarReader)
				if err != nil {
					log.Printf("Unable to rewrite file %v\n", name)
					panic(err)
				}
				json, err := xj.Convert(tarReader)
				if err != nil {
					panic("That's embarrassing...")
				}

				fmt.Println(json.String())
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
}
