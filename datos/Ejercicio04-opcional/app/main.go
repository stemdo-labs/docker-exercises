// golang program to print ip address of the host or x-forwarded-for when behind a proxy
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!\n", r.RemoteAddr)
		if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
			fmt.Fprintf(w, "Your IP address is %s\n", r.Header.Get("X-Forwarded-For"))
		} else {
			fmt.Fprintf(w, "Undetected X-Forwarded-For header!\n")
		}
		fmt.Fprintf(w, "--------------------\nAll headers:\n--------------------\n")
		for k, v := range r.Header {
			fmt.Fprintf(w, "%s: %s\n", k, v)
		}
	})

	if port := os.Getenv("PORT"); port == "" {
		fmt.Println("PORT environment variable not set, defaulting to 8080")
		os.Setenv("PORT", "8080")
	}

	fmt.Println("Listening on port " + os.Getenv("PORT"))
	fmt.Println("Try it out: curl http://localhost:" + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}