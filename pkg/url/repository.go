package url

// Repository defines storage interface for url
type Repository interface {
	Save(url *URL) (*URL, error)
	Get(code string) (*URL, error)
	GetAll() ([]*URL, error)
	Update(code string, url *URL) (*URL, error)
	Delete(code string) error
}
