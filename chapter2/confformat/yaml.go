package confformat

import (
	"bytes"
	"github.com/go-yaml/yaml"
)

// YAMLData is our common data struct
// with YAML struct tags
type YAMLData struct {
	Name string `yaml:"name"`
	Age int `yaml:"age"`
}

// ToYAML dumps the YAMLData struct to
// a YAML format bytes.Buffer
func (y *YAMLData) ToYAML() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(y)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)
	return b, nil
}

// Decode will decode into YAMLData
func (y *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, y)
}
