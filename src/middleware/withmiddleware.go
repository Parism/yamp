package middleware

import "net/http"

/*
Middleware interface
is used to provide additional functionality
to the view functions
*/
type Middleware func(http.HandlerFunc) http.HandlerFunc

/*
WithMiddleware function
Instead of going straight to the view function
we provide this entry function
that iterates along a list of middleware functions
before ending up in the actual function
E.g.
User tries to reach /secret
instead of writing http.HandleFunc("/secret",secret)
i will write
http.HandleFunc("/secret",middleware.WithMiddleware(
	views.SecretView,
middleware.WithLogin()))
***we can provide additional middlewares by writing
http.HandleFunc("/secret",middleware.WithMiddleware(
	views.SecretView,
middleware.WithLogin(),middleware.WithRateLimit()))
and so on and so forth
*/
func WithMiddleware(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
