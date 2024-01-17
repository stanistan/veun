package handler

import "net/http"

// OnlyRoot will only apply the provided handler on the root URL path.
//
// This is usefull for mounting a handler at the '/' path.
//
//	http.Handle("/", OnlyRoot(...))
//
// It will 404 for anything else.
var OnlyRoot = MatchesPath(func(path string) bool {
	return path == "/"
})

var ExceptRoot = MatchesPath(func(path string) bool {
	return path != "/"
})

func MatchesPath(matches func(string) bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if matches(r.URL.Path) {
				next.ServeHTTP(w, r)
			} else {
				http.NotFound(w, r)
			}
		})
	}
}
