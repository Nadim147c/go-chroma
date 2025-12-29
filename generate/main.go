package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const file = `material.tar.gz`

func main() {
	file, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decompresser, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer decompresser.Close()

	arc := tar.NewReader(decompresser)

	buf := bytes.NewBuffer(nil)

	if err := os.MkdirAll("num", 0o755); err != nil {
		log.Fatal(err)
	}

	for {
		header, err := arc.Next()
		if err == io.EOF {
			break // end of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		buf.Reset()

		if header.Typeflag != tar.TypeReg {
			continue
		}

		root := strings.IndexRune(header.Name, '/')
		if root == -1 {
			continue
		}
		fmt.Println(header.Name[root+1:])
		mod, name, found := strings.Cut(header.Name[root+1:], "/")

		if !found || mod != "color" && mod != "num" {
			continue
		}

		fmt.Println("processing", header.Name)

		if _, err := io.Copy(buf, arc); err != nil {
			log.Fatal(err)
		}

		if mod == "color" {
			process(buf, name, "material/v2/num", "go-chroma/num")
		}

		if mod == "num" {
			out, err := os.Create("num/" + name)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			if _, err := io.Copy(out, buf); err != nil {
				log.Fatal(err)
			}
		}
	}
}
