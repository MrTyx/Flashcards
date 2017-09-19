package main

import (
  "appengine"
  "appengine/datastore"
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "strconv"
  // "math"
  "net/http"
  "time"
)

func getTimestamp(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%d", time.Now().Unix())
}

func createNewProgress(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  ctx := appengine.NewContext(r)
  due := time.Now().Unix()
  due += 86400
  p := PROGRESS{
    UID: vars["uid"],
    Code: vars["code"],
    Due: due,
    Time: 86400,
  }
  key := datastore.NewIncompleteKey(ctx, "Progress", nil)
  if _, err := datastore.Put(ctx, key, &p); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }


  var user []USER
  q := datastore.NewQuery("User").Filter("UID =", vars["uid"]).Limit(1)
  k, err := q.GetAll(ctx, &user)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if len(user) == 1 {
    datastore.Delete(ctx, k[0])
    user[0].Order = user[0].Order + 1
  } else {
    user = append(user, USER{
      UID: vars["uid"],
      Order: 2,
    })
  }
  nik := datastore.NewIncompleteKey(ctx, "User", nil)
  if _, err := datastore.Put(ctx, nik, &user[0]); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Fprintf(w, "Done")
}

func updateProgress(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  ctx := appengine.NewContext(r)
  // code, ratio, uid
  q := datastore.NewQuery("Progress").
        Filter("Code =", vars["code"]).
        Filter("UID =", vars["uid"]).
        Limit(1)
  var progress []PROGRESS
  k, _ := q.GetAll(ctx, &progress)
  ratio, _ := strconv.ParseFloat(vars["ratio"], 64)
  ratio = ratio * float64(progress[0].Time)
  due := time.Now().Unix() + int64(ratio)
  progress[0].Time = int64(ratio)
  progress[0].Due = due
  datastore.Delete(ctx, k[0])
  key := datastore.NewIncompleteKey(ctx, "Progress", nil)
  datastore.Put(ctx, key, &progress[0])
  json.NewEncoder(w).Encode(progress)
}

func getDueFlashcards(w http.ResponseWriter, r *http.Request) {
  var progress []PROGRESS
  var flags []FLAG
  vars := mux.Vars(r)
  ctx := appengine.NewContext(r)
  now := time.Now().Unix()
  query := datastore.NewQuery("Progress").Filter("UID =", vars["uid"]).Filter("Due <=", now).Limit(10)
  query.GetAll(ctx, &progress)
  for _, p := range progress {
      var temp []FLAG
    q := datastore.NewQuery("Flag").Filter("Code =", p.Code)
    q.GetAll(ctx, &temp)
    flags = append(flags, temp[0])
  }
  json.NewEncoder(w).Encode(flags)
}

func getFlagsByOrder(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  order, _  := strconv.ParseInt(vars["order"], 10, 64)
  ctx := appengine.NewContext(r)
  q := datastore.NewQuery("Flag").
        Filter("Order >=", order).
        Order("Order").
        Limit(10)
  flags := make([]FLAG, 0, 10)
  if _, err := q.GetAll(ctx, &flags); err != nil {
    //
  }
  json.NewEncoder(w).Encode(flags)
}
