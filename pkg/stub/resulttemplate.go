package stub

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
	"math/rand"
	"text/template"
	"time"
)

var funcMap = template.FuncMap{
	"randomString": randomString,
	"currentTime":  currentTime,
	"randomUUID":   randomUUID,
	"randomInt":    randomInt,
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func currentTime() string {
	return time.Now().Format(time.RFC3339)
}

func randomUUID() string {
	return uuid.New().String()
}

func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// TODO: Extract the above in an interface

type ResultTemplate struct {
	ResultSpec interface{}
	template   *template.Template
}

func (rt *ResultTemplate) init() error {
	// Serialize the unprocessedResult into Yaml
	rawYaml, err := yaml.Marshal(rt.ResultSpec)
	if err != nil {
		return fmt.Errorf("failed to serialize JSON: %w", err)
	}

	// Process the Yaml
	tmpl, err := template.New("ResultTemplate").Funcs(funcMap).Parse(string(rawYaml))
	if err != nil {
		return err
	}

	rt.template = tmpl
	return nil
}

func (rt *ResultTemplate) process(data interface{}) (interface{}, error) {
	var yamlResult bytes.Buffer
	err := rt.template.Execute(&yamlResult, data)
	if err != nil {
		return nil, err
	}

	// Deserialize the unprocessedResult back to an interface
	var result interface{}
	err = yaml.Unmarshal(yamlResult.Bytes(), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize Yaml: %w", err)
	}

	return result, nil
}
