package filedb

type Config struct {
	Path            string
	DefaultFileData string
	ModelObj        model
	Translator      Translator
}
