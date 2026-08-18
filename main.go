package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", redHandler)

    log.Printf("Listening on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}

func redHandler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        hostname = "unknown"
    }
    html := fmt.Sprintf(`<html>
<body style="background-color:red">
<h1>Welcome to GKE! Version 2.0</h1>
<p>Hostname: %s</p>
</body>
</html>`, hostname)
    fmt.Fprint(w, html)
}
