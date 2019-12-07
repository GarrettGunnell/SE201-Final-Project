package main

import (
    "log"
    "net/http"
	"os"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
	mux.HandleFunc("/players", players)
    mux.HandleFunc("/game", game)
    mux.HandleFunc("/move", move)
    mux.HandleFunc("/chat", wsEndpoint)

    fileServer := http.FileServer(http.Dir("ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Println("Starting server on :4000")
    err := http.ListenAndServe(getPort(), mux)
    log.Fatal(err)
}

func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":4000"
}
