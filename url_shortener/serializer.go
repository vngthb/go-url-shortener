package url_shortener

type ISerializer interface {
	Deserialize(input *Entry) ([]byte, error)
	Serialize(input []byte) (*Entry, error)
}
