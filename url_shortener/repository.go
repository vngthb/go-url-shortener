package url_shortener

type Repository interface {
	Save(entry Entry) error
	Find(path string) (Entry, error)
}