package filedb

type Translator interface {
	Save(path string, obj any) error
	Load(path string, obj any) error
}
