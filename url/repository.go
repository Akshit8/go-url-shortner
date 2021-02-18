package url

// RedirectRepository defines persistence methods for a Redirect
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
