package shortener

type Entry struct {
	url       string
	path        string
	dateAdded int64
}

func (e *Entry) Url() string {
	return e.url
}

func (e *Entry) SetUrl(url string) {
	e.url = url
}

func (e *Entry) Path() string {
	return e.path
}

func (e *Entry) SetPath(path string) {
	e.path = path
}

func (e *Entry) DateAdded() int64 {
	return e.dateAdded
}

func (e *Entry) SetDateAdded(dateAdded int64) {
	e.dateAdded = dateAdded
}
