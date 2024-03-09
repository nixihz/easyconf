package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

func main() {
	// input params
	runTimeEnv := flag.String("r", "", "env: local, k8s")
	templateFile := flag.String("t", "", "path to template file")
	dataFilePath := flag.String("f", "", "path to data file")
	outputFilePath := flag.String("o", "", "path to output file")

	flag.Parse()

	// params check
	if *runTimeEnv == "" || *templateFile == "" || *dataFilePath == "" || *outputFilePath == "" {
		flag.Usage()
		log.Println("example:")
		log.Println("easyconf -r local -t template.yaml -f data.yaml -o generateFile.yaml")
		return
	}

	// read template file
	templateContent, err := os.ReadFile(*templateFile)
	if err != nil {
		log.Println("Failed to read template file:", err)
		return
	}

	data := parseValue(dataFilePath)
	if *runTimeEnv == "k8s" {
		// k8s output configmap.yaml
		err = generateFile(templateContent, data[*runTimeEnv], *outputFilePath)
	} else {
		// local convert configmap.yaml to multiple files
		local(templateContent, data[*runTimeEnv], *outputFilePath)
	}

}

type ConfigMap struct {
	Data map[string]string "yaml:`data`"
}

// local env
func local(templateContent []byte, dataMap map[string]string, outputFilePath string) {
	// parse yaml file
	var configMap ConfigMap
	err := yaml.Unmarshal(templateContent, &configMap)
	if err != nil {
		log.Println("parse configmap.yaml error:", err)
		return
	}

	for fname, value := range configMap.Data {
		err := generateFile([]byte(value), dataMap, outputFilePath+fname)
		if err != nil {
			log.Println(err)
		}
	}
}

// parse template and combine data
func generateFile(templateContent []byte, dataMap map[string]string, outputFilePath string) error {
	// parse template
	tmpl := template.Must(template.New("yaml").Parse(string(templateContent)))

	// create generateFile file
	output, err := os.Create(outputFilePath)
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to create generateFile file:", err))
	}
	defer output.Close()

	// Render the data using a template and write the results to an output file
	err = tmpl.Execute(output, dataMap)
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to execute generateFile file:", err))
	}

	log.Println("generate success:", outputFilePath)
	return nil
}

// parseValue parse configuration data file
func parseValue(dataFilePath *string) map[string]map[string]string {
	// read content
	dataContent, err := os.ReadFile(*dataFilePath)
	if err != nil {
		log.Println("Failed to read data file:", err)
		return nil
	}
	// parse YAML data
	var dataMap map[string]map[string]string
	err = yaml.Unmarshal(dataContent, &dataMap)
	if err != nil {
		log.Fatal("Failed to unmarshal YAML data:", err)
	}

	return dataMap
}
