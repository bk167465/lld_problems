package main

import "fmt"

func main() {
	urlShortner := NewUrlShortner(NewInMemoryStorage(), NewBase62IDGenerator())

	// Saving flow
	fmt.Println("Enter the long URL")

	var longUrl string
	fmt.Scan(&longUrl)

	shortUrl := urlShortner.ShortenURL(longUrl)

	fmt.Printf("Your short URL is %s\n", shortUrl)

	// Fetching Flow
	fmt.Println("Enter the short URL")

	var newShortUrl string
	fmt.Scan(&newShortUrl)

	status, newLongUrl := urlShortner.Redirect(newShortUrl)
	fmt.Println(status, newLongUrl)
}
