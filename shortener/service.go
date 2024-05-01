package shortener

type Service interface {
	Shorten(url string) (*Entry, error)
	Redirect(code string) (string, error)
}
