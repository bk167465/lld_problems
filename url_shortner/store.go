package main

type Store interface {
	save(shortUrl string, entry UrlEntry)
	get(shortUrl string) (UrlEntry, bool)
}

type InMemoryStorage struct {
	mp map[string]UrlEntry
}

func NewInMemoryStorage() InMemoryStorage {
	return InMemoryStorage{
		mp: make(map[string]UrlEntry),
	}
}

func (i InMemoryStorage) save(shortUrl string, entry UrlEntry) {
	i.mp[shortUrl] = entry
}

func (i InMemoryStorage) get(shortUrl string) (UrlEntry, bool) {
	entry, ok := i.mp[shortUrl]
	return entry, ok
}
