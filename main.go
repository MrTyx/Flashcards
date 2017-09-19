package main

import (
  "appengine"
  "appengine/datastore"
  "appengine/user"
  "html/template"
  "github.com/gorilla/mux"
  "net/http"
)

func init() {
  r := mux.NewRouter()

  // main.go
  r.HandleFunc("/", welcome)
  r.HandleFunc("/study", study)
  r.HandleFunc("/review", review)

  // api.go
  r.HandleFunc("/time", getTimestamp)
  r.HandleFunc("/study/{code}/{uid}", createNewProgress)
  r.HandleFunc("/review/{code}/{ratio}/{uid}", updateProgress)
  r.HandleFunc("/due/{uid}", getDueFlashcards)
  // r.HandleFunc("/flag/{code}", getFlagByCode)
  r.HandleFunc("/order/{order}", getFlagsByOrder)

  // seed.go
  r.HandleFunc("/seed", seedDatastore)

  http.Handle("/", r)
}

func welcome(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  u := user.Current(ctx)
  if u == nil {
    url, _ := user.LoginURL(ctx, "/")
    t := template.Must(template.ParseFiles("templates/login.html"))
    t.Execute(w, url);
    return
  }
  url, _ := user.LogoutURL(ctx, "/")
  d := Data {
    LogoutURL: url,
    Email: u.Email,
    ID: u.ID,
  }
  t := template.Must(template.ParseFiles("templates/index.html"))
  t.Execute(w, d);
}

func study(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  u := user.Current(ctx)
  if u == nil {
    url, _ := user.LoginURL(ctx, "/")
    t := template.Must(template.ParseFiles("templates/login.html"))
    t.Execute(w, url);
    return
  }

  url, _ := user.LogoutURL(ctx, "/")
  d := Data {
    LogoutURL: url,
    Email: u.Email,
    ID: u.ID,
    Order: 1,
  }

  var user []USER
  query := datastore.NewQuery("User").Filter("UID =", u.ID).Limit(1)
  query.GetAll(ctx, &user)
  if len(user) == 1 {
    d.Order = user[0].Order
  }

  t := template.Must(template.ParseFiles("templates/study.html"))
  t.Execute(w, d);
}

func review(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  u := user.Current(ctx)
  if u == nil {
    url, _ := user.LoginURL(ctx, "/")
    t := template.Must(template.ParseFiles("templates/login.html"))
    t.Execute(w, url)
    return
  }

  url, _ := user.LogoutURL(ctx, "/")
  d := Data {
    LogoutURL: url,
    Email: u.Email,
    ID: u.ID,
  }

  t := template.Must(template.ParseFiles("templates/review.html"))
  t.Execute(w, d);
}
