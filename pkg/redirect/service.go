package redirect

import "github.com/Akshit8/url-shortner/pkg/url"

// Service defines actions available on redirect
type Service interface {
	Find(code string) (*url.URL, error)
}
