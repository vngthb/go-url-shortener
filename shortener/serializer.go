package shortener

type Serializer interface {
	Deserialize(input *Entry) ([]byte, error)
	Serialize(input []byte) (*Entry, error)
}
