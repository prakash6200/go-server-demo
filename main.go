

package main

import (
    "fmt"
    "net/http"
    "log"
    "html"
)

func main() {
    
   

    http.HandleFunc("/wallet", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "wallet server is working, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
