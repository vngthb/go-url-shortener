package url_shortener

type Generator func(string) string

type Timestamper func() int64

type service struct {
	Repository   Repository
	NewId        Generator
	NewTimestamp Timestamper
}

func NewService(options ...func(*service)) *service {
	service := &service{}
	for _, option := range options {
		option(service)
	}
	return service
}

func WithRepo(repository Repository) func(*service) {
	return func(service *service) {
		service.Repository = repository
	}
}

func WithGeneratorFunc(generator Generator) func(*service) {
	return func(service *service) {
		service.NewId = generator
	}
}

func WithTimestampFunc(timestamper Timestamper) func(*service) {
	return func(service *service) {
		service.NewTimestamp = timestamper
	}
}

func (service *service) Shorten(url string) (Entry, error) {
	entry := NewEntry(url, service.NewId(url), service.NewTimestamp())
	if err := service.Repository.Save(entry); err != nil {
		return Entry{}, err
	}
	return entry, nil
}

func (service *service) Redirect(path string) (string, error) {
	entry, err := service.Repository.Find(path)
	if err != nil {
		return "", err
	}
	return entry.Url(), nil
}
