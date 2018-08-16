package v1

import (
	"net/http"
	"strings"
)

// Handle parses the request route and transfers the request to
// the respective controller.
func Handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case "/api/v1/package/register":
		switch method {
		case "POST":
			packageRegisterPOST(w, r)
		default:
			Error405(w, r)
		}
	default:
		Error404(w, r)
	}
}

// Here param path and route should be cleaned.
func matchRoute(path, route string) bool {
	if route == "/" {
		return path == "/"
	}
	return route == path || strings.HasPrefix(path, route+"/")
}
