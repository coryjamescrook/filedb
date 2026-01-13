package translators

import (
	"encoding/json"
)

type JsonTranslator struct{}

func (j JsonTranslator) Serialize(obj any) ([]byte, error) {
	return json.MarshalIndent(obj, "", "  ")
}

func (j JsonTranslator) Deserialize(data []byte, obj any) error {
	return json.Unmarshal(data, obj)
}
