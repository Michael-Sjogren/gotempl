package mytypes

import "github.com/a-h/templ"

type MenuLink struct {
	Url    templ.SafeURL
	Name   string
	Active bool
}
