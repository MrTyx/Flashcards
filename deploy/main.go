package main

import (
  "github.com/gorilla/mux"
  "net/http"
)

func init() {
  r := mux.NewRouter()

  // main.go
  r.HandleFunc("/", home)
  r.HandleFunc("/study", study)
  r.HandleFunc("/review", review)
  r.HandleFunc("/progress", progress)
  r.HandleFunc("/stats", stats)
  r.HandleFunc("/about", about)

  // api.go
  r.HandleFunc("/time", getTimestamp)
  r.HandleFunc("/study/{code}/{uid}", createNewProgress)
  r.HandleFunc("/review/{code}/{ratio}/{uid}", updateProgress)
  r.HandleFunc("/due/{uid}", getDueFlashcards)
  // r.HandleFunc("/flag/{code}", getFlagByCode)
  r.HandleFunc("/order/{order}", getFlagsByOrder)

  // seed.go
  // r.HandleFunc("/seed", seedDatastore)

  http.Handle("/", r)
}
