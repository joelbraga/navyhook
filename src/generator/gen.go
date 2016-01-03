package generator

import (

	"text/template"
	"os"
	"fmt"
)

func WriteTemplate(t interface{}, templateFile, outputFile string) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("Error parsing template file %s: %v", templateFile, err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("Error creating file %s: %v", outputFile, err)
	}

	err = tmpl.Execute(file, t)
	if err != nil {
		return fmt.Errorf("Error executing template file %s: %v", templateFile, err)
	}

	return nil
}