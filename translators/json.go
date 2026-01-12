package translators

import (
	"encoding/json"
	"os"
)

type JsonTranslator struct{}

func (j JsonTranslator) Save(path string, obj any) error {
	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (j JsonTranslator) Load(path string, obj any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, obj); err != nil {
		return err
	}

	return nil
}
