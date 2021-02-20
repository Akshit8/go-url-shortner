package redirect

import "github.com/Akshit8/url-shortner/pkg/url"

type redirectService struct {
	redirectRepository Repository
}

// NewRedirectService creates instance of redirectService
func NewRedirectService(redirectRepository Repository) Service {
	return &redirectService{redirectRepository: redirectRepository}
}

func (r *redirectService) Find(code string) (*url.URL, error) {
	return nil, nil
}
