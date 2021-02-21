package redirect

// Repository defines storage interface
type Repository interface {
	Find(code string) (string, error)
}
