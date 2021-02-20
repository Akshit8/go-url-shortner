package url

import "time"

// URL defines structure for entity url
type URL struct {
	Code      string    `json:"code" bson:"code" msgpack:"code"`
	URL       string    `json:"url" bson:"url" msgpack:"url" validate:"required,url"`
	ShortURL  string    `json:"shortUrl" bson:"shortUrl"  msgpack:"shortUrl" validate:"url"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt" msgpack:"createdAt"`
}
