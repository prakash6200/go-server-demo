

package main

import (
    "fmt"
    "net/http"
    "log"
    "html"
)

func main() {
//     port := os.Getenv("PORT")
//     if port == "" {
//         port = "8080"
//     }

//     router := gin.Default()
//     router.GET("/", func(c *gin.Context) {
//         c.String(http.StatusOK, "Hello, Dude!")
//     })

//     fmt.Printf("Server listening on port %s...\n", port)
//     err := router.Run(":" + port)
//     if err != nil {
//         panic(err)
//     }
    
    http.HandleFunction("/", func(w http.ResponseWriter, r * http.Request){
        fmt.Fprintf(w, "Go server is working, %q ", html.EsceapeString(r.URL.Path))
    })
    
    http.HandleFunc("/wallet", func(w http.ResponseWriter, r* http.Request){
        fmt.Fprintf(w, "Wallet server is workin")   
    })
    
    log.fatal(http.ListenAndServe(":8081", nil))
}
