package redirect

// Service defines actions available on redirect
type Service interface {
	Find(code string) (string, error)
}
