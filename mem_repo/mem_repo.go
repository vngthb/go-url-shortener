package mem_repo

import (
	"errors"
	"go-url-shortener/shortener"
)

type MemoryRepo struct {
	entries []shortener.Entry
}

func New() *MemoryRepo {
	return &MemoryRepo{}
}

func (memoryRepo *MemoryRepo) Save(entry *shortener.Entry) error {
	memoryRepo.entries = append(memoryRepo.entries, *entry)
	return nil
}

func (memoryRepo *MemoryRepo) Find(path string) (*shortener.Entry, error) {
	for _, entry := range memoryRepo.entries {
		if entry.Path == path {
			return &entry, nil
		}
	}
	return nil, errors.New("path not found")
}
