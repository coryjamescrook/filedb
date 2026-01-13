package filedb

import "os"

func read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func write(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
