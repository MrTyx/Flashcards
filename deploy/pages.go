package main

import (
  "appengine"
  "appengine/datastore"
  "appengine/user"
  "html/template"
  "math"
  "net/http"
  "time"
)

const BucketURL string = "https://s3-ap-southeast-2.amazonaws.com/s3394330-flashcards"

func home(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  u := user.Current(ctx)
  if u == nil {
    url, _ := user.LoginURL(ctx, "/")
    t := template.Must(template.ParseFiles("templates/login.html"))
    t.Execute(w, url);
    return
  }
  url, _ := user.LogoutURL(ctx, "/")

  var user []USER
  var progress []PROGRESS
  userQuery := datastore.NewQuery("User").Filter("UID =", u.ID).Limit(1)
  dueQuery := datastore.NewQuery("Progress").Filter("UID = ", u.ID).Filter("Due <=", time.Now().Unix())
  userQuery.GetAll(ctx, &user)
  dueQuery.GetAll(ctx, &progress)
  d := Data {
    LogoutURL: url,
    Email: u.Email,
    ID: u.ID,
    Due: len(progress),
  }
  if len(user) == 0 {
    d.Order = 0
  } else {
    d.Order = user[0].Order - 1
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
  d.BucketURL = BucketURL
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
  d.BucketURL = BucketURL
  t := template.Must(template.ParseFiles("templates/review.html"))
  t.Execute(w, d);
}

func progress(w http.ResponseWriter, r *http.Request) {
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
  var progress []PROGRESS
  q := datastore.NewQuery("Progress").Filter("UID =", u.ID).Order("Due")
  q.GetAll(ctx, &progress)
  for i, _ := range progress {
    progress[i].Due = int64(math.Ceil(float64(progress[i].Due - time.Now().Unix()) / 60))
    progress[i].Time = int64(math.Ceil(float64(progress[i].Time) / 60))
  }
  d.Progress = progress
  t := template.Must(template.ParseFiles("templates/progress.html"))
  t.Execute(w, d)
}


func stats(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  u := user.Current(ctx)
  if u == nil {
    url, _ := user.LoginURL(ctx, "/")
    t := template.Must(template.ParseFiles("templates/login.html"))
    t.Execute(w, url)
    return
  }
  url, _ := user.LogoutURL(ctx, "/")
  s := STATS {
    LogoutURL: url,
    Email: u.Email,
    ID: u.ID,
    Studies: 0,
    Reviews: 0,
    RCount: 0,
  }
  var users []USER
  var progress []PROGRESS
  q := datastore.NewQuery("User");
  q.GetAll(ctx, &users)
  s.Users = len(users);
  for i, _ := range users {
    s.Studies = s.Studies + (users[i].Order - 1)
  }
  s.SAverage = int(s.Studies / s.Users)
  query := datastore.NewQuery("Progress");
  query.GetAll(ctx, &progress);
  s.RCount = len(progress)
  for i, _ := range progress {
    s.Reviews += progress[i].Reviews;
  }
  s.RAverage = int(s.Reviews / s.Users);

  var hardest []PROGRESS
  var easiest []PROGRESS
  hardestQuery := datastore.NewQuery("Progress").Order("Time").Limit(1);
  hardestQuery.GetAll(ctx, &hardest);
  easiestQuery := datastore.NewQuery("Progress").Order("-Time").Limit(1);
  easiestQuery.GetAll(ctx, &easiest);
  s.Hardest = hardest;
  s.Easiest = easiest;
  t := template.Must(template.ParseFiles("templates/stats.html"))
  t.Execute(w, s)
}

func about(w http.ResponseWriter, r *http.Request) {
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
  t := template.Must(template.ParseFiles("templates/about.html"))
  t.Execute(w, d)
}
