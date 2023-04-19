package middleware

// a place to declare all middleware functions
// every middleware function accepts http.HandlerFunc as an input and returns the same by calling passed function with ServeHTTP method

// example :

// func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
//	 return func(w http.ResponseWriter, r *http.Request) {
//		 log.Printf("[%s] %s\n", r.Method, r.URL.Path)
//		 next.ServeHTTP(w, r)
//	 }
// }
