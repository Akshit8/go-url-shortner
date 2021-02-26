package resolver

import "github.com/Akshit8/url-shortner/pkg/url"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UrlService urls.Service
}
