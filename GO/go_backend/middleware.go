package main

import (
	"fmt"
	"net/http"
	"strings"
)
//--------- Handler Example --------
func HelloHandler(w http.ResponseWriter, r *http.Request) {	
	if r.URL.Path == "/panic"{
		panic("something bad happened")
	}

	w.Write([]byte(`Hello, world!`))
	fmt.Println("URL Path",r.URL.Path)
}

//----------- LogMiddleware -------------
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requsted URL:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

//----------- ErrorMiddleware -------------
func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, `something went wrong`, http.StatusInternalServerError)
				fmt.Println(`ERROR:`, err)

			}
		}()
		next.ServeHTTP(w, r)
	})

}

// -----------auth middleware -------------
//Authorization: ApiKey {insert apiKey here}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get("Authorization")
		
		if val == "" {
			http.Error(w,`no Authorization info found`,http.StatusUnauthorized)
			return  
		}

		vals := strings.Split(val, " ")
		if len(vals) != 2{
			http.Error(w,`malformed auth header`,http.StatusUnauthorized)
			return 
		}

		if vals[0] != "apiKey"{
			http.Error(w,`malformed first part of the auth header`,http.StatusUnauthorized)
			return 
		}

		next.ServeHTTP(w,r)
	})
}
