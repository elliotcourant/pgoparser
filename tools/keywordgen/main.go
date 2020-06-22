package main

import (
	"bytes"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"text/template"
)

var (
	inputFile           = flag.String("input", "keywords.yaml", "Keywords YAML file.")
	keywordTemplateFile = flag.String("keyword-template", "keywords.template", "Keywords struct template file")
	keywordMapFile      = flag.String("keyword-map-template", "keyword_map.template", "Keywords map template file.")
	outputFile          = flag.String("output", "keywords.generated.go", "Output keywords file.")
)

type (
	KeywordFile struct {
		Keywords []string `yaml:"keywords"`
	}
)

func main() {
	flag.Parse()

	keywordsYamlData, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(err)
	}

	templateFile, err := ioutil.ReadFile(*keywordTemplateFile)
	if err != nil {
		panic(err)
	}

	keywordMapTemplatefile, err := ioutil.ReadFile(*keywordMapFile)
	if err != nil {
		panic(err)
	}

	var keywords KeywordFile
	if err := yaml.Unmarshal(keywordsYamlData, &keywords); err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)

	buf.WriteString(`package keywords

import (
    "strings"
)

`)

	keywordMapTemplate, err := template.New("map").Parse(string(keywordMapTemplatefile))
	if err != nil {
		panic(err)
	}

	if err := keywordMapTemplate.Execute(buf, keywords.Keywords); err != nil {
		panic(err)
	}

	keywordTemplate, err := template.New("keywords").Parse(string(templateFile))
	if err != nil {
		panic(err)
	}

	for _, keyword := range keywords.Keywords {
		if err := keywordTemplate.Execute(buf, struct {
			Keyword string
		}{
			Keyword: keyword,
		}); err != nil {
			panic(err)
		}
	}

	output, err := os.Create(*outputFile)
	if err != nil {
		panic(err)
	}

	output.Truncate(0)

	output.Write(buf.Bytes())

	output.Sync()

	output.Close()

	fmt.Println(buf.String())
}