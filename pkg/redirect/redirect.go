package redirect

type redirectService struct {
	redirectRepository Repository
}

// NewRedirectService creates instance of redirectService
func NewRedirectService(redirectRepository Repository) Service {
	return &redirectService{redirectRepository: redirectRepository}
}

func (r *redirectService) Find(code string) (string, error) {
	return "", nil
}
