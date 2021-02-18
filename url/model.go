package url

import "time"

// Redirect struct defines model url
type Redirect struct {
	Code      string    `json:"code" bson:"code" msgpack:"code"`
	URL       string    `json:"url" bson:"url" msgpack:"url" validate:"required,url"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" msgpack:"created_at"`
}
