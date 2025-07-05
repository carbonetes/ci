package helper

import (
	"encoding/xml"

	"gopkg.in/yaml.v2"
)

func ToXML(data interface{}) ([]byte, error) {
	return xml.Marshal(data)
}

func ToYAML(t interface{}) ([]byte, error) {
	return yaml.Marshal(t)
}

func WriteYAML(t interface{}, path string) error {
	b, err := ToYAML(t)
	if err != nil {
		return err
	}
	return WriteFile(b, path)
}
