package main

import (
    "fmt"
    "log"
    "net/http"

    "around/backend"
    "around/handler"
    // jwtmiddleware "github.com/auth0/go-jwt-middleware"
    // jwt "github.com/form3tech-oss/jwt-go"

)

func main() {
    fmt.Println("started-service")

    backend.InitElasticsearchBackend()
    backend.InitGCSBackend()


    log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}



