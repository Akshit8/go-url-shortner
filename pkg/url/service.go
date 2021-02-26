package url

// Service defines available actions on url
type Service interface {
	Save(url *URL) (*URL, error)
	Get(code string) (*URL, error)
}
