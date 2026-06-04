package parsers
import (
	"fmt"
	"path/filepath"
	"strings"
)

func Parser interface {
	Parse(data []byte) (map[string]interface{}, error)
}

func GetParserByExtension(filename string) (Parser, error) {
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".json":
		return &JSONParser{}, nil
	case ".yaml", ".yml":
		return &YAMLParser{}, nil
	default:
		return nil, fmt.Errorf("Неподдерживаемое расширение файла: %s", ext)
	}
}