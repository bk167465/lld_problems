package main

import "time"

type UrlEntry struct {
	longUrl        string
	expirationTime time.Time
	createdAt      time.Time
}
