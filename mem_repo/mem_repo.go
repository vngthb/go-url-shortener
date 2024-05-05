package mem_repo

import (
	"errors"
	"go-url-shortener/url_shortener"
)

type MemoryRepo struct {
	entries []url_shortener.Entry
}

func NewRepo() *MemoryRepo {
	return &MemoryRepo{}
}

func (memoryRepo *MemoryRepo) Save(entry url_shortener.Entry) error {
	memoryRepo.entries = append(memoryRepo.entries, entry)
	return nil
}

func (memoryRepo *MemoryRepo) Find(path string) (url_shortener.Entry, error) {
	for _, entry := range memoryRepo.entries {
		if entry.Path() == path {
			return entry, nil
		}
	}
	return url_shortener.Entry{}, errors.New("path not found")
}
