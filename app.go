package main

import (
    "fmt"
    "os"
    "net/http"
)

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
    myhostname, _ := os.Hostname()
    fmt.Fprintln(w, "Hostname:", myhostname)
}

func healtcheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Write([]byte("ok"))
}

func fivehunderHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(500)
    w.Write([]byte("error"))
}

func main() {
    const port string = "8000"
    fmt.Println("Server listening on port", port)
    http.HandleFunc("/", hostnameHandler)
    http.HandleFunc("/health", healtcheckHandler)
    http.HandleFunc("/bad", fivehunderHandler)
    http.ListenAndServe(":" + port, nil)
}
