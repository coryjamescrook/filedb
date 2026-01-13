package filedb

import (
	"os"
	"path/filepath"
	"sync"
)

type model interface {
	Path() string
	Load() error
	Save() error
}

type Model struct {
	cfg *Config
	mu  *sync.Mutex
}

func (db *Model) Save() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := db.cfg.Translator.Serialize(db.cfg.ModelObj)
	if err != nil {
		return err
	}

	return write(db.Path(), data)
}

func (db *Model) Load() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := read(db.Path())
	if err != nil {
		return err
	}

	return db.cfg.Translator.Deserialize(data, db.cfg.ModelObj)
}

func (db *Model) Path() string {
	return db.cfg.Path
}

func (db *Model) initFile() error {
	path := db.cfg.Path

	// create dir path for db file
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// create the file if it doesn't exist already
	var isFileNew bool = true
	file, err := os.OpenFile(db.cfg.Path, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666)
	if err != nil {
		if os.IsExist(err) {
			isFileNew = false
		} else {
			return err
		}
	}
	defer file.Close()

	if isFileNew {
		_, err = file.WriteString(db.cfg.DefaultFileData)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Model) Configure(config *Config) error {
	db.cfg = config
	db.mu = new(sync.Mutex)

	// ensure file exists for this model
	if err := db.initFile(); err != nil {
		return err
	}

	// load the most recent data into the model
	return db.Load()
}
