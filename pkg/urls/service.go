package urls

// Service defines available actions on url
type Service interface {
	Save(url *URL) (*URL, error)
	Get(code string) (*URL, error)
}

type urlService struct {
	urlRepository Repository
}

// NewURLService is
func NewURLService(urlRepo Repository) Service {
	return &urlService{urlRepository: urlRepo}
}

func (u *urlService) Save(url *URL) (*URL, error) {
	return u.urlRepository.Save(url)
}

func (u *urlService) Get(code string) (*URL, error) {
	return u.urlRepository.Get(code)
}
