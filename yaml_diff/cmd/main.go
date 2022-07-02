package main

import (
	"fmt"
	diff "go-play/yaml_diff"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {
	srcYaml, err := diff.ReadFile(diff.Source)
	if err != nil {
		log.Fatalf("error loading src file: %s", err)
	}

	var src interface{}
	err = yaml.Unmarshal(srcYaml, &src)
	if err != nil {
		log.Fatalf("error unmarshalling src file: %s", err)
	}

	destYaml, err := diff.ReadFile(diff.Sample)
	if err != nil {
		log.Fatalf("error loading dest file: %s", err)
	}
	var dest interface{}

	err = yaml.Unmarshal(destYaml, &dest)
	if err != nil {
		log.Fatalf("error unmarshalling dest file: %s", err)
	}

	fmt.Println(diff.Compare(src, dest))
}
