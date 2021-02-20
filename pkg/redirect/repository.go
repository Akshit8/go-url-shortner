package redirect

import "github.com/Akshit8/url-shortner/pkg/url"

// Repository defines storage interface
type Repository interface {
	Find(code string) (*url.URL, error)
}
