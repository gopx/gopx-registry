package route

import (
	"net/http"
	"path"
	"strings"

	"gopx.io/gopx-registry/api/v1"
	"gopx.io/gopx-registry/pkg/log"
)

// GoPXRegistryRouter handles HTTP requests from GoPX Registry API service.
type GoPXRegistryRouter struct{}

func (gr GoPXRegistryRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("%s %s", strings.ToUpper(r.Method), r.RequestURI)
	processRoute(w, r)
}

func processRoute(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = sanitizeRoute(r.URL.Path)
	path := r.URL.Path

	switch {
	case strings.HasPrefix(path, "/api/v1/"):
		v1.Handle(w, r)
	default:
		Error404(w, r)
	}
}

// Here requested route needs to be converted to lower case,
// which enables "/api/V1" is equivalent to "/api/v1" etc.
// and finally cleans the path e.g. end slashes would be removed from path
// e.g. "/api/" -> "/api" etc.
func sanitizeRoute(route string) string {
	return path.Clean(strings.ToLower(route))
}

// NewGoPXRegistryRouter returns a new GoPXRegistryRouter instance.
func NewGoPXRegistryRouter() *RegistryRouter {
	return &RegistryRouter{}
}
