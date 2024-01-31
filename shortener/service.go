package shortener

type ShortenerService interface {
	Shorten(url string) (*Entry, error)
	Redirect(code string) (string, error)
}

type Generator func(string) string

type Timestamper func() int64

type shortenerService struct {
	shortenerRepository ShortenerRepository
	NewId               Generator
	NewTimestamp        Timestamper
}

func New(options ...func(*shortenerService)) *shortenerService {
	service := &shortenerService{}
	for _, option := range options {
		option(service)
	}
	return service
}

func WithRepo(shortenerRepository ShortenerRepository) func(*shortenerService) {
	return func(service *shortenerService) {
		service.shortenerRepository = shortenerRepository
	}
}

func WithGeneratorFunc(generator Generator) func(*shortenerService) {
	return func(service *shortenerService) {
		service.NewId = generator
	}
}

func WithTimestampFunc(timestamper Timestamper) func(*shortenerService) {
	return func(service *shortenerService) {
		service.NewTimestamp = timestamper
	}
}

func (service *shortenerService) Shorten(url string) (*Entry, error) {
	entry := &Entry{}
	entry.SetUrl(url)
	entry.SetPath(service.NewId(url))
	entry.SetDateAdded(service.NewTimestamp())
	if err := service.shortenerRepository.Save(entry); err != nil {
		return nil, err
	}
	return entry, nil
}

func (service *shortenerService) Redirect(path string) (string, error) {
	entry, err := service.shortenerRepository.Find(path)
	if err != nil {
		return "", err
	}
	return entry.Url(), nil
}