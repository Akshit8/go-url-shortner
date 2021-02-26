package urls

// Serializer defines methods available for url serializer
type Serializer interface {
	Decode(input []byte) (*URL, error)
	Encode(input *URL) ([]byte, error)
}
