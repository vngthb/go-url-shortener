package shortener

type ShortenerSerializer interface {
	Deserialize(input *Entry) ([]byte, error)
	Serialize(input []byte) (*Entry, error)
}
