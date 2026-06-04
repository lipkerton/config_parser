package parsers
import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type YAMLParser struct{}

func (p *YAMLParser) Parse(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}

	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("Ошибка парсинга YAML: %w", err)
	}

	return result, nil
}