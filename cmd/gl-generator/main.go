package main

import (
	"embed"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed assets
var assets embed.FS

func main() {
	tpl, err := template.ParseFS(assets, "assets/*")
	must(err)
	tplFiles, err := assets.ReadDir("assets")
	must(err)
	systems := []string{
		"windows",
		"linux",
		"darwin",
	}

	for _, system := range systems {
		for _, tplFile := range tplFiles {
			must(generateGL(tpl, system, tplFile.Name()))

		}
	}
}

func generateGL(tpl *template.Template, system string, tplFileName string) error {
	targetFileName := strings.TrimSuffix(tplFileName, ".go.tpl")
	targetFilePath := filepath.Join("gl", targetFileName+"_"+system+".go")
	os.MkdirAll(filepath.Dir(targetFilePath), os.ModePerm)
	f, err := os.Create(targetFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return tpl.ExecuteTemplate(f, tplFileName, map[string]string{"OS": system})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
