package main

import (
	"net/http"
	"time"
)

type UrlShortner struct {
	store       Store
	idGenerator IdGenerator
}

func NewUrlShortner(store Store, idGenerator IdGenerator) *UrlShortner {
	return &UrlShortner{
		store:       store,
		idGenerator: idGenerator,
	}
}

func (u *UrlShortner) ShortenURL(longUrl string) string {
	id := u.idGenerator.generate()

	urlEntry := UrlEntry{
		longUrl:        longUrl,
		expirationTime: time.Now().AddDate(1, 0, 0),
		createdAt:      time.Now(),
	}
	u.store.save(id, urlEntry)

	return id
}

func (u *UrlShortner) Redirect(shortUrl string) (int, string) {
	urlEntry, ok := u.store.get(shortUrl)
	if !ok {
		return http.StatusNotFound, ""
	}
	if time.Now().After(urlEntry.expirationTime) {
		return http.StatusGone, ""
	}

	return http.StatusMovedPermanently, urlEntry.longUrl
}
