package shortener

type ShortenerRepository interface {
	Save(entry *Entry) error
	Find(path string) (*Entry, error)
}