package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func parseTemplates(p string) (*template.Template, error) {
	files := []string{}

	templ := template.New("")

	err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			files = append(files, path)
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	_, err = templ.ParseFiles(files...)

	if err != nil {
		return nil, err
	}

	return templ, nil
}
