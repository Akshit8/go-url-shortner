package url

import "time"

// URL defines structure for entity url
type URL struct {
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	ShortURL  string    `json:"shortUrl"`
	CreatedAt time.Time `json:"createdAt"`
}
