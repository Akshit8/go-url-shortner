package urls

// Repository defines storage interface for url
type Repository interface {
	Save(url *URL) (*URL, error)
	Get(code string) (*URL, error)
}
