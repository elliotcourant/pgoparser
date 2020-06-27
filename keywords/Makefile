.PHONY: generated

generated:
	keywordgen --input=./keywords.yaml --keyword-template=./templates/keywords.template --keyword-map-template=./templates/keyword_map.template
	go fmt ./keywords.generated.go