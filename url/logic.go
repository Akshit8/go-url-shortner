package url

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
)

var (
	// ErrRedirectNotFound error when redirect not found
	ErrRedirectNotFound = errors.New("redirect not found")
	// ErrRedirectInvalid error for invalid code
	ErrRedirectInvalid = errors.New("redirect invalid")
)

type redirectService struct {
	redirectRepository RedirectRepository
}

// NewRedirectService creates a new instance of redirect service
func NewRedirectService(redirectRepository RedirectRepository) RedirectService {
	return &redirectService{redirectRepository: redirectRepository}
}

func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepository.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	validate := validator.New()
	if err := validate.Struct(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now()
	return r.redirectRepository.Store(redirect)
}
