package filedb

type Translator interface {
	Serialize(obj any) ([]byte, error)
	Deserialize(data []byte, obj any) error
}
