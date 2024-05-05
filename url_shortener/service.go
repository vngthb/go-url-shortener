package url_shortener

type IService interface {
	Shorten(url string) (Entry, error)
	Redirect(code string) (string, error)
}
