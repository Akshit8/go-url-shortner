package url

type urlService struct {
	urlRepository Repository
}

// NewURLService is
func NewURLService(urlRepo Repository) Service {
	return &urlService{urlRepository: urlRepo}
}

func (u *urlService) Save(url *URL) (*URL, error) {
	return nil, nil
}

func (u *urlService) Get(code string) (*URL, error) {
	return nil, nil
}
