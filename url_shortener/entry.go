package url_shortener

type Entry struct {
	url       string
	path      string
	dateAdded int64
}

func NewEntry(url, path string, dateAdded int64) Entry {
	return Entry{	
		url: url,
		path: path,
		dateAdded: dateAdded,
	}
}

func(entry Entry) Url() string {
	return entry.url
}

func(entry Entry) Path() string {
	return entry.path
}

func(entry Entry) DateAdded() int64 {
	return entry.dateAdded
}
